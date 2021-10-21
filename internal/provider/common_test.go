package provider

import (
	"net/http"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stretchr/testify/require"
)

func checkAPIServerConnection(t *testing.T) {
	apiURL, ok := os.LookupEnv("KTAPI_URL")
	require.True(t, ok, "KTAPI_URL env variable not set")

	_, err := http.Get(apiURL) //nolint: bodyclose, gosec, noctx
	require.NoErrorf(t, err, "failed to connect to the API Server on URL %q", apiURL)
}

// providerFactories are used to instantiate a provider during acceptance testing.
// The factory function will be invoked for every Terraform CLI command executed
// to create a provider server to which the CLI can reattach.
func providerFactories() map[string]func() (*schema.Provider, error) {
	return map[string]func() (*schema.Provider, error){
		"kentik-cloudexport": func() (*schema.Provider, error) {
			return New(), nil
		},
	}
}
