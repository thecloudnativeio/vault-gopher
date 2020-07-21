package client

import (
	"crypto/tls"
	"crypto/x509"
	"github.com/trx35479/vault-gopher/pkg/log"
	"net/http"
	"time"
)

type Client struct {
	HttpClient *http.Client
}

var logger = log.NewLogger()

// Https config for the client + the ca cert
func (c *Client) Https(cacrt []byte) *http.Client {
	caCertPool := x509.NewCertPool()
	if ok := caCertPool.AppendCertsFromPEM(cacrt); !ok {
		logger.Println("could not decode ca.cert")
	}
	c.HttpClient = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs: caCertPool,
			},
		},
		Timeout: 2 * time.Second,
	}
	return c.HttpClient
}

// Http config for the client
// Set the insecure to true if the endpoint is using self signed tls
func (c *Client) Http(insecure bool) *http.Client {
	c.HttpClient = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: insecure,
			},
		},
		Timeout: 2 * time.Second,
	}
	return c.HttpClient
}