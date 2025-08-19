package main

import (
	"fmt"
	"C"
	"encoding/json"
	"github.com/vanilla-os/vib/api"
)

type ExampleModule struct {
	// Mandatory values, your plugin will not work if these are not present!
	Name string `json:"name"`
	Type string `json:"type"`

	// Additional values such as Source can be added here
	Source api.Source `json:"Source"`
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


// This function returns information about the plugin, most notably the type of plugin.
// Plugins that do not define this function are considered deprecated
// while they still work, support may be dropped in future releases.
//export PlugInfo
func PlugInfo() *C.char {
	plugininfo := &api.PluginInfo{
		Name: "PLUGINAME", // The name of the plugin
		Type: api.BuildPlugin, // The type of plugin. This plugin template does NOT function as a FinalizePlugin, so unless you have manually modified it accordingly, this value should stay as api.BuildPlugin
		UseContainerCmds: False, // If the plugin adds its own Containerfile directives, set it to False if the plugin only generates shell commands
	}
	pluginjson, err := json.Marshal(plugininfo)
	if err != nil {
		return C.CString(fmt.Sprintf("ERROR: %s", err.Error()))
	}
	return C.CString(string(pluginjson))
}

// This is the entry point for plugins that vib calls
// The arguments are required to be (*C.char, *C.char) => string
// Make sure NOT to remove the "export BuildModule"

//export BuildModule
func BuildModule(moduleInterface *C.char, recipeInterface *C.char, arch *C.char) *C.Char {
	// It is advisable to convert the interface to an actual struct
	// The use of json.Unmarshal for this is recommended, but not required
	var module *ExampleModule
	var recipe *api.Recipe
	
	err := json.Unmarshal([]byte(C.GoString(moduleInterface)), &module)
	if err != nil {

		// Due to the way c-shared builds work, the function has to return a CString
		// CGO includes a proper way to ensure strings are converted to CStrings
		// Any return in BuildModule will have to be wrapped with C.CString
		//
		// Since only one value can be returned, vib checks if the return
		// starts with "ERROR:", if this is the case, it assumes that the
		// plugin has failed and takes anything after the "ERROR:" as the
		// error message.
		return C.CString(fmt.Sprintf("ERROR: %s", err.Error()))
	}
	
	err = json.Unmarshal([]byte(C.GoString(recipeInterface)), &recipe)
	if err != nil {
		return C.CString(fmt.Sprintf("ERROR: %s", err.Error()))
	}
	
	err = fetchSources(module.Source, module.Name, recipe)
	if err != nil {
		return C.CString(fmt.Sprintf("ERROR: %s", err.Error()))
	}

	// The sources will be made available at /sources/ during build
	// if the plugins requires manually downloaded sources, they will
	// be available in /sources/<modulename>
	cmd := fmt.Sprintf("cd /sources/%s && cp * /etc/%s", module.Name, module.Name)

	return C.CString(cmd)
}

// Keep this function empty
// It will never be triggered. Any code in here is dead code.
func main() {}
