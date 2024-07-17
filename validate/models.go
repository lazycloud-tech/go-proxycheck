package validate

import (
	"errors"
	"net/url"
	"time"
)

const (
	apiV2Address         = "https://proxycheck.io/v2/"
	defaultClientTimeout = 30 * time.Second

	statusOK     = "ok"
	statusDenied = "denied"
	statusError  = "error"

	VPNOptionProxyOnly = "0"
	VPNOptionAny       = "1"
	VPNOptionVPNOnly   = "2"
	VPNOptionBoth      = "3"

	ASNOptionInactive = "0"
	ASNOptionActive   = "1"

	CurrencyOptionInactive = "0"
	CurrencyOptionActive   = "1"

	NodeOptionInactive = "0"
	NodeOptionActive   = "1"

	TimeOptionInactive = "0"
	TimeOptionActive   = "1"

	RiskOptionIncative  = "0"
	RiskOptionScoreOnly = "1"
	RiskOptionFull      = "2"

	PortOptionInactive = "0"
	PortOptionActive   = "1"

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
)

type APIResponse struct {
	Status    string                        `mapstructure:"status"`
	Message   string                        `mapstructure:"message"`
	QueryTime string                        `mapstructure:"query time"`
	Node      string                        `mapstructure:"node"`
	Data      map[string]IPValidationResult `mapstructure:",remain"`
}

// Options is the struct that holds the options for the validation.
// Full list of flags can be found at https://proxycheck.io/api/#query_flags.
type IPValidationOptions struct {
	APIAddress string // Keep empty to use default Proxycheck address. Replace in case you need to mock the API.
	Timeout    time.Duration
	APIKey     string
	VPN        string
	ASN        string
	Currency   string
	Node       string
	Time       string
	Risk       string
	Port       string
	Seen       string
	Days       string
	Tag        string
}

func (opts *IPValidationOptions) MakeQuery() string {
	payload := url.Values{}
	// p is always 0 to not waste resources - we don't need to prettify the output.
	payload.Add("p", "0")

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
	// Error will be filled if there is an error fot this specific entity. All other fields will be empty.
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
	// Error will be filled if there is an error fot this specific entity. All other fields will be empty.
	Error      string `json:"error"`
	Disposable string `json:"disposable"`
}
