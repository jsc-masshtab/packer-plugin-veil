package main

import (
	"fmt"
	"os"

	"github.com/hashicorp/packer-plugin-sdk/plugin"

	//"github.com/hashicorp/packer-plugin-veil/builder/veil"
	//"github.com/hashicorp/packer-plugin-veil/version"

	"packer-plugin-veil/builder/veil"
	"packer-plugin-veil/version"
)

func main() {
	pps := plugin.NewSet()
	pps.RegisterBuilder(plugin.DEFAULT_NAME, new(veil.Builder))
	pps.SetVersion(version.PluginVersion)
	err := pps.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
