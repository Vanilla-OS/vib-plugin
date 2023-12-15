
# vib-plugin

A template to create plugins for [vib](https://github.com/vanilla-os/vib)

## Usage

To use this template, fork it and change the module name in `go.mod`.

Then you can modify the source in `src/plugin.go` to create your plugin.

We recommend adding the `vib-plugin` tag to your repo so other people can discover it.

## Plugin requirements

`src/plugin.go` explains all requirements for plugins with code examples.

A short tldr:

- vib requires `BuildModule(interface{}, recipe *api.Recipe*) (string, error)` to be available as it is used as the entry point for the plugin. Any other functions can be freely declared and will not be used by vib
- Each plugin needs to have a custom struct for the module, with at least two mandatory values: `Name string` and `Type string`
- It is recommended, but not required, to use the api functions for source definition or downloading sources

## Building

NOTE: Plugins will have to be compiled with the same go version as Vib, for builds from vanilla-os, this version is 1.21, it is recommended to use `podman`/`docker` or the GitHub workflow to build the plugins.

Plugins can be built with `go build -trimpath -buildmode=plugin`, producing a `.so` file.
To use the plugin, the `.so` file has to be moved into the `plugins/` directory of the vib recipe.
Plugins are only loaded when required, if a recipe never uses a plugin example, `example.so` will never be loaded.

This template contains a GitHub workflow that can build the plugin with the right arguments automatically.
Otherwise one can use podman/docker to build the plugin:
`docker run --rm -v "$PWD":/usr/src/myapp -w /usr/src/myapp golang:1.19 go build -trimpath -buildmode=plugin -v`
