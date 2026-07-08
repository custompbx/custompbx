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

func TestNormalizeAndValidateWebSocketSettings(t *testing.T) {
	tests := []struct {
		name      string
		cfg       WebServer
		wantWrite int
		wantRead  int
		wantPing  int
		wantQueue int
	}{
		{
			name:      "legacy defaults",
			cfg:       WebServer{},
			wantWrite: DefaultWSWriteTimeoutSeconds,
			wantRead:  DefaultWSReadTimeoutSeconds,
			wantPing:  DefaultWSPingIntervalSeconds,
			wantQueue: DefaultWebSocketQueueSize,
		},
		{
			name:      "configured values",
			cfg:       WebServer{WriteTimeoutSeconds: 5, ReadTimeoutSeconds: 20, PingIntervalSeconds: 10, WebSocketQueueSize: 128},
			wantWrite: 5,
			wantRead:  20,
			wantPing:  10,
			wantQueue: 128,
		},
		{
			name:      "invalid values normalize safely",
			cfg:       WebServer{WriteTimeoutSeconds: -1, ReadTimeoutSeconds: 1, PingIntervalSeconds: -1, WebSocketQueueSize: -1},
			wantWrite: DefaultWSWriteTimeoutSeconds,
			wantRead:  DefaultWSReadTimeoutSeconds,
			wantPing:  DefaultWSPingIntervalSeconds,
			wantQueue: DefaultWebSocketQueueSize,
		},
		{
			name:      "ping not lower than read is adjusted",
			cfg:       WebServer{WriteTimeoutSeconds: 1, ReadTimeoutSeconds: 10, PingIntervalSeconds: 10, WebSocketQueueSize: 1},
			wantWrite: 1,
			wantRead:  10,
			wantPing:  5,
			wantQueue: 1,
		},
		{
			name:      "queue clamps to maximum",
			cfg:       WebServer{WebSocketQueueSize: MaxWebSocketQueueSize + 1},
			wantWrite: DefaultWSWriteTimeoutSeconds,
			wantRead:  DefaultWSReadTimeoutSeconds,
			wantPing:  DefaultWSPingIntervalSeconds,
			wantQueue: MaxWebSocketQueueSize,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.cfg.NormalizeAndValidateOrigins(); err != nil {
				t.Fatal(err)
			}
			if tt.cfg.WriteTimeoutSeconds != tt.wantWrite || tt.cfg.ReadTimeoutSeconds != tt.wantRead || tt.cfg.PingIntervalSeconds != tt.wantPing || tt.cfg.WebSocketQueueSize != tt.wantQueue {
				t.Fatalf("got write=%d read=%d ping=%d queue=%d", tt.cfg.WriteTimeoutSeconds, tt.cfg.ReadTimeoutSeconds, tt.cfg.PingIntervalSeconds, tt.cfg.WebSocketQueueSize)
			}
		})
	}
}

func TestCreateConfigIncludesWebSocketDefaults(t *testing.T) {
	conf := createConfig()
	if conf.Web.WriteTimeoutSeconds != DefaultWSWriteTimeoutSeconds ||
		conf.Web.ReadTimeoutSeconds != DefaultWSReadTimeoutSeconds ||
		conf.Web.PingIntervalSeconds != DefaultWSPingIntervalSeconds ||
		conf.Web.WebSocketQueueSize != DefaultWebSocketQueueSize {
		t.Fatalf("unexpected websocket defaults: %+v", conf.Web)
	}
}
