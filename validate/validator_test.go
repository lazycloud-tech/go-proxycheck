package validate

import (
	"errors"
	"reflect"
	"testing"
	"time"
)

func TestValidateDecodeResponse(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		data    []byte
		want    *APIResponse
		wantErr error
	}{
		{
			name:    "empty response",
			data:    []byte(`{}`),
			want:    &APIResponse{},
			wantErr: nil,
		},
		{
			name: "valid response",
			data: []byte(`{"status":"ok","node":"QUASAR","98.111.111.4":{"asn":"AS6389","range":"98.111.111.0/20","hostname":"adsl-098-075-002-004.sip.owb.bellsouth.net","provider":"Some Corp.","continent":"North America","continentcode":"NA","country":"United States","isocode":"US","timezone":"America/Chicago","latitude":37.751,"longitude":-97.822,"currency":{"code":"USD","name":"Dollar","symbol":"$"},"proxy":"no","type":"Residential","risk":0},"query time":"0.001s"}`),
			want: &APIResponse{
				Status:    statusOK,
				Node:      "QUASAR",
				QueryTime: "0.001s",
				Data: map[string]IPValidationResult{
					"98.111.111.4": {
						ASN:           "AS6389",
						Range:         "98.111.111.0/20",
						Hostname:      "adsl-098-075-002-004.sip.owb.bellsouth.net",
						Provider:      "Some Corp.",
						Continent:     "North America",
						ContinentCode: "NA",
						Country:       "United States",
						ISOCode:       "US",
						Timezone:      "America/Chicago",
						Latitude:      37.751,
						Longitude:     -97.822,
						Currency: Currency{
							Code:   "USD",
							Name:   "Dollar",
							Symbol: "$",
						},
						Proxy: "no",
						Type:  "Residential",
						Risk:  0,
					},
				},
			},
			wantErr: nil,
		},
		{
			name: "denied",
			data: []byte(`{"status":"denied"}`),
			want: &APIResponse{
				Status: statusDenied,
			},
			wantErr: nil,
		},
		{
			name: "error",
			data: []byte(`{"status":"error"}`),
			want: &APIResponse{
				Status: statusError,
			},
			wantErr: nil,
		},
		{
			name:    "invalid JSON",
			data:    []byte(`{"status":"ok",`),
			want:    nil,
			wantErr: ErrUnmarshallingResponse,
		},
		{
			name:    "unexpected data type",
			data:    []byte(`{"status":123}`),
			want:    nil,
			wantErr: ErrDecodingResponse,
		},
		{
			name:    "missing status field",
			data:    []byte(`{"node":"QUASAR"}`),
			want:    &APIResponse{Node: "QUASAR"},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := decodeResponse(tt.data)
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("decodeResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("decodeResponse() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewValidator(t *testing.T) {
	tests := []struct {
		name           string
		opts           IPValidationOptions
		wantTimeout    time.Duration
		wantAPIAddress string
	}{
		{
			name: "custom options",
			opts: IPValidationOptions{
				Timeout:    5 * time.Second,
				APIAddress: "http://example.com",
				VPN:        VPNOptionBoth,
				ASN:        ASNOptionActive,
				Currency:   CurrencyOptionActive,
				Node:       NodeOptionActive,
				Time:       TimeOptionActive,
				Risk:       RiskOptionFull,
				Port:       PortOptionActive,
				Seen:       SeenOptionActive,
				Days:       "30",
				Tag:        "test",
			},
			wantTimeout:    5 * time.Second,
			wantAPIAddress: "http://example.com",
		},
		{
			name: "default timeout",
			opts: IPValidationOptions{
				APIAddress: "http://example.com",
			},
			wantTimeout:    defaultClientTimeout,
			wantAPIAddress: "http://example.com",
		},
		{
			name: "default API address",
			opts: IPValidationOptions{
				Timeout: 5 * time.Second,
			},
			wantTimeout:    5 * time.Second,
			wantAPIAddress: apiV2Address,
		},
		{
			name:           "default options",
			opts:           IPValidationOptions{},
			wantTimeout:    defaultClientTimeout,
			wantAPIAddress: apiV2Address,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			validator := NewValidator(tt.opts)

			if validator.client == nil {
				t.Error("NewValidator should initialize the client")
			}

			if validator.client.Timeout != tt.wantTimeout {
				t.Errorf("NewValidator client timeout = %v, want %v", validator.client.Timeout, tt.wantTimeout)
			}

			if validator.APIAddress != tt.wantAPIAddress {
				t.Errorf("NewValidator APIAddress = %s, want %s", validator.APIAddress, tt.wantAPIAddress)
			}

			if validator.OptionsQuery != tt.opts.MakeQuery() {
				t.Errorf("NewValidator OptionsQuery = %s, want %s", validator.OptionsQuery, tt.opts.MakeQuery())
			}
		})
	}
}
