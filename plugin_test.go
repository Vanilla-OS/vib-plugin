package main

import (
	"encoding/json"
	"testing"

	"github.com/vanilla-os/vib/api"
)

type testCases struct {
	module   interface{}
	expected string
}

// This slice is most likely the only thing that needs to be modified
// It should be filled with the cases that the function should run through
// it consists of the module that should be passed and the output that should result from the module
var test = []testCases{
	{ExampleModule{Name: "Case Descrition", Type: "PluginName", Source: api.Source{Type: "tar"}}, "expected result here"},
}

// Simply run through all the test cases defined above
func TestBuildModule(t *testing.T) {
	for _, testCase := range test {
		moduleInterface, err := json.Marshal(testCase.module)
		if err != nil {
			t.Errorf("Error in json %s", err.Error())
		}
		if output := BuildModule(convertToCString(string(moduleInterface)), convertToCString("")); convertToGoString(output) != testCase.expected {
			t.Errorf("Output %s not equivalent to expected %s", convertToGoString(output), testCase.expected)
		}
	}

}
