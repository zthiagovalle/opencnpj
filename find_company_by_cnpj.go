package opencnpj

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// FindByCNPJ finds a company by its CNPJ.
func (c *Client) FindCompanyByCNPJ(cnpj string) (*Company, error) {
	url := fmt.Sprintf(c.baseURL, cnpj)
	res, err := c.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer res.Body.Close()

	switch res.StatusCode {
	case http.StatusOK:
		break
	case http.StatusTooManyRequests:
		return nil, ErrTooManyRequests
	case http.StatusNotFound:
		return nil, ErrNotFound
	default:
		return nil, fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}

	var company Company
	if err := json.NewDecoder(res.Body).Decode(&company); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return &company, nil
}
