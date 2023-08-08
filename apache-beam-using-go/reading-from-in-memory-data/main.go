package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/apache/beam/sdks/v2/go/pkg/beam"
	_ "github.com/apache/beam/sdks/v2/go/pkg/beam/io/filesystem/local"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/io/textio"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/runners/direct"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/transforms/stats"
	"regexp"
)

var (
	//input = flag.String("input", "data/*", "File(s) to read.")
	lines = []string{
		"To be, or not to be: that is the question: ",
		"Whether 'tis nobler in the mind to suffer ",
		"The slings and arrows of outrageous fortune, ",
		"Or to take arms against a sea of troubles, ",
	}
	output = flag.String("output", "outputs/wordcounts.txt", "Output filename.")
	wordRE = regexp.MustCompile(`[a-zA-Z]+('[a-z])?`)
)

func main() {
	flag.Parse()

	beam.Init()

	pipeline, root := beam.NewPipelineWithRoot()
	//root := pipeline.Root()

	linesPCol := beam.CreateList(root, lines)
	words := beam.ParDo(root, func(line string, emit func(string)) {
		for _, word := range wordRE.FindAllString(line, -1) {
			emit(word)
		}
	}, linesPCol)
	counted := stats.Count(root, words)

	formatted := beam.ParDo(root, func(word string, count int) string {
		return fmt.Sprintf("%s: %v", word, count)
	}, counted)
	textio.Write(root, *output, formatted)

	direct.Execute(context.Background(), pipeline)
}
