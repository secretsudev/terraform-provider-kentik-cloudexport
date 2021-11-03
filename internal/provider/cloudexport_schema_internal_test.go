package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func TestCloudExportSerializerAlwaysSetsTheEnabledField(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		configuredInput interface{}
		expectedOutput  bool
	}{
		{
			configuredInput: nil,
			expectedOutput:  false,
		},
		{
			configuredInput: false,
			expectedOutput:  false,
		},
		{
			configuredInput: true,
			expectedOutput:  true,
		},
	}

	for _, tc := range testCases {
		subtestName := fmt.Sprintf("Configured input: %v, expected output: %v", tc.configuredInput, tc.expectedOutput)
		t.Run(subtestName, func(t *testing.T) {
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

	pd := make(ProviderDefinition)
	pd["bucket"] = "dummy"
	if err := d.Set(provider, []ProviderDefinition{pd}); err != nil {
		t.Fatal(err)
	}

	return d
}
