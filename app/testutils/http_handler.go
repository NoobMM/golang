package testutils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http/httptest"
	"os"
	"path"
	"reflect"
	"sort"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

// HandlerArgs ...
type HandlerArgs struct {
	RequestMethod            *string
	RequestPath              *string
	RequestHeaders           map[string]string
	RequestQueries           map[string]interface{}
	RequestFormFileFieldName *string
	RequestFormFileBody      map[string]string
	RequestFormFieldBody     map[string]string
	RequestJSONBody          interface{}
	RequestParams            []gin.Param
}

func getAPIPath(requestPath string, requestQueries map[string]interface{}) string {
	if len(requestQueries) > 0 {
		var queries []string
		for key, value := range requestQueries {
			queries = append(queries, fmt.Sprintf("%v=%v", key, value))
		}
		return fmt.Sprintf("%s?%s", requestPath, strings.Join(queries, "&"))
	}
	return requestPath
}

func getFormFileRequestBody(fileNameField string, fields map[string]string) (io.Reader, string, error) {
	requestBody := new(bytes.Buffer)
	multipartWriter := multipart.NewWriter(requestBody)
	for key, value := range fields {
		if key == fileNameField {
			currentPath, _ := os.Getwd()
			file := path.Join(currentPath, value)
			if _, err := multipartWriter.CreateFormFile(key, file); err != nil {
				return nil, "", err
			}
			continue
		}
		if err := multipartWriter.WriteField(key, value); err != nil {
			return nil, "", err
		}
	}
	_ = multipartWriter.Close()
	return requestBody, multipartWriter.FormDataContentType(), nil
}

func getFormRequestBody(formFields map[string]string) (io.Reader, string, error) {
	requestBody := new(bytes.Buffer)
	multipartWriter := multipart.NewWriter(requestBody)
	for key, value := range formFields {
		err := multipartWriter.WriteField(key, value)
		if err != nil {
			return nil, "", err
		}
	}
	_ = multipartWriter.Close()
	return requestBody, multipartWriter.FormDataContentType(), nil
}

func getJSONRequestBody(JSON interface{}) (io.Reader, error) {
	jsonBytes, err := json.Marshal(JSON)
	if err != nil {
		return nil, err
	}
	return bytes.NewBuffer(jsonBytes), nil
}

func getRequestBodyAndType(args HandlerArgs) (io.Reader, *string, error) {
	contentType := gin.MIMEJSON
	if args.RequestFormFileBody != nil {
		if args.RequestFormFileFieldName == nil {
			return nil, nil, errors.New("args.RequestFormFileFieldName is not set")
		}
		body, contentType, err := getFormFileRequestBody(*args.RequestFormFileFieldName, args.RequestFormFileBody)
		if err != nil {
			return nil, nil, err
		}
		return body, &contentType, nil
	}
	if args.RequestJSONBody != nil {
		body, err := getJSONRequestBody(args.RequestJSONBody)
		if err != nil {
			return nil, nil, err
		}
		return body, &contentType, nil
	}
	if args.RequestFormFieldBody != nil {
		body, contentType, err := getFormRequestBody(args.RequestFormFieldBody)
		if err != nil {
			return nil, nil, err
		}
		return body, &contentType, nil
	}
	return nil, &contentType, nil
}

// AssertRequestURI is a function for assert HTTP URI
func AssertRequestURI(t *testing.T, prefix, gotURI, wantURI string) bool {
	gotURIMark := strings.Index(gotURI, "?")
	wantURIMark := strings.Index(wantURI, "?")
	gotBaseURI := gotURI[:gotURIMark]
	wantBaseURI := wantURI[:wantURIMark]
	if strings.Compare(gotBaseURI, wantBaseURI) != 0 {
		t.Errorf(prefix, "request base URI", gotBaseURI, wantBaseURI)
		return false
	}
	gotQueries := strings.Split(gotURI[gotURIMark+1:], "&")
	wantQueries := strings.Split(wantURI[gotURIMark+1:], "&")
	sort.Strings(gotQueries)
	sort.Strings(wantQueries)
	if !reflect.DeepEqual(gotQueries, wantQueries) {
		t.Errorf(prefix, "request URI queries", gotQueries, wantQueries)
		return false
	}
	return true
}

// SetUpContext is a function for setup http context
func SetUpContext(respRecorder *httptest.ResponseRecorder, args HandlerArgs) (*gin.Context, error) {
	if args.RequestMethod == nil {
		return nil, errors.New("args.RequestMethod is not set")
	}
	if args.RequestPath == nil {
		return nil, errors.New("args.RequestPath is not set")
	}
	checkBody := 0
	for _, b := range []bool{args.RequestJSONBody != nil, args.RequestFormFieldBody != nil, args.RequestFormFileBody != nil} {
		if b {
			checkBody++
		}
	}
	if checkBody > 1 {
		return nil, errors.New("cannot set multiple request body")
	}
	requestBody, contentType, err := getRequestBodyAndType(args)
	if err != nil {
		return nil, err
	}
	ctx, _ := gin.CreateTestContext(respRecorder)
	ctx.Request = httptest.NewRequest(*args.RequestMethod, getAPIPath(*args.RequestPath, args.RequestQueries), requestBody)
	ctx.Request.Header.Set("Content-Type", *contentType)
	for key, value := range args.RequestHeaders {
		ctx.Request.Header.Set(key, value)
	}
	if args.RequestParams != nil {
		ctx.Params = args.RequestParams
	}
	return ctx, nil
}
