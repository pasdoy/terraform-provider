package main

import (
	"errors"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceAlarm() *schema.Resource {
	return &schema.Resource{
		Create: resourceAlarmCreate,
		Read:   resourceAlarmRead,
		//Update: resourceUpdate,
		Delete: resourceDelete,
		Schema: map[string]*schema.Schema{
			"alarm_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Alarm ID",
			},
			"apikey": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The cluster API key",
			},
			"type": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The type of alarm to set (cpu, memory, disk)",
			},
			"value_threshold": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "What value to trigger the alarm for",
			},
			"time_threshold": {
				Type:        schema.TypeInt,
				Required:    true,
				ForceNew:    true,
				Description: "For how long (in seconds) the value_threshold should be active before trigger alarm",
			},
		},
	}
}

func resourceAlarmCreate(d *schema.ResourceData, meta interface{}) error {
	api := getAPIClient(d.Get("apikey").(string))

	keys := []string{"type", "value_threshold", "time_threshold"}
	params := make(map[string]interface{})
	for _, k := range keys {
		if v := d.Get(k); v != nil {
			params[k] = v
		}
	}

	data := make(map[string]interface{})
	_, err := api.Post("/api/alarms").BodyJSON(params).ReceiveSuccess(&data)
	if err != nil {
		return err
	}

	d.SetId(idToString(data["id"]))
	for k, v := range data {
		d.Set(k, v)
	}

	return nil

}

func resourceAlarmRead(d *schema.ResourceData, meta interface{}) error {
	api := getAPIClient(d.Get("apikey").(string))

	data := make([]map[string]interface{}, 0)
	_, err := api.Get("/api/alarms").ReceiveSuccess(&data)
	if err != nil {
		return err
	}

	for _, alarm := range data {
		if idToString(alarm["id"]) == d.Id() {
			for k, v := range alarm {
				d.Set(k, v)
			}
			return nil
		}

	}
	return errors.New("Cannot read alarms")

}

func resourceAlarmDelete(d *schema.ResourceData, meta interface{}) error {
	api := getAPIClient(d.Get("apikey").(string))

	params := map[string]string{
		"id": d.Id(),
	}

	_, err := api.Post("/api/alarms").BodyJSON(params).ReceiveSuccess(nil)
	return err
}
