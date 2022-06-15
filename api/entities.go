package api

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

type entitiesService service

func (e *entitiesService) Get(entityId string) (*EnvV2Entity, *resty.Response, error) {
	entity := new(EnvV2Entity)

	uri := fmt.Sprintf("/api/v2/entities/%s", entityId)
	apiResponse, err := e.client.Do("GET", uri, nil, entity, nil)

	if err != nil {
		return nil, apiResponse, err
	}

	if apiResponse.StatusCode()/100 == 2 {
		return entity, apiResponse, nil
	}

	return nil, apiResponse, StatusError(apiResponse.StatusCode())

}
