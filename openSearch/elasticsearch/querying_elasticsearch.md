# Understanding Elasticsearch

## 3. Querying Elasticsearch

Elasticsearch provides a powerful and flexible query language to search and analyze your data. Here are some basic concepts and examples to get you started.

### Basic Queries

#### Match Query
The `match` query is used to search for a specific value in one or more fields.

```json
{
    "query": {
        "match": {
            "field_name": "search_value"
        }
    }
}
```

#### Term Query
The `term` query is used to search for exact values in a field.

```json
{
    "query": {
        "term": {
            "field_name": "exact_value"
        }
    }
}
```

#### Range Query
The `range` query is used to search for values within a specified range.

```json
{
    "query": {
        "range": {
            "field_name": {
                "gte": "start_value",
                "lte": "end_value"
            }
        }
    }
}
```

### Full-Text Search
Elasticsearch excels at full-text search. The `match` query is often used for this purpose, as it analyzes the text before searching.

```json
{
    "query": {
        "match": {
            "content": "full text search example"
        }
    }
}
```

### Filtering and Sorting Results
You can filter and sort your search results using the `bool` query and the `sort` parameter.

#### Filtering
```json
{
    "query": {
        "bool": {
            "filter": {
                "term": {
                    "status": "active"
                }
            }
        }
    }
}
```

#### Sorting
```json
{
    "sort": [
        {
            "date": {
                "order": "desc"
            }
        }
    ]
}
```

[Next: Analyzers and Tokenizers](analyzers-and-tokenizers.md)