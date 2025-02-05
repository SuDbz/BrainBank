

## Table of Contents

- [Listing All Indices in OpenSearch](#listing-all-indices-in-opensearch)
- [Retrieving All Data from an Index](#retrieving-all-data-from-an-index)
- [Paginating Search Results](#paginating-search-results)
- [Retrieving a Specific Document by Name](#retrieving-a-specific-document-by-name)
- [Retrieving Documents with Multiple Conditions](#retrieving-documents-with-multiple-conditions)
- [Retrieving Documents with Not Equal Conditions](#retrieving-documents-with-not-equal-conditions)
- [Extracting a Specific List of Fields from Documents](#extracting-a-specific-list-of-fields-from-documents)
- [Excluding Specific Fields from Documents](#excluding-specific-fields-from-documents)
- [Combining Multiple Conditions and Aggregations](#combining-multiple-conditions-and-aggregations)
- [Retrieving a Specific Record and Its Details](#retrieving-a-specific-record-and-its-details)


# Listing All Indices in OpenSearch

To list all the indices in OpenSearch, you can use the `_cat/indices` API. This API provides a simple way to get information about all the indices in your OpenSearch cluster.

## Steps

1. **Open your terminal or command prompt.**
2. **Use the following `curl` command to list all indices:**

    ```sh
    curl -X GET "localhost:9200/_cat/indices?v"
    ```


## Retrieving All Data from an Index

To retrieve all the data from a specific index in OpenSearch, you can use the `_search` API. This API allows you to query and retrieve documents from an index.

### Steps

1. **Open your terminal or command prompt.**
 2. **Use the following `curl` command to retrieve all data from an index:**

    ```sh
        curl -X GET "localhost:9200/your_index/_search?pretty" -H 'Content-Type: application/json' -d'
        {
          "query": {
            "match_all": {}
          }
        }
        '
    ```



## Paginating Search Results

To paginate search results in OpenSearch, you can use the `from` and `size` parameters in the `_search` API. This allows you to control the number of documents returned and the starting point of the search results.

### Steps

1. **Open your terminal or command prompt.**
2. **Use the following `curl` command to paginate search results:**

    ```sh
    curl -X GET "localhost:9200/your_index/_search?pretty" -H 'Content-Type: application/json' -d'
    {
      "from": 0,
      "size": 10,
      "query": {
        "match_all": {}
      }
    }
    '
    ```


## Retrieving a Specific Document by Name

To retrieve a specific document by its name in OpenSearch, you can use the `_search` API with a `match` query. This allows you to search for documents that match a specific field value.

### Steps

1. **Open your terminal or command prompt.**
2. **Use the following `curl` command to retrieve a document with the name `testName`:**

    ```sh
    curl -X GET "localhost:9200/your_index/_search?pretty" -H 'Content-Type: application/json' -d'
    {
      "query": {
        "match": {
          "name": "testName"
        }
      }
    }
    '
    ```


## Retrieving Documents with Multiple Conditions

To retrieve documents that match multiple conditions, such as a specific name and content, you can use the `_search` API with a `bool` query that combines multiple `must` conditions.

### Steps

1. **Open your terminal or command prompt.**
2. **Use the following `curl` command to retrieve documents with the name `testXyz` and content `abcd`:**

    ```sh
    curl -X GET "localhost:9200/your_index/_search?pretty" -H 'Content-Type: application/json' -d'
    {
      "query": {
        "bool": {
          "must": [
            { "match": { "name": "testXyz" }},
            { "match": { "content": "abcd" }}
          ]
        }
      }
    }
    '
    ```


## Retrieving Documents with Not Equal Conditions

To retrieve documents that do not match specific conditions, such as a name not equal to `testXyz` and content equal to `abcd`, you can use the `_search` API with a `bool` query that combines `must_not` and `must` conditions.

### Steps

1. **Open your terminal or command prompt.**
2. **Use the following `curl` command to retrieve documents where the name is not `testXyz` and the content is `abcd`:**

    ```sh
    curl -X GET "localhost:9200/your_index/_search?pretty" -H 'Content-Type: application/json' -d'
    {
      "query": {
        "bool": {
          "must_not": [
            { "match": { "name": "testXyz" }}
          ],
          "must": [
            { "match": { "content": "abcd" }}
          ]
        }
      }
    }
    '
    ```


## Extracting a Specific List of Fields from Documents

To extract a specific list of fields from documents in OpenSearch, you can use the `_search` API with the `_source` parameter. This allows you to specify which fields you want to include in the search results.

### Steps

1. **Open your terminal or command prompt.**
2. **Use the following `curl` command to extract specific fields from documents:**

    ```sh
    curl -X GET "localhost:9200/your_index/_search?pretty" -H 'Content-Type: application/json' -d'
    {
      "_source": ["field1", "field2"],
      "query": {
        "match_all": {}
      }
    }
    '
    ```


## Excluding Specific Fields from Documents

To exclude specific fields from documents in OpenSearch, you can use the `_search` API with the `_source` parameter and the `excludes` option. This allows you to specify which fields you want to exclude from the search results.

### Steps

1. **Open your terminal or command prompt.**
2. **Use the following `curl` command to exclude specific fields from documents:**

    ```sh
        curl -X GET "localhost:9200/your_index/_search?pretty" -H 'Content-Type: application/json' -d'
        {
          "_source": {
            "excludes": ["field1", "field2"]
          },
          "query": {
            "match_all": {}
          }
        }
        '
        ```

## Combining Multiple Conditions and Aggregations

To combine multiple conditions and aggregations in OpenSearch, you can use the `_search` API with a `bool` query for the conditions and an `aggs` section for the aggregations. This allows you to filter documents based on multiple criteria and perform aggregations on the filtered results.

### Steps

1. **Open your terminal or command prompt.**
 2. **Use the following `curl` command to combine multiple conditions and aggregations:**

            ```sh
            curl -X GET "localhost:9200/your_index/_search?pretty" -H 'Content-Type: application/json' -d'
            {
              "query": {
                "bool": {
                  "must": [
                    { "match": { "status": "active" }},
                    { "range": { "created_at": { "gte": "2022-01-01" }}}
                  ]
                }
              },
              "aggs": {
                "status_count": {
                  "terms": {
                    "field": "status.keyword"
                  }
                },
                "recent_docs": {
                  "top_hits": {
                    "sort": [
                      { "created_at": { "order": "desc" }}
                    ],
                    "_source": {
                      "includes": ["title", "created_at"]
                    },
                    "size": 5
                  }
                }
              }
            }
            '
            ```

### Explanation
- **`query.bool.must`**: This section specifies the conditions that documents must match. In this example, documents must have a `status` field equal to `active` and a `created_at` field greater than or equal to `2022-01-01`.
- **`aggs`**: This section specifies the aggregations to perform on the filtered documents.
- **`status_count`**: This aggregation creates buckets of documents based on unique values of the `status.keyword` field.
- **`recent_docs`**: This aggregation retrieves the top 5 most recent documents based on the `created_at` field, including only the `title` and `created_at` fields in the results.

By combining multiple conditions and aggregations, you can filter documents based on specific criteria and perform detailed analysis on the filtered results.


## Retrieving a Specific Record and Its Details

To retrieve a specific record and its details in OpenSearch, you can use the `_doc` API. This API allows you to fetch a document by its ID.

### Steps

1. **Open your terminal or command prompt.**
2. **Use the following `curl` command to retrieve a specific record by its ID:**

    ```sh
    curl -X GET "localhost:9200/your_index/_doc/your_document_id?pretty"
    ```

### Explanation
- **`your_index`**: Replace this with the name of your index.
- **`your_document_id`**: Replace this with the ID of the document you want to retrieve.

This command will return the details of the specified document, including all its fields and values.
