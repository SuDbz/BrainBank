# OpenSearch Guide
This guide provides a basic overview of OpenSearch setup and usage. For more detailed information, refer to the [OpenSearch documentation](https://opensearch.org/docs/).

## Table of Contents
- [Introduction](#introduction)
- [Setup Using Docker](#setup-using-docker)
- [Accessing OpenSearch](#accessing-opensearch)
- [Configuring OpenSearch](#configuring-opensearch)
- [Setting Up with Helm](#setting-up-with-helm)
- [Enabling OpenSearch Dashboard](#enabling-opensearch-dashboard)
- [Pushing Data](#pushing-data)
- [Structuring Data](#structuring-data)
- [Creating an Index](#creating-an-index)
- [What is an Index?](#what-is-an-index)
- [Specifying Fields](#specifying-fields)
- [Querying](#querying)
- [Different Options or Parameters](#different-options-or-parameters)
- [Create an Index with Mapping](#create-an-index-with-mapping)
- [Query OpenSearch](#query-opensearch)
- [Query with Parameters](#query-with-parameters)
- [Term Query](#term-query)
- [Range Query](#range-query)
- [Shards and Replicas](#shards-and-replicas)
- [Finding Unique Elements](#finding-unique-elements)
- [Top Elements](#top-elements)
- [Basics](#features)
- [Elasticsearch in detail](#elasticsearch)

## Introduction
OpenSearch is a community-driven, open-source search and analytics suite derived from Elasticsearch. It provides a distributed, RESTful search engine capable of handling large volumes of data.

## Setup Using Docker
To set up OpenSearch using Docker, use the following command:
```sh
docker run -d --name opensearch -p 9200:9200 -p 9600:9600 -e "discovery.type=single-node" opensearchproject/opensearch:latest
```
This command pulls the latest OpenSearch image and runs it in a Docker container.

## Accessing OpenSearch
Once OpenSearch is running, you can access it via `http://localhost:9200`. Use tools like `curl` or Postman to interact with the OpenSearch API.

Example:
```sh
curl -X GET "localhost:9200"
```

## Configuring OpenSearch
Configuration can be done via the `opensearch.yml` file or through API calls. For example, to set the number of shards:
```yaml
index.number_of_shards: 3
```

## Setting Up with Helm
To set up OpenSearch using Helm, add the OpenSearch Helm repository and install the chart:
```sh
helm repo add opensearch https://opensearch-project.github.io/helm-charts/
helm install my-opensearch opensearch/opensearch
```
This command installs OpenSearch using Helm charts.

## Enabling OpenSearch Dashboard
To enable the OpenSearch Dashboard, use the following command:
```sh
helm install my-opensearch-dashboards opensearch/opensearch-dashboards
```
This command installs the OpenSearch Dashboard, which can be accessed via `http://localhost:5601`.

## Pushing Data
Data can be pushed to OpenSearch using the `_bulk` API or individual document indexing.

Example:
```sh
curl -X POST "localhost:9200/my-index/_doc/1" -H 'Content-Type: application/json' -d'
{
    "field1": "value1",
    "field2": "value2"
}'
```

## Structuring Data
Data in OpenSearch is stored in JSON format. Proper structuring of data ensures efficient querying and indexing.

Example:
```json
{
    "user": "john_doe",
    "post_date": "2023-10-01",
    "message": "Hello, OpenSearch!"
}
```

## Creating an Index
An index in OpenSearch is created using the `PUT` request.

Example:
```sh
curl -X PUT "localhost:9200/my-index"
```

## What is an Index?
An index is a collection of documents that share similar characteristics. It is the primary unit of data storage in OpenSearch.

## Specifying Fields
Fields in OpenSearch are specified in the document structure. They can be of various types like text, keyword, date, etc.

Example:
```json
{
    "properties": {
        "field1": { "type": "text" },
        "field2": { "type": "keyword" }
    }
}
```

## Querying
Querying in OpenSearch is done using the `_search` API. Queries can be simple or complex, depending on the requirements.

## Different Options or Parameters
OpenSearch queries support various options and parameters like `size`, `from`, `sort`, etc.

Example:
```sh
curl -X GET "localhost:9200/my-index/_search?size=10&from=0&sort=field1:asc"
```

## Create an Index with Mapping
Mappings define the structure of the documents in an index.

Example:
```sh
curl -X PUT "localhost:9200/my-index" -H 'Content-Type: application/json' -d'
{
    "mappings": {
        "properties": {
            "field1": { "type": "text" },
            "field2": { "type": "keyword" }
        }
    }
}'
```

## Query OpenSearch
To query OpenSearch, use the `_search` endpoint with a query body.

Example:
```sh
curl -X GET "localhost:9200/my-index/_search" -H 'Content-Type: application/json' -d'
{
    "query": {
        "match_all": {}
    }
}'
```

## Query with Parameters
Queries can include parameters to refine the search results.

Example:
```sh
curl -X GET "localhost:9200/my-index/_search" -H 'Content-Type: application/json' -d'
{
    "query": {
        "match": {
            "field1": "value1"
        }
    }
}'
```

## Term Query
A term query searches for exact matches in the specified field.

Example:
```sh
curl -X GET "localhost:9200/my-index/_search" -H 'Content-Type: application/json' -d'
{
    "query": {
        "term": {
            "field1": "value1"
        }
    }
}'
```

## Range Query
A range query searches for values within a specified range.

Example:
```sh
curl -X GET "localhost:9200/my-index/_search" -H 'Content-Type: application/json' -d'
{
    "query": {
        "range": {
            "post_date": {
                "gte": "2023-01-01",
                "lte": "2023-12-31"
            }
        }
    }
}'
```

## Shards and Replicas
Shards and replicas are fundamental concepts in OpenSearch for distributing and replicating data.

- **Shards**: Subdivisions of an index that allow for parallel processing.
- **Replicas**: Copies of shards that provide redundancy and high availability.

Example configuration:
```yaml
index.number_of_shards: 3
index.number_of_replicas: 2
```

## Finding Unique Elements
To find unique elements in a field, you can use the `terms` aggregation.

Example:
```sh
curl -X GET "localhost:9200/my-index/_search" -H 'Content-Type: application/json' -d'
{
    "aggs": {
        "unique_field1": {
            "terms": {
                "field": "field1.keyword"
            }
        }
    },
    "size": 0
}'
```

## Top Elements
To find the top elements based on a specific field, you can use the `terms` aggregation with a size parameter.

Example:
```sh
curl -X GET "localhost:9200/my-index/_search" -H 'Content-Type: application/json' -d'
{
    "aggs": {
        "top_field1": {
            "terms": {
                "field": "field1.keyword",
                "size": 5
            }
        }
    },
    "size": 0
}'
```

## Basics 
- [Configure opensearch with helm](basic/opensearch_helm.md)
- [Basic open-search queries](basic/query_basics.md)
- [Data types](basic/data_types_in_es.md)
- [Dynamic Mapping](basic/dynamic_mapping.md)

## Elasticsearch
- [Elasticsearch](elasticsearch/init.md)
- [Playbook](https://github.com/ImadSaddik/ElasticSearch_Python_Course/tree/main/notebooks)
   
