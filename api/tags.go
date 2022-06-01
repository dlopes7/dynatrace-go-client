package api

import (
	"github.com/go-resty/resty/v2"
)

type tagService service

type Tags struct {
	Tags []Tag `json:"tags"`
}
type Tag struct {
	Key   string `json:"key"`
	Value string `json:"value,omitempty"`
}

type TagResponse struct {
	AppliedTags          []AppliedTag `json:"appliedTags"`
	MatchedEntitiesCount int          `json:"matchedEntitiesCount"`
}
type AppliedTag struct {
	Context              string `json:"context"`
	Key                  string `json:"key"`
	StringRepresentation string `json:"stringRepresentation"`
}

//
// Create adds a tag to the specified entities
func (t *tagService) Create(entitySelector string, tags []Tag) (*TagResponse, *resty.Response, error) {
	tagRequest := new(Tags)
	tagRequest.Tags = tags

	tagResponse := new(TagResponse)

	params := map[string]string{
		"entitySelector": entitySelector,
	}

	apiResponse, err := t.client.Do("POST", "/api/v2/tags", tagRequest, tagResponse, params)

	if err != nil {
		return nil, apiResponse, err
	}

	if apiResponse.StatusCode()/100 == 2 {
		return tagResponse, apiResponse, nil
	}

	return nil, apiResponse, StatusError(apiResponse.StatusCode())

}
