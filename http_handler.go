package zdpgo_log

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/zhangdapeng520/zdpgo_log/core"
)

// ServeHTTP is a simple JSON endpoint that can report on or change the current
// logging level.
//
// GET
//
// The GET request returns a JSON description of the current logging level like:
//   {"level":"info"}
//
// PUT
//
// The PUT request changes the logging level. It is perfectly safe to change the
// logging level while a program is running. Two content types are supported:
//
//    Content-Type: application/x-www-form-urlencoded
//
// With this content type, the level can be provided through the request body or
// a query parameter. The log level is URL encoded like:
//
//    level=debug
//
// The request body takes precedence over the query parameter, if both are
// specified.
//
// This content type is the default for a curl PUT request. Following are two
// example curl requests that both set the logging level to debug.
//
//    curl -X PUT localhost:8080/log/level?level=debug
//    curl -X PUT localhost:8080/log/level -d level=debug
//
// For any other content type, the payload is expected to be JSON encoded and
// look like:
//
//   {"level":"info"}
//
// An example curl request could look like this:
//
//    curl -X PUT localhost:8080/log/level -H "Content-Type: application/json" -d '{"level":"debug"}'
//
func (lvl AtomicLevel) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	type errorResponse struct {
		Error string `json:"error"`
	}
	type payload struct {
		Level core.Level `json:"level"`
	}

	enc := json.NewEncoder(w)

	switch r.Method {
	case http.MethodGet:
		enc.Encode(payload{Level: lvl.Level()})
	case http.MethodPut:
		requestedLvl, err := decodePutRequest(r.Header.Get("Content-Type"), r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			enc.Encode(errorResponse{Error: err.Error()})
			return
		}
		lvl.SetLevel(requestedLvl)
		enc.Encode(payload{Level: lvl.Level()})
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		enc.Encode(errorResponse{
			Error: "Only GET and PUT are supported.",
		})
	}
}

// Decodes incoming PUT requests and returns the requested logging level.
func decodePutRequest(contentType string, r *http.Request) (core.Level, error) {
	if contentType == "application/x-www-form-urlencoded" {
		return decodePutURL(r)
	}
	return decodePutJSON(r.Body)
}

func decodePutURL(r *http.Request) (core.Level, error) {
	lvl := r.FormValue("level")
	if lvl == "" {
		return 0, fmt.Errorf("must specify logging level")
	}
	var l core.Level
	if err := l.UnmarshalText([]byte(lvl)); err != nil {
		return 0, err
	}
	return l, nil
}

func decodePutJSON(body io.Reader) (core.Level, error) {
	var pld struct {
		Level *core.Level `json:"level"`
	}
	if err := json.NewDecoder(body).Decode(&pld); err != nil {
		return 0, fmt.Errorf("malformed request body: %v", err)
	}
	if pld.Level == nil {
		return 0, fmt.Errorf("must specify logging level")
	}
	return *pld.Level, nil

}