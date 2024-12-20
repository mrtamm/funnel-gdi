package elastic

import (
	"context"
	"encoding/json"
	"time"

	"github.com/ohsu-comp-bio/funnel/config"
	"github.com/ohsu-comp-bio/funnel/tes"
	"google.golang.org/protobuf/encoding/protojson"
	elastic "gopkg.in/olivere/elastic.v5"
)

var minimal = elastic.NewFetchSourceContext(true).Include("id", "state", "owner")
var basic = elastic.NewFetchSourceContext(true).
	Exclude("logs.logs.stderr", "logs.logs.stdout", "inputs.content", "logs.system_logs")
var unmarshalOpts = protojson.UnmarshalOptions{DiscardUnknown: true}

// Elastic provides an elasticsearch database server backend.
type Elastic struct {
	client    *elastic.Client
	conf      config.Elastic
	taskIndex string
	nodeIndex string
}

// NewElastic returns a new Elastic instance.
func NewElastic(conf config.Elastic) (*Elastic, error) {
	client, err := elastic.NewClient(
		elastic.SetURL(conf.URL),
		elastic.SetSniff(false),
		elastic.SetRetrier(
			elastic.NewBackoffRetrier(
				elastic.NewExponentialBackoff(time.Millisecond*50, time.Minute),
			),
		),
	)
	if err != nil {
		return nil, err
	}
	es := &Elastic{
		client,
		conf,
		conf.IndexPrefix + "-tasks",
		conf.IndexPrefix + "-nodes",
	}
	return es, nil
}

// Close closes the database client.
func (es *Elastic) Close() {
	es.client.Stop()
}

func (es *Elastic) initIndex(ctx context.Context, name, body string) error {
	exists, err := es.client.
		IndexExists(name).
		Do(ctx)

	if err != nil {
		return err
	} else if !exists {
		if _, err := es.client.CreateIndex(name).Body(body).Do(ctx); err != nil {
			return err
		}
	}
	return nil
}

// Init creates the Elasticsearch indices.
func (es *Elastic) Init() error {
	ctx := context.Background()
	taskMappings := `{
    "mappings": {
      "task":{
      "properties":{
        "id": {
          "type": "keyword"
        },
        "state": {
          "type": "keyword"
        },
        "owner": {
          "type": "keyword"
        },
        "inputs": {
          "type": "nested"
        },
        "logs": {
          "type": "nested",
          "properties": {
            "logs": {
              "type": "nested"
              }
            }
          }
        }
      }
    }
  }`
	if err := es.initIndex(ctx, es.taskIndex, taskMappings); err != nil {
		return err
	}
	if err := es.initIndex(ctx, es.nodeIndex, ""); err != nil {
		return err
	}
	return nil
}

func marshalTask(task *tes.Task) ([]byte, error) {
	return protojson.Marshal(task)
}

func unmarshalTask(msg json.RawMessage) (*tes.Task, error) {
	t := &tes.Task{}
	err := unmarshalOpts.Unmarshal(msg, t)
	return t, err
}
