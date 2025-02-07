# Understanding Elasticsearch

Elasticsearch is a distributed, RESTful search and analytics engine capable of solving a growing number of use cases. As the heart of the Elastic Stack, it centrally stores your data so you can discover the expected and uncover the unexpected.

## 11. Hands-on Projects

### Building a Search Engine
Building a search engine involves setting up an Elasticsearch cluster, indexing documents, and querying the index.

#### Example Code in Go
```go
// The code is generated hence it might not be proper
package main

import (
    "context"
    "fmt"
    "github.com/olivere/elastic/v7"
)

func main() {
    client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"))
    if err != nil {
        panic(err)
    }

    ctx := context.Background()
    _, err = client.Index().
        Index("example").
        BodyJson(map[string]string{"title": "Hello World"}).
        Do(ctx)
    if err != nil {
        panic(err)
    }

    fmt.Println("Document indexed")
}
```

#### Example Code in Java
```java
// The code is generated hence it might not be proper
import org.elasticsearch.action.index.IndexRequest;
import org.elasticsearch.action.index.IndexResponse;
import org.elasticsearch.client.RequestOptions;
import org.elasticsearch.client.RestHighLevelClient;
import org.elasticsearch.common.xcontent.XContentType;

import java.io.IOException;

public class ElasticsearchExample {
    public static void main(String[] args) throws IOException {
        RestHighLevelClient client = new RestHighLevelClient(
                RestClient.builder(new HttpHost("localhost", 9200, "http")));

        IndexRequest request = new IndexRequest("example");
        request.source("{\"title\":\"Hello World\"}", XContentType.JSON);

        IndexResponse response = client.index(request, RequestOptions.DEFAULT);
        System.out.println("Document indexed: " + response.getId());

        client.close();
    }
}
```

### Implementing Autocomplete
Autocomplete functionality can be implemented using Elasticsearch's suggesters.

#### Example Code in Go
```go
// The code is generated hence it might not be proper
package main

import (
    "context"
    "fmt"
    "github.com/olivere/elastic/v7"
)

func main() {
    client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"))
    if err != nil {
        panic(err)
    }

    ctx := context.Background()
    suggest := elastic.NewCompletionSuggester("suggest").Field("suggest").Text("hel")
    searchResult, err := client.Search().
        Index("example").
        Suggester(suggest).
        Do(ctx)
    if err != nil {
        panic(err)
    }

    fmt.Println("Suggestions:", searchResult.Suggest)
}
```

#### Example Code in Java
```java
// The code is generated hence it might not be proper
import org.elasticsearch.action.search.SearchRequest;
import org.elasticsearch.action.search.SearchResponse;
import org.elasticsearch.client.RequestOptions;
import org.elasticsearch.client.RestHighLevelClient;
import org.elasticsearch.index.query.QueryBuilders;
import org.elasticsearch.search.builder.SearchSourceBuilder;
import org.elasticsearch.search.suggest.SuggestBuilder;
import org.elasticsearch.search.suggest.SuggestionBuilder;
import org.elasticsearch.search.suggest.completion.CompletionSuggestionBuilder;

import java.io.IOException;

public class ElasticsearchAutocomplete {
    public static void main(String[] args) throws IOException {
        RestHighLevelClient client = new RestHighLevelClient(
                RestClient.builder(new HttpHost("localhost", 9200, "http")));

        SuggestionBuilder<?> completionSuggestion = new CompletionSuggestionBuilder("suggest").text("hel");
        SuggestBuilder suggestBuilder = new SuggestBuilder();
        suggestBuilder.addSuggestion("suggest", completionSuggestion);

        SearchSourceBuilder searchSourceBuilder = new SearchSourceBuilder();
        searchSourceBuilder.suggest(suggestBuilder);

        SearchRequest searchRequest = new SearchRequest("example");
        searchRequest.source(searchSourceBuilder);

        SearchResponse searchResponse = client.search(searchRequest, RequestOptions.DEFAULT);
        System.out.println("Suggestions: " + searchResponse.getSuggest());

        client.close();
    }
}
```

### Analyzing Log Data
Elasticsearch can be used to analyze log data by indexing logs and running aggregations.

#### Example Code in Go
```go
// The code is generated hence it might not be proper
package main

import (
    "context"
    "fmt"
    "github.com/olivere/elastic/v7"
)

func main() {
    client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"))
    if err != nil {
        panic(err)
    }

    ctx := context.Background()
    agg := elastic.NewTermsAggregation().Field("status.keyword")
    searchResult, err := client.Search().
        Index("logs").
        Aggregation("status_counts", agg).
        Do(ctx)
    if err != nil {
        panic(err)
    }

    fmt.Println("Aggregation result:", searchResult.Aggregations)
}
```

#### Example Code in Java
```java
// The code is generated hence it might not be proper
import org.elasticsearch.action.search.SearchRequest;
import org.elasticsearch.action.search.SearchResponse;
import org.elasticsearch.client.RequestOptions;
import org.elasticsearch.client.RestHighLevelClient;
import org.elasticsearch.index.query.QueryBuilders;
import org.elasticsearch.search.aggregations.AggregationBuilders;
import org.elasticsearch.search.aggregations.bucket.terms.TermsAggregationBuilder;
import org.elasticsearch.search.builder.SearchSourceBuilder;

import java.io.IOException;

public class ElasticsearchLogAnalysis {
    public static void main(String[] args) throws IOException {
        RestHighLevelClient client = new RestHighLevelClient(
                RestClient.builder(new HttpHost("localhost", 9200, "http")));

        TermsAggregationBuilder aggregation = AggregationBuilders.terms("status_counts").field("status.keyword");

        SearchSourceBuilder searchSourceBuilder = new SearchSourceBuilder();
        searchSourceBuilder.aggregation(aggregation);

        SearchRequest searchRequest = new SearchRequest("logs");
        searchRequest.source(searchSourceBuilder);

        SearchResponse searchResponse = client.search(searchRequest, RequestOptions.DEFAULT);
        System.out.println("Aggregation result: " + searchResponse.getAggregations());

        client.close();
    }
}
```