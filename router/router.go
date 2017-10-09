package router

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/eapache/go-resiliency/breaker"
	"github.com/pressly/chi"
	"github.com/prometheus/client_golang/prometheus"
)

/*
Example of router package, the router wrapper is based on pressly/chi
*/

// Router struct
type Router struct {
	opt Option
	r   chi.Router
	cb  *breaker.Breaker
}

// Option for router
type Option struct {
	Timeout time.Duration
	// Instrumenet white-box monitoring from the application, provide data about handler
	Instrument bool
}

// New router
func New(opt Option) *Router {
	chiRouter := chi.NewRouter()
	rtr := &Router{
		r:   chiRouter,
		opt: opt,
		// this is only an example
		cb: breaker.New(10, 2, time.Second*15),
	}
	return rtr
}

const (
	BreakerOpen = "rtr_breaker_open"
)

func OpenBreaker(r *http.Request) {
	ctx := context.WithValue(r.Context(), BreakerOpen, 1)
	r = r.WithContext(ctx)
}

func (rtr *Router) circuitbreaker(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := rtr.cb.Run(func() error {
			h(w, r)
			v := r.Context().Value(BreakerOpen)
			if v == nil || v != 1 {
				return nil
			}
			return errors.New("Circuit breaker open")
		})
		if err == nil || err != breaker.ErrBreakerOpen {
			return
		}
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte("Service unavailable"))
	}
}

// timeout middleware
func (rtr *Router) timeout(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		opt := rtr.opt
		// cancel context
		if opt.Timeout > 0 {
			ctx, cancel := context.WithTimeout(r.Context(), opt.Timeout*time.Second)
			defer cancel()
			r = r.WithContext(ctx)
		}

		doneChan := make(chan bool)
		go func() {
			h(w, r)
			doneChan <- true
		}()
		select {
		case <-r.Context().Done():
			// only an example response
			resp := map[string]string{
				"error": "Request timed out",
			}
			jsonResp, _ := json.Marshal(resp)
			w.WriteHeader(http.StatusRequestTimeout)
			// only an example response
			w.Write(jsonResp)
			return
		case <-doneChan:
			return
		}
	}
}

// URLParam get param from rest request
func URLParam(r *http.Request, key string) string {
	return chi.URLParam(r, key)
}

// URLParamFromCtx get param from rest request from context
func URLParamFromCtx(ctx context.Context, key string) string {
	return chi.URLParamFromCtx(ctx, key)
}

// Get function
func (rtr *Router) Get(pattern string, h http.HandlerFunc) {
	rtr.r.Get(pattern, prometheus.InstrumentHandlerFunc(pattern, rtr.timeout(h)))
}

// Post function
func (rtr *Router) Post(pattern string, h http.HandlerFunc) {
	rtr.r.Post(pattern, prometheus.InstrumentHandlerFunc(pattern, rtr.timeout(h)))
}

// Put function
func (rtr *Router) Put(pattern string, h http.HandlerFunc) {
	rtr.r.Put(pattern, prometheus.InstrumentHandlerFunc(pattern, rtr.timeout(h)))
}

// Delete function
func (rtr *Router) Delete(pattern string, h http.HandlerFunc) {
	rtr.r.Delete(pattern, prometheus.InstrumentHandlerFunc(pattern, rtr.timeout(h)))
}

// Patch function
func (rtr *Router) Patch(pattern string, h http.HandlerFunc) {
	rtr.r.Patch(pattern, prometheus.InstrumentHandlerFunc(pattern, rtr.timeout(h)))
}
