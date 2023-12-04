package main

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/vanilla-os/vib/api"
)

type ExampleModule struct {
	// Mandatory values, your plugin will not work if these are not present!
	Name string `json:"name"`
	Type string `json:"type"`

	// Additional values such as Source can be added here
	Source api.Source
}

// Plugins can define extra functions that are used internally
// Vib will never call any function other than BuildModule
func fetchSources(source api.Source, name string, recipe *api.Recipe) error {
	// The plugin api offers functions to download sources
	// To be able to use them, the use of api.Source for
	// source definition is recommended
	// Using these functions to fetch sources is not required
	// but highly recommended to ensure sources are always in the right directory
	err := api.DownloadSource(recipe.DownloadsPath, source, name)
	if err != nil {
		return err
	}
	err = api.MoveSource(recipe.DownloadsPath, recipe.SourcesPath, source, name)
	return err
}

// This is the entry point for plugins that vib calls
// The arguments are required to be (interface{}, recipe) => (string, error)
func BuildModule(moduleInterface interface{}, recipe *api.Recipe) (string, error) {
	// It is advisable to convert the interface to an actual struct
	// The use of mapstructure for this is recommended, but not required
	var module ExampleModule
	err := mapstructure.Decode(moduleInterface, &module)
	if err != nil {
		return "", err
	}
	err = fetchSources(module.Source, module.Name, recipe)
	if err != nil {
		return "", err
	}

	// The sources will be made available at /sources/ during build
	// if the plugins requires manually downloaded sources, they will
	// be available in /sources/<modulename>
	cmd := fmt.Sprintf("cd /sources/%s && cp * /etc/%s", module.Name, module.Name)

	return cmd, nil
}
