# Data Types in Elasticsearch

## Table of Contents

1. [String Data Types](#1-string-data-types)
    - [`text`](#text)
    - [`keyword`](#keyword)
2. [Numeric Data Types](#2-numeric-data-types)
    - [`integer`](#integer)
    - [`float`](#float)
3. [Date Data Type](#3-date-data-type)
    - [`date`](#date)
4. [Boolean Data Type](#4-boolean-data-type)
    - [`boolean`](#boolean)
5. [Array Data Type](#5-array-data-type)
    - [`array`](#array)
6. [Object Data Type](#6-object-data-type)
    - [`object`](#object)
7. [Note](#note)

Elasticsearch supports a variety of data types for the fields in your documents. Understanding these data types is crucial for effective indexing and querying. Below are some of the primary data types in Elasticsearch with relevant examples.

## 1. String Data Types

### `text`
The `text` type is used for full-text search. It is analyzed, meaning it is broken down into individual terms for searching.

```json
{
    "mappings": {
        "properties": {
            "description": {
                "type": "text"
            }
        }
    }
}
```

### `keyword`
The `keyword` type is used for structured content such as IDs, email addresses, or tags. It is not analyzed.

```json
{
    "mappings": {
        "properties": {
            "status": {
                "type": "keyword"
            }
        }
    }
}
```

## 2. Numeric Data Types

### `integer`
The `integer` type is used for whole numbers.

```json
{
    "mappings": {
        "properties": {
            "age": {
                "type": "integer"
            }
        }
    }
}
```

### `float`
The `float` type is used for floating-point numbers.

```json
{
    "mappings": {
        "properties": {
            "price": {
                "type": "float"
            }
        }
    }
}
```

## 3. Date Data Type

### `date`
The `date` type is used for dates and times. It supports various formats.

```json
{
    "mappings": {
        "properties": {
            "created_at": {
                "type": "date",
                "format": "yyyy-MM-dd HH:mm:ss"
            }
        }
    }
}
```

## 4. Boolean Data Type

### `boolean`
The `boolean` type is used for true/false values.

```json
{
    "mappings": {
        "properties": {
            "is_active": {
                "type": "boolean"
            }
        }
    }
}
```

## 5. Array Data Type

### `array`
Elasticsearch does not have a dedicated array type. Instead, you can define an array by specifying multiple values for a field.

```json
{
    "mappings": {
        "properties": {
            "tags": {
                "type": "keyword"
            }
        }
    }
}
```

## 6. Object Data Type

### `object`
The `object` type is used for JSON objects.

```json
{
    "mappings": {
        "properties": {
            "address": {
                "type": "object",
                "properties": {
                    "street": {
                        "type": "text"
                    },
                    "city": {
                        "type": "keyword"
                    }
                }
            }
        }
    }
}
```


## 7. Note

In Elasticsearch, `keyword` and `text` are two different data types used for different purposes:

- **Keyword**: 
    - Used for structured data that can be filtered, sorted, and aggregated.
    - Not analyzed, meaning the exact value is indexed as is.
    - Suitable for fields like tags, IDs, email addresses, etc.

- **Text**: 
    - Used for full-text search capabilities.
    - Analyzed, meaning the text is processed (e.g., tokenized, lowercased) before indexing.
    - Suitable for fields like body of an email, product descriptions, etc.

Example mapping:
```json
{
    "mappings": {
        "properties": {
            "title": {
                "type": "text"
            },
            "category": {
                "type": "keyword"
            }
        }
    }
}
```