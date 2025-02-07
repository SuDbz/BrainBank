# Understanding Elasticsearch

## 5. Aggregations

### What are aggregations?
Aggregations in Elasticsearch allow you to analyze and extract statistics from your data. They enable you to perform complex queries and generate insights by grouping and summarizing your data.

### Metric aggregations
Metric aggregations are used to calculate metrics, such as the sum, average, minimum, and maximum, from your data. Examples include:
- `avg`: Calculates the average value of a numeric field.
    ```json
    {
        "aggs": {
            "average_price": {
                "avg": {
                    "field": "price"
                }
            }
        }
    }
    ```
- `sum`: Computes the sum of numeric field values.
    ```json
    {
        "aggs": {
            "total_sales": {
                "sum": {
                    "field": "sales"
                }
            }
        }
    }
    ```
- `min`: Finds the minimum value of a numeric field.
    ```json
    {
        "aggs": {
            "min_age": {
                "min": {
                    "field": "age"
                }
            }
        }
    }
    ```
- `max`: Determines the maximum value of a numeric field.
    ```json
    {
        "aggs": {
            "max_age": {
                "max": {
                    "field": "age"
                }
            }
        }
    }
    ```

### Bucket aggregations
Bucket aggregations group documents into buckets based on certain criteria. Each bucket can contain multiple documents. Examples include:
- `terms`: Groups documents by unique values of a field.
    ```json
    {
        "aggs": {
            "genres": {
                "terms": {
                    "field": "genre"
                }
            }
        }
    }
    ```
- `range`: Creates buckets for documents that fall within specified ranges.
    ```json
    {
        "aggs": {
            "price_ranges": {
                "range": {
                    "field": "price",
                    "ranges": [
                        { "to": 50 },
                        { "from": 50, "to": 100 },
                        { "from": 100 }
                    ]
                }
            }
        }
    }
    ```
- `date_histogram`: Groups documents by date intervals.
    ```json
    {
        "aggs": {
            "sales_over_time": {
                "date_histogram": {
                    "field": "sale_date",
                    "calendar_interval": "month"
                }
            }
        }
    }
    ```

### Pipeline aggregations
Pipeline aggregations take the output of other aggregations and perform additional processing. Examples include:
- `derivative`: Calculates the derivative of a specified metric.
    ```json
    {
        "aggs": {
            "sales_over_time": {
                "date_histogram": {
                    "field": "sale_date",
                    "calendar_interval": "month"
                },
                "aggs": {
                    "sales_derivative": {
                        "derivative": {
                            "buckets_path": "sales"
                        }
                    }
                }
            }
        }
    }
    ```
- `moving_avg`: Computes a moving average of a specified metric.
    ```json
    {
        "aggs": {
            "sales_over_time": {
                "date_histogram": {
                    "field": "sale_date",
                    "calendar_interval": "month"
                },
                "aggs": {
                    "sales_moving_avg": {
                        "moving_avg": {
                            "buckets_path": "sales"
                        }
                    }
                }
            }
        }
    }
    ```
- `cumulative_sum`: Calculates the cumulative sum of a specified metric.
    ```json
    {
        "aggs": {
            "sales_over_time": {
                "date_histogram": {
                    "field": "sale_date",
                    "calendar_interval": "month"
                },
                "aggs": {
                    "sales_cumulative_sum": {
                        "cumulative_sum": {
                            "buckets_path": "sales"
                        }
                    }
                }
            }
        }
    }
    ```

[Next: Advanced Search Features](advanced_search_features.md)