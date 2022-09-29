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
	"encoding/json"
	"strconv"
	"time"

	"github.com/volcengine/volc-sdk-golang/service/tls/pb"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/plog"
)

func resourceToLogContents(resource pcommon.Resource) []*pb.LogContent {
	logContents := make([]*pb.LogContent, 1)
	attrs := resource.Attributes()

	fields := map[string]string{}
	attrs.Range(func(k string, v pcommon.Value) bool {
		fields[k] = v.AsString()
		return true
	})
	attributeBuffer, _ := json.Marshal(fields)
	logContents[0] = &pb.LogContent{
		Key:   "resource",
		Value: string(attributeBuffer),
	}

	return logContents
}

func instrumentationScopeToLogContents(instrumentationScope pcommon.InstrumentationScope) []*pb.LogContent {
	logContents := make([]*pb.LogContent, 1)

	var fields = map[string]interface{}{}
	fields["name"] = instrumentationScope.Name()
	fields["version"] = instrumentationScope.Version()
	var attributes = map[string]string{}
	instrumentationScope.Attributes().Range(func(k string, v pcommon.Value) bool {
		attributes[k] = v.AsString()
		return true
	})
	fields["attributes"] = attributes
	data, _ := json.Marshal(fields)
	logContents[0] = &pb.LogContent{
		Key:   "instrumentation",
		Value: string(data),
	}
	return logContents
}

func mapLogRecordToLogService(lr plog.LogRecord,
	resourceContents,
	instrumentationLibraryContents []*pb.LogContent) *pb.Log {
	if lr.Body().Type() == pcommon.ValueTypeEmpty {
		return nil
	}
	var tlsLog pb.Log
	preAllocCount := 7
	tlsLog.Contents = make([]*pb.LogContent, 0, preAllocCount+len(resourceContents)+len(instrumentationLibraryContents))

	tlsLog.Contents = append(tlsLog.Contents, resourceContents...)
	tlsLog.Contents = append(tlsLog.Contents, instrumentationLibraryContents...)

	tlsLog.Contents = append(tlsLog.Contents, &pb.LogContent{
		Key:   "severity_number",
		Value: strconv.FormatInt(int64(lr.SeverityNumber()), 10),
	})

	tlsLog.Contents = append(tlsLog.Contents, &pb.LogContent{
		Key:   "severity_text",
		Value: lr.SeverityText(),
	})

	fields := map[string]string{}
	lr.Attributes().Range(func(k string, v pcommon.Value) bool {
		fields[k] = v.AsString()
		return true
	})
	attributeBuffer, _ := json.Marshal(fields)
	tlsLog.Contents = append(tlsLog.Contents, &pb.LogContent{
		Key:   "attribute",
		Value: string(attributeBuffer),
	})

	tlsLog.Contents = append(tlsLog.Contents, &pb.LogContent{
		Key:   tlsLogContentKey,
		Value: lr.Body().AsString(),
	})

	tlsLog.Contents = append(tlsLog.Contents, &pb.LogContent{
		Key:   "flags",
		Value: strconv.FormatUint(uint64(lr.FlagsStruct()), 16),
	})

	tlsLog.Contents = append(tlsLog.Contents, &pb.LogContent{
		Key:   "trace_id",
		Value: lr.TraceID().HexString(),
	})

	tlsLog.Contents = append(tlsLog.Contents, &pb.LogContent{
		Key:   "span_id",
		Value: lr.SpanID().HexString(),
	})

	if lr.Timestamp() > 0 {
		tlsLog.Time = lr.Timestamp().AsTime().UnixMilli()
	} else if lr.ObservedTimestamp() > 0 {
		tlsLog.Time = lr.ObservedTimestamp().AsTime().UnixMilli()
	} else {
		tlsLog.Time = time.Now().UnixMilli()
	}

	return &tlsLog
}
