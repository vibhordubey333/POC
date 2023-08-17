package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/io/textio"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/runners/direct"
	"log"
	"sort"

	"github.com/apache/beam/sdks/v2/go/pkg/beam"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/register"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/testing/passert"
	// Imports to enable correct filesystem access and runner setup in LOOPBACK mode
	_ "github.com/apache/beam/sdks/v2/go/pkg/beam/io/filesystem/gcs"
	_ "github.com/apache/beam/sdks/v2/go/pkg/beam/io/filesystem/local"
	_ "github.com/apache/beam/sdks/v2/go/pkg/beam/runners/universal"
)

var (
	output = flag.String("output", "./output/groupby-output.txt", "Output filename.")
)

// formatFn is a DoFn that formats a word and its count as a string.
func formatFn(w string, c []int) string {
	sort.Ints(c)
	return fmt.Sprintf("%v:%v", w, c)
}

// KV used to represent KV PCollection values
type KV struct {
	X string
	Y int64
}

func getKV(kv KV, emit func(string, int64)) {
	emit(kv.X, kv.Y)
}

func collectValues(key string, iter func(*int64) bool) (string, []int) {
	var count int64
	var values []int
	for iter(&count) {
		values = append(values, int(count))
	}
	return key, values
}

func init() {
	register.Function2x1(formatFn)
	register.Function2x0(getKV)
	register.Function2x2(collectValues)

	register.Emitter2[string, int64]()
	register.Iter1[int64]()
}

func main() {
	flag.Parse()
	beam.Init()

	if *output == "" {
		log.Fatal("No output path provided.")
	}

	p := beam.NewPipeline()
	s := p.Root()

	// Using the cross-language transform
	kvs := beam.Create(s, KV{X: "0", Y: 1}, KV{X: "0", Y: 2}, KV{X: "1", Y: 3})
	in := beam.ParDo(s, getKV, kvs)
	out := beam.GroupByKey(s, in)

	vals := beam.ParDo(s, collectValues, out)
	formatted := beam.ParDo(s, formatFn, vals)
	passert.Equals(s, formatted, "0:[1 2]", "1:[3]")

	/*if err := beamx.Run(context.Background(), p); err != nil {
		log.Fatalf("Failed to execute job: %v", err)
	}*/
	textio.Write(s, *output, formatted)
	direct.Execute(context.Background(), p)
}
