0. Import 
        `"github.com/apache/beam/sdks/v2/go/pkg/beam"`<br>
       or
`github.com/apache/beam/sdks/v2`
1. Other libraries <br>
```
"github.com/apache/beam/sdks/v2/go/pkg/beam/io/textio"
"github.com/apache/beam/sdks/v2/go/pkg/beam/runners/direct"
```

2. Main program. <br>
   A. Create a Beam pipeline. <br>
   B. Create an initial PCollection. <br>
   C. Apply transforms. <br>
   D.Run the pipeline, using the Direct Runner. <br>
3. To run: ``
4. Common errors:

    - Inspect Your DoFn: Inside the DoFn being applied by the ParDo, review the ProcessElement method. Ensure that the method expects and processes the correct number and type of input elements.
````
    vagrant@vibhor-virtualbox:~/p-repos/POC/apache-beam-using-go/side-output (master)$ go run main.go
panic: 	inserting ParDo in scope root
creating new DoFn in scope root
binding fn main.main.func3
binding params [{Value int} {Value int}] to input [int]
too few inputs: forgot an input or to annotate options?

goroutine 1 [running]:
github.com/apache/beam/sdks/v2/go/pkg/beam.MustN(...)
/home/vagrant/p-repos/POC/apache-beam-using-go/vendor/github.com/apache/beam/sdks/v2/go/pkg/beam/util.go:104
github.com/apache/beam/sdks/v2/go/pkg/beam.ParDo({0xc00064ff20?, 0xc0000b28c0?}, {0x189af80, 0x20718f8}, {0x0?}, {0x0?, 0x443991?, 0xc0ffffffff?})
/home/vagrant/p-repos/POC/apache-beam-using-go/vendor/github.com/apache/beam/sdks/v2/go/pkg/beam/pardo.go:419 +0x91
main.main()
/home/vagrant/p-repos/POC/apache-beam-using-go/side-output/main.go:44 +0x310
exit status 2
````
