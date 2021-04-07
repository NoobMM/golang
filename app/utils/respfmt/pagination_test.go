package respfmt_test

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/AlekSi/pointer"
	"github.com/NoobMM/golang/app/environments"
	"github.com/NoobMM/golang/app/testutils"
	"github.com/NoobMM/golang/app/utils/respfmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestPagination_New(t *testing.T) {
	httpContext := gin.Context{
		Request: &http.Request{
			Method: "Get",
			URL: &url.URL{
				Path: "/api/data",
			},
			Proto:      "HTTP/1.1",
			ProtoMajor: 1,
			ProtoMinor: 1,
			Header: map[string][]string{
				"Content-Type": {gin.MIMEJSON},
			},
			Host: "localhost",
		},
	}
	type args struct {
		c        *gin.Context
		dataSize uint64
		count    uint64
		limit    uint64
		start    uint64
	}
	tests := []struct {
		name string
		args args
		want *respfmt.Pagination
	}{
		{
			name: "it should return proper information",
			args: args{
				c:        &httpContext,
				dataSize: 10,
				count:    40,
				limit:    10,
				start:    0,
			},
			want: &respfmt.Pagination{
				Total: pointer.ToUint64(40),
				Size:  pointer.ToUint64(10),
				Limit: pointer.ToUint64(10),
				Start: pointer.ToUint64(0),
				Links: &respfmt.Link{
					Previous: nil,
					Self:     pointer.ToString("http://localhost/api/data?limit=10&start=0"),
					Next:     pointer.ToString("http://localhost/api/data?limit=10&start=10"),
				},
			},
		},
		{
			name: "it should return proper information when start not zero",
			args: args{
				c:        &httpContext,
				dataSize: 10,
				count:    40,
				limit:    10,
				start:    10,
			},
			want: &respfmt.Pagination{
				Total: pointer.ToUint64(40),
				Size:  pointer.ToUint64(10),
				Limit: pointer.ToUint64(10),
				Start: pointer.ToUint64(10),
				Links: &respfmt.Link{
					Previous: pointer.ToString("http://localhost/api/data?limit=10&start=0"),
					Self:     pointer.ToString("http://localhost/api/data?limit=10&start=10"),
					Next:     pointer.ToString("http://localhost/api/data?limit=10&start=20"),
				},
			},
		},
		{
			name: "it should return proper information when start not zero and it is last page",
			args: args{
				c:        &httpContext,
				dataSize: 10,
				count:    20,
				limit:    10,
				start:    10,
			},
			want: &respfmt.Pagination{
				Total: pointer.ToUint64(20),
				Size:  pointer.ToUint64(10),
				Limit: pointer.ToUint64(10),
				Start: pointer.ToUint64(10),
				Links: &respfmt.Link{
					Previous: pointer.ToString("http://localhost/api/data?limit=10&start=0"),
					Self:     pointer.ToString("http://localhost/api/data?limit=10&start=10"),
					Next:     nil,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			environments.BaseURL = "http://localhost"
			got := new(respfmt.Pagination).New(tt.args.c, tt.args.dataSize, tt.args.count, tt.args.limit, tt.args.start)
			if !assert.Equal(t, got.Total, tt.want.Total) {
				t.Errorf("Pagination.New().Total = %v, want %v", got.Total, tt.want.Total)
			}
			if !assert.Equal(t, got.Size, tt.want.Size) {
				t.Errorf("Pagination.New().Size = %v, want %v", got.Size, tt.want.Size)
			}
			if !assert.Equal(t, got.Limit, tt.want.Limit) {
				t.Errorf("Pagination.New().Limit = %v, want %v", got.Limit, tt.want.Limit)
			}
			if !assert.Equal(t, got.Start, tt.want.Start) {
				t.Errorf("Pagination.New().Start = %v, want %v", got.Start, tt.want.Start)
			}
			if tt.want.Links.Previous != nil && !testutils.AssertRequestURI(t, "Pagination.New().Links.Previous %s = %v, want %v", *got.Links.Previous, *tt.want.Links.Previous) {
				return
			}
			if tt.want.Links.Self != nil && !testutils.AssertRequestURI(t, "Pagination.New().Links.Self %s = %v, want %v", *got.Links.Self, *tt.want.Links.Self) {
				return
			}
			if tt.want.Links.Next != nil && !testutils.AssertRequestURI(t, "Pagination.New().Links.Next %s = %v, want %v", *got.Links.Next, *tt.want.Links.Next) {
				return
			}
		})
	}
}
