package toolbox

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/coreos/go-oidc/v3/oidc"
)

type forwardedTransport struct {
	base        http.RoundTripper
	internalURL *url.URL
	externalURL *url.URL
}

func (ft *forwardedTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if strings.Contains(req.URL.Host, ft.externalURL.Host) {
		req.URL.Scheme = ft.internalURL.Scheme
		req.URL.Host = ft.internalURL.Host
	}
	// Add Forwarded header
	req.Header.Add("Forwarded", fmt.Sprintf("host=%s;proto=%s", ft.externalURL.Host, ft.externalURL.Scheme))
	return ft.base.RoundTrip(req)
}

// ContextWithForwarded returns a context that allows http client to use a Forwarded header
func ContextWithForwarded(ctx context.Context, internalURL *url.URL, externalURL *url.URL) context.Context {
	if internalURL == nil || externalURL == nil {
		return ctx
	}
	return oidc.ClientContext(ctx, &http.Client{
		Transport: &forwardedTransport{
			base:        http.DefaultTransport,
			internalURL: internalURL,
			externalURL: externalURL,
		},
	})

}
