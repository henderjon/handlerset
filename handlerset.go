package handlerset

import (
	"context"
	"net/http"
)

type ctxKnobErrorKey int

var knobErrorKey ctxKnobErrorKey

// CtxSetError sets the error for a set of HandlerSet
func ctxSetError(ctx context.Context, e error) context.Context {
	return context.WithValue(ctx, knobErrorKey, e)
}

// CtxGetError gets the error for a set of HandlerSet
func ctxGetError(ctx context.Context) (bool, error) {
	b, ok := ctx.Value(knobErrorKey).(error)
	return ok, b
}

// Cancel places an Error in the requests Context
func Cancel(r *http.Request, e error) {
	ctx := r.Context()
	ctx = ctxSetError(ctx, e)
	*r = *r.WithContext(ctx)
}

// HandlerSet wraps a variable number of http.Handlers that are executed in order
type HandlerSet struct {
	handlers []http.Handler
}

// New creates a new HandlerSet
func New(handlers ...http.Handler) HandlerSet {
	return HandlerSet{
		handlers: handlers,
	}
}

// Append adds a handler to a specific HandlerSet after it's been initialized
func (h *HandlerSet) Append(handler http.Handler) {
	h.handlers = append(h.handlers, handler)
}

// Prepend adds a handler to a specific HandlerSet after it's been initialized
func (h *HandlerSet) Prepend(handler http.Handler) {
	h.handlers = append([]http.Handler{handler}, h.handlers...)
}

// ServeHTTP fullfills the http.Hander interface
func (h HandlerSet) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, handler := range h.handlers {
		handler.ServeHTTP(w, r)
		if ok, _ := ctxGetError(r.Context()); ok {
			return
		}
	}
}
