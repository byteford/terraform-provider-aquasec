package aquasec

import (
	"log"

	"github.com/aquasecurity/terraform-provider-aquasec/client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceRegistry() *schema.Resource {
	return &schema.Resource{
		Read: dataRegistryRead,
		Schema: map[string]*schema.Schema{
			"username": {
				Type:        schema.TypeString,
				Description: "The username for registry authentication.",
				Computed:    true,
			},
			"password": {
				Type:        schema.TypeString,
				Description: "The password for registry authentication",
				Computed:    true,
			},
			"type": {
				Type:        schema.TypeString,
				Description: "Registry type (HUB / V1 / V2 / ENGINE / AWS / GCR).",
				Computed:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Description: "The name of the registry; string, required - this will be treated as the registry's ID, so choose a simple alphanumerical name without special signs and spaces",
				Required:    true,
			},
			"url": {
				Type:        schema.TypeString,
				Description: "The URL, address or region of the registry",
				Computed:    true,
			},
			"auto_pull": {
				Type:        schema.TypeBool,
				Description: "Whether to automatically pull images from the registry on creation and daily",
				Computed:    true,
			},
			"auto_pull_max": {
				Type:        schema.TypeInt,
				Description: "Maximum number of repositories to pull every day, defaults to 100",
				Computed:    true,
			},
			"auto_pull_time": {
				Type:        schema.TypeString,
				Description: "The time of day to start pulling new images from the registry, in the format HH:MM (24-hour clock), defaults to 03:00",
				Computed:    true,
			},
			"scanner_type": {
				Type:        schema.TypeString,
				Description: "Scanner type",
				Optional:    true,
			},
			"prefixes": {
				Type:        schema.TypeList,
				Description: "List of possible prefixes to image names pulled from the registry",
				Computed:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataRegistryRead(d *schema.ResourceData, m interface{}) error {
	log.Println("[DEBUG]  inside dataRegistryRead")
	ac := m.(*client.Client)
	name := d.Get("name").(string)
	reg, err := ac.GetRegistry(name)
	if err == nil {
		prefixes := d.Get("prefixes").([]interface{})
		d.Set("username", reg.Username)
		d.Set("password", reg.Password)
		d.Set("name", reg.Name)
		d.Set("type", reg.Type)
		d.Set("url", reg.URL)
		d.Set("auto_pull", reg.AutoPull)
		d.Set("auto_pull_max", reg.AutoPullMax)
		d.Set("auto_pull_time", reg.AutoPullTime)
		d.Set("scanner_type", reg.ScannerType)
		d.Set("prefixes", convertStringArr(prefixes))
		d.SetId(name)
	} else {
		return err
	}

	return nil
}
