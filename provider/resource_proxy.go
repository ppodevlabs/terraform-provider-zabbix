package provider

import (
	"errors"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/ppodevlabs/go-zabbix-api"
)

// proxySchemaBase base proxy schema
var proxySchemaBase = map[string]*schema.Schema{
	"name": &schema.Schema{
		Type:         schema.TypeString,
		Description:  "FQDN of proxy",
		ValidateFunc: validation.StringIsNotWhiteSpace,
		Required:     true,
	},
}

// dataProxy terraform proxy resource entrypoint
func dataProxy() *schema.Resource {
	return &schema.Resource{
		Read:   dataProxyRead,
		Schema: proxySchemaBase,
	}
}

// dataProxyRead read handler for data resource
func dataProxyRead(d *schema.ResourceData, m interface{}) error {

	api := m.(*zabbix.API)
	hostName := "name"
	if api.Config.Version < 70000 {
		hostName = "host"
	}
	log.Trace(fmt.Sprintf("version: %d", api.Config.Version))
	return proxyRead(d, m, zabbix.Params{
		"filter": map[string]interface{}{
			hostName: d.Get("name"),
		},
	})
}

// proxyRead common proxy read function
func proxyRead(d *schema.ResourceData, m interface{}, params zabbix.Params) error {
	api := m.(*zabbix.API)

	proxys, err := api.ProxiesGet(params)

	if err != nil {
		return err
	}

	if len(proxys) < 1 {
		return errors.New("proxy not found")
	}
	if len(proxys) > 1 {
		return errors.New("multiple proxys found")
	}
	proxy := proxys[0]

	log.Debug("Got proxy: %+v", proxy)

	d.SetId(proxy.ProxyID)
	d.Set("name", proxy.Name)

	return nil
}
