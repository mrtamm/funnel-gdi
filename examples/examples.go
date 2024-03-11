// Package examples bundles example tasks into the Funnel CLI.
package examples

import (
	"path"
	"strings"

	intern "github.com/ohsu-comp-bio/funnel/examples/internal"
)

var examples = buildExamples()

func buildExamples() map[string][]byte {
	examples := map[string][]byte{}
	for _, n := range intern.AssetNames() {
		sn := path.Base(n)
		sn = strings.TrimSuffix(sn, path.Ext(sn))
		b := intern.MustAsset(n)
		examples[sn] = b
	}
	return examples
}

// Examples returns a set of example tasks.
func Examples() map[string][]byte {
	return examples
}
