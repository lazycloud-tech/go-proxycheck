package validate

import (
	"errors"
	"net/url"
	"time"
)

const (
	apiV2Address         = "https://proxycheck.io/v2/"
	defaultClientTimeout = 30 * time.Second
)

const (
	statusOK     = "ok"
	statusDenied = "denied"
	statusError  = "error"
)

const (
	VPNOptionProxyOnly = "0"
	VPNOptionAny       = "1"
	VPNOptionVPNOnly   = "2"
	VPNOptionBoth      = "3"
)

const (
	ASNOptionInactive = "0"
	ASNOptionActive   = "1"
)

const (
	CurrencyOptionInactive = "0"
	CurrencyOptionActive   = "1"
)

const (
	NodeOptionInactive = "0"
	NodeOptionActive   = "1"
)

const (
	TimeOptionInactive = "0"
	TimeOptionActive   = "1"
)

const (
	RiskOptionInactive  = "0"
	RiskOptionScoreOnly = "1"
	RiskOptionFull      = "2"
)

const (
	PortOptionInactive = "0"
	PortOptionActive   = "1"
)

const (
	SeenOptionInactive = "0"
	SeenOptionActive   = "1"
)

var (
	ErrEmptyValues           = errors.New("empty values")
	ErrNoValidationDataFound = errors.New("no validation data found in response")
	ErrDecodingResponse      = errors.New("error decoding response")
	ErrUnmarshallingResponse = errors.New("error unmarshalling response")
	ErrRequestDenied         = errors.New("request denied")
	ErrRequestError          = errors.New("request error")
	ErrBadHTTPStatusCode     = errors.New("unexpected status code")
	ErrReadingResponseBody   = errors.New("error reading response body")
	ErrPreparingRequest      = errors.New("error preparing request")
	ErrSendingRequest        = errors.New("error sending request")
	ErrUnexpectedStatus      = errors.New("unexpected status")
)

type APIResponse struct {
	Status    string                        `mapstructure:"status"`
	Message   string                        `mapstructure:"message"`
	QueryTime string                        `mapstructure:"query time"`
	Node      string                        `mapstructure:"node"`
	Data      map[string]IPValidationResult `mapstructure:",remain"`
}

// IPValidationOptions holds the options for the validation.
// Full list of flags can be found at https://proxycheck.io/api/#query_flags.
type IPValidationOptions struct {
	APIAddress string        // Keep empty to use default Proxycheck address. Replace in case you need to mock the API.
	Timeout    time.Duration // Timeout for the HTTP client.
	APIKey     string        // API key for authentication.
	VPN        string        // VPN option.
	ASN        string        // ASN option.
	Currency   string        // Currency option.
	Node       string        // Node option.
	Time       string        // Time option.
	Risk       string        // Risk option.
	Port       string        // Port option.
	Seen       string        // Seen option.
	Days       string        // Days option.
	Tag        string        // Tag option.
}

// MakeQuery constructs the query string from the IPValidationOptions.
func (opts *IPValidationOptions) MakeQuery() string {
	payload := url.Values{}
	payload.Add("p", "0") // p is always 0 to not waste resources - we don't need to prettify the output.

	if opts.APIKey != "" {
		payload.Add("key", opts.APIKey)
	}

	if opts.VPN != "" {
		payload.Add("vpn", opts.VPN)
	}

	if opts.ASN != "" {
		payload.Add("asn", opts.ASN)
	}

	if opts.Currency != "" {
		payload.Add("cur", opts.Currency)
	}

	if opts.Node != "" {
		payload.Add("node", opts.Node)
	}

	if opts.Time != "" {
		payload.Add("time", opts.Time)
	}

	if opts.Risk != "" {
		payload.Add("risk", opts.Risk)
	}

	if opts.Port != "" {
		payload.Add("port", opts.Port)
	}

	if opts.Seen != "" {
		payload.Add("seen", opts.Seen)
	}

	if opts.Days != "" {
		payload.Add("days", opts.Days)
	}

	if opts.Tag != "" {
		payload.Add("tag", opts.Tag)
	}

	return payload.Encode()
}

type IPValidationResult struct {
	// Error will be filled if there is an error for this specific entity.
	// All other fields will be empty.
	Error         string        `json:"error"`
	ASN           string        `json:"asn"`
	Range         string        `json:"range"`
	Hostname      string        `json:"hostname"`
	Provider      string        `json:"provider"`
	Organisation  string        `json:"organisation"`
	Continent     string        `json:"continent"`
	ContinentCode string        `json:"continentcode"`
	Country       string        `json:"country"`
	ISOCode       string        `json:"isocode"`
	Region        string        `json:"region"`
	RegionCode    string        `json:"regioncode"`
	Timezone      string        `json:"timezone"`
	City          string        `json:"city"`
	Postcode      string        `json:"postcode"`
	Latitude      float64       `json:"latitude"`
	Longitude     float64       `json:"longitude"`
	Currency      Currency      `json:"currency"`
	Proxy         string        `json:"proxy"`
	Type          string        `json:"type"`
	Risk          float64       `json:"risk"`
	AttackHistory AttackHistory `json:"attack history"`
	LastSeenHuman string        `json:"last seen human"`
	LastSeenUnix  string        `json:"last seen unix"`
}

type Currency struct {
	Code   string `json:"code"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

type AttackHistory struct {
	Total                string `json:"Total"`
	VulnerabilityProbing string `json:"Vulnerability Probing"`
	ForumSpam            string `json:"Forum Spam"`
	LoginAttempt         string `json:"Login Attempt"`
	RegistrationAttempt  string `json:"Registration Attempt"`
	CommentSpam          string `json:"Comment Spam"`
	DenialofService      string `json:"Denial of Service"`
	FormSubmission       string `json:"Form Submission"`
}

type EmailValidationResult struct {
	// Error will be filled if there is an error for this specific entity.
	// All other fields will be empty.
	Error      string `json:"error"`
	Disposable string `json:"disposable"`
}
