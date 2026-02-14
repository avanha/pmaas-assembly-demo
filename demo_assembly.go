package main

import (
	"fmt"

	"github.com/avanha/pmaas-core"
	"github.com/avanha/pmaas-core/config"
	basicwebui "github.com/avanha/pmaas-plugin-basicwebui"
	dblog "github.com/avanha/pmaas-plugin-dblog"
	dblogconfig "github.com/avanha/pmaas-plugin-dblog/config"
	environment "github.com/avanha/pmaas-plugin-environment"
	gotexttemplate "github.com/avanha/pmaas-plugin-gotexttemplate"
	netmon "github.com/avanha/pmaas-plugin-netmon"
	netmonconfig "github.com/avanha/pmaas-plugin-netmon/config"
	porkbun "github.com/avanha/pmaas-plugin-porkbun"
	porkbunconfig "github.com/avanha/pmaas-plugin-porkbun/config"
)

func main() {
	fmt.Printf("Starting Demo PMAAS assembly\n")
	conf := config.NewConfig()
	conf.HttpPort = 8090

	addBasicWebUI(conf)
	addDbLog(conf)
	addNetmon(conf)
	addEnvironment(conf)
	addPorkBun(conf)

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

func addDbLog(serverConfig *config.Config) {
	conf := dblogconfig.NewPluginConfig()
	conf.DriverName = "postgres"
	conf.DataSourceName = localDataSource
	serverConfig.AddPlugin(dblog.NewPlugin(conf), config.PluginConfig{
		//ContentPathOverride: localProjectRoot + "/plugins/dblog/internal/http/content",
	})
}

func addNetmon(serverConfig *config.Config) {
	conf := netmonconfig.NewPluginConfig()
	serverConfig.AddPlugin(netmon.NewPlugin(conf), config.PluginConfig{
		//ContentPathOverride: localProjectRoot + "/plugins/netmon/internal/http/content",
	})
}

func addEnvironment(serverConfig *config.Config) {
	environmentConfig := environment.NewPluginConfig()
	serverConfig.AddPlugin(environment.NewPlugin(environmentConfig), config.PluginConfig{
		//ContentPathOverride: srcRoot + "/plugins/environment/content",
	})
}

func addPorkBun(serverConfig *config.Config) {
	conf := porkbunconfig.NewPluginConfig()
	conf.ApiKey = "porkbunApiKey"
	conf.ApiSecret = "porkbunApiSecret"
	//exampledotcom := conf.AddDomain("example.com")
	//hostDnsARecord := exampledotcom.AddDnsRecord("A", "host")
	//hostDnsAAAARecord := exampledotcom.AddDnsRecord("AAAA", "host")
	serverConfig.AddPlugin(porkbun.NewPlugin(conf), config.PluginConfig{
		//ContentPathOverride: localProjectRoot + "/plugins/porkbun/internal/http/content",
	})
}
