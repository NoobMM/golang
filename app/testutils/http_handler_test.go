package testutils

import (
	"bytes"
	"github.com/AlekSi/pointer"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/textproto"
	"os"
	"path"
	"reflect"
	"strings"
	"testing"
)

func TestSetUpContext(t *testing.T) {
	const testFilePath = "../domain/fixtures/merchant_upload_fixture.xlsx"
	type json struct {
		Foo string `json:"foo"`
		Bar string `json:"bar"`
	}
	var testJSONBody = json{
		Foo: "foo",
		Bar: "bar",
	}
	jsonBytes, _ := getJSONRequestBody(testJSONBody)
	type formFile struct {
		File *multipart.FileHeader `form:"file"`
	}
	const formFileKey = "file"
	currentPath, _ := os.Getwd()
	testFileAbsolutePath := path.Join(currentPath, testFilePath)
	testFormFileBody := map[string]string{
		formFileKey: testFilePath,
	}
	formFileBytes, _, _ := getFormFileRequestBody(formFileKey, testFormFileBody)
	type formField struct {
		Foo string `form:"foo"`
		Bar string `form:"bar"`
	}
	var testFormFieldStruct = formField{
		Foo: "foo",
		Bar: "bar",
	}
	testFormFieldBody := map[string]string{
		"foo": testFormFieldStruct.Foo,
		"bar": testFormFieldStruct.Bar,
	}
	formFieldBytes, _, _ := getFormRequestBody(testFormFieldBody)
	type args struct {
		respRecorder *httptest.ResponseRecorder
		args         HandlerArgs
	}
	tests := []struct {
		name           string
		args           args
		wantJSON       *json
		wantFromFile   *formFile
		wantFromFields *formField
		wantRequest    *http.Request
		wantErr        error
	}{
		{
			name: "it should throw error when request method is not set",
			args: args{
				respRecorder: httptest.NewRecorder(),
				args: HandlerArgs{
					RequestMethod: nil,
					RequestPath:   pointer.ToString("/api/test"),
				},
			},
			wantErr: errors.New("args.RequestMethod is not set"),
		},
		{
			name: "it should throw error when request path is not set",
			args: args{
				respRecorder: httptest.NewRecorder(),
				args: HandlerArgs{
					RequestMethod: pointer.ToString("GET"),
					RequestPath:   nil,
				},
			},
			wantErr: errors.New("args.RequestPath is not set"),
		},
		{
			name: "it should throw error when set multiple request body",
			args: args{
				respRecorder: httptest.NewRecorder(),
				args: HandlerArgs{
					RequestMethod:       pointer.ToString("POST"),
					RequestPath:         pointer.ToString("/api/test"),
					RequestJSONBody:     testJSONBody,
					RequestFormFileBody: testFormFileBody,
				},
			},
			wantErr: errors.New("cannot set multiple request body"),
		},
		{
			name: "it should throw error when set form file request body without set file field name",
			args: args{
				respRecorder: httptest.NewRecorder(),
				args: HandlerArgs{
					RequestMethod:       pointer.ToString("POST"),
					RequestPath:         pointer.ToString("/api/test"),
					RequestFormFileBody: testFormFileBody,
				},
			},
			wantErr: errors.New("args.RequestFormFileFieldName is not set"),
		},
		{
			name: "it should not throw error when no request body",
			args: args{
				respRecorder: httptest.NewRecorder(),
				args: HandlerArgs{
					RequestMethod: pointer.ToString("GET"),
					RequestPath:   pointer.ToString("/api/test"),
				},
			},
			wantRequest: &http.Request{
				Method:     "GET",
				RequestURI: "/api/test",
				Header:     map[string][]string{"Content-Type": {gin.MIMEJSON}},
			},
		},
		{
			name: "it should not throw error when no request body with query string",
			args: args{
				respRecorder: httptest.NewRecorder(),
				args: HandlerArgs{
					RequestMethod: pointer.ToString("GET"),
					RequestPath:   pointer.ToString("/api/test"),
					RequestQueries: map[string]interface{}{
						"query": "value",
						"max":   100,
					},
				},
			},
			wantRequest: &http.Request{
				Method: "GET",
				RequestURI: getAPIPath(
					"/api/test",
					map[string]interface{}{
						"query": "value",
						"max":   100,
					},
				),
				Header: map[string][]string{"Content-Type": {gin.MIMEJSON}},
			},
		},
		{
			name: "it should not throw error when set JSON request body",
			args: args{
				respRecorder: httptest.NewRecorder(),
				args: HandlerArgs{
					RequestMethod:   pointer.ToString("POST"),
					RequestPath:     pointer.ToString("/api/test"),
					RequestJSONBody: testJSONBody,
				},
			},
			wantRequest: &http.Request{
				Method:        "POST",
				RequestURI:    "/api/test",
				Header:        map[string][]string{"Content-Type": {gin.MIMEJSON}},
				Body:          ioutil.NopCloser(jsonBytes.(*bytes.Buffer)),
				ContentLength: int64(jsonBytes.(*bytes.Buffer).Len()),
			},
			wantJSON: &testJSONBody,
		},
		{
			name: "it should not throw error when set form file request body",
			args: args{
				respRecorder: httptest.NewRecorder(),
				args: HandlerArgs{
					RequestMethod:            pointer.ToString("POST"),
					RequestPath:              pointer.ToString("/api/test"),
					RequestFormFileFieldName: pointer.ToString(formFileKey),
					RequestFormFileBody:      testFormFileBody,
				},
			},
			wantRequest: &http.Request{
				Method:        "POST",
				RequestURI:    "/api/test",
				Header:        map[string][]string{"Content-Type": {gin.MIMEMultipartPOSTForm}},
				Body:          ioutil.NopCloser(formFileBytes.(*bytes.Buffer)),
				ContentLength: int64(formFileBytes.(*bytes.Buffer).Len()),
			},
			wantFromFile: &formFile{File: &multipart.FileHeader{
				Filename: testFileAbsolutePath,
				Header: textproto.MIMEHeader{
					"Content-Disposition": {"form-data; name=\"" + formFileKey + "\"; filename=\"" + testFileAbsolutePath + "\""},
					"Content-Type":        {"application/octet-stream"},
				},
				Size: 0,
			}},
		},
		{
			name: "it should not throw error when set form field request body",
			args: args{
				respRecorder: httptest.NewRecorder(),
				args: HandlerArgs{
					RequestMethod:        pointer.ToString("POST"),
					RequestPath:          pointer.ToString("/api/test"),
					RequestFormFieldBody: testFormFieldBody,
				},
			},
			wantRequest: &http.Request{
				Method:        "POST",
				RequestURI:    "/api/test",
				Header:        map[string][]string{"Content-Type": {gin.MIMEMultipartPOSTForm}},
				Body:          ioutil.NopCloser(formFieldBytes.(*bytes.Buffer)),
				ContentLength: int64(formFieldBytes.(*bytes.Buffer).Len()),
			},
			wantFromFields: &testFormFieldStruct,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			got, err := SetUpContext(tt.args.respRecorder, tt.args.args)
			if tt.wantErr != nil && err != nil && !strings.Contains(err.Error(), tt.wantErr.Error()) {
				t.Errorf("SetUpContext() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil {
				const errorPrefix = "SetUpContext() %s = %v, want %v"
				// Request Method
				if !assert.Equal(t, got.Request.Method, tt.wantRequest.Method) {
					t.Errorf(errorPrefix, "request method", got.Request.Method, tt.wantRequest.Method)
				}
				// Content-Type
				if !strings.Contains(got.Request.Header.Get("Content-Type"), tt.wantRequest.Header.Get("Content-Type")) {
					t.Errorf(errorPrefix, "request content-type", got.Request.Header.Get("Content-Type"), tt.wantRequest.Header.Get("Content-Type"))
				}
				// Request URI
				if strings.Index(tt.wantRequest.RequestURI, "?") > 0 {
					if !AssertRequestURI(t, errorPrefix, got.Request.RequestURI, tt.wantRequest.RequestURI) {
						return
					}
				} else {
					if !assert.Equal(t, got.Request.RequestURI, tt.wantRequest.RequestURI) {
						t.Errorf(errorPrefix, "request base URI", got.Request.RequestURI, tt.wantRequest.RequestURI)
					}
				}
				// Request Content Length
				if !assert.Equal(t, got.Request.ContentLength, tt.wantRequest.ContentLength) {
					t.Errorf(errorPrefix, "request content length", got.Request.ContentLength, tt.wantRequest.ContentLength)
				}
				// JSON Body
				if tt.wantJSON != nil {
					var gotJSON json
					errBind := got.ShouldBindJSON(&gotJSON)
					if errBind != nil {
						t.Errorf("SetUpContext() failed to bind json")
					}
					if !reflect.DeepEqual(gotJSON, *tt.wantJSON) {
						t.Errorf(errorPrefix, "request json body", gotJSON, *tt.wantJSON)
					}
				}
				// Form File Body
				if tt.wantFromFile != nil {
					var gotFormFile formFile
					errBind := got.ShouldBind(&gotFormFile)
					if errBind != nil {
						t.Errorf("SetUpContext() failed to bind form file")
					}
					if !assert.Equal(t, gotFormFile.File.Filename, tt.wantFromFile.File.Filename) {
						t.Errorf(errorPrefix, "request form file body file.Filename", gotFormFile.File.Filename, tt.wantFromFile.File.Filename)
					}
					if !reflect.DeepEqual(gotFormFile.File.Header, tt.wantFromFile.File.Header) {
						t.Errorf(errorPrefix, "request form file body file.Header", gotFormFile.File.Header, tt.wantFromFile.File.Header)
					}
				}
				// Form Fields Body
				if tt.wantFromFields != nil {
					var gotFormField formField
					errBind := got.ShouldBind(&gotFormField)
					if errBind != nil {
						t.Errorf("SetUpContext() failed to bind form field(s)")
					}
					if !reflect.DeepEqual(gotFormField, *tt.wantFromFields) {
						t.Errorf(errorPrefix, "request form field body", gotFormField, *tt.wantJSON)
					}
				}
			}
		})
	}
}
