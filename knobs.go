package knobs

import (
	"net/http"
)

// OrderedHandlers wraps a variable number of http.Handlers that are executed in order
type OrderedHandlers struct {
	handlers []http.Handler
}

// New creates a new OrderedHandlers
func New(handlers ...http.Handler) OrderedHandlers {
	return OrderedHandlers{
		handlers: handlers,
	}
}

func (h OrderedHandlers) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, handler := range h.handlers {
		handler.ServeHTTP(w, r)
	}
}
