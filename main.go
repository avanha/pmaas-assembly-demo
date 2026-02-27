package main

import (
	"fmt"
	"runtime"

	"github.com/avanha/pmaas-core"
	"github.com/avanha/pmaas-core/config"
	basicwebui "github.com/avanha/pmaas-plugin-basicwebui"
	bluetooth "github.com/avanha/pmaas-plugin-bluetooth"
	dblog "github.com/avanha/pmaas-plugin-dblog"
	environment "github.com/avanha/pmaas-plugin-environment"
	gotexttemplate "github.com/avanha/pmaas-plugin-gotexttemplate"
	tunnelbroker "github.com/avanha/pmaas-plugin-hetunnelbroker"
	netmon "github.com/avanha/pmaas-plugin-netmon"
	porkbun "github.com/avanha/pmaas-plugin-porkbun"
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
	addHeTunnelBroker(conf)

	if runtime.GOOS == "linux" {
		// The server will run on MacOS, but the scanner will continue to fail
		addBluetooth(conf)
	}

	var pmaas = core.NewPMAAS(conf)
	err := pmaas.Run()

	if err != nil {
		fmt.Printf("pmaas.Run completed with error: %s\n", err)
	}

	fmt.Printf("Demo PMAAS assembly terminated\n")
}

func addBasicWebUI(serverConfig *config.Config) {
	templateEngineConfig := gotexttemplate.NewPluginConfig()
	serverConfig.AddPlugin(gotexttemplate.NewPlugin(templateEngineConfig), config.PluginConfig{})

	conf := basicwebui.NewPluginConfig()
	serverConfig.AddPlugin(basicwebui.NewPlugin(conf), config.PluginConfig{
		//ContentPathOverride: localProjectRoot + "/plugins/webrender/content",
	})
}

func addBluetooth(coreConfig *config.Config) {
	conf := bluetooth.NewPluginConfig()
	conf.EnableTestDevices = true
	conf.AddThermometer("A4:C1:38:A1:B2:C3", "Basement")
	coreConfig.AddPlugin(bluetooth.NewPlugin(conf), config.PluginConfig{
		//ContentPathOverride: localProjectRoot + "/plugins/bluetooth/content",
	})
}

func addDbLog(serverConfig *config.Config) {
	conf := dblog.NewPluginConfig()
	conf.DriverName = "postgres"
	// This is how you define a data source.
	conf.DataSourceName = fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"localhost", 5432, "pmaas_user", "pmaas_user_password", "pmaas_db")
	serverConfig.AddPlugin(dblog.NewPlugin(conf), config.PluginConfig{
		//ContentPathOverride: localProjectRoot + "/plugins/dblog/internal/http/content",
	})
}

func addNetmon(serverConfig *config.Config) {
	conf := netmon.NewPluginConfig()
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
	conf := porkbun.NewPluginConfig()
	conf.ApiKey = "porkbunApiKey"
	conf.ApiSecret = "porkbunApiSecret"
	//exampledotcom := conf.AddDomain("example.com")
	//hostDnsARecord := exampledotcom.AddDnsRecord("A", "host")
	//hostDnsAAAARecord := exampledotcom.AddDnsRecord("AAAA", "host")
	serverConfig.AddPlugin(porkbun.NewPlugin(conf), config.PluginConfig{
		//ContentPathOverride: localProjectRoot + "/plugins/porkbun/internal/http/content",
	})
}

func addHeTunnelBroker(serverConfig *config.Config) {
	conf := tunnelbroker.NewPluginConfig()
	serverConfig.AddPlugin(tunnelbroker.NewPlugin(conf), config.PluginConfig{})
}
