package http

import (
	"net/http"

	"github.com/2637309949/micro/v3/service/debug/trace"
	"github.com/2637309949/micro/v3/service/errors"
	"github.com/2637309949/micro/v3/service/logger"
)

func WriteError(w http.ResponseWriter, r *http.Request, err error) {
	// trace request
	traceID, _, _ := trace.FromContext(r.Context())

	// response content type
	w.Header().Set("Content-Type", "application/json")

	// parse out the error code
	ce := errors.Parse(err.Error(), traceID)
	if cd2 := errors.Parse(ce.Detail, traceID); cd2.Code != 0 {
		ce = cd2
	}

	switch ce.Code {
	case 0:
		ce.Code = 500
		ce.Status = http.StatusText(500)
		ce.Detail = "error during request: " + ce.Detail
		w.WriteHeader(http.StatusOK)
	default:
		w.WriteHeader(http.StatusOK)
	}

	_, werr := w.Write([]byte(ce.Error()))
	if werr != nil {
		if logger.V(logger.ErrorLevel, logger.DefaultLogger) {
			logger.Error(werr)
		}
	}
}
