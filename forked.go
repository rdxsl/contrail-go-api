package contrail

import "net/http"

// ----------------------------------------------------------------------------
// NOTE(jehwang): This files contains any RDXSL custom logic that
// we had to add in our fork.
// ----------------------------------------------------------------------------

// NewClientWithHTTP allocates and initializes a Contrail API client.
// Exposed new method to allow us to specify an http client with a proper timeout.
func NewClientWithHTTP(server string, port int, httpClient *http.Client) *Client {
	client := NewClient(server, port)
	if httpClient != nil {
		client.httpClient = httpClient
	}
	return client
}
