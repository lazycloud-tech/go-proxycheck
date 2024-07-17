package validate

import "testing"

func TestMakeQuery(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		opts IPValidationOptions
		want string
	}{
		{
			name: "empty options",
			opts: IPValidationOptions{},
			want: "p=0",
		},
		{
			name: "all options",
			opts: IPValidationOptions{
				VPN:      VPNOptionBoth,
				ASN:      ASNOptionActive,
				Currency: CurrencyOptionActive,
				Node:     NodeOptionActive,
				Time:     TimeOptionActive,
				Risk:     RiskOptionFull,
				Port:     PortOptionActive,
				Seen:     SeenOptionActive,
				Days:     "30",
				Tag:      "test",
			},
			want: "asn=1&cur=1&days=30&node=1&p=0&port=1&risk=2&seen=1&tag=test&time=1&vpn=3",
		}, {
			name: "VPN 0",
			opts: IPValidationOptions{
				VPN: VPNOptionProxyOnly,
			},
			want: "p=0&vpn=0",
		},
		{
			name: "VPN 1",
			opts: IPValidationOptions{
				VPN: VPNOptionAny,
			},
			want: "p=0&vpn=1",
		},
		{
			name: "VPN 2",
			opts: IPValidationOptions{
				VPN: VPNOptionVPNOnly,
			},
			want: "p=0&vpn=2",
		},
		{
			name: "VPN 3",
			opts: IPValidationOptions{
				VPN: VPNOptionBoth,
			},
			want: "p=0&vpn=3",
		},
		{
			name: "ASN active",
			opts: IPValidationOptions{
				ASN: ASNOptionActive,
			},
			want: "asn=1&p=0",
		},
		{
			name: "ASN inactive",
			opts: IPValidationOptions{
				ASN: ASNOptionInactive,
			},
			want: "asn=0&p=0",
		},
		{
			name: "currency active",
			opts: IPValidationOptions{
				Currency: CurrencyOptionActive,
			},
			want: "cur=1&p=0",
		},
		{
			name: "currency inactive",
			opts: IPValidationOptions{
				Currency: CurrencyOptionInactive,
			},
			want: "cur=0&p=0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.opts.MakeQuery()
			if got != tt.want {
				t.Errorf("MakeQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}
