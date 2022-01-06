// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Original source: github.com/micro/go-micro/v3/internal/ctx/ctx.go

package ctx

import (
	"context"
	"net/http"
	"net/textproto"
	"strings"

	"github.com/2637309949/micro/v3/service/context/metadata"
	"github.com/2637309949/micro/v3/service/debug/trace"
	"github.com/google/uuid"
)

func FromContext(ctx context.Context, patchMd ...metadata.Metadata) context.Context {
	if _, _, ok := trace.FromContext(ctx); !ok {
		ctx = trace.ToContext(ctx, uuid.New().String(), uuid.New().String())
	}

	for i := range patchMd {
		ctx = metadata.MergeContext(ctx, patchMd[i], true)
	}
	return ctx
}

func FromRequest(r *http.Request) context.Context {
	md, ok := metadata.FromContext(r.Context())
	if !ok {
		md = make(metadata.Metadata)
	}

	for k, v := range r.Header {
		md[textproto.CanonicalMIMEHeaderKey(k)] = strings.Join(v, ",")
	}

	md["Host"] = r.Host
	md["Method"] = r.Method
	if r.URL != nil {
		md["URL"] = r.URL.String()
	}

	ctx := FromContext(r.Context(), md)
	traceID, parentSpanID, _ := trace.FromContext(ctx)

	if v := r.Header.Get(trace.TraceIDKey); len(v) <= 0 {
		r.Header.Set(trace.TraceIDKey, traceID)
	}

	if v := r.Header.Get(trace.SpanIDKey); len(v) <= 0 {
		r.Header.Set(trace.SpanIDKey, parentSpanID)
	}

	return ctx
}
