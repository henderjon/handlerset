package knobs

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestOrderHandlerOrder(t *testing.T) {
	var b bytes.Buffer
	set := New(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			b.WriteString("a")
		}),
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			b.WriteString("b")
		}),
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			b.WriteString("c")
		}),
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			b.WriteString("d")
		}),
	)

	req := httptest.NewRequest("GET", "http://127.0.0.1", nil)
	w := httptest.NewRecorder()
	set.ServeHTTP(w, req)

	if diff := cmp.Diff(b.String(), "abcd"); diff != "" {
		t.Errorf("incoming data does not match expected data: get data (/get/users): (-got +want)\n%s", diff)
	}
}

// func TestOrderHandlerCtxCancellation(t *testing.T) {
// 	var b bytes.Buffer
// 	set := New(
// 		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 			b.WriteString("a")
// 		}),
// 		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 			b.WriteString("b")
// 		}),
// 		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 			_, cancel := context.WithCancel(r.Context())
// 			cancel()
// 			return
// 		}),
// 		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 			b.WriteString("c")
// 			b.WriteString("d")
// 		}),
// 	)

// 	req := httptest.NewRequest("GET", "http://127.0.0.1", nil)
// 	w := httptest.NewRecorder()
// 	set.ServeHTTP(w, req)

// 	if diff := cmp.Diff(b.String(), "ab"); diff != "" {
// 		t.Errorf("incoming data does not match expected data: get data (/get/users): (-got +want)\n%s", diff)
// 	}
// }
