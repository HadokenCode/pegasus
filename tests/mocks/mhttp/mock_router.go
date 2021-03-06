package mhttp

import (
	"github.com/gorilla/mux"
	"net/http"
)

// MockRouter mock mux.Router
type MockRouter struct {
	HandleFuncMock func(path string, f func(http.ResponseWriter, *http.Request)) *mux.Route
}

// HandleFunc mock for mux.Router HandleFunc
func (m MockRouter) HandleFunc(path string, f func(http.ResponseWriter, *http.Request)) *mux.Route {
	if m.HandleFuncMock != nil {
		return m.HandleFuncMock(path, f)
	}
	return &mux.Route{}
}
