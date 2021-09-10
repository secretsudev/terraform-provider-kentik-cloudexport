package provider

import (
	"context"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/kentik/community_sdk_golang/apiv6/kentikapi"
)

func dataSourceCloudExportList() *schema.Resource {
	return &schema.Resource{
		Description: "DataSource representing list of cloud exports",
		ReadContext: dataSourceCloudExportListRead,
		Schema: map[string]*schema.Schema{
			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: makeCloudExportSchema(READ_LIST),
				},
			},
		},
	}
}

func dataSourceCloudExportListRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	listResp, httpResp, err := m.(*kentikapi.Client).CloudExportAdminServiceAPI.
		ExportList(ctx).
		Execute()
	if err != nil {
		return detailedDiagError("Failed to read cloud export list", err, httpResp)
	}

	if listResp.Exports != nil {
		numExports := len(*listResp.Exports)
		exports := make([]interface{}, numExports, numExports)
		for i, e := range *listResp.Exports {
			exports[i] = cloudExportToMap(&e)
		}

		if err = d.Set("items", exports); err != nil {
			return diag.FromErr(err)
		}
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10)) // use unixtime as id to force list update every time terraform asks for the list

	return nil
}
