package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/takebayashi/go-dozens/dozens"
)

func resourceRecord() *schema.Resource {
	return &schema.Resource{
		Create: createRecord,
		Read:   readRecord,
		Update: updateRecord,
		Delete: deleteRecord,
		Schema: map[string]*schema.Schema{
			"domain": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"address": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"ttl": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func createRecord(d *schema.ResourceData, m interface{}) error {
	client := m.(*dozens.Client)
	domain, err := client.GetDomain(d.Get("domain").(string))
	if err != nil {
		return err
	}
	name := d.Get("name").(string)
	record := &dozens.Record{SName: name, Type: d.Get("type").(string), Prio: d.Get("priority").(string), Content: d.Get("address").(string), Ttl: d.Get("ttl").(string)}
	record, err = client.AddRecord(domain, record)
	if err != nil {
		return err
	}
	applyRecord(record, d)
	return nil
}

func readRecord(d *schema.ResourceData, m interface{}) error {
	client := m.(*dozens.Client)
	domain, err := client.GetDomain(d.Get("domain").(string))
	if err != nil {
		return err
	}
	name := d.Get("name").(string)
	fqname := name + "." + domain.Name
	if name == "" {
		fqname = domain.Name
	}

	list, err := client.ListRecords(domain)
	if err != nil {
		return err
	}

	for _, record := range list {
		if record.FQName == fqname {
			applyRecord(record, d)
			return nil
		}
	}
	d.SetId("")
	return nil
}

func updateRecord(d *schema.ResourceData, m interface{}) error {
	client := m.(*dozens.Client)
	record := &dozens.Record{Id: d.Id(), Prio: d.Get("priority").(string), Content: d.Get("address").(string), Ttl: d.Get("ttl").(string)}
	_, err := client.EditRecord(record)
	return err
}

func deleteRecord(d *schema.ResourceData, m interface{}) error {
	client := m.(*dozens.Client)
	domain, err := client.GetDomain(d.Get("domain").(string))
	if err != nil {
		return err
	}
	name := d.Get("name").(string)
	list, err := client.ListRecords(domain)
	if err != nil {
		return nil
	}
	var rdel *dozens.Record
	fqname := name + "." + domain.Name
	if name == "" {
		fqname = domain.Name
	}
	for _, record := range list {
		if record.FQName == fqname {
			rdel = record
		}
	}
	err = m.(*dozens.Client).DeleteRecord(rdel)
	d.SetId("")
	return err
}

func applyRecord(r *dozens.Record, d *schema.ResourceData) {
	d.SetId(r.Id)
	d.Set("type", r.Type)
	d.Set("priority", r.Prio)
	d.Set("address", r.Content)
	d.Set("ttl", r.Ttl)
}
