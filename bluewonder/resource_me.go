// resource_me.go
package bluewonder

import (
	"fmt"

	"example.com/client/getme"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceMe() *schema.Resource {
	fmt.Println("Creating resourceMe resource...")
	return &schema.Resource{
		Read:   resourceMeRead,
		Schema: map[string]*schema.Schema{
			// Define your resource schema here if applicable
		},
	}
}

func resourceMeRead(d *schema.ResourceData, meta interface{}) error {
	fmt.Println("Reading resourceMe resource...")
	// Call the GetMe function
	response, err := getme.GetMe("me")
	if err != nil {
		return err
	}

	// Set resource data attributes based on the response
	d.Set("name", response.Name)
	d.Set("type", response.Type)
	d.Set("groups", response.Groups)

	return nil
}
