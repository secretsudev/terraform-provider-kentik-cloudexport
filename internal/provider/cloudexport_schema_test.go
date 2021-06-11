package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestCloudExportSerializerAlwaysSetsTheEnabledField(t *testing.T) {
	test_cases := []struct{
		configuredInput interface{}
		expectedOutput bool
	}{
		{
			configuredInput: nil,
			expectedOutput: false,
		},
		{
			configuredInput: false,
			expectedOutput: false,
		},
		{
			configuredInput: true,
			expectedOutput: true,
		},
	}

	for _, tc := range test_cases {
		subtestName := fmt.Sprintf("Configured input: %v, expected output: %v", tc.configuredInput, tc.expectedOutput)
		t.Run(subtestName, func(t* testing.T) {
			// given
			d := makeDummyResourceData(t)
			if err := d.Set("enabled", tc.configuredInput); err != nil {
				t.Fatal(err)
			}

			// when
			export, err := resourceDataToCloudExport(d)

			// then
			if err != nil {
				t.Fatal(err)
			}

			if export.Enabled == nil {
				t.Fatal("Got unexpected output: nil")
			}

			if *export.Enabled != tc.expectedOutput {
				t.Errorf("Got unexpected output: %v", *export.Enabled)
			}
		})
	}
}

func makeDummyResourceData(t *testing.T) *schema.ResourceData {
	type ProviderDefinition = map[string]interface{} // groups provider's attributes
	const provider = "ibm"

	d := resourceCloudExport().Data(nil)
	if err := d.Set("cloud_provider", provider); err != nil {
		t.Fatal(err)
	}
	provider_definition := make(ProviderDefinition)
	provider_definition["bucket"] = "dummy"
	d.Set(provider, []ProviderDefinition{provider_definition})
	return d
}