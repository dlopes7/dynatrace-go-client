package api

import "fmt"

// Get all application naming rules in order
func (c *Client) AllApplicationNameDetectionRules() ([]NameDetectionRule, error) {
	resp := NameDetectionResponse{}

	apiResponse, err := c.Do("GET", "/api/config/v1/applicationDetectionRules", nil, &resp, nil)

	if err != nil {
		return resp.Values, err
	}

	if apiResponse.StatusCode()/100 == 2 {
		return resp.Values, err
	}

	return resp.Values, StatusError(apiResponse.StatusCode())
}

func (c *Client) UpdateApplicationNameDetectionRuleOrder(ids []string) error {
	rules := make([]NameDetectionRule, len(ids))

	for i, id := range ids {
		rules[i] = NameDetectionRule{
			Id: id,
		}
	}

	body := NameDetectionRuleOrderRequest{
		Values: rules,
	}

	apiResponse, err := c.Do("PUT", "/api/config/v1/applicationDetectionRules/order", body, nil, nil)

	if err != nil {
		return err
	}

	if apiResponse.StatusCode()/100 != 2 {
		return StatusError(apiResponse.StatusCode())
	}

	return nil
}

// Get an application name detection rule. If the rule does not exist, an empty NameDetectionRuleDetail
// will be returned with an Id of ""
func (c *Client) GetApplicationNameDetectionRule(id string) (NameDetectionRuleDetail, error) {
	resp := NameDetectionRuleDetail{}

	if id == "" {
		return resp, fmt.Errorf("Empty string is not a valid id")
	}

	apiResponse, err := c.Do("GET", fmt.Sprintf("/api/config/v1/applicationDetectionRules/%s", id), nil, &resp, nil)

	if err != nil {
		return resp, err
	}

	if apiResponse.StatusCode()/100 == 2 {
		return resp, nil
	}

	if apiResponse.StatusCode() == 404 {
		return resp, nil
	}

	return resp, StatusError(apiResponse.StatusCode())
}

// Delete an application name detection rule.
func (c *Client) DeleteApplicationNameDetectionRule(id string) error {
	if id == "" {
		return fmt.Errorf("Empty string is not a valid id")
	}

	_, err := c.Do("DELETE", fmt.Sprintf("/api/config/v1/applicationDetectionRules/%s", id), nil, nil, nil)

	return err
}

// Create an application name detection rule. If the API responds with a non-2xx status code, an error is returned.
func (c *Client) CreateApplicationNameDetectionRule(body NameDetectionRuleDetail) (NameDetectionRule, error) {
	resp := NameDetectionRule{}

	apiResponse, err := c.Do("POST", "/api/config/v1/applicationDetectionRules", body, &resp, nil)

	if err != nil {
		return resp, err
	}

	if apiResponse.StatusCode()/100 == 2 {
		return resp, err
	}

	return resp, StatusError(apiResponse.StatusCode())
}

// Update an application name detection rule. If the API responds with a non-2xx status code, an error is returned.
func (c *Client) UpdateApplicationNameDetectionRule(id string, body NameDetectionRuleDetail) error {
	if id == "" {
		return fmt.Errorf("Empty string is not a valid id")
	}

	apiResponse, err := c.Do("PUT", fmt.Sprintf("/api/config/v1/applicationDetectionRules/%s", id), body, nil, nil)

	if err != nil {
		return err
	}

	if apiResponse.StatusCode()/100 == 2 {
		return nil
	}

	return StatusError(apiResponse.StatusCode())
}
