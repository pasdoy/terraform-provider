package main

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"apikey": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CLOUDKARAFKA_APIKEY", nil),
				Description: "Key used to authentication to the CloudKarafka API",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"cloudkarafka_instance":       resourceInstance(),
			"cloudkarafka_instance_alarm": resourceAlarm(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	return getCustomerClient(d.Get("apikey").(string)), nil
}
