// provider.go
package bluewonder

import (
	"context"
	"fmt"

	"example.com/client/getme"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	fmt.Println("Creating bluewonder provider...")
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"bluewonder_me": providerResourceMe(),
		},
	}
}

func providerResourceMe() *schema.Resource {
	fmt.Println("Creating bluewonder_me resource...")
	return &schema.Resource{
		ReadContext: providerResourceMeRead,
	}
}

func providerResourceMeRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	fmt.Println("Reading bluewonder_me resource...")
	// Call your getme.GetMe function here
	path := "me"
	data, err := getme.GetMe(path)
	if err != nil {
		return diag.FromErr(err)
	}

	// Set the resource data
	d.SetId(data.Name)
	d.Set("name", data.Name)
	d.Set("type", data.Type)
	d.Set("groups", data.Groups)

	return nil
}
