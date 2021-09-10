package provider

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/kentik/community_sdk_golang/apiv6/kentikapi"
)

const (
	apiURLKey      = "apiurl"
	emailKey       = "email"
	tokenKey       = "token"
	logPayloadsKey = "log_payloads"
)

func init() {
	// Set descriptions to support Markdown syntax, this will be used in document generation and the language server.
	schema.DescriptionKind = schema.StringMarkdown
}

func New() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			apiURLKey: {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("KTAPI_URL", nil),
				Description: "CloudExport API server URL (optional). Can also be specified with KTAPI_URL environment variable.",
			},
			emailKey: {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("KTAPI_AUTH_EMAIL", nil),
				Description: "Authorization email (required). Can also be specified with KTAPI_AUTH_EMAIL environment variable.",
			},
			tokenKey: {
				Type:        schema.TypeString,
				Sensitive:   true,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("KTAPI_AUTH_TOKEN", nil),
				Description: "Authorization token (required). Can also be specified with KTAPI_AUTH_TOKEN environment variable.",
			},
			logPayloadsKey: {
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("KTAPI_LOG_PAYLOADS", false),
				Description: "Log payloads flag enables verbose debug logs of requests and responses (optional). " +
					"Can also be specified with KTAPI_LOG_PAYLOADS environment variable.",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"kentik-cloudexport_item": resourceCloudExport(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"kentik-cloudexport_list": dataSourceCloudExportList(),
			"kentik-cloudexport_item": dataSourceCloudExportItem(),
		},
		ConfigureContextFunc: configure,
	}
}

func configure(_ context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	cfg := kentikapi.Config{
		CloudExportAPIURL: getURL(d),
		AuthEmail:         d.Get(emailKey).(string),
		AuthToken:         d.Get(tokenKey).(string),
		LogPayloads:       d.Get(logPayloadsKey).(bool),
	}
	log.Printf("[DEBUG] Creating Kentik API client with config: %+v", stripSensitiveData(cfg))

	return kentikapi.NewClient(cfg), nil
}

func getURL(d *schema.ResourceData) string {
	var url string
	apiURL, ok := d.GetOk(apiURLKey)
	if ok {
		url = apiURL.(string)
	}
	return url
}

func stripSensitiveData(cfg kentikapi.Config) kentikapi.Config {
	cfg.AuthToken = "<stripped>"
	return cfg
}
