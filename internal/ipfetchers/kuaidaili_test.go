package ipfetchers

import (
	"gitlab.wizmacau.com/jack/proxypool/internal/configs"
	"testing"
)

func TestFetchIPsFromPrivateProxy(t *testing.T) {
	config, err := configs.NewConfigs()
	if err != nil {
		t.Fatalf("failed to load configs: %v", err)
	}

	ips, err := FetchIPsFromPrivateProxy(&config.KuaiDaiLi)

	if err != nil {
		t.Errorf("FetchIPsFromPrivateProxy() error = %v", err)
	}

	if len(ips) == 0 {
		t.Errorf("FetchIPsFromPrivateProxy() got = %v, want > 0", len(ips))
	}

	for _, ip := range ips {
		t.Logf("IP: %s, Source: %s", ip.Data, ip.Source)
	}
}
