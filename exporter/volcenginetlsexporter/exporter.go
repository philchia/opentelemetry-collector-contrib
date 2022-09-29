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
	"context"

	"github.com/volcengine/volc-sdk-golang/service/tls/pb"
	"github.com/volcengine/volc-sdk-golang/service/tls/producer"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/config"
	"go.opentelemetry.io/collector/exporter/exporterhelper"
	"go.opentelemetry.io/collector/pdata/plog"
	"go.uber.org/zap"
)

// newLogsExporter return a new LogService logs exporter.
func newLogsExporter(set component.ExporterCreateSettings, cfg config.Exporter) (component.LogsExporter, error) {
	l := &tlsLogsPusher{
		logger: set.Logger,
		cfg:    cfg.(*Config),
	}

	var err error
	if l.producer, err = newTLSProducer(cfg.(*Config)); err != nil {
		return nil, err
	}

	return exporterhelper.NewLogsExporter(
		context.TODO(),
		set,
		cfg,
		l.pushLogsData,
		exporterhelper.WithStart(func(ctx context.Context, host component.Host) error {
			l.producer.Start()
			return nil
		}),
		exporterhelper.WithShutdown(func(ctx context.Context) error {
			l.producer.Close()
			return nil
		}),
	)
}

type tlsLogsPusher struct {
	cfg      *Config
	logger   *zap.Logger
	producer producer.Producer
}

func (s *tlsLogsPusher) pushLogsData(
	_ context.Context,
	md plog.Logs) error {
	groups := s.convertPlogsToLogGroupList(md)

	for _, logGroup := range groups {
		if err := s.producer.SendLogs(s.cfg.HashKey, s.cfg.TopicID, logGroup.GetSource(), logGroup.GetFileName(), logGroup, s); err != nil {
			s.logger.Error("put logGroup data", zap.Error(err))
			return err
		}
	}
	return nil
}

func (s *tlsLogsPusher) convertPlogsToLogGroupList(md plog.Logs) []*pb.LogGroup {
	md.LogRecordCount()
	rls := md.ResourceLogs()
	var groups = make([]*pb.LogGroup, 0, rls.Len())
	for i := 0; i < rls.Len(); i++ {
		rl := rls.At(i)
		sl := rl.ScopeLogs()
		resource := rl.Resource()
		group := &pb.LogGroup{
			Source:   "",
			FileName: "",
		}
		resourceContents := resourceToLogContents(resource)
		for j := 0; j < sl.Len(); j++ {
			ils := sl.At(j)
			instrumentationLibraryContents := instrumentationScopeToLogContents(ils.Scope())
			logs := ils.LogRecords()
			for j := 0; j < logs.Len(); j++ {
				slsLog := mapLogRecordToLogService(logs.At(j), resourceContents, instrumentationLibraryContents)
				if slsLog != nil {
					group.Logs = append(group.Logs, slsLog)
				}
			}
		}
		groups = append(groups, group)
	}

	return groups
}

func (s *tlsLogsPusher) Success(_ *producer.Result) {}

func (s *tlsLogsPusher) Fail(result *producer.Result) {
	s.logger.Error("send log to tls failed",
		zap.String("endpoint", s.cfg.Endpoint),
		zap.String("region", s.cfg.Region),
		zap.String("topic_id", s.cfg.TopicID),
		zap.String("hash_key", s.cfg.HashKey),
		zap.Any("result", result),
	)
}

func newTLSProducer(cfg *Config) (producer.Producer, error) {
	producerCfg := producer.GetDefaultProducerConfig()
	producerCfg.Endpoint = cfg.Endpoint
	producerCfg.Region = cfg.Region
	producerCfg.AccessKeyID = cfg.AccessKeyID
	producerCfg.AccessKeySecret = cfg.AccessKeySecret
	prod := producer.NewProducer(producerCfg)
	if cfg.SecurityToken != "" {
		prod.ResetAccessKeyToken(cfg.AccessKeyID, cfg.AccessKeySecret, cfg.SecurityToken)
	}

	return prod, nil
}
