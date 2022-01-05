package web

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/2637309949/micro/v3/service/api/handler"
	"github.com/2637309949/micro/v3/service/api/resolver"
	"github.com/2637309949/micro/v3/service/api/resolver/subdomain"
	cors "github.com/2637309949/micro/v3/service/api/server/http"
	"github.com/2637309949/micro/v3/service/client"
	"github.com/2637309949/micro/v3/service/errors"
	"github.com/2637309949/micro/v3/util/ctx"
	"github.com/2637309949/micro/v3/util/encoding"
	uhttp "github.com/2637309949/micro/v3/util/http"
)

type rpcRequest struct {
	Service  string
	Endpoint string
	Method   string
	Address  string
	Request  interface{}
}

type rpcHandler struct {
	resolver resolver.Resolver
	client   client.Client
}

func (h *rpcHandler) String() string {
	return "internal/rpc"
}

// ServeHTTP passes on a JSON or form encoded RPC request to a service.
func (h *rpcHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		cors.SetHeaders(w, r)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	defer r.Body.Close()

	var service, endpoint, address string
	var request interface{}

	// response content type
	w.Header().Set("Content-Type", "application/json")

	ct := r.Header.Get("Content-Type")

	// Strip charset from Content-Type (like `application/json; charset=UTF-8`)
	if idx := strings.IndexRune(ct, ';'); idx >= 0 {
		ct = ct[:idx]
	}

	switch ct {
	case "application/json":
		var rpcReq rpcRequest

		d := json.NewDecoder(r.Body)
		d.UseNumber()

		if err := d.Decode(&rpcReq); err != nil {
			uhttp.WriteError(w, r, err)
			return
		}

		service = rpcReq.Service
		endpoint = rpcReq.Endpoint
		address = rpcReq.Address
		request = rpcReq.Request
		if len(endpoint) == 0 {
			endpoint = rpcReq.Method
		}

		// JSON as string
		if req, ok := rpcReq.Request.(string); ok {
			d := json.NewDecoder(strings.NewReader(req))
			d.UseNumber()

			if err := d.Decode(&request); err != nil {
				uhttp.WriteError(w, r, err)
				return
			}
		}
	default:
		r.ParseForm()
		service = r.Form.Get("service")
		endpoint = r.Form.Get("endpoint")
		address = r.Form.Get("address")
		if len(endpoint) == 0 {
			endpoint = r.Form.Get("method")
		}

		d := json.NewDecoder(strings.NewReader(r.Form.Get("request")))
		d.UseNumber()

		if err := d.Decode(&request); err != nil {
			uhttp.WriteError(w, r, err)
			return
		}
	}

	if len(service) == 0 {
		uhttp.WriteError(w, r, errors.InternalServerError("invalid service"))
		return
	}

	if len(endpoint) == 0 {
		uhttp.WriteError(w, r, errors.InternalServerError("invalid endpoint"))
		return
	}

	// create request/response
	var response json.RawMessage
	req := client.NewRequest(service, endpoint, request, client.WithContentType("application/json"))

	var opts []client.CallOption

	timeout, _ := strconv.Atoi(r.Header.Get("Timeout"))
	// set timeout
	if timeout > 0 {
		opts = append(opts, client.WithRequestTimeout(time.Duration(timeout)*time.Second))
	}

	// remote call
	if len(address) > 0 {
		opts = append(opts, client.WithAddress(address))
	}

	// since services can be running in many domains, we'll use the resolver to determine the domain
	// which should be used on the call
	if resolver, ok := h.resolver.(*subdomain.Resolver); ok {
		if dom := resolver.Domain(r); len(dom) > 0 {
			opts = append(opts, client.WithNetwork(dom))
		}
	}

	// create context
	ctx := ctx.FromRequest(r)

	// remote call
	err := h.client.Call(ctx, req, &response, opts...)
	if err != nil {
		uhttp.WriteError(w, r, err)
		return
	}

	rsp, err := response.MarshalJSON()
	if err != nil {
		uhttp.WriteError(w, r, err)
		return
	}
	rsp = encoding.JSONMarshal(ctx, rsp)
	w.Header().Set("Content-Length", strconv.Itoa(len(rsp)))
	w.Write(rsp)
}

// NewRPCHandler returns an initialized RPC handler
func NewRPCHandler(r resolver.Resolver, c client.Client) handler.Handler {
	return &rpcHandler{r, c}
}
