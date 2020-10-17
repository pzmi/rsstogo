# Filters

Filters that process input from feed, enriching them or putting into
canonical form. Filters can be combined into filter chains.
Consider different type of filters for different processing stages.

Example:

```yaml
feeds:
  - address: http://blog.acolyer.org/feed/
    filters:
      - link-content # fetches link and puts it into content (or something like that)
      - readability # passes content through readability, so that it's not a full website but only the article
```
Note: actually readability already can resolve link, so it could work without link-content

Think about providing configuration for filters as well.

# Presets

Presets are predefined filter chains.

Example:

```yaml
preset:
    name: reddit
    filters:
      - link-to-content
      - json-fields:
          json-path: .data.comments
        
```

# Stages (components)

Think about extracting separate components like:

- feed combining
- ebook creation
- distribution
  - local file
  - api
  - kindle
  - s3
