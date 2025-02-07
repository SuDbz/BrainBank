# Understanding Elasticsearch

## 4. Analyzers and Tokenizers

### Understanding Analyzers
Analyzers in Elasticsearch are used to break down text into tokens or terms, which are then indexed for search. An analyzer performs three functions:
1. Character filters: Preprocess the text.
2. Tokenizer: Splits the text into tokens.
3. Token filters: Modify the tokens.

Analyzers are crucial for full-text search as they determine how text is processed and indexed. They help in normalizing the text, removing unwanted characters, and breaking it down into searchable terms.

### Built-in Analyzers
Elasticsearch provides several built-in analyzers, including:
- **Standard Analyzer**: The default analyzer which tokenizes text based on Unicode text segmentation. Use this for general-purpose text analysis.
- **Simple Analyzer**: Tokenizes text by non-letter characters and lowercases the tokens. Use this for basic text analysis where case sensitivity is not required.
- **Whitespace Analyzer**: Tokenizes text based on whitespace. Use this when you need to preserve the original text structure.
- **Stop Analyzer**: Similar to the Simple Analyzer but also removes stop words. Use this to improve search relevance by ignoring common words.
- **Keyword Analyzer**: Treats the entire input as a single token. Use this for exact match scenarios.

### Custom Analyzers and Tokenizers
You can create custom analyzers by combining different character filters, tokenizers, and token filters. Custom analyzers are useful when the built-in analyzers do not meet your specific requirements. Here is an example of a custom analyzer:


[Next: Aggregations](aggregations.md)