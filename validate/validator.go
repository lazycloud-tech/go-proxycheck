package validate

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/mitchellh/mapstructure"
)

type Validator struct {
	client       *http.Client
	APIAddress   string
	OptionsQuery string
}

func NewValidator(opts IPValidationOptions) *Validator {
	timeout := opts.Timeout
	if timeout == 0 {
		timeout = defaultClientTimeout
	}

	apiAddress := opts.APIAddress
	if apiAddress == "" {
		apiAddress = apiV2Address
	}

	return &Validator{
		client:       &http.Client{Timeout: timeout, Transport: http.DefaultTransport},
		APIAddress:   apiAddress,
		OptionsQuery: opts.MakeQuery(),
	}
}

func (v *Validator) CheckIPAddress(ctx context.Context, ips []string) (*APIResponse, error) {
	if len(ips) == 0 {
		return nil, ErrEmptyValues
	}

	data, err := v.MakeRequest(ctx, ips)
	if err != nil {
		return nil, err
	}

	resp, err := decodeResponse(data)
	if err != nil {
		return nil, err
	}

	switch resp.Status {
	case statusOK:
		if len(resp.Data) == 0 {
			return nil, ErrNoValidationDataFound
		}

		return resp, nil
	case statusDenied:
		return nil, fmt.Errorf("%w: %s", ErrRequestDenied, resp.Message)
	case statusError:
		return nil, fmt.Errorf("%w: %s", ErrRequestError, resp.Message)
	default:
		return nil, fmt.Errorf("%w: %s", ErrUnexpectedStatus, resp.Status)
	}
}

func (v *Validator) MakeRequest(ctx context.Context, ips []string) ([]byte, error) {
	url := fmt.Sprintf("%s%s?%s", v.APIAddress, strings.Join(ips, ","), v.OptionsQuery)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrPreparingRequest, err.Error())
	}

	res, err := v.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrSendingRequest, err.Error())
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrReadingResponseBody, err.Error())
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%w: %d", ErrBadHTTPStatusCode, res.StatusCode)
	}

	return data, nil
}

func decodeResponse(data []byte) (*APIResponse, error) {
	var result map[string]interface{}

	if err := json.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("%w: %s", ErrUnmarshallingResponse, err.Error())
	}

	var resp APIResponse
	if err := mapstructure.Decode(result, &resp); err != nil {
		return nil, fmt.Errorf("%w: %s", ErrDecodingResponse, err.Error())
	}

	return &resp, nil
}
