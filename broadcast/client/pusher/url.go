package pusher

import (
	"fmt"
	"net/url"
)

type Options struct {
	Cluster       string
	Client        string
	ClientVersion string
	APIVersion    string
	Protocol      string
	AppId         string
}

func NewURL(options Options) *url.URL {
	u := &url.URL{}
	u.Host = getHost(options)
	u.Scheme = options.Protocol

	q := u.Query()
	q.Set("client", options.Client)
	q.Set("protocol", options.APIVersion)
	q.Set("version", options.ClientVersion)

	u.RawQuery = q.Encode()
	u.Path = fmt.Sprintf("app/%v", options.AppId)

	return u
}

func getHost(options Options) string {
	port := 80

	if options.Protocol == "wss" {
		port = 443
	}

	return fmt.Sprintf("ws-%v.pusher.com:%v", options.Cluster, port)
}
