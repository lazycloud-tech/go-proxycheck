# go-proxycheck

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)]

go-proxucheck is a simple library to check general IP address information and verify if it is a proxy or VPN. It's also possible to check email address for signs of being a disposable.

The lib uses the [proxycheck.io](https://proxycheck.io) API to check the IP address and email address. You can use free tier or perform a registration to get a free API key and enjoy bigger limits.

In plans creationg of a simple web server to perform the checks and return results in XML or JSON.

## Usage

Single request for occasional checks. Useful when your app needs to check some info about IP once in a while.

```
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
```

A validator instance for regurlar checks (e.g. in a web server). Useful when you need to check multiple IPs in a short period of time.

```
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
```