package badger

import (
	"bytes"
	"context"
	"fmt"

	badger "github.com/dgraph-io/badger/v2"
	"github.com/ohsu-comp-bio/funnel/server"
	"github.com/ohsu-comp-bio/funnel/tes"
	"google.golang.org/protobuf/proto"
)

// GetTask gets a task, which describes a running task
func (db *Badger) GetTask(ctx context.Context, req *tes.GetTaskRequest) (*tes.Task, error) {
	var task *tes.Task

	err := db.db.View(func(txn *badger.Txn) error {
		t, err := getTask(txn, req.Id, ctx)
		task = t
		return err
	})
	if err != nil {
		return nil, err
	}

	switch req.View {
	case tes.Minimal:
		task = task.GetMinimalView()
	case tes.Basic:
		task = task.GetBasicView()
	}
	return task, nil
}

// ListTasks returns a list of tasks.
func (db *Badger) ListTasks(ctx context.Context, req *tes.ListTasksRequest) (*tes.ListTasksResponse, error) {
	var tasks []*tes.Task
	pageSize := tes.GetPageSize(req.GetPageSize())

	err := db.db.View(func(txn *badger.Txn) error {

		it := txn.NewIterator(badger.IteratorOptions{
			// Keys (task IDs) are in ascending order, and we want the first page
			// to be the most recent task, so that's at the end of the list.
			Reverse:        true,
			PrefetchValues: true,
			PrefetchSize:   pageSize,
		})
		defer it.Close()

		i := 0

		// For pagination, figure out the starting key.
		if req.PageToken != "" {
			it.Seek(taskKey(req.PageToken))
			// Seek moves to the key, but the start of the page is the next key.
			it.Next()
		} else {
			it.Rewind()
		}

	taskLoop:
		for ; it.Valid() && len(tasks) < pageSize; it.Next() {
			if !it.Valid() || !bytes.HasPrefix(it.Item().Key(), taskKeyPrefix) {
				break
			}

			if !isAccessible(txn, ownerKeyFromTaskKey(it.Item().Key()), ctx) {
				continue
			}

			var val []byte
			err := it.Item().Value(func(d []byte) error {
				val = copyBytes(d)
				return nil
			})
			if err != nil {
				return fmt.Errorf("loading item value: %s", err)
			}

			// Load task.
			task := &tes.Task{}
			err = proto.Unmarshal(val, task)
			if err != nil {
				return fmt.Errorf("unmarshaling data: %s", err)
			}

			// Filter tasks by tag.
			for k, v := range req.Tags {
				tval, ok := task.Tags[k]
				if !ok || tval != v {
					continue taskLoop
				}
			}

			// Filter tasks by state.
			if req.State != tes.Unknown && req.State != task.State {
				continue taskLoop
			}

			switch req.View {
			case tes.Minimal:
				task = task.GetMinimalView()
			case tes.Basic:
				task = task.GetBasicView()
			}

			tasks = append(tasks, task)
			i++
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	out := tes.ListTasksResponse{
		Tasks: tasks,
	}

	if len(tasks) == pageSize {
		out.NextPageToken = tasks[len(tasks)-1].Id
	}

	return &out, nil
}

func getTask(txn *badger.Txn, id string, ctx context.Context) (*tes.Task, error) {
	if !isAccessible(txn, ownerKey(id), ctx) {
		return nil, tes.ErrNotPermitted
	}

	item, err := txn.Get(taskKey(id))
	if err == badger.ErrKeyNotFound {
		return nil, tes.ErrNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("loading item: %s", err)
	}

	var val []byte
	err = item.Value(func(d []byte) error {
		val = copyBytes(d)
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("loading item value: %s", err)
	}

	task := &tes.Task{}
	err = proto.Unmarshal(val, task)
	if err != nil {
		return nil, fmt.Errorf("unmarshaling data: %s", err)
	}
	return task, nil
}

func isAccessible(txn *badger.Txn, ownerKey []byte, ctx context.Context) bool {
	userInfo, ok := ctx.Value(server.UserInfoKey).(*server.UserInfo)
	if !ok || userInfo.IsAdmin {
		return true
	}

	owner := ""
	if item, err := txn.Get(ownerKey); err == nil {
		_ = item.Value(func(d []byte) error {
			owner = string(d)
			return nil
		})
	}

	return userInfo.IsAccessible(owner)
}

func copyBytes(in []byte) []byte {
	out := make([]byte, len(in))
	copy(out, in)
	return out
}
