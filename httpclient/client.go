package httpclient

import (
	"net/http"

	"github.com/go-resty/resty/v2"
)

func NewRestyWithClient(baseURL string, client *http.Client) *resty.Client {
	return resty.NewWithClient(client).SetBaseURL(baseURL)
}
