package worker

import (
	"context"
	"fmt"
	"github.com/ohsu-comp-bio/funnel/config"
	"github.com/ohsu-comp-bio/funnel/events"
	"github.com/ohsu-comp-bio/funnel/logger"
	"github.com/ohsu-comp-bio/funnel/proto/tes"
	"github.com/ohsu-comp-bio/funnel/server/boltdb"
	"github.com/ohsu-comp-bio/funnel/server/dynamodb"
	"github.com/ohsu-comp-bio/funnel/server/elastic"
	"github.com/ohsu-comp-bio/funnel/server/mongodb"
	"github.com/ohsu-comp-bio/funnel/storage"
	"github.com/ohsu-comp-bio/funnel/worker"
	"path"
	"strings"
)

// Run runs the "worker run" command.
func Run(ctx context.Context, conf config.Config, taskID string, log *logger.Logger) error {
	w, err := NewWorker(conf, taskID, log)
	if err != nil {
		return err
	}
	w.Run(ctx)
	return nil
}

// NewWorker returns a new Funnel worker based on the given config.
func NewWorker(conf config.Config, taskID string, log *logger.Logger) (*worker.DefaultWorker, error) {
	if log == nil {
		log = logger.NewLogger("worker", conf.Logger)
	}
	log.Debug("NewWorker", "config", conf, "taskID", taskID)

	ctx := context.Background()
	var err error
	var db tes.ReadOnlyServer
	var reader worker.TaskReader
	var writer events.Writer

	switch strings.ToLower(conf.Worker.TaskReader) {
	case "dynamodb":
		db, err = dynamodb.NewDynamoDB(conf.DynamoDB)
	case "elastic":
		db, err = elastic.NewElastic(ctx, conf.Elastic)
	case "mongodb":
		db, err = mongodb.NewMongoDB(conf.MongoDB)
	case "rpc":
		reader, err = worker.NewRPCTaskReader(conf.RPC, taskID)
	default:
		err = fmt.Errorf("unknown TaskReader")
	}
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate TaskReader: %v", err)
	}

	if reader == nil {
		reader = worker.NewGenericTaskReader(db.GetTask, taskID)
	}

	writers := []events.Writer{}
	eventWriterSet := make(map[string]interface{})
	if strings.ToLower(conf.Database) == "boltdb" {
		eventWriterSet["rpc"] = nil
	} else {
		eventWriterSet[strings.ToLower(conf.Database)] = nil
	}
	for _, w := range conf.EventWriters {
		eventWriterSet[strings.ToLower(w)] = nil
	}

	for w := range eventWriterSet {
		switch w {
		case "log":
			writer = &events.Logger{Log: log}
		case "boltdb":
			writer, err = boltdb.NewBoltDB(conf.BoltDB)
		case "dynamodb":
			writer, err = dynamodb.NewDynamoDB(conf.DynamoDB)
		case "elastic":
			writer, err = elastic.NewElastic(ctx, conf.Elastic)
		case "kafka":
			k, kerr := events.NewKafkaWriter(conf.Kafka)
			defer k.Close()
			err = kerr
			writer = k
		case "mongodb":
			writer, err = mongodb.NewMongoDB(conf.MongoDB)
		case "rpc":
			writer, err = events.NewRPCWriter(conf.RPC)
		default:
			err = fmt.Errorf("unknown EventWriter")
		}
		if err != nil {
			return nil, fmt.Errorf("failed to instantiate EventWriter: %v", err)
		}
		writers = append(writers, writer)
	}

	m := events.MultiWriter(writers)
	ew := &events.ErrLogger{Writer: &m, Log: log}

	s, err := storage.NewStorage(conf)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate Storage backend: %v", err)
	}

	w := &worker.DefaultWorker{
		Conf:       conf.Worker,
		Mapper:     worker.NewFileMapper(path.Join(conf.Worker.WorkDir, taskID)),
		Store:      s,
		TaskReader: reader,
		Event:      events.NewTaskWriter(taskID, 0, conf.Logger.Level, ew),
	}
	return w, nil
}
