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
	domain := &dozens.Domain{Name: d.Get("name").(string)}
	mail := d.Get("mail").(string)
	domain, err := m.(*dozens.Client).AddDomain(domain, mail)
	if err != nil {
		return err
	}
	applyDomain(domain, d)
	return nil
}

func readDomain(d *schema.ResourceData, m interface{}) error {
	list, err := m.(*dozens.Client).ListDomains()
	if err != nil {
		return err
	}
	for _, e := range list {
		if e.Id == d.Id() {
			applyDomain(e, d)
		}
	}
	return nil
}

func deleteDomain(d *schema.ResourceData, m interface{}) error {
	name := d.Get("name").(string)
	ddel, err := m.(*dozens.Client).GetDomain(name)
	if err != nil {
		return nil
	}
	err = m.(*dozens.Client).DeleteDomain(ddel)
	if err == nil {
		d.SetId("")
	}
	return err
}

func applyDomain(d *dozens.Domain, r *schema.ResourceData) {
	r.SetId(d.Id)
	r.Set("name", d.Name)
}
