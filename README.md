### Go server + Elastic search index

This is an exampleof how to connect a Golang web server to an ElasticSearch index.

## ElasticSearch:

### Create an index and a type:
`curl -XPOST http://localhost:9200/documents/document/<id>` (id is optional)

### Delete an index: 
`curl -XDELETE http://localhost:9200/documents`

#### Create a record:
`curl --header "Content-Type: application/json" \
  --request POST \
  --data '{<your-json-object>}' \
  http://localhost:9200/documents/document/`

### Query all types in index:
`curl -XGET http://localhost:9200/documents/_search`