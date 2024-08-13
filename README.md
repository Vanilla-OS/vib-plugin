
# vib-plugin

A template to create plugins for [vib](https://github.com/vanilla-os/vib)

## Usage

To use this template, fork it and change the module name in `go.mod`.

Then you can modify the source in `src/plugin.go` to create your plugin.

A couple of test cases should also be added in `src/plugin_test.go`, view [vib-fsguard](https://github.com/vanilla-os/vib) and [vib-pacman](https://github.com/axtloss/vib-plugins/blob/main/pacman/plugin_test.go) for examples.

We recommend adding the `vib-plugin` tag to your repo so other people can discover it.

## Plugin requirements

`src/plugin.go` explains all requirements for plugins with code examples.

A short tldr:

- vib requires `BuildModule(moduleInterface *C.char, recipeInterface *C.char) *C.char` to be available as it is used as the entry point for the plugin. Any other functions can be freely declared and will not be used by vib
- plugins need to pass their name and plugin type through the `PlugInfo() *C.Char` function, the api contains a struct of `api.PluginInfo` which makes this easier, vib expects this function to return a json marshalled version of this struct
- Each plugin needs to have a custom struct for the module, with at least two mandatory values: `Name string` and `Type string`
- It is recommended, but not required, to use the api functions for source definition or downloading sources

## Building
Plugins can be built with `go build -buildmode=c-shared -o plugin.so`, producing a `.so` file.
To use the plugin, the `.so` file has to be moved into the `plugins/` directory of the vib recipe.
Plugins are only loaded when required, if a recipe never uses a plugin example, `example.so` will never be loaded.

This template contains a GitHub workflow that can build the plugin with the right arguments automatically.
