package elastic

import (
	"fmt"

	"github.com/ohsu-comp-bio/funnel/compute/scheduler"
	"github.com/ohsu-comp-bio/funnel/tes"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	elastic "gopkg.in/olivere/elastic.v5"
)

// ReadQueue returns a slice of queued Tasks. Up to "n" tasks are returned.
func (es *Elastic) ReadQueue(n int) []*tes.Task {
	ctx := context.Background()

	q := elastic.NewTermQuery("state", tes.State_QUEUED.String())
	res, err := es.client.Search().
		Index(es.taskIndex).
		Type("task").
		Size(n).
		Sort("id", true).
		Query(q).
		Do(ctx)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	var tasks []*tes.Task
	for _, hit := range res.Hits.Hits {
		t, err := unmarshalTask(*hit.Source)
		if err != nil {
			continue
		}

		t = t.GetBasicView()
		tasks = append(tasks, t)
	}

	return tasks
}

// GetNode gets a node
func (es *Elastic) GetNode(ctx context.Context, req *scheduler.GetNodeRequest) (*scheduler.Node, error) {
	res, err := es.client.Get().
		Index(es.nodeIndex).
		Type("node").
		Id(req.Id).
		Do(ctx)

	if elastic.IsNotFound(err) {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("%v: nodeID: %s", err.Error(), req.Id))
	}
	if err != nil {
		return nil, err
	}

	node := &scheduler.Node{}
	err = protojson.Unmarshal(*res.Source, node)
	if err != nil {
		return nil, err
	}
	// Must happen after the unmarshal
	node.Version = *res.Version
	return node, nil
}

// PutNode puts a node in the database.
//
// For optimisic locking, if the node already exists and node.Version
// doesn't match the version in the database, an error is returned.
func (es *Elastic) PutNode(ctx context.Context, node *scheduler.Node) (*scheduler.PutNodeResponse, error) {
	g := es.client.Get().
		Index(es.nodeIndex).
		Type("node").
		Preference("_primary").
		Id(node.Id)

		// If the version is 0, then this should be creating a new node.
	if node.GetVersion() != 0 {
		g = g.Version(node.GetVersion())
	}

	res, err := g.Do(ctx)

	if err != nil && !elastic.IsNotFound(err) {
		return nil, err
	}

	existing := &scheduler.Node{}
	if err == nil {
		if err2 := protojson.Unmarshal(*res.Source, existing); err2 != nil {
			fmt.Printf("Detected error while unmarshaling node info from HTTP response: %s\n", err2)
		}
	}

	err = scheduler.UpdateNode(ctx, es, node, existing)
	if err != nil {
		return nil, err
	}

	mar := protojson.MarshalOptions{}
	b, err := mar.Marshal(node)
	if err != nil {
		return nil, err
	}

	i := es.client.Index().
		Index(es.nodeIndex).
		Type("node").
		Id(node.Id).
		Refresh("true").
		BodyString(string(b))

	if node.GetVersion() != 0 {
		i = i.Version(node.GetVersion())
	}
	_, err = i.Do(ctx)

	return &scheduler.PutNodeResponse{}, err
}

// DeleteNode deletes a node by ID.
func (es *Elastic) DeleteNode(ctx context.Context, node *scheduler.Node) (*scheduler.DeleteNodeResponse, error) {
	_, err := es.client.Delete().
		Index(es.nodeIndex).
		Type("node").
		Id(node.Id).
		Version(node.Version).
		Refresh("true").
		Do(ctx)
	return &scheduler.DeleteNodeResponse{}, err
}

// ListNodes is an API endpoint that returns a list of nodes.
func (es *Elastic) ListNodes(ctx context.Context, req *scheduler.ListNodesRequest) (*scheduler.ListNodesResponse, error) {
	res, err := es.client.Search().
		Index(es.nodeIndex).
		Type("node").
		Version(true).
		Size(1000).
		Do(ctx)

	if err != nil {
		return nil, err
	}

	resp := &scheduler.ListNodesResponse{}
	for _, hit := range res.Hits.Hits {
		node := &scheduler.Node{}
		err = protojson.Unmarshal(*hit.Source, node)
		if err != nil {
			return nil, err
		}
		node.Version = *hit.Version
		resp.Nodes = append(resp.Nodes, node)
	}

	return resp, nil
}
