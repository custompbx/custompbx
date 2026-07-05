package cfg

import "testing"

func TestNormalizeAndValidateOrigins(t *testing.T) {
	tests := []struct {
		name string
		cfg  WebServer
		want string
		err  bool
	}{
		{name: "legacy defaults secure", cfg: WebServer{}, want: OriginPolicySameOrigin},
		{name: "development allow all", cfg: WebServer{OriginPolicy: OriginPolicyAllowAll}, want: OriginPolicyAllowAll},
		{name: "allow list", cfg: WebServer{OriginPolicy: OriginPolicyAllowList, AllowedOrigins: []string{"HTTPS://Example.COM/"}}, want: OriginPolicyAllowList},
		{name: "empty allow list", cfg: WebServer{OriginPolicy: OriginPolicyAllowList}, err: true},
		{name: "invalid policy", cfg: WebServer{OriginPolicy: "anything"}, err: true},
		{name: "origin with path", cfg: WebServer{OriginPolicy: OriginPolicyAllowList, AllowedOrigins: []string{"https://example.com/path"}}, err: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.cfg.NormalizeAndValidateOrigins()
			if (err != nil) != tt.err {
				t.Fatalf("error = %v", err)
			}
			if !tt.err && tt.cfg.OriginPolicy != tt.want {
				t.Fatalf("policy = %q, want %q", tt.cfg.OriginPolicy, tt.want)
			}
		})
	}
}
