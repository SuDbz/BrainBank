# Understanding Elasticsearch

## 6. Advanced Search Features

### Multi-index, Multi-type Search
Elasticsearch allows you to search across multiple indices and types. This can be useful when you have related data spread across different indices.

**Example:**
```json
GET /index1,index2/_search
{
    "query": {
        "match": {
            "field": "value"
        }
    }
}
```
This example demonstrates how to perform a search query across multiple indices (`index1` and `index2`). The query matches documents where the specified `field` contains the `value`.

### Highlighting Search Results
Highlighting helps to emphasize the parts of the text that match the search query. This is particularly useful in search result pages to show users why a particular document was matched.

**Example:**
```json
GET /index/_search
{
    "query": {
        "match": {
            "content": "Elasticsearch"
        }
    },
    "highlight": {
        "fields": {
            "content": {}
        }
    }
}
```
This example demonstrates how to use the highlighting feature in Elasticsearch. The search query matches documents where the `content` field contains the term "Elasticsearch". The `highlight` section specifies that the matching parts of the `content` field should be highlighted in the search results. This helps users quickly identify why a document was included in the search results by emphasizing the matching terms.

### Suggesters
Suggesters are used to provide auto-complete or spell correction functionality. They help improve the user experience by suggesting possible completions for partial queries.

**Example:**
```json
POST /index/_search
{
    "suggest": {
        "text": "Elastcsearch",
        "simple_phrase": {
            "phrase": {
                "field": "content",
                "size": 1,
                "real_word_error_likelihood": 0.95,
                "max_errors": 1,
                "gram_size": 3,
                "direct_generator": [{
                    "field": "content",
                    "suggest_mode": "always"
                }]
            }
        }
    }
}
```

In this example:
- `text`: The text to get suggestions for. Here, "Elastcsearch" is a misspelled version of "Elasticsearch".
- `simple_phrase`: The name of the suggester.
- `phrase`: The type of suggester used. In this case, a phrase suggester.
- `field`: The field to run the suggester on. Here, it is the `content` field.
- `size`: The maximum number of suggestions to return. Here, it is set to 1.
- `real_word_error_likelihood`: The likelihood of a term being a real word. This helps in filtering out suggestions that are less likely to be correct.
- `max_errors`: The maximum number of errors allowed in the suggestions. Here, it is set to 1.
- `gram_size`: The size of the n-grams used for generating suggestions. Here, it is set to 3.
- `direct_generator`: A list of generators that produce suggestions.
  - `field`: The field to generate suggestions from. Here, it is the `content` field.
  - `suggest_mode`: The mode for generating suggestions. `always` means suggestions are always generated.

### Completion Suggester
The completion suggester is used for auto-complete functionality. It is optimized for speed and can provide instant suggestions as the user types.

**Example:**
```json
POST /index/_search
{
    "suggest": {
        "song-suggest": {
            "prefix": "nir",
            "completion": {
                "field": "suggest"
            }
        }
    }
}
```

In this example:
- `song-suggest`: The name of the suggester.
- `prefix`: The prefix text to get suggestions for. Here, it is "nir".
- `completion`: The type of suggester used. In this case, a completion suggester.
- `field`: The field to run the suggester on. Here, it is the `suggest` field.

### Term Suggester
The term suggester is used for spell correction. It suggests terms that are similar to the input term.

**Example:**
```json
POST /index/_search
{
    "suggest": {
        "term-suggest": {
            "text": "nirvana",
            "term": {
                "field": "suggest"
            }
        }
    }
}
```

In this example:
- `term-suggest`: The name of the suggester.
- `text`: The text to get suggestions for. Here, it is "nirvana".
- `term`: The type of suggester used. In this case, a term suggester.
- `field`: The field to run the suggester on. Here, it is the `suggest` field.

[Next: Performance Tuning](performance_tuning.md)