package main

import (
	"time"

	"github.com/avanha/pmaas-core/config"
	fakes "github.com/avanha/pmaas-plugin-fakes"
	"github.com/avanha/pmaas-spi/alert"
	"github.com/avanha/pmaas-spi/device"
	"github.com/avanha/pmaas-spi/environment"
)

// The fake devices the demo automations use. Set by addFakes.
var (
	holidayLights            device.ISwitch
	fakeSwitch1, fakeSwitch2 device.ISwitch
	fakeFanSwitch            device.ISwitch
	fakeThermometer          environment.IWirelessThermometer

	fakeAtticFanSwitch                           device.ISwitch
	fakeAtticThermometer, fakeOutsideThermometer environment.IWirelessThermometer

	fakeGarageDoor device.IDoorSensor
	fakeRouter     *fakes.HostHandle
)

// addFakes adds the plugin that has the fake devices, so that the demo has things to automate and to alert about without
// any hardware: holiday lights, two switches the toggle automation flips, a fan switch and thermometer, an attic fan with an
// attic and an outside thermometer, a garage door sensor, and a router the network monitor polls. Open the Fakes page
// to change them: press a switch, set a temperature, open the garage door, take the router down.
//
// Replace a fake with the real thing by swapping the line that makes it, say for an ESPHome plug's switch: the
// automations only ask for a device.ISwitch. Add this before the plugin that runs the automations, so that it is stopped after.
func addFakes(coreConfig *config.Config) {
	conf := fakes.NewPluginConfig()
	holidayLights = conf.AddSwitch("Holiday lights")
	fakeSwitch1 = conf.AddSwitch("Fake switch 1")
	fakeSwitch2 = conf.AddSwitch("Fake switch 2")
	fakeFanSwitch = conf.AddSwitch("Fake Fan Switch")
	fakeThermometer = conf.AddThermometer("Fake Thermometer")
	fakeAtticFanSwitch = conf.AddSwitch("Fake Attic Fan Switch")
	fakeAtticThermometer = conf.AddThermometer("Fake Attic Thermometer")
	fakeOutsideThermometer = conf.AddThermometer("Fake Outside Thermometer")

	// A warning once the sensor has been out of touch for five minutes: a setting on the device, not an automation.
	// With no setting there is no alert about it.
	fakeGarageDoor = conf.AddDoorSensor("Fake Garage Door").AlertWhenOffline(alert.WhenOffline(5*time.Minute, nil))
	fakeRouter = conf.AddHost("fake-router")

	coreConfig.AddPlugin(fakes.NewPlugin(conf), config.PluginConfig{})
}
