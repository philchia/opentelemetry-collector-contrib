# VolcEngine TLS Exporter

| Status                   |                    |
| ------------------------ |--------------------|
| Stability                | [alpha]            |
| Supported pipeline types | logs               |
| Distributions            | [contrib]          |

This exporter supports sending Log data to [TLS](https://www.volcengine.com/product/tls).

# Configuration options:

- `endpoint` (required): tls [Endpoint](https://www.volcengine.com/docs/6470/73641).
- `region` (required): tls region name.
- `topic_id` (required): tls topic id.
- `access_key_id` (optional): volcengine access key id.
- `access_key_secret` (optional): volcengine access key secret.
- `hash_key` (optional): key used to determine which shard to push

# Example:

```yaml
receivers:
  filelog:
    include: [ /var/log/containers/*.log ]
    operators:
      - type: json_parser
        timestamp:
          parse_from: attributes.time
          layout: '%Y-%m-%d %H:%M:%S'
exporters:
  volcenginetls:
    endpoint: "tls-cn-guilin.volces.com"
    region: "cn-guilin"
    topic_id: "otel-collector"
    access_key_id: "access-key-id"
    access_key_secret: "access-key-secret"
    hash_key: ""
    compress_type: ""

service:
  pipelines:
    logs:
      receivers: [filelog]
      exporters: [volcenginetls]
```
