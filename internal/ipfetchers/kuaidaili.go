package ipfetchers

import (
	"gitlab.wizmacau.com/jack/proxypool/internal/configs"
	"gitlab.wizmacau.com/jack/proxypool/internal/models"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// FetchIPsFromPrivateProxy fetches IP addresses from the private proxy service 快代理.
// It takes a pointer to a KuaiDaiLi configuration object as input and returns a slice of IP addresses and any error encountered.
// The function constructs the API request URL using the configuration details, sends a GET request to the API, and processes the response.
// It reads the response body, splits it by newline characters to get individual IP addresses, and appends them to the IPs slice.
// Each IP address is wrapped in an IP model object with the source set as "快代理-私密代理".
// The function returns the IPs slice and any error encountered.
func FetchIPsFromPrivateProxy(cfg *configs.KuaiDaiLi) ([]*models.IP, error) {
	api := "https://dps.kdlapi.com/api/getdps"
	params := url.Values{}
	params.Set("secret_id", cfg.SecretId)
	params.Set("signature", cfg.SecretKey)
	params.Set("num", strconv.Itoa(cfg.Num))

	fullApi := api + "?" + params.Encode()

	response, err := http.Get(fullApi)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(response.Body)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	r := strings.Split(string(body), "\n")
	var ips []*models.IP
	for _, ip := range r {
		ips = append(
			ips, &models.IP{
				Data:   ip,
				Source: "快代理-私密代理",
			},
		)
	}

	return ips, nil
}
