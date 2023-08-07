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
