# Understanding Elasticsearch

## 7. Performance Tuning

### Indexing Performance
Indexing performance is crucial for ensuring that data is ingested into Elasticsearch efficiently. Here are some tips to improve indexing performance:
- **Bulk Indexing**: Use the `_bulk` API to index multiple documents in a single request.
- **Refresh Interval**: Increase the `refresh_interval` to reduce the frequency of refresh operations.
- **Replica Count**: Temporarily reduce the number of replicas during heavy indexing operations.

Example:
```json
POST _bulk
{ "index" : { "_index" : "my_index", "_id" : "1" } }
{ "field1" : "value1" }
{ "index" : { "_index" : "my_index", "_id" : "2" } }
{ "field1" : "value2" }
```

### Search Performance
Optimizing search performance ensures that queries are executed quickly and efficiently. Consider the following strategies:

- **Query Caching**: Enable query caching for frequently run queries. This can significantly reduce the load on your cluster by reusing the results of previous queries. Note that query caching is most effective for queries that are identical and run frequently.

- **Filter Context**: Use filter context for parts of the query that do not need scoring. Filters are faster than queries because they do not calculate relevance scores. This is particularly useful for boolean queries where certain conditions are used to include or exclude documents without affecting the score.

- **Index Sorting**: Pre-sort index data to speed up search queries. By sorting the data at index time, Elasticsearch can quickly retrieve sorted results without the need for expensive runtime sorting. This is especially beneficial for fields that are frequently used in sort operations.

Example:
```json
GET /my_index/_search
{
    "query": {
        "bool": {
            "filter": [
                { "term": { "status": "active" } }
            ]
        }
    }
}
```

In this example, the `filter` context is used to filter documents where the `status` field is `active`. This ensures that the filter is applied efficiently without affecting the scoring of the query.

### Managing Shards and Replicas
Properly managing shards and replicas is essential for both performance and reliability:

- **Shard Count**: Choose an appropriate number of primary shards based on the expected index size and query load. Too few shards can lead to large shard sizes, which can impact performance. Conversely, too many shards can lead to overhead and resource contention.

Example:
```json
PUT /my_index
{
    "settings": {
        "number_of_shards": 5
    }
}
```
In this example, the index is configured with 5 primary shards.

- **Replica Count**: Set the number of replicas to ensure high availability and fault tolerance. Replicas provide redundancy and can improve search performance by allowing queries to be distributed across multiple nodes.

Example:
```json
PUT /my_index/_settings
{
    "number_of_replicas": 2
}
```
Here, the index is configured with 2 replicas, ensuring that each primary shard has two copies.

- **Shard Allocation**: Use shard allocation awareness to distribute shards across different nodes. This helps in preventing data loss and ensures high availability in case of node failures. You can configure shard allocation based on attributes like rack, zone, or any custom attribute.

Example:
```json
PUT /_cluster/settings
{
    "persistent": {
        "cluster.routing.allocation.awareness.attributes": "rack_id"
    }
}
```
In this example, shards are allocated based on the `rack_id` attribute, ensuring that primary and replica shards are distributed across different racks.

By carefully managing shards and replicas, you can optimize the performance and reliability of your Elasticsearch cluster.

[Next: Monitoring and Maintenance](monitoring_and_maintenance.md)