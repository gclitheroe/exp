# exp

Experiments and learning.  Also stuff and ting.

## Protobufs

Generate Go libs from protobuf files:

```
protoc --proto_path=protobuf --go_out=internal --go_opt=paths=source_relative quake/quake.proto quake/key.proto
```

## sc3ml2quake

Converts earthquake information in SeismComPML format to Quake protobufs.  See also [Protobufs With Go](https://blog.geoffc.nz/protobufs-go/)

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

## Kafka

See [https://blog.geoffc.nz/kafka-ksqldb-quakes/](https://blog.geoffc.nz/kafka-ksqldb-quakes/) for more information about working with the Kafka examples.

Use the Docker Compose file in the `kafka` directory.  This is from https://ksqldb.io/quickstart-platform.html#quickstart-content

```
docker-compose up -d
...
docker-compose down
```

### quake-producer-kafka

Sends Quake protobufs to a Kafka topic using schema registry and key and quake protobuf schemas from `protobuf/quake`. 
Protobufs for two quakes are included in `cmd/quake-producer-kafka/demo-data`.  

For experimenting with more data the [quake-protobufs](https://github.com/gclitheroe/exp/releases/tag/quake-protobuf) release on this repo has a tar file
`quake-2020.tar.gz`.  It contains 304510 update files for 22355 earthquakes from New Zealand from the year 2020.  
Download and extract this file and then run:

```
quake-producer-kafka path/quake-2020
```

See sc3ml2quake for creating more data.

### quake-consumer-kafka

Reads Quake protobufs from a Kafka topic.

## Acknowledgement 

The New Zealand GeoNet programme and its sponsors EQC, GNS Science, LINZ, NEMA and MBIE are acknowledged for providing data used in this repo.
