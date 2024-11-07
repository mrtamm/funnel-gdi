package server

import (
	"errors"
	"fmt"
	"strings"

	"github.com/ohsu-comp-bio/funnel/events"
	"github.com/ohsu-comp-bio/funnel/logger"
	"github.com/ohsu-comp-bio/funnel/metrics"
	"github.com/ohsu-comp-bio/funnel/tes"
	"github.com/ohsu-comp-bio/funnel/version"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// TaskService is a wrapper which handles common TES Task Service operations,
// such as initializing a task when CreateTask is called. The TaskService is backed by
// two parts: a read API which provides the GetTask and ListTasks endpoints, and a write
// API which implements the events.Writer interface. Task creation and cancelation is
// managed by writing events to underlying event writer.
//
// This makes it easier to define task service backends for new databases, and ensures
// that common operations are handled consistently, such as setting IDs, handling 404s,
// GetServiceInfo, etc.
type TaskService struct {
	Name    string
	Event   events.Writer
	Compute events.Writer
	Read    tes.ReadOnlyServer
	Log     *logger.Logger
}

// CreateTask provides an HTTP/gRPC endpoint for creating a task.
// This is part of the TES implementation.
func (ts *TaskService) CreateTask(ctx context.Context, task *tes.Task) (*tes.CreateTaskResponse, error) {
	if err := tes.InitTask(task, true); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if err := ReplaceInputBearerToken(ctx, task); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if err := ts.Event.WriteEvent(ctx, events.NewTaskCreated(task)); err != nil {
		return nil, fmt.Errorf("error creating task: %s", err)
	}

	// dispatch to compute backend
	go func() {
		err := ts.Compute.WriteEvent(ctx, events.NewTaskCreated(task))
		if err != nil {
			ts.Log.Error("error submitting task to compute backend", "taskID", task.Id, "error", err)
		}
	}()

	return &tes.CreateTaskResponse{Id: task.Id}, nil
}

// GetTask calls GetTask on the underlying tes.ReadOnlyServer. If the underlying server
// returns tes.ErrNotFound, TaskService will handle returning the appropriate gRPC error.
func (ts *TaskService) GetTask(ctx context.Context, req *tes.GetTaskRequest) (*tes.Task, error) {
	task, err := ts.Read.GetTask(ctx, req)
	if err == tes.ErrNotFound {
		err = status.Errorf(codes.NotFound, fmt.Sprintf("%v: taskID: %s", err.Error(), req.Id))
	}
	return task, err
}

// ListTasks calls ListTasks on the underlying tes.ReadOnlyServer.
func (ts *TaskService) ListTasks(ctx context.Context, req *tes.ListTasksRequest) (*tes.ListTasksResponse, error) {
	return ts.Read.ListTasks(ctx, req)
}

// CancelTask cancels a task
func (ts *TaskService) CancelTask(ctx context.Context, req *tes.CancelTaskRequest) (*tes.CancelTaskResponse, error) {
	// dispatch to compute backend
	err := ts.Compute.WriteEvent(ctx, events.NewState(req.Id, tes.Canceled))
	if err != nil {
		ts.Log.Error("compute backend failed to cancel task", "taskID", req.Id, "error", err)
	}

	// updated database and other event streams
	err = ts.Event.WriteEvent(ctx, events.NewState(req.Id, tes.Canceled))
	if err == tes.ErrNotFound {
		err = status.Errorf(codes.NotFound, fmt.Sprintf("%v: taskID: %s", err.Error(), req.Id))
	}
	return &tes.CancelTaskResponse{}, err
}

// GetServiceInfo returns service metadata.
func (ts *TaskService) GetServiceInfo(ctx context.Context, info *tes.ServiceInfoRequest) (*tes.ServiceInfo, error) {
	resp := &tes.ServiceInfo{
		Name: ts.Name,
		Doc:  version.String(),
	}

	if c, ok := ts.Read.(metrics.TaskStateCounter); ok {
		resp.TaskStateCounts = make(map[string]int32)
		// Ensure that all states are present in the response, even if zero.
		for key := range tes.State_value {
			resp.TaskStateCounts[key] = 0
		}
		cs, err := c.TaskStateCounts(ctx)
		if err != nil {
			ts.Log.Error("counting task states", "error", err)
		}
		// Override the zero values in the response.
		for key, count := range cs {
			resp.TaskStateCounts[key] = count
		}
	}

	return resp, nil
}

func ReplaceInputBearerToken(ctx context.Context, task *tes.Task) error {
	userInfo, ok := ctx.Value(UserInfoKey).(*UserInfo)
	noToken := !ok || userInfo.Token == ""

	for _, input := range task.Inputs {
		if !strings.HasPrefix(input.Url, "sda://") &&
			!strings.HasPrefix(input.Url, "htsget://") ||
			strings.Contains(input.Url, "#") {
			continue
		}
		if noToken {
			scheme, _, _ := strings.Cut(input.Url, "://")
			if scheme == "htsget" {
				continue
			}
			return errors.New("Task input from SDA requires a Bearer token " +
				"to be used for fetching the data, however, current " +
				"authentication-context has no information about the token " +
				"to use. If necessary, please provide an explicit Bearer " +
				"token in the URL right after the hash-sign ('#'): " +
				"sda://dataset-id/file/path#bearer-token")
		}
		input.Url = input.Url + "#" + userInfo.Token
	}

	return nil
}
