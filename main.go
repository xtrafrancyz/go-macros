package main

import (
	"embed"
	"fmt"
	"github.com/haroflow/go-macros/fakerinput"
	"html/template"
	"os"
)

//go:embed www
var fs embed.FS

//go:embed templates/help.template.html
var helpTemplateStr string
var helpTemplate *template.Template

var conf *Config

func main() {
	conf = loadConfig()

	if conf.FakerInput {
		defer fakerinput.InitGlobal().Close()
	}

	var err error
	helpTemplate, err = template.New("help").Parse(helpTemplateStr)
	if err != nil {
		panic(fmt.Sprintf("error parsing template: %s", err))
	}

	err = startUI()
	if err != nil {
		fmt.Println("error initializing UI:", err)
		os.Exit(1)
	}
}
