package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/apache/beam/sdks/v2/go/pkg/beam"
	_ "github.com/apache/beam/sdks/v2/go/pkg/beam/io/filesystem/local"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/io/textio"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/register"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/runners/direct"
	"sort"
)

/*
Reference:https://beam.apache.org/documentation/programming-guide/#cogroupbykey
*/
var (
	output = flag.String("output", "output.txt", "Output file (required).")
)

type stringPair struct {
	K, V string
}

func splitStringPair(e stringPair) (string, string) {
	return e.K, e.V
}

func init() {
	// Register DoFn.
	register.Function3x1(formatCoGBKResults)
	register.Function1x2(splitStringPair)
}

// CreateAndSplit is a helper function that creates
func CreateAndSplit(s beam.Scope, input []stringPair) beam.PCollection {
	initial := beam.CreateList(s, input)
	return beam.ParDo(s, splitStringPair, initial)
}

var emailSlice = []stringPair{
	{"amy", "amy@example.com"},
	{"carl", "carl@example.com"},
	{"julia", "julia@example.com"},
	{"carl", "carl@email.com"},
}

var phoneSlice = []stringPair{
	{"amy", "111-222-3333"},
	{"james", "222-333-4444"},
	{"amy", "333-444-5555"},
	{"carl", "444-555-6666"},
}

func main() {
	p, s := beam.NewPipelineWithRoot()

	emails := CreateAndSplit(s.Scope("CreateEmails"), emailSlice)
	phones := CreateAndSplit(s.Scope("CreatePhones"), phoneSlice)
	results := beam.CoGroupByKey(s, emails, phones)

	contactLines := beam.ParDo(s, formatCoGBKResults, results)
	textio.Write(s, *output, contactLines)
	direct.Execute(context.Background(), p)
}

func formatCoGBKResults(key string, emailIter, phoneIter func(*string) bool) string {
	var s string
	var emails, phones []string
	for emailIter(&s) {
		emails = append(emails, s)
	}
	for phoneIter(&s) {
		phones = append(phones, s)
	}
	// Values have no guaranteed order, sort for deterministic output.
	sort.Strings(emails)
	sort.Strings(phones)
	return fmt.Sprintf("%s; %s; %s", key, formatFn(emails), formatFn(phones))
}

func formatFn(w []string) string {
	return fmt.Sprintf("%s", w)
}

/*
func main() {
	beam.Init()
	p := beam.NewPipeline()
	s := p.Root()
	// Read input data from a text file on GCP (replace with your GCS path).
	lines := textio.Read(s, "gs://data-flow-bucket-vd/CoGroupBy-AKA-FlatMap-input.txt")
	// Split each line into words.
	words := beam.ParDo(s, func(line string, emit func(string)) {
		for _, word := range strings.Fields(line) {
			emit(word)
		}
	}, lines)
	// Define the words you want to count.
	selectedWords := beam.CreateList(s, []string{"Peter", "piper"})
	// Filter the words to include only the selected ones.
	selectedWordCounts := beam.ParDo(s, func(word string, emit func(string, int)) {
		emit(word, 1)
	}, selectedWords)
	// Count the occurrences of each selected word.
	wordCounts := stats.CountPerElement(s, words)
	// Combine the selected word counts.
	combined := beam.CombinePerKey(s, sumInts, beam.CoGroupByKey(s, selectedWordCounts, wordCounts))
	// Format and write the combined results to GCS (replace with your GCS path).
	formattedResults := beam.ParDo(s, func(word string, count int) string {
		return fmt.Sprintf("%s: %d", word, count)
	}, combined)
	textio.Write(s, "gs://data-flow-bucket-vd/CoGroupBy-AKA-FlatMap-output.txt", formattedResults)
	if err := beamx.Run(context.Background(), p); err != nil {
		log.Fatalf("Pipeline failed: %v", err)
	}
}
func sumInts(a, b int) int {
	return a + b
}
*/
