package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/apache/beam/sdks/v2/go/pkg/beam"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/io/textio"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/x/beamx"
	"log"
)

var (
	outputEven = flag.String("output1", "output/output-even", "even output file")
	outputOdd  = flag.String("output2", "output/output-odd", "od output file")
)

func main() {
	flag.Parse()
	beam.Init()

	p := beam.NewPipeline()
	s := p.Root()

	input := beam.Create(s, 1, 2, 3, 4, 5)

	// Define side outputs.
	evenOutput := beam.ParDo(s, &EvenDoFn{}, input)
	oddOutput := beam.ParDo(s, &OddDoFn{}, input)

	// Continue processing or saving the data as needed.
	// For example, you can write the side outputs to different sinks.

	// Print the even numbers.
	beam.ParDo0(s, func(i int) {
		fmt.Println("Even:", i)
	}, evenOutput)

	// Print the odd numbers.
	beam.ParDo0(s, func(i int) {
		fmt.Println("Odd:", i)
	}, oddOutput)

	formattedEven := beam.ParDo(s, func(count int) string {
		return fmt.Sprintf("%v", count)
	}, evenOutput)

	formattedOdd := beam.ParDo(s, func(count int) string {
		return fmt.Sprintf("%v", count)
	}, oddOutput)
	textio.Write(s, *outputEven, formattedEven)
	textio.Write(s, *outputOdd, formattedOdd)

	if err := beamx.Run(context.Background(), p); err != nil {
		log.Fatalf("Failed to execute job: %v", err)
	}
}

type EvenDoFn struct{}

func (fn *EvenDoFn) ProcessElement(i int, emit func(int)) {
	if i%2 == 0 {
		emit(i) // Emit even numbers to the evenOutput side output.
	}
}

type OddDoFn struct{}

func (fn *OddDoFn) ProcessElement(i int, emit func(int)) {
	if i%2 != 0 {
		emit(i) // Emit odd numbers to the oddOutput side output.
	}
}
