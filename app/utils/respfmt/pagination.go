package respfmt

import (
	"strconv"

	"github.com/AlekSi/pointer"
	"github.com/NoobMM/golang/app/environments"
	"github.com/gin-gonic/gin"
)

// Pagination is a response model that define how pagination should look like
type Pagination struct {
	Total *uint64 `json:"total"`
	Size  *uint64 `json:"size"`
	Limit *uint64 `json:"limit"`
	Start *uint64 `json:"start"`
	Links *Link   `json:"links"`
}

// Link is a sub model of Pagination that define how to call previous or next page
type Link struct {
	Previous *string `json:"prev"`
	Self     *string `json:"self"`
	Next     *string `json:"next"`
}

// New is a function to create respfmt.Pagination
func (p *Pagination) New(c *gin.Context, dataSize, count, limit, start uint64) *Pagination {
	prevURL, selfURL, nextURL := *c.Request.URL, *c.Request.URL, *c.Request.URL
	prevQuery, selfQuery, nextQuery := selfURL.Query(), nextURL.Query(), prevURL.Query()
	nextOffset := start + limit
	prevOffset := start - limit
	if start != 0 && start < limit {
		prevOffset = 0
	}

	if limit != 0 {
		strLimit := strconv.FormatUint(limit, 10)
		prevQuery.Set("limit", strLimit)
		selfQuery.Set("limit", strLimit)
		nextQuery.Set("limit", strLimit)
	}

	prevQuery.Set("start", strconv.FormatUint(prevOffset, 10)) // start: -10, limit: 10
	selfQuery.Set("start", strconv.FormatUint(start, 10))      // start: 0, limit: 10
	nextQuery.Set("start", strconv.FormatUint(nextOffset, 10)) // start: 10, limit: 10
	prevURL.RawQuery = prevQuery.Encode()
	selfURL.RawQuery = selfQuery.Encode()
	nextURL.RawQuery = nextQuery.Encode()

	nextURLAddr := nextURL.String()
	var nextPage *string = nil
	if nextOffset < count {
		nextPage = pointer.ToString(environments.BaseURL + nextURLAddr)
	}

	prevURLAddr := prevURL.String()
	var prevPage *string = nil
	if prevOffset >= 0 && prevOffset < count {
		prevPage = pointer.ToString(environments.BaseURL + prevURLAddr)
	}

	pagination := &Pagination{
		Total: &count,
		Size:  &dataSize,
		Limit: &limit,
		Start: &start,
		Links: &Link{
			Previous: prevPage,
			Self:     pointer.ToString(environments.BaseURL + selfURL.String()),
			Next:     nextPage,
		},
	}

	return pagination
}
