package main

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/takebayashi/go-dozens/dozens"
)

func resourceDomain() *schema.Resource {
	return &schema.Resource{
		Create: createDomain,
		Read:   readDomain,
		Update: nil,
		Delete: deleteDomain,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"mail": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func createDomain(d *schema.ResourceData, m interface{}) error {
	name := d.Get("name").(string)
	mail := d.Get("mail").(string)
	list, err := m.(*dozens.Client).AddDomain(name, mail)
	for _, domain := range list {
		if domain.Name == name {
			d.SetId(domain.Id)
		}
	}
	return err
}

func readDomain(d *schema.ResourceData, m interface{}) error {
	name := d.Get("name").(string)
	list, err := m.(*dozens.Client).ListDomains()
	for _, domain := range list {
		if domain.Name == name {
			d.SetId(domain.Id)
		}
	}
	return err
}

func deleteDomain(d *schema.ResourceData, m interface{}) error {
	name := d.Get("name").(string)
	ddel, err := m.(*dozens.Client).GetDomain(name)
	if err != nil {
		return nil
	}
	_, err = m.(*dozens.Client).DeleteDomain(ddel)
	d.SetId("")
	return err
}
