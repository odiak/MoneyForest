// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "MoneyForest": category TestHelpers
//
// Command:
// $ goagen
// --design=github.com/odiak/MoneyForest/design
// --out=$(GOPATH)/src/github.com/odiak/MoneyForest
// --version=v1.3.0

package test

import (
	"bytes"
	"context"
	"fmt"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/goatest"
	"github.com/odiak/MoneyForest/app"
	uuid "github.com/satori/go.uuid"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
)

// CreateCategoryBadRequest runs the method Create of the given controller with the given parameters and payload.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func CreateCategoryBadRequest(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.CategoryController, payload *app.CategoryPayload) (http.ResponseWriter, error) {
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

	// Validate payload
	err := payload.Validate()
	if err != nil {
		e, ok := err.(goa.ServiceError)
		if !ok {
			panic(err) // bug
		}
		return nil, e
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/api/categories"),
	}
	req, _err := http.NewRequest("POST", u.String(), nil)
	if _err != nil {
		panic("invalid test " + _err.Error()) // bug
	}
	prms := url.Values{}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "CategoryTest"), rw, req, prms)
	createCtx, __err := app.NewCreateCategoryContext(goaCtx, req, service)
	if __err != nil {
		_e, _ok := __err.(goa.ServiceError)
		if !_ok {
			panic("invalid test data " + __err.Error()) // bug
		}
		return nil, _e
	}
	createCtx.Payload = payload

	// Perform action
	__err = ctrl.Create(createCtx)

	// Validate response
	if __err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", __err, logBuf.String())
	}
	if rw.Code != 400 {
		t.Errorf("invalid response status code: got %+v, expected 400", rw.Code)
	}
	var mt error
	if resp != nil {
		var __ok bool
		mt, __ok = resp.(error)
		if !__ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of error", resp, resp)
		}
	}

	// Return results
	return rw, mt
}

// CreateCategoryOK runs the method Create of the given controller with the given parameters and payload.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func CreateCategoryOK(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.CategoryController, payload *app.CategoryPayload) (http.ResponseWriter, *app.CategoryMedia) {
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

	// Validate payload
	err := payload.Validate()
	if err != nil {
		e, ok := err.(goa.ServiceError)
		if !ok {
			panic(err) // bug
		}
		t.Errorf("unexpected payload validation error: %+v", e)
		return nil, nil
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/api/categories"),
	}
	req, _err := http.NewRequest("POST", u.String(), nil)
	if _err != nil {
		panic("invalid test " + _err.Error()) // bug
	}
	prms := url.Values{}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "CategoryTest"), rw, req, prms)
	createCtx, __err := app.NewCreateCategoryContext(goaCtx, req, service)
	if __err != nil {
		_e, _ok := __err.(goa.ServiceError)
		if !_ok {
			panic("invalid test data " + __err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", _e)
		return nil, nil
	}
	createCtx.Payload = payload

	// Perform action
	__err = ctrl.Create(createCtx)

	// Validate response
	if __err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", __err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt *app.CategoryMedia
	if resp != nil {
		var __ok bool
		mt, __ok = resp.(*app.CategoryMedia)
		if !__ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of app.CategoryMedia", resp, resp)
		}
		__err = mt.Validate()
		if __err != nil {
			t.Errorf("invalid response media type: %s", __err)
		}
	}

	// Return results
	return rw, mt
}

// CreateCategoryOKWithChildren runs the method Create of the given controller with the given parameters and payload.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func CreateCategoryOKWithChildren(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.CategoryController, payload *app.CategoryPayload) (http.ResponseWriter, *app.CategoryMediaWithChildren) {
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

	// Validate payload
	err := payload.Validate()
	if err != nil {
		e, ok := err.(goa.ServiceError)
		if !ok {
			panic(err) // bug
		}
		t.Errorf("unexpected payload validation error: %+v", e)
		return nil, nil
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/api/categories"),
	}
	req, _err := http.NewRequest("POST", u.String(), nil)
	if _err != nil {
		panic("invalid test " + _err.Error()) // bug
	}
	prms := url.Values{}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "CategoryTest"), rw, req, prms)
	createCtx, __err := app.NewCreateCategoryContext(goaCtx, req, service)
	if __err != nil {
		_e, _ok := __err.(goa.ServiceError)
		if !_ok {
			panic("invalid test data " + __err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", _e)
		return nil, nil
	}
	createCtx.Payload = payload

	// Perform action
	__err = ctrl.Create(createCtx)

	// Validate response
	if __err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", __err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt *app.CategoryMediaWithChildren
	if resp != nil {
		var __ok bool
		mt, __ok = resp.(*app.CategoryMediaWithChildren)
		if !__ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of app.CategoryMediaWithChildren", resp, resp)
		}
		__err = mt.Validate()
		if __err != nil {
			t.Errorf("invalid response media type: %s", __err)
		}
	}

	// Return results
	return rw, mt
}

// CreateCategoryOKWithParent runs the method Create of the given controller with the given parameters and payload.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func CreateCategoryOKWithParent(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.CategoryController, payload *app.CategoryPayload) (http.ResponseWriter, *app.CategoryMediaWithParent) {
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

	// Validate payload
	err := payload.Validate()
	if err != nil {
		e, ok := err.(goa.ServiceError)
		if !ok {
			panic(err) // bug
		}
		t.Errorf("unexpected payload validation error: %+v", e)
		return nil, nil
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/api/categories"),
	}
	req, _err := http.NewRequest("POST", u.String(), nil)
	if _err != nil {
		panic("invalid test " + _err.Error()) // bug
	}
	prms := url.Values{}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "CategoryTest"), rw, req, prms)
	createCtx, __err := app.NewCreateCategoryContext(goaCtx, req, service)
	if __err != nil {
		_e, _ok := __err.(goa.ServiceError)
		if !_ok {
			panic("invalid test data " + __err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", _e)
		return nil, nil
	}
	createCtx.Payload = payload

	// Perform action
	__err = ctrl.Create(createCtx)

	// Validate response
	if __err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", __err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt *app.CategoryMediaWithParent
	if resp != nil {
		var __ok bool
		mt, __ok = resp.(*app.CategoryMediaWithParent)
		if !__ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of app.CategoryMediaWithParent", resp, resp)
		}
		__err = mt.Validate()
		if __err != nil {
			t.Errorf("invalid response media type: %s", __err)
		}
	}

	// Return results
	return rw, mt
}

// DeleteCategoryNoContent runs the method Delete of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func DeleteCategoryNoContent(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.CategoryController, categoryID uuid.UUID) http.ResponseWriter {
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
		Path: fmt.Sprintf("/api/categories/%v", categoryID),
	}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["categoryID"] = []string{fmt.Sprintf("%v", categoryID)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "CategoryTest"), rw, req, prms)
	deleteCtx, _err := app.NewDeleteCategoryContext(goaCtx, req, service)
	if _err != nil {
		e, ok := _err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + _err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", e)
		return nil
	}

	// Perform action
	_err = ctrl.Delete(deleteCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 204 {
		t.Errorf("invalid response status code: got %+v, expected 204", rw.Code)
	}

	// Return results
	return rw
}

// DeleteCategoryNotFound runs the method Delete of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func DeleteCategoryNotFound(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.CategoryController, categoryID uuid.UUID) http.ResponseWriter {
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
		Path: fmt.Sprintf("/api/categories/%v", categoryID),
	}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["categoryID"] = []string{fmt.Sprintf("%v", categoryID)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "CategoryTest"), rw, req, prms)
	deleteCtx, _err := app.NewDeleteCategoryContext(goaCtx, req, service)
	if _err != nil {
		e, ok := _err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + _err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", e)
		return nil
	}

	// Perform action
	_err = ctrl.Delete(deleteCtx)

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

// ListCategoryOK runs the method List of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ListCategoryOK(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.CategoryController, count int, page int) (http.ResponseWriter, *app.CategoryListMedia) {
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
	query := url.Values{}
	{
		sliceVal := []string{strconv.Itoa(count)}
		query["count"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(page)}
		query["page"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/api/categories"),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	{
		sliceVal := []string{strconv.Itoa(count)}
		prms["count"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(page)}
		prms["page"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "CategoryTest"), rw, req, prms)
	listCtx, _err := app.NewListCategoryContext(goaCtx, req, service)
	if _err != nil {
		e, ok := _err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + _err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", e)
		return nil, nil
	}

	// Perform action
	_err = ctrl.List(listCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt *app.CategoryListMedia
	if resp != nil {
		var _ok bool
		mt, _ok = resp.(*app.CategoryListMedia)
		if !_ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of app.CategoryListMedia", resp, resp)
		}
		_err = mt.Validate()
		if _err != nil {
			t.Errorf("invalid response media type: %s", _err)
		}
	}

	// Return results
	return rw, mt
}

// UpdateCategoryBadRequest runs the method Update of the given controller with the given parameters and payload.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func UpdateCategoryBadRequest(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.CategoryController, categoryID uuid.UUID, payload *app.UpdateCategoryPayload) (http.ResponseWriter, error) {
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

	// Validate payload
	err := payload.Validate()
	if err != nil {
		e, ok := err.(goa.ServiceError)
		if !ok {
			panic(err) // bug
		}
		return nil, e
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/api/categories/%v", categoryID),
	}
	req, _err := http.NewRequest("PUT", u.String(), nil)
	if _err != nil {
		panic("invalid test " + _err.Error()) // bug
	}
	prms := url.Values{}
	prms["categoryID"] = []string{fmt.Sprintf("%v", categoryID)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "CategoryTest"), rw, req, prms)
	updateCtx, __err := app.NewUpdateCategoryContext(goaCtx, req, service)
	if __err != nil {
		_e, _ok := __err.(goa.ServiceError)
		if !_ok {
			panic("invalid test data " + __err.Error()) // bug
		}
		return nil, _e
	}
	updateCtx.Payload = payload

	// Perform action
	__err = ctrl.Update(updateCtx)

	// Validate response
	if __err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", __err, logBuf.String())
	}
	if rw.Code != 400 {
		t.Errorf("invalid response status code: got %+v, expected 400", rw.Code)
	}
	var mt error
	if resp != nil {
		var __ok bool
		mt, __ok = resp.(error)
		if !__ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of error", resp, resp)
		}
	}

	// Return results
	return rw, mt
}

// UpdateCategoryNotFound runs the method Update of the given controller with the given parameters and payload.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func UpdateCategoryNotFound(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.CategoryController, categoryID uuid.UUID, payload *app.UpdateCategoryPayload) http.ResponseWriter {
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

	// Validate payload
	err := payload.Validate()
	if err != nil {
		e, ok := err.(goa.ServiceError)
		if !ok {
			panic(err) // bug
		}
		t.Errorf("unexpected payload validation error: %+v", e)
		return nil
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/api/categories/%v", categoryID),
	}
	req, _err := http.NewRequest("PUT", u.String(), nil)
	if _err != nil {
		panic("invalid test " + _err.Error()) // bug
	}
	prms := url.Values{}
	prms["categoryID"] = []string{fmt.Sprintf("%v", categoryID)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "CategoryTest"), rw, req, prms)
	updateCtx, __err := app.NewUpdateCategoryContext(goaCtx, req, service)
	if __err != nil {
		_e, _ok := __err.(goa.ServiceError)
		if !_ok {
			panic("invalid test data " + __err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", _e)
		return nil
	}
	updateCtx.Payload = payload

	// Perform action
	__err = ctrl.Update(updateCtx)

	// Validate response
	if __err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", __err, logBuf.String())
	}
	if rw.Code != 404 {
		t.Errorf("invalid response status code: got %+v, expected 404", rw.Code)
	}

	// Return results
	return rw
}

// UpdateCategoryOK runs the method Update of the given controller with the given parameters and payload.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func UpdateCategoryOK(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.CategoryController, categoryID uuid.UUID, payload *app.UpdateCategoryPayload) (http.ResponseWriter, *app.CategoryMedia) {
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

	// Validate payload
	err := payload.Validate()
	if err != nil {
		e, ok := err.(goa.ServiceError)
		if !ok {
			panic(err) // bug
		}
		t.Errorf("unexpected payload validation error: %+v", e)
		return nil, nil
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/api/categories/%v", categoryID),
	}
	req, _err := http.NewRequest("PUT", u.String(), nil)
	if _err != nil {
		panic("invalid test " + _err.Error()) // bug
	}
	prms := url.Values{}
	prms["categoryID"] = []string{fmt.Sprintf("%v", categoryID)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "CategoryTest"), rw, req, prms)
	updateCtx, __err := app.NewUpdateCategoryContext(goaCtx, req, service)
	if __err != nil {
		_e, _ok := __err.(goa.ServiceError)
		if !_ok {
			panic("invalid test data " + __err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", _e)
		return nil, nil
	}
	updateCtx.Payload = payload

	// Perform action
	__err = ctrl.Update(updateCtx)

	// Validate response
	if __err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", __err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt *app.CategoryMedia
	if resp != nil {
		var __ok bool
		mt, __ok = resp.(*app.CategoryMedia)
		if !__ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of app.CategoryMedia", resp, resp)
		}
		__err = mt.Validate()
		if __err != nil {
			t.Errorf("invalid response media type: %s", __err)
		}
	}

	// Return results
	return rw, mt
}

// UpdateCategoryOKWithChildren runs the method Update of the given controller with the given parameters and payload.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func UpdateCategoryOKWithChildren(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.CategoryController, categoryID uuid.UUID, payload *app.UpdateCategoryPayload) (http.ResponseWriter, *app.CategoryMediaWithChildren) {
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

	// Validate payload
	err := payload.Validate()
	if err != nil {
		e, ok := err.(goa.ServiceError)
		if !ok {
			panic(err) // bug
		}
		t.Errorf("unexpected payload validation error: %+v", e)
		return nil, nil
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/api/categories/%v", categoryID),
	}
	req, _err := http.NewRequest("PUT", u.String(), nil)
	if _err != nil {
		panic("invalid test " + _err.Error()) // bug
	}
	prms := url.Values{}
	prms["categoryID"] = []string{fmt.Sprintf("%v", categoryID)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "CategoryTest"), rw, req, prms)
	updateCtx, __err := app.NewUpdateCategoryContext(goaCtx, req, service)
	if __err != nil {
		_e, _ok := __err.(goa.ServiceError)
		if !_ok {
			panic("invalid test data " + __err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", _e)
		return nil, nil
	}
	updateCtx.Payload = payload

	// Perform action
	__err = ctrl.Update(updateCtx)

	// Validate response
	if __err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", __err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt *app.CategoryMediaWithChildren
	if resp != nil {
		var __ok bool
		mt, __ok = resp.(*app.CategoryMediaWithChildren)
		if !__ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of app.CategoryMediaWithChildren", resp, resp)
		}
		__err = mt.Validate()
		if __err != nil {
			t.Errorf("invalid response media type: %s", __err)
		}
	}

	// Return results
	return rw, mt
}

// UpdateCategoryOKWithParent runs the method Update of the given controller with the given parameters and payload.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func UpdateCategoryOKWithParent(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.CategoryController, categoryID uuid.UUID, payload *app.UpdateCategoryPayload) (http.ResponseWriter, *app.CategoryMediaWithParent) {
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

	// Validate payload
	err := payload.Validate()
	if err != nil {
		e, ok := err.(goa.ServiceError)
		if !ok {
			panic(err) // bug
		}
		t.Errorf("unexpected payload validation error: %+v", e)
		return nil, nil
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/api/categories/%v", categoryID),
	}
	req, _err := http.NewRequest("PUT", u.String(), nil)
	if _err != nil {
		panic("invalid test " + _err.Error()) // bug
	}
	prms := url.Values{}
	prms["categoryID"] = []string{fmt.Sprintf("%v", categoryID)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "CategoryTest"), rw, req, prms)
	updateCtx, __err := app.NewUpdateCategoryContext(goaCtx, req, service)
	if __err != nil {
		_e, _ok := __err.(goa.ServiceError)
		if !_ok {
			panic("invalid test data " + __err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", _e)
		return nil, nil
	}
	updateCtx.Payload = payload

	// Perform action
	__err = ctrl.Update(updateCtx)

	// Validate response
	if __err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", __err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt *app.CategoryMediaWithParent
	if resp != nil {
		var __ok bool
		mt, __ok = resp.(*app.CategoryMediaWithParent)
		if !__ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of app.CategoryMediaWithParent", resp, resp)
		}
		__err = mt.Validate()
		if __err != nil {
			t.Errorf("invalid response media type: %s", __err)
		}
	}

	// Return results
	return rw, mt
}
