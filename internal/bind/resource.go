package bind

import (
	"fmt"
	"net/http"

	"terraform-provider-haproxy/internal/transaction"
	"terraform-provider-haproxy/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func ResourceHaproxyBind() *schema.Resource {
	return &schema.Resource{
		Create: resourceHaproxyBindCreate,
		Read:   resourceHaproxyBindRead,
		Update: resourceHaproxyBindUpdate,
		Delete: resourceHaproxyBindDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"parent_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"parent_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"port": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"address": {
				Type:     schema.TypeString,
				Required: true,
			},
			"mode": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"maxconn": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"user": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"group": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceHaproxyBindRead(d *schema.ResourceData, m interface{}) error {
	bindName := d.Get("name").(string)
	parentName := d.Get("parent_name").(string)
	parentType := d.Get("parent_type").(string)

	configMap := m.(map[string]interface{})
	BindConfig := configMap["bind"].(*ConfigBind)
	tranConfig := configMap["transaction"].(*transaction.ConfigTransaction)

	resp, err := tranConfig.Transaction(func(transactionID string) (*http.Response, error) {
		return BindConfig.GetABindConfiguration(bindName, transactionID, parentName, parentType)
	})

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return utils.HandleError(bindName, "error reading Bind configuration", fmt.Errorf("response status: %s , err: %s", resp.Status, err))
	}

	d.SetId(bindName)
	return nil
}

func resourceHaproxyBindCreate(d *schema.ResourceData, m interface{}) error {
	bindName := d.Get("name").(string)
	port := d.Get("port").(int)
	address := d.Get("address").(string)
	parentName := d.Get("parent_name").(string)
	parentType := d.Get("parent_type").(string)
	maxconn := d.Get("maxconn").(int)
	user := d.Get("user").(string)
	group := d.Get("group").(string)
	mode := d.Get("mode").(string)

	payload := []byte(fmt.Sprintf(`
	{
		"name": "%s",
		"address": "%s",
		"port": %d,
		"maxconn": %d,
		"user": "%s",
		"group": "%s",
		"mode": "%s"
	}
	`, bindName, address, port, maxconn, user, group, mode))

	configMap := m.(map[string]interface{})
	BindConfig := configMap["bind"].(*ConfigBind)
	tranConfig := configMap["transaction"].(*transaction.ConfigTransaction)

	resp, err := tranConfig.Transaction(func(transactionID string) (*http.Response, error) {
		return BindConfig.AddBindConfiguration(payload, transactionID, parentName, parentType)
	})

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return utils.HandleError(bindName, "error creating Bind configuration", fmt.Errorf("response status: %s , err: %s", resp.Status, err))
	}

	d.SetId(bindName)
	return nil
}

func resourceHaproxyBindUpdate(d *schema.ResourceData, m interface{}) error {
	bindName := d.Get("name").(string)
	port := d.Get("port").(int)
	address := d.Get("address").(string)
	parentName := d.Get("parent_name").(string)
	parentType := d.Get("parent_type").(string)
	maxconn := d.Get("maxconn").(int)
	user := d.Get("user").(string)
	group := d.Get("group").(string)
	mode := d.Get("mode").(string)

	payload := []byte(fmt.Sprintf(`
	{
		"name": "%s",
		"address": "%s",
		"port": %d,
		"maxconn": %d,
		"user": "%s",
		"group": "%s",
		"mode": "%s"
	}
	`, bindName, address, port, maxconn, user, group, mode))

	configMap := m.(map[string]interface{})
	BindConfig := configMap["bind"].(*ConfigBind)
	tranConfig := configMap["transaction"].(*transaction.ConfigTransaction)

	resp, err := tranConfig.Transaction(func(transactionID string) (*http.Response, error) {
		return BindConfig.UpdateBindConfiguration(bindName, payload, transactionID, parentName, parentType)
	})

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return utils.HandleError(bindName, "error updating Bind configuration", fmt.Errorf("response status: %s , err: %s", resp.Status, err))
	}

	d.SetId(bindName)
	return nil
}

func resourceHaproxyBindDelete(d *schema.ResourceData, m interface{}) error {
	bindName := d.Get("name").(string)
	parentName := d.Get("parent_name").(string)
	parentType := d.Get("parent_type").(string)

	configMap := m.(map[string]interface{})
	BindConfig := configMap["bind"].(*ConfigBind)
	tranConfig := configMap["transaction"].(*transaction.ConfigTransaction)

	resp, err := tranConfig.Transaction(func(transactionID string) (*http.Response, error) {
		return BindConfig.DeleteBindConfiguration(bindName, transactionID, parentName, parentType)
	})

	if resp.StatusCode != 200 && resp.StatusCode != 202 {
		return utils.HandleError(bindName, "error deleting Bind configuration", fmt.Errorf("response status: %s , err: %s", resp.Status, err))
	}

	d.SetId("")
	return nil
}