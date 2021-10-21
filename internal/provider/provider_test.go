package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/stretchr/testify/assert"
)

func TestProvider(t *testing.T) {
	t.Parallel()
	err := New().InternalValidate()
	assert.NoError(t, err)
}

func TestProvider_Configure_MinimalConfig(t *testing.T) {
	t.Parallel()
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { checkAPIServerConnection(t) },
		ProviderFactories: providerFactories(),
		Steps: []resource.TestStep{{
			Config: minimalConfig,
		}},
	})
}

func TestProvider_Configure_CustomRetryConfig(t *testing.T) {
	t.Parallel()
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { checkAPIServerConnection(t) },
		ProviderFactories: providerFactories(),
		Steps: []resource.TestStep{{
			Config: customRetryConfig,
		}},
	})
}

func TestProvider_Configure_InvalidRetryConfig(t *testing.T) {
	t.Parallel()
	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { checkAPIServerConnection(t) },
		ProviderFactories: providerFactories(),
		Steps: []resource.TestStep{{
			Config:      invalidRetryConfig,
			ExpectError: regexp.MustCompile("parse max_delay duration"),
		}},
	})
}

const (
	minimalConfig = `
		provider "kentik-cloudexport" {}
		
		// Trigger arbitrary action
		data "kentik-cloudexport_item" "aws" {
			id = "1"
		}
	`
	customRetryConfig = `
		provider "kentik-cloudexport" {
			retry {
				max_attempts = 66
				min_delay = "100ms"
				max_delay = "1m"
			}
			log_payloads = true
		}
		
		// Trigger arbitrary action
		data "kentik-cloudexport_item" "aws" {
			id = "1"
		}
	`
	invalidRetryConfig = `
		provider "kentik-cloudexport" {
			retry {
				max_attempts = 66
				min_delay = "100ms"
				max_delay = "invalid-delay"
			}
			log_payloads = true
		}
		
		// Trigger arbitrary action
		data "kentik-cloudexport_item" "aws" {
			id = "1"
		}
	`
)
