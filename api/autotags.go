package api

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

type autoTagsService service

// GetAll lists all configured auto tags
func (s *autoTagsService) GetAll() ([]AutoTag, *resty.Response, error) {

	autoTags := new(AutoTagResponse)

	apiResponse, err := s.client.Do("GET", "/api/config/v1/autoTags", nil, autoTags, nil)

	if err != nil {
		return nil, apiResponse, err
	}

	if apiResponse.StatusCode()/100 == 2 {
		return autoTags.Values, apiResponse, nil
	}

	return nil, apiResponse, StatusError(apiResponse.StatusCode())

}

// Create creates a new auto tag
func (s *autoTagsService) Create(autoTag AutoTag) (*AutoTag, *resty.Response, error) {
	autoTagResp := new(AutoTag)

	apiResponse, err := s.client.Do("POST", "/api/config/v1/autoTags", autoTag, autoTagResp, nil)

	if err != nil {
		return nil, apiResponse, err
	}

	if apiResponse.StatusCode()/100 == 2 {
		return autoTagResp, apiResponse, nil
	}

	return nil, apiResponse, StatusError(apiResponse.StatusCode())

}

// Get gets the properties of the specified auto tag
func (s *autoTagsService) Get(ID string, includeProcessGroupReferences bool) (*AutoTag, *resty.Response, error) {

	autoTag := new(AutoTag)

	path := fmt.Sprintf("/api/config/v1/autoTags/%s?includeProcessGroupReferences=%t", ID, includeProcessGroupReferences)

	apiResponse, err := s.client.Do("GET", path, nil, autoTag, nil)

	if err != nil {
		return nil, apiResponse, err
	}

	if apiResponse.StatusCode()/100 == 2 {
		return autoTag, apiResponse, nil
	}

	return nil, apiResponse, StatusError(apiResponse.StatusCode())

}

// Update updates an existing auto tag or creates a new one
func (s *autoTagsService) Update(ID string, autoTag AutoTag) (*AutoTag, *resty.Response, error) {
	autoTagResp := new(AutoTag)

	url := fmt.Sprintf("/api/config/v1/autoTags/%s", ID)
	apiResponse, err := s.client.Do("PUT", url, autoTag, autoTagResp, nil)

	if err != nil {
		return nil, apiResponse, err
	}

	if apiResponse.StatusCode() == 204 {
		return nil, apiResponse, nil
	}

	if apiResponse.StatusCode()/100 == 2 {
		return autoTagResp, apiResponse, nil
	}

	return nil, apiResponse, StatusError(apiResponse.StatusCode())

}

// Delete deletes the specified auto tag
func (s *autoTagsService) Delete(ID string) (*resty.Response, error) {

	url := fmt.Sprintf("/api/config/v1/autoTags/%s", ID)
	apiResponse, err := s.client.Do("DELETE", url, nil, nil, nil)

	if err != nil {
		return nil, err
	}

	if apiResponse.StatusCode()/100 == 2 {
		return apiResponse, nil
	}

	return apiResponse, StatusError(apiResponse.StatusCode())

}

// ValidateUpdate validates update of existing auto tags for the `PUT /autoTags/{id}` request
func (s *autoTagsService) ValidateUpdate(ID string, autoTag AutoTag) (*ErrorDetail, *resty.Response, error) {

	url := fmt.Sprintf("/api/config/v1/autoTags/%s/validator", ID)

	apiResponse, err := s.client.Do("POST", url, autoTag, nil, nil)

	if apiResponse.StatusCode() == 400 {
		return apiResponse.Error().(*ErrorResponse).Detail, apiResponse, err
	}

	if apiResponse.StatusCode()/100 == 2 {
		return nil, apiResponse, nil
	}

	if err != nil {
		return nil, nil, err
	}

	return nil, apiResponse, StatusError(apiResponse.StatusCode())

}

// ValidateCreate validates new auto tags for the `POST /autoTags` request
func (s *autoTagsService) ValidateCreate(autoTag AutoTag) (*ErrorDetail, *resty.Response, error) {

	apiResponse, err := s.client.Do("POST", "/api/config/v1/autoTags/", autoTag, nil, nil)

	if apiResponse.StatusCode() == 400 {
		return apiResponse.Error().(*ErrorResponse).Detail, apiResponse, err
	}

	if apiResponse.StatusCode()/100 == 2 {
		return nil, apiResponse, nil
	}

	if err != nil {
		return nil, nil, err
	}

	return nil, apiResponse, StatusError(apiResponse.StatusCode())

}
