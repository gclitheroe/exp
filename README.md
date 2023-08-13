# exp

Experiments and learning.  Also stuff and ting.

---

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

---

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

---

## NATS

* [https://docs.nats.io/](https://docs.nats.io/)
* [https://natsbyexample.com/](https://natsbyexample.com/)

`docker run -p 4222:4222 -p 8222:8222 -p 6222:6222 -ti nats:latest`

### quake-producer-nats

### quake-consumer-nats

---

## pgvector

`cmd/search-pgvector` contains scripts and data for experimenting with vector search in Postgres using the [pgvector](https://github.com/pgvector/pgvector) extension. 

See [https://blog.geoffc.nz/pgvector/](https://blog.geoffc.nz/pgvector/) for more information.

Container with test data preloaded

```shell
docker pull gclitheroe/pgvector-exp
docker run --name pgvector-test -e POSTGRES_PASSWORD=password -d -p 5432:5432 gclitheroe/pgvector-exp
```

### Setup

Or see prebuilt Docker image later on.

Download input data from https://github.com/gclitheroe/exp/releases/download/pgvector/data.tar.gz
Run the Docker container with pgvector already installed.

```shell
docker pull ankane/pgvector
docker run --name pgvector-test -e POSTGRES_PASSWORD=password -d -p 5432:5432 ankane/pgvector

# Copy the postgres CSV files into the running docker container.
docker cp data/items.csv pgvector-test:/tmp
docker cp data/search.csv pgvector-test:/tmp
docker cp data/issues.csv pgvector-test:/tmp
```

Create the extensions and tables and load the test data.

```postgresql
psql -h 127.0.0.1 -p 5432 -U postgres
CREATE EXTENSION vector;
CREATE TABLE items (id bigserial PRIMARY KEY, embedding vector(384), description text);
CREATE TABLE issues (id bigserial PRIMARY KEY, embedding vector(384), description text);
CREATE TABLE search (id bigserial PRIMARY KEY, embedding vector(384), term text);
COPY items FROM '/tmp/items.csv' (FORMAT csv);
COPY search FROM '/tmp/search.csv' (FORMAT csv);
COPY issues FROM '/tmp/issues.csv' (FORMAT csv);
```

### Queries

Using extended output format for readability `\x`.


#### Support Tickets

There are 8469 support tickets with associated embeddings loaded into the issues table.  Queries perform well without an index.  

Results are about software problems even though the exact phrases may be different.

```postgresql
SELECT description, embedding <=> (SELECT embedding FROM search WHERE term = 'software problem') AS cos FROM items ORDER BY cos ASC;

-[ RECORD 1 ]--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
description | I'm having an issue with the {product_purchased}. Please assist.                                                                                                                                                                                                                                                                                                                                             +
            |                                                                                                                                                                                                                                                                                                                                                                                                              +
            | 1) If you want new (not already installed) software, you may need to use:                                                                                                                                                                                                                                                                                                                                    +
            |                                                                                                                                                                                                                                                                                                                                                                                                              +
            | 1.) Windows 7 Professional.                                                                                                                                                                                                                                                                                                                                                                                  +
            |                                                                                                                                                                                                                                                                                                                                                                                                              +
            | 2.) This problem started occurring after the recent software update. I haven't made any other changes to the device.
cos         | 0.10983020430459844
-[ RECORD 2 ]--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
description | I'm having an issue with the {product_purchased}. Please assist.                                                                                                                                                                                                                                                                                                                                             +
            |                                                                                                                                                                                                                                                                                                                                                                                                              +
            |                                                                                                                                                                                                                                                                                                                                                                                                              +
            | I have the product purchased as a full time job. I have used the software and it has worked so far and I am satisfied! A few months ago I need assistance as soon as possible because it's affecting my work and productivity.
cos         | 0.11112324098668158
-[ RECORD 3 ]--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
description | I'm having an issue with the {product_purchased}. Please assist.                                                                                                                                                                                                                                                                                                                                             +
            |                                                                                                                                                                                                                                                                                                                                                                                                              +
            | I want to give your company a free demo program. Please help me create this program.                                                                                                                                                                                                                                                                                                                         +
            |                                                                                                                                                                                                                                                                                                                                                                                                              +
            | I want your support. Please add your name. I'm worried that the issue might be hardware-related and might require repair or replacement.
cos         | 0.11770975873161138
-[ RECORD 4 ]----
...
```

Using a LIKE query on the raw text doesn't yield any results.

```postgresql
SELECT description FROM items WHERE description LIKE '%software problem%'; 
(0 rows)
```

```postgresql
SELECT description, embedding <=> (SELECT embedding FROM search WHERE term = 'original charger') AS cos FROM items ORDER BY cos ASC;
-[ RECORD 1 ]--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
description | I'm having an issue with the {product_purchased}. Please assist.                                                                                                                                                                                                                                                                                                                                             +
            |                                                                                                                                                                                                                                                                                                                                                                                                              +
            | To do I'm using the original charger that came with my {product_purchased}, but it's not charging properly.
cos         | 0.1284701526158636
-[ RECORD 2 ]--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
description | I'm having an issue with the {product_purchased}. Please assist. I'm using the original charger that came with my {product_purchased}, but it's not charging properly.
cos         | 0.130345908872733
-[ RECORD 3 ]--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
description | I'm having an issue with the {product_purchased}. Please assist. I'm using the original charger that came with my {product_purchased}, but it's not charging properly.
cos         | 0.130345908872733
-[ RECORD 4 ]--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
description | I'm having an issue with the {product_purchased}. Please assist. I'm using the original charger that came with my {product_purchased}, but it's not charging properly.
cos         | 0.130345908872733
-[ RECORD 5 ]--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
description | I'm having an issue with the {product_purchased}. Please assist. I'm using the original charger that came with my {product_purchased}, but it's not charging properly.
cos         | 0.130345908872733
-[ RECORD 6 ]
...
```

There is no specific mention of dogs laying in the sun but the search does find some possible matches including pet in an email address.

```postgresql
SELECT description, embedding <=> (SELECT embedding FROM search WHERE term = 'dog laying in the sun') AS cos FROM items ORDER BY cos ASC;

-[ RECORD 1 ]--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
description | I'm having an issue with the {product_purchased}. Please assist.                                                                                                                                                                                                                                                                                                                                             +
            |                                                                                                                                                                                                                                                                                                                                                                                                              +
            | I want a picture of your dog. Please come and visit me soon.                                                                                                                                                                                                                                                                                                                                                 +
            |                                                                                                                                                                                                                                                                                                                                                                                                              +
            | I'll keep the pictures. Please come to me soon. I've checked for any available software updates for my {product_purchased}, but there are none.
cos         | 0.20796826662189327
-[ RECORD 2 ]--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
description | I'm having an issue with the {product_purchased}. Please assist. 1-800-859-7267 2 e-mail us at tips@pet-babe.us for questions or to try out this product if you I've tried different settings and configurations on my {product_purchased}, but the issue persists.
cos         | 0.2115174712052278
-[ RECORD 3 ]--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
description | I'm having an issue with the {product_purchased}. Please assist.                                                                                                                                                                                                                                                                                                                                             +
            |                                                                                                                                                                                                                                                                                                                                                                                                              +
            | We have two customers:                                                                                                                                                                                                                                                                                                                                                                                       +
            |                                                                                                                                                                                                                                                                                                                                                                                                              +
            | Carnivorous pet!                                                                                                                                                                                                                                                                                                                                                                                             +
            |                                                                                                                                                                                                                                                                                                                                                                                                              +
            | Grenadine (Grizzly)                                                                                                                                                                                                                                                                                                                                                                                          +
            |                                                                                                                                                                                                                                                                                                                                                                                                              +
            | Kelica ( I've tried different settings and configurations on my {product_purchased}, but the issue persists.
cos         | 0.2263596774337866
-[ RECORD 4 ]
...
```

#### Issues

There are 1,000,000 GitHub issues with associated embeddings loaded into the issues table.  Performance can be greatly
improved by adding an index to group the embeddings into lists that are probed at query time.  This increases performance
but can reduce recall.  

```postgresql    
EXPLAIN ANALYSE SELECT description, embedding <=> (SELECT embedding FROM search WHERE term = 'software problem') AS cos FROM issues ORDER BY cos ASC LIMIT 100;
                                                                  QUERY PLAN                                                                  
----------------------------------------------------------------------------------------------------------------------------------------------
 Limit  (cost=217132.44..217144.11 rows=100 width=365) (actual time=735.190..743.605 rows=100 loops=1)
   InitPlan 1 (returns $0)
     ->  Seq Scan on search  (cost=0.00..20.12 rows=4 width=32) (actual time=0.031..0.032 rows=1 loops=1)
           Filter: (term = 'software problem'::text)
           Rows Removed by Filter: 3
   ->  Gather Merge  (cost=217112.32..314241.30 rows=832476 width=365) (actual time=728.648..737.053 rows=100 loops=1)
         Workers Planned: 2
         Params Evaluated: $0
         Workers Launched: 2
         ->  Sort  (cost=216112.29..217152.89 rows=416238 width=365) (actual time=702.454..702.461 rows=77 loops=3)
               Sort Key: ((issues.embedding <=> $0))
               Sort Method: top-N heapsort  Memory: 108kB
               Worker 0:  Sort Method: top-N heapsort  Memory: 115kB
               Worker 1:  Sort Method: top-N heapsort  Memory: 108kB
               ->  Parallel Seq Scan on issues  (cost=0.00..200203.97 rows=416238 width=365) (actual time=3.581..642.596 rows=333333 loops=3)
 Planning Time: 0.220 ms
 JIT:
   Functions: 19
   Options: Inlining false, Optimization false, Expressions true, Deforming true
   Timing: Generation 4.064 ms, Inlining 0.000 ms, Optimization 1.492 ms, Emission 15.255 ms, Total 20.811 ms
 Execution Time: 746.185 ms
```

Add an index to issues to use approximate nearest neighbor search, which trades some recall for performance.  Need
to create the index after the table has some data in it.
 
```postgresql
SET maintenance_work_mem TO '512 MB';

CREATE INDEX ON issues USING ivfflat (embedding vector_cosine_ops) WITH (lists = 1000);
```

Set search probes to ~sqrt the number of lists.  If set equal the number of lists the index won't be used.

```postgresql
SET ivfflat.probes = 35;

EXPLAIN ANALYSE SELECT description, embedding <=> (SELECT embedding FROM search WHERE term = 'software problem') AS cos FROM issues ORDER BY cos ASC LIMIT 100;
QUERY PLAN                                                                       
--------------------------------------------------------------------------------------------------------------------------------------------------------
 Limit  (cost=17835.12..17848.26 rows=100 width=365) (actual time=74.305..74.560 rows=100 loops=1)
   InitPlan 1 (returns $0)
->  Seq Scan on search  (cost=0.00..20.12 rows=4 width=32) (actual time=0.026..0.029 rows=1 loops=1)
Filter: (term = 'software problem'::text)
Rows Removed by Filter: 3
->  Index Scan using issues_embedding_idx on issues  (cost=17815.00..149137.00 rows=1000000 width=365) (actual time=74.302..74.547 rows=100 loops=1)
Order By: (embedding <=> $0)
Planning Time: 0.260 ms
Execution Time: 74.620 ms
```

### Data Generation

Data from [https://www.kaggle.com/](https://www.kaggle.com/):
* 8469 tickets from [Customer Support Ticket Dataset](https://www.kaggle.com/datasets/suraj520/customer-support-ticket-dataset)
* 5332153 github issues from [GitHub Issues](https://www.kaggle.com/datasets/davidshinn/github-issues)

Install dependencies:

```python
pip install transformers datasets
pip install torch
pip install tensorflow
pip install psycopg2-binary
pip install pgvector
```

Uses [gte-small](https://huggingface.co/thenlper/gte-small) model:  
* Doing well in the [Massive Text Embedding Benchmark (MTEB) Leaderboard](https://huggingface.co/spaces/mteb/leaderboard).
* Trained on English text only.
* Input will be truncated to 512 tokens.
* Embeddings have 384 dimensions.
* Storing an embedding in a Postgres vector type takes 4 * dimensions + 8 bytes.  In this case 1544 bytes per embedding.
* Other models are available that have been trained on multi-lingual input.  They generate embeddings with more dimensions.  

Use Python scripts to generate vectors and save them to the database:
* `issues.py` loads 1,000,000 of the github issues.
* `support.py` loads the support tickets.
* `search.py` loads vectors for a few search terms for testing queries. In an application these would be created on the fly from user queries.

Optionally add the following in the scripts to normalise the embedding and then query using L2 distance (inner product) <#> instead of cosine similarity.

```python
import torch.nn.functional as F
...
embeddings = F.normalize(embeddings, p=2, dim=1)
```

### Docker build

https://cadu.dev/creating-a-docker-image-with-database-preloaded/

Using the dumpfile created with

```shell
pg_dump -h localhost -p 5432 -U postgres postgres > data/pgvector_dump.sql
```

Edit the dumpfile and add:

```postgresql
SET maintenance_work_mem TO '512 MB';
```

```shell
docker image build . -t gclitheroe/pgvector-exp:latest
```


---

## Acknowledgement 

The New Zealand GeoNet programme and its sponsors EQC, GNS Science, LINZ, NEMA and MBIE are acknowledged for providing data used in this repo.
