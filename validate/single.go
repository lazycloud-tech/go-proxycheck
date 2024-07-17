package validate

import (
	"context"
)

// CheckIPAddress is a standalone function that creates new validator and sends a request to the proxycheck.io API.
//
// It's reasonable to use it you want to perform a single check (even for many addresses at once).
// If you want to perform multiple checks, it's better to create a new validator manually and use it for future checks.
func CheckIPAddress(ctx context.Context, ips []string, opts IPValidationOptions) (*APIResponse, error) {
	return NewValidator(opts).CheckIPAddress(ctx, ips)
}
