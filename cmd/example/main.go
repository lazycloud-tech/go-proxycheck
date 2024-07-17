package main

import (
	"context"

	"log"

	"github.com/lazycloud-tech/go-proxycheck/validate"
)

func main() {
	// Single validator call for one-time operations.
	result, err := validate.CheckIPAddress(context.Background(),
		[]string{"8.8.8.8", "8.8.4.4"},
		validate.IPValidationOptions{
			VPN:      validate.VPNOptionBoth,
			ASN:      validate.ASNOptionActive,
			Currency: validate.CurrencyOptionActive,
			Node:     validate.NodeOptionActive,
			Time:     validate.TimeOptionActive,
			Risk:     validate.RiskOptionFull,
			Port:     validate.PortOptionActive,
			Seen:     validate.SeenOptionActive,
			Days:     "30",
			Tag:      "test",
		})
	if err != nil {
		panic(err)
	}

	log.Printf("%+v", result)

	// Multiple validator calls for multiple operations (using in middleware, for example).
	validator := validate.NewValidator(validate.IPValidationOptions{
		VPN:      validate.VPNOptionBoth,
		ASN:      validate.ASNOptionActive,
		Currency: validate.CurrencyOptionActive,
		Node:     validate.NodeOptionActive,
		Time:     validate.TimeOptionActive,
		Risk:     validate.RiskOptionFull,
		Port:     validate.PortOptionActive,
		Seen:     validate.SeenOptionActive,
		Days:     "30",
		Tag:      "test",
	})

	result, err = validator.CheckIPAddress(context.Background(), []string{"8.8.8.8"})
	if err != nil {
		panic(err)
	}

	log.Printf("%+v", result)

	result, err = validator.CheckIPAddress(context.Background(), []string{"8.8.4.4"})
	if err != nil {
		panic(err)
	}

	log.Printf("%+v", result)
}
