module github.com/open-telemetry/opentelemetry-collector-contrib/internal/aws/proxy

go 1.18

require (
	github.com/aws/aws-sdk-go v1.44.93
	github.com/open-telemetry/opentelemetry-collector-contrib/internal/common v0.59.0
	github.com/stretchr/testify v1.8.1
	go.opentelemetry.io/collector v0.59.1-0.20220907232555-bef9373cfc20
	go.uber.org/zap v1.26.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/kr/pretty v0.3.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	go.uber.org/multierr v1.10.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/open-telemetry/opentelemetry-collector-contrib/internal/common => ../../../internal/common
