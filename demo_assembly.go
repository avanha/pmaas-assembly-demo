package main

import (
	"fmt"

	"github.com/avanha/pmaas-core"
	"github.com/avanha/pmaas-core/config"
	basicwebui "github.com/avanha/pmaas-plugin-basicwebui"
	gotexttemplate "github.com/avanha/pmaas-plugin-gotexttemplate"
)

func main() {
	fmt.Printf("Starting Demo PMAAS assembly\n")
	conf := config.NewConfig()
	conf.HttpPort = 8090

	addBasicWebUI(conf)

	var pmaas = core.NewPMAAS(conf)
	err := pmaas.Run()

	if err != nil {
		fmt.Printf("pmaas.Run completed with error: %s\n", err)
	}

	fmt.Printf("Demo PMAAS assembly terminated\n")
}

func addBasicWebUI(serverConfig *config.Config) {
	templateEngineConfig := gotexttemplate.NewGoLangTextTemplatePluginConfig()
	serverConfig.AddPlugin(gotexttemplate.NewGoLangTextTemplatePlugin(templateEngineConfig), config.PluginConfig{})

	conf := basicwebui.NewPluginConfig()
	serverConfig.AddPlugin(basicwebui.NewPlugin(conf), config.PluginConfig{
		//ContentPathOverride: localProjectRoot + "/plugins/webrender/content",
	})
}
