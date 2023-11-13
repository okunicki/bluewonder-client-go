package main

import (
	"example.com/client/bluewonder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: bluewonder.Provider,
	})
}
