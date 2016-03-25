# exp

Experiments, learning, and testing.  Also stuff and ting.

# Quake Protocol Buffers

A test of encoding SeisComPML as a [Google Protocol Buffer](https://developers.google.com/protocol-buffers/) using Golang.

## Packages

### quake

A Quake protobuf is defined in `protobuf/quake/quake.proto`.  The Go code in `quake/quake.pb.go` is generated using the [protobuf compiler](https://developers.google.com/protocol-buffers/docs/proto3#generating) with [Go support](https://github.com/golang/protobuf):

```
protoc --proto_path=protobuf/quake/ --go_out=quake protobuf/quake/quake.proto
```

### seiscompml07

Defines structs and methods for reading [SeisComPML 0.7](http://geofon.gfz-potsdam.de/schema/0.7/sc3ml_0.7.xsd) (XML).  The event format used with [SeisComP3](http://www.seiscomp3.org/).

## Tests

### File Size

There are tests to unmarshal a SeisComPML file and marshal it as a Quake protobuf, XML, and JSON files.  Not all information in the SeisComPML is present in the Quake message.  

```
go test ./quake ./seiscompml07
ok  	github.com/gclitheroe/exp/quake	0.041s
ok  	github.com/gclitheroe/exp/seiscompml07	0.050s

ls -l seiscompml07/etc quake/etc

495917 seiscompml07/etc/2015p768477.xml - the complete SeisComPML file.
113830 quake/etc/2015p768477.xml - Quake file as XML
 99615 quake/etc/2015p768477.json - Quake file as JSON
 14181 quake/etc/2015p768477.pb - Quake file as protobuf
```

### Unmarshal Performance

There are benchmark tests to unmarshal a SeisComPML and Quake protobuf, XML, and JSON files to their related Go types.  The benchmarks preform the unmarshal on `[]byte` so as to remove any i/o bias.

```
go test -bench=.  ./quake ./seiscompml07

BenchmarkUnmarshalSeiscompml-4	        50	  30269773 ns/op - SeisComPML XML
BenchmarkUnmarshalQuakeXML-4     	     200	   8545983 ns/op - Quake XML
BenchmarkUnmarshalQuakeJSON-4    	    1000	   1800593 ns/op - Quake JSON
BenchmarkUnmarshalQuakeProtobuf-4	   10000	    163473 ns/op - Quake protobuf
```  

### Integration Tests

All SeisComPML files in the dir `/work/seismcompml07-test` can be unmarshalled and basic checks made using the integration tag:

```
go test -tags=integration ./...
```
