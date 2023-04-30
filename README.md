# exp

Experiments, learning, and testing.  Also stuff and ting.

### Unmarshal Performance

Benchmark tests to unmarshal a SeisComPML and Quake protobuf, XML, and JSON files to their related Go types.  
The benchmarks preform the unmarshal on `[]byte` to remove any i/o bias.

```
go test -bench=.  ./...

BenchmarkUnmarshalSeiscompml-4	        50	  30269773 ns/op - SeisComPML XML
BenchmarkUnmarshalQuakeXML-4     	     200	   8545983 ns/op - Quake XML
BenchmarkUnmarshalQuakeJSON-4    	    1000	   1800593 ns/op - Quake JSON
BenchmarkUnmarshalQuakeProtobuf-4	   10000	    163473 ns/op - Quake protobuf
```  

