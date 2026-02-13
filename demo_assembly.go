package main

import (
	"fmt"

	"github.com/avanha/pmaas-core"
	"github.com/avanha/pmaas-core/config"
	basicwebui "github.com/avanha/pmaas-plugin-basicwebui"
)

func main() {
	fmt.Printf("Starting Demo PMAAS assembly\n")
	conf := config.NewConfig()
	conf.HttpPort = 8090

	addSimpleWebUI(conf)

	var pmaas = core.NewPMAAS(conf)
	err := pmaas.Run()

	if err != nil {
		fmt.Printf("pmaas.Run completed with error: %s\n", err)
	}

	fmt.Printf("Demo PMAAS assembly terminated\n")
}

func addSimpleWebUI(serverConfig *config.Config) {
	conf := basicwebui.NewPluginConfig()
	serverConfig.AddPlugin(basicwebui.NewPlugin(conf), config.PluginConfig{
		//ContentPathOverride: localProjectRoot + "/plugins/webrender/content",
	})
}
