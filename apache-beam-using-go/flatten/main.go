package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/apache/beam/sdks/v2/go/pkg/beam"
	_ "github.com/apache/beam/sdks/v2/go/pkg/beam/io/filesystem/local"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/io/textio"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/runners/direct"
)

var (
	output = flag.String("output", "./flatten/flatten-output.txt", "Output filename.")
)

func main() {
	flag.Parse()
	beam.Init()
	p, s := beam.NewPipelineWithRoot()

	even := []int64{2, 4, 6, 8}
	odd := []int64{1, 3, 5, 7}
	//TODO: Need to figure out how to append different datatypes like integer and string
	//names := []string{"Alpha", "Tango", "Charlie"}
	//name := []string{"John", "Alex", "Murphy"}

	evenPC := beam.CreateList(s, even)
	oddPC := beam.CreateList(s, odd)
	//namePC := beam.CreateList(s, names)

	//result := beam.Flatten(s, evenPC, oddPC, namePC)
	result := beam.Flatten(s, evenPC, oddPC)
	formatted := beam.ParDo(s, formatFn, result)
	fmt.Println("Formatted:", formatted)
	textio.Write(s, *output, formatted)
	direct.Execute(context.Background(), p)
}

// formatFn is a DoFn that formats a word and its count as a string.
func formatFn(c any) string {
	return fmt.Sprintf("%v", c)
}
