package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestCloudExportSerializerAlwaysSetsEnabledField(t *testing.T) {
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
		// given
		d := makeDummyResourceData(t)
		if err := d.Set("enabled", tc.configuredInput); err != nil {
			t.Fatal(err)
		}

		// when
		export, err := resourceDataToCloudExport(d)

		// then
		if err != nil {
			t.Error(err)
			continue
		}
		if export.Enabled == nil {
			t.Errorf("Expected export.Enabled != nil, got nil for input %v", tc.configuredInput)
			continue
		}

		if *export.Enabled != tc.expectedOutput {
			t.Errorf("Expected *export.Enabled == %v, got %v for input %v", tc.expectedOutput, *export.Enabled, tc.configuredInput)
		}
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