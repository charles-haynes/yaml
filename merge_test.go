package main

import (
	"testing"

	yaml "gopkg.in/yaml.v2"
)

func TestMerge(t *testing.T) {
	result, _ := mergeYaml([]string{"examples/data1.yaml", "examples/data2.yaml", "examples/data3.yaml"})
	expected := yaml.MapSlice{
		yaml.MapItem{Key: "a", Value: "simple"},
		yaml.MapItem{Key: "b", Value: []interface{}{1, 2}},
		yaml.MapItem{Key: "c", Value: yaml.MapSlice{yaml.MapItem{Key: "other", Value: true}, yaml.MapItem{Key: "test", Value: 1}}},
		yaml.MapItem{Key: "d", Value: false},
	}
	assertResultComplex(t, expected, result)
}

func TestMergeWithOverwrite(t *testing.T) {
	overwriteFlag = true
	result, _ := mergeYaml([]string{"examples/data1.yaml", "examples/data2.yaml", "examples/data3.yaml"})
	expected := yaml.MapSlice{
		yaml.MapItem{Key: "a", Value: "other"},
		yaml.MapItem{Key: "b", Value: []interface{}{2, 3, 4}},
		yaml.MapItem{Key: "c", Value: yaml.MapSlice{yaml.MapItem{Key: "other", Value: true}, yaml.MapItem{Key: "test", Value: 2}}},
		yaml.MapItem{Key: "d", Value: false},
	}
	assertResultComplex(t, expected, result)
}

func TestMergeYAML(t *testing.T) {
	overwriteFlag = true
	result, _ := mergeYaml([]string{"examples/up.yaml", "examples/gr.yaml"})
	expected := yaml.MapSlice{
		yaml.MapItem{Key: "image", Value: "https://images.gr-assets.com/books/1328834743m/2865998.jpg"},
		yaml.MapItem{Key: "pages", Value: 240},
		yaml.MapItem{Key: "title", Value: "Fashioning Technology: A DIY Intro to Smart Crafting"},
		yaml.MapItem{Key: "authors", Value: []interface{}{"Syuzi Pakhchyan"}},
	}
	assertResultComplex(t, expected, result)
}
