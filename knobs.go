package knobs

import (
	"context"
	"net/http"
)

type ctxKnobErrorKey int

var knobErrorKey ctxKnobErrorKey

// CtxSetError sets the error for a set of OrderedHandlers
func ctxSetError(ctx context.Context, e error) context.Context {
	return context.WithValue(ctx, knobErrorKey, e)
}

// CtxGetError gets the error for a set of OrderedHandlers
func ctxGetError(ctx context.Context) (bool, error) {
	b, ok := ctx.Value(knobErrorKey).(error)
	return ok, b
}

// SetError places an Error in the requests Context
func SetError(r *http.Request, e error) {
	ctx := r.Context()
	ctx = ctxSetError(ctx, e)
	*r = *r.WithContext(ctx)
}

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
		if ok, _ := ctxGetError(r.Context()); ok {
			return
		}
	}
}
