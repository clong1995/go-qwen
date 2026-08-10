package conf

import (
	pcolor "github.com/clong1995/go-ansi-color"
	"github.com/clong1995/go-config"
)

var ApiKey string
var ImageModel string
var TextModel string
var Prefix = "qwen"

func init() {
	var exists bool
	if ApiKey, exists = config.Value[string]("API KEY"); !exists {
		pcolor.PrintFatal(Prefix, "API KEY not found")
		return
	}
	if ImageModel, exists = config.Value[string]("IMAGE MODEL"); !exists {
		pcolor.PrintFatal(Prefix, "IMAGE MODEL not found")
		return
	}
	if TextModel, exists = config.Value[string]("TEXT MODEL"); !exists {
		pcolor.PrintFatal(Prefix, "TEXT MODEL not found")
		return
	}
}
