# Elasticsearch Dynamic Mapping

## Table of Contents
1. [Introduction](#introduction)
2. [What is Dynamic Mapping?](#what-is-dynamic-mapping)
    - [Enabling Dynamic Mapping](#enabling-dynamic-mapping)
    - [Example: Enabling Dynamic Mapping](#example-enabling-dynamic-mapping)
3. [Dynamic Templates](#dynamic-templates)
    - [Example: Dynamic Templates](#example-dynamic-templates)
4. [Advantages of Dynamic Mapping](#advantages-of-dynamic-mapping)
5. [Disadvantages of Dynamic Mapping](#disadvantages-of-dynamic-mapping)
6. [Controlling Dynamic Mapping](#controlling-dynamic-mapping)
    - [Example: Disabling Dynamic Mapping](#example-disabling-dynamic-mapping)
    - [Example: Strict Dynamic Mapping](#example-strict-dynamic-mapping)


## Introduction

Elasticsearch is a powerful search and analytics engine that allows you to store, search, and analyze large volumes of data quickly and in near real-time. One of the key features of Elasticsearch is its ability to dynamically map fields in your documents. This feature, known as dynamic mapping, allows Elasticsearch to automatically detect and add new fields to the index without requiring you to define the schema upfront.

## What is Dynamic Mapping?

Dynamic mapping is a feature in Elasticsearch that automatically detects and adds new fields to the index as they are encountered in the documents being indexed. This allows you to index documents without having to define the schema in advance. Elasticsearch will infer the data type of each field based on the content of the documents.

### Enabling Dynamic Mapping

Dynamic mapping is enabled by default in Elasticsearch. You can control dynamic mapping at the index level or at the field level using the `dynamic` setting.

### Example: Enabling Dynamic Mapping

```json
{
  "mappings": {
    "dynamic": true,
    "properties": {
      "name": {
        "type": "text"
      },
      "age": {
        "type": "integer"
      }
    }
  }
}
```

In this example, dynamic mapping is enabled for the index. Any new fields encountered in the documents will be automatically added to the index.

## Dynamic Templates

Dynamic templates allow you to define custom mappings for dynamically added fields based on their names or data types. This gives you more control over how new fields are mapped.

### Example: Dynamic Templates

```json
{
  "mappings": {
    "dynamic_templates": [
      {
        "strings_as_keywords": {
          "match_mapping_type": "string",
          "mapping": {
            "type": "keyword"
          }
        }
      }
    ],
    "properties": {
      "name": {
        "type": "text"
      },
      "age": {
        "type": "integer"
      }
    }
  }
}
```

In this example, any new fields with a data type of `string` will be mapped as `keyword` instead of the default `text`.

## Advantages of Dynamic Mapping

1. **Flexibility**: Dynamic mapping allows you to index documents without having to define the schema upfront. This is useful when dealing with unstructured or semi-structured data.
2. **Ease of Use**: Dynamic mapping simplifies the process of indexing documents by automatically detecting and adding new fields.
3. **Rapid Development**: Dynamic mapping allows you to quickly prototype and develop applications without worrying about the schema.

## Disadvantages of Dynamic Mapping

1. **Unpredictable Mappings**: Dynamic mapping can lead to unpredictable mappings if the data types of the fields are not consistent. For example, if a field has different data types in different documents, Elasticsearch may create multiple fields with different suffixes (e.g., `field`, `field.keyword`).
2. **Index Bloat**: Dynamic mapping can lead to index bloat if there are many unique field names in the documents. Each unique field name will be added to the index, increasing the size of the index.
3. **Performance Overhead**: Dynamic mapping adds a performance overhead as Elasticsearch needs to analyze and infer the data types of new fields.

## Controlling Dynamic Mapping

You can control dynamic mapping at the index level or at the field level using the `dynamic` setting. The `dynamic` setting can be set to `true`, `false`, or `strict`.

### Example: Disabling Dynamic Mapping

```json
{
  "mappings": {
    "dynamic": false,
    "properties": {
      "name": {
        "type": "text"
      },
      "age": {
        "type": "integer"
      }
    }
  }
}
```

In this example, dynamic mapping is disabled for the index. Any new fields encountered in the documents will be ignored.

### Example: Strict Dynamic Mapping

```json
{
  "mappings": {
    "dynamic": "strict",
    "properties": {
      "name": {
        "type": "text"
      },
      "age": {
        "type": "integer"
      }
    }
  }
}
```

In this example, strict dynamic mapping is enabled for the index. Any new fields encountered in the documents will result in an error.

