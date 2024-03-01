package cmd

import (
	"bytes"
	"os"
	"strings"
	"testing"

	taskCmd "github.com/ohsu-comp-bio/funnel/cmd/task"
	"github.com/ohsu-comp-bio/funnel/tests"
)

func TestCreateStdin(t *testing.T) {
	tests.ParseConfig()
	conf := tests.DefaultConfig()
	conf.Compute = "noop"
	fun := tests.NewFunnel(conf)
	fun.StartServer()

	a, _ := os.ReadFile("hello-world.json")
	b, _ := os.ReadFile("hello-world.json")

	in := &bytes.Buffer{}
	out := &bytes.Buffer{}
	in.Write(a)
	in.Write(b)

	err := taskCmd.Create(conf.Server.HTTPAddress(), []string{"hello-world.json"}, in, out)
	if err != nil {
		t.Fatal(err)
	}

	outStr := out.String()
	ids := strings.Split(strings.TrimSpace(outStr), "\n")
	if len(ids) != 3 {
		t.Fatalf("Expected 3, got %d task ID value(s) from stdout buffer %q", len(ids), outStr)
	}
}
