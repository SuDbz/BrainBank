# Elasticsearch Guide

## 2. Core Concepts

### Understanding JSON Documents
Elasticsearch uses JSON (JavaScript Object Notation) as the format for documents. JSON is a lightweight data interchange format that's easy for humans to read and write, and easy for machines to parse and generate.

Example of a JSON document:
```json
{
    "user": "john_doe",
    "post_date": "2023-10-01",
    "message": "Exploring Elasticsearch!"
}
```

### Indexing and Retrieving Documents
Indexing is the process of storing documents in Elasticsearch. Each document is stored in an index and assigned a unique identifier.

Example of indexing a document:
```bash
PUT /my_index/_doc/1
{
    "user": "john_doe",
    "post_date": "2023-10-01",
    "message": "Exploring Elasticsearch!"
}
```

Retrieving a document by ID:
```bash
GET /my_index/_doc/1
```

### CRUD Operations
CRUD stands for Create, Read, Update, and Delete. These are the four basic operations you can perform on documents in Elasticsearch.

- **Create**: Index a new document.
    ```bash
    POST /my_index/_doc/
    {
        "user": "jane_doe",
        "post_date": "2023-10-02",
        "message": "Learning CRUD operations!"
    }
    ```

- **Read**: Retrieve a document by ID.
    ```bash
    GET /my_index/_doc/1
    ```

- **Update**: Update an existing document.
    ```bash
    POST /my_index/_update/1
    {
        "doc": {
            "message": "Updated message content"
        }
    }
    ```

- **Delete**: Delete a document by ID.
    ```bash
    DELETE /my_index/_doc/1
    ```

This guide provides a basic understanding of core concepts in Elasticsearch. For more detailed information, refer to the official Elasticsearch documentation.

[Next: Querying Elasticsearch](querying_elasticsearch.md)