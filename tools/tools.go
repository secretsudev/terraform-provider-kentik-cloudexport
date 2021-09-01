// +build tools

package tools

import (
	// document generation
	_ "github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs"
	_ "github.com/kentik/community_sdk_golang/apiv6/localhost_apiserver" // used for tests
)
