// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

module github.com/open-telemetry/opentelemetry-collector-contrib/receiver/solacereceiver

go 1.20

require (
	github.com/Azure/go-amqp v1.0.2
	github.com/stretchr/testify v1.9.0
	go.opencensus.io v0.24.0
	go.opentelemetry.io/collector/component v0.89.0
	go.opentelemetry.io/collector/config/configopaque v1.19.0
	go.opentelemetry.io/collector/config/configtls v1.19.0
	go.opentelemetry.io/collector/confmap v0.89.0
	go.opentelemetry.io/collector/consumer v0.89.0
	go.opentelemetry.io/collector/pdata v1.0.0-rcv0018
	go.opentelemetry.io/collector/receiver v0.89.0
	go.opentelemetry.io/otel v1.20.0
	go.uber.org/zap v1.26.0
	google.golang.org/protobuf v1.31.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/fsnotify/fsnotify v1.8.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/hashicorp/go-version v1.6.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/knadh/koanf/maps v0.1.1 // indirect
	github.com/knadh/koanf/providers/confmap v0.1.0 // indirect
	github.com/knadh/koanf/v2 v2.0.1 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/mapstructure v1.5.1-0.20220423185008-bf980b35cac4 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	go.opentelemetry.io/collector/config/configtelemetry v0.89.0 // indirect
	go.opentelemetry.io/collector/featuregate v1.0.0-rcv0018 // indirect
	go.opentelemetry.io/otel/metric v1.20.0 // indirect
	go.opentelemetry.io/otel/trace v1.20.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/net v0.18.0 // indirect
	golang.org/x/sys v0.14.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230822172742-b8732ec3820d // indirect
	google.golang.org/grpc v1.59.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

retract (
	v0.76.2
	v0.76.1
	v0.65.0
)
