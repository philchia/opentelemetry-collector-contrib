// Copyright 2020, OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package volcenginetlsexporter ...
package volcenginetlsexporter // import "github.com/open-telemetry/opentelemetry-collector-contrib/exporter/volcenginetlsexporter"

import (
	"errors"

	"go.opentelemetry.io/collector/config"
)

// Config defines configuration for AlibabaCloud Log Service exporter.
type Config struct {
	config.ExporterSettings `mapstructure:",squash"`
	// required, tls [Endpoint](https://www.volcengine.com/docs/6470/73641).
	Endpoint string `mapstructure:"endpoint"`
	// required, tls region, should match endpoint.
	Region string `mapstructure:"region"`
	// required, VolcanoEngine access key id
	AccessKeyID string `mapstructure:"access_key_id"`
	// required, VolcanoEngine access key secret
	AccessKeySecret string `mapstructure:"access_key_secret"`
	// optional, VolcanoEngine security token
	SecurityToken string `mapstructure:"security_token"`
	// required, tls topic id
	TopicID string `mapstructure:"topic_id"`
	// optional, key used to determine which shard to push
	HashKey string `mapstructure:"hash_key"`
}

func (c Config) Validate() error {
	if c.Endpoint == "" || c.AccessKeyID == "" || c.AccessKeySecret == "" || c.Region == "" || c.TopicID == "" {
		return errors.New("invalid config")
	}

	return nil
}
