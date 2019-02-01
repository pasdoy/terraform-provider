package main

import (
	"github.com/dghubble/sling"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceInstance() *schema.Resource {
	return &schema.Resource{
		Create: resourceCreate,
		Read:   resourceRead,
		Update: resourceUpdate,
		Delete: resourceDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the instance",
			},
			"plan": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the plan",
			},
			"tags": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Tags",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"region": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Name of the region you want to create your instance in",
			},
			"vpc_subnet": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Dedicated VPC subnet, shouldn't overlap with your current VPC's subnet",
			},
			"ca": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Broker CA",
			},
			"brokers": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Comma separated list of Kafka broker urls",
			},
			"username": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Username for accessing the Kafka cluster",
			},
			"password": {
				Type:        schema.TypeString,
				Computed:    true,
				Sensitive:   true,
				Description: "Password for accessing the Kafka cluster",
			},
			"apikey": {
				Type:        schema.TypeString,
				Computed:    true,
				Sensitive:   true,
				Description: "API key for the Kafka cluster",
			},
		},
	}
}

func resourceCreate(d *schema.ResourceData, meta interface{}) error {
	api := meta.(*sling.Sling)

	keys := []string{"name", "plan", "region", "vpc_subnet"}
	params := make(map[string]interface{})
	for _, k := range keys {
		if v := d.Get(k); v != nil {
			params[k] = v
		}
	}

	data := make(map[string]interface{})
	_, err := api.Post("/api/instances").BodyJSON(params).ReceiveSuccess(&data)
	if err != nil {
		return err
	}

	d.SetId(idToString(data["id"]))
	for k, v := range data {
		d.Set(k, v)
	}

	return nil

}

func resourceRead(d *schema.ResourceData, meta interface{}) error {
	api := meta.(*sling.Sling)

	data := make(map[string]interface{})
	_, err := api.Get("/api/instances/" + d.Id()).ReceiveSuccess(&data)
	if err != nil {
		return err
	}

	for k, v := range data {
		d.Set(k, v)
	}

	return nil

}

func resourceUpdate(d *schema.ResourceData, meta interface{}) error {
	api := meta.(*sling.Sling)

	keys := []string{"name", "plan"}
	params := make(map[string]interface{})
	for _, k := range keys {
		params[k] = d.Get(k)
	}

	_, err := api.Put("/api/instances/" + d.Id()).BodyJSON(params).ReceiveSuccess(nil)
	return err
}

func resourceDelete(d *schema.ResourceData, meta interface{}) error {
	api := meta.(*sling.Sling)
	_, err := api.Delete("/api/instances/" + d.Id()).ReceiveSuccess(nil)
	return err
}
