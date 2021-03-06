// Code generated by goagen v1.2.0-dirty, DO NOT EDIT.
//
// API "CSAPI": health TestHelpers
//
// Command:
// $ goagen
// --design=github.com/negronjl/csapi/design
// --out=$(GOPATH)/src/github.com/negronjl/csapi
// --version=v1.2.0-dirty

package test

import (
	"bytes"
	"context"
	"fmt"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/goatest"
	"github.com/negronjl/csapi/app"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
)

// HealthHealthNotFound runs the method Health of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func HealthHealthNotFound(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.HealthController, dc string, hgroup string, hostname string) http.ResponseWriter {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/api/v1/health/%v/%v/%v", dc, hgroup, hostname),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["dc"] = []string{fmt.Sprintf("%v", dc)}
	prms["hgroup"] = []string{fmt.Sprintf("%v", hgroup)}
	prms["hostname"] = []string{fmt.Sprintf("%v", hostname)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "HealthTest"), rw, req, prms)
	healthCtx, _err := app.NewHealthHealthContext(goaCtx, req, service)
	if _err != nil {
		panic("invalid test data " + _err.Error()) // bug
	}

	// Perform action
	_err = ctrl.Health(healthCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 404 {
		t.Errorf("invalid response status code: got %+v, expected 404", rw.Code)
	}

	// Return results
	return rw
}

// HealthHealthOK runs the method Health of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func HealthHealthOK(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.HealthController, dc string, hgroup string, hostname string) http.ResponseWriter {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/api/v1/health/%v/%v/%v", dc, hgroup, hostname),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["dc"] = []string{fmt.Sprintf("%v", dc)}
	prms["hgroup"] = []string{fmt.Sprintf("%v", hgroup)}
	prms["hostname"] = []string{fmt.Sprintf("%v", hostname)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "HealthTest"), rw, req, prms)
	healthCtx, _err := app.NewHealthHealthContext(goaCtx, req, service)
	if _err != nil {
		panic("invalid test data " + _err.Error()) // bug
	}

	// Perform action
	_err = ctrl.Health(healthCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}

	// Return results
	return rw
}
