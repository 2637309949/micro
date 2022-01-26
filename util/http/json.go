package http

import (
	"bytes"
	"context"
	"strings"

	"github.com/2637309949/micro/v3/service/debug/trace"
)

// see https://stackoverflow.com/questions/28595664/how-to-stop-json-marshal-from-escaping-and/28596225
func Marshal(ctx context.Context, t []byte) []byte {
	traceID, _, _ := trace.FromContext(ctx)
	rsp := bytes.TrimRight(t, "\n")
	if strings.HasPrefix(string(rsp), "{") {
		if !strings.Contains(string(rsp), "code") {
			rsp = []byte(strings.Replace(string(rsp), "{", "{\"code\": 200,", 1))
		}
		if !strings.Contains(string(rsp), "request_id") {
			rsp = []byte(strings.Replace(string(rsp), "{", "{\"request_id\": \""+traceID+"\",", 1))
		}
	}
	rsp = []byte(strings.ReplaceAll(string(rsp), ",}", "}"))
	return rsp
}
