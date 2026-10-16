package main

import (
	"time"

	pmaasautomations "github.com/avanha/pmaas-automations"
	"github.com/avanha/pmaas-automations/sun"
	"github.com/avanha/pmaas-automations/trigger"
	"github.com/avanha/pmaas-core/config"
	automations "github.com/avanha/pmaas-plugin-automations"
	"github.com/avanha/pmaas-spi/alert"
	"github.com/avanha/pmaas-spi/device"
)

// home is where the sun is worked out for: change it to wherever you are. Sunrise, sunset and the dusk and dawn the garage
// door alert goes by all depend on it.
var home = sun.Location{Latitude: 40.7128, Longitude: -74.0060} // New York, NY

// addAutomations adds the plugin that runs automations, and the demo automations over the fake devices. Add it after the
// plugins the automations use, so that it is stopped before them.
func addAutomations(coreConfig *config.Config) {
	conf := automations.NewPluginConfig()

	conf.Add(pmaasautomations.New(
		"holiday-lights-on", "Turn on the holiday lights 15 minutes before sunset",
		trigger.Sun(home, sun.Sunset, -15*time.Minute),
		pmaasautomations.TurnOn(holidayLights),
			))

	conf.Add(pmaasautomations.New(
		"holiday-lights-off", "Turn off the holiday lights at 12:30 AM",
		trigger.AtTimeOfDay(0, 30),
		pmaasautomations.TurnOff(holidayLights),
				pmaasautomations.WithCatchUp(6*time.Hour),
	))

	// Something to watch: the two fake switches flip together every five minutes, on the five. Open the Fakes page to
	// see them change and what PMAAS asked of them, and the Automations page to see it run.
	conf.Add(pmaasautomations.New(
		"fake-switches-toggle", "Toggle the fake switches every five minutes",
		trigger.Every(5*time.Minute),
		pmaasautomations.Toggle(fakeSwitch1, fakeSwitch2),
		pmaasautomations.WithDescription("Both switches are put in the same state, the opposite of the time before."),
		// It runs again in five minutes anyway, so a late run is not worth making.
		pmaasautomations.WithCatchUp(4*time.Minute),
	))

	// Watch the fan: on the Fakes page, turn the fan switch on ("Press: on") and set the thermometer above 75°F. At 1:00 AM
	// the automation waits, and the Automations page says why. Cool the room and it turns the fan off, once. Turn the
	// fan back on and it stays on, as someone took control; Re-enable on the Automations page hands it back. It does
	// nothing from 8:00.
	conf.Add(pmaasautomations.New(
		"fan-off-when-cool", "Turn off the fan after 1:00 AM once the room is below 75°F",
		trigger.AtTimeOfDay(1, 0),
		pmaasautomations.When(
			pmaasautomations.TurnOff(fakeFanSwitch),
			pmaasautomations.SwitchIs(fakeFanSwitch, device.SwitchStateOn),
			pmaasautomations.TemperatureBelow(fakeThermometer, pmaasautomations.Fahrenheit(75), 10*time.Minute),
		),
		pmaasautomations.WithDescription("Waits, from 1:00 AM, for the fan to be on and the room to be cool. Does nothing from 8:00."),
		// Due at 1:00, so this is up to 7:59.
		pmaasautomations.WithCatchUp(6*time.Hour+59*time.Minute),
	))

	// The attic fan. During the day, from 8:00 AM until an hour after sunset, the fan comes on when the attic is over 90°F
	// and warmer than outside, and goes off when the attic is under 88°F or outside is as warm. Between 88 and 90 it
	// stays as it is, so a temperature wobbling across 90 doesn't make it start and stop, and it is never changed more
	// than once in five minutes whatever the temperatures do. The fan, the attic and the outside thermometer are all
	// fakes: set the temperatures on the Fakes page to watch it. A reading more than ten minutes old can't be gone by,
	// and the fan is left as it is. It is checked every minute, and has no catch-up, so it never acts after the window has
	// closed. If it keeps failing, say the fan can't be reached, an alert is raised after ten failures in a row.
	conf.Add(pmaasautomations.New(
		"attic-fan", "Run the attic fan when the attic is over 90°F and warmer than outside, until it is under 88°F",
		trigger.Windowed(
			trigger.Every(time.Minute),
			trigger.AtTimeOfDay(8, 0),
			trigger.Sun(home, sun.Sunset, time.Hour)),
		pmaasautomations.Regulate(
			fakeAtticFanSwitch,
			pmaasautomations.AllOf(
				pmaasautomations.TemperatureAbove(fakeAtticThermometer, pmaasautomations.Fahrenheit(90), 10*time.Minute),
				pmaasautomations.Warmer(fakeAtticThermometer, fakeOutsideThermometer, 10*time.Minute)),
			pmaasautomations.AnyOf(
				pmaasautomations.TemperatureBelow(fakeAtticThermometer, pmaasautomations.Fahrenheit(88), 10*time.Minute),
				pmaasautomations.Not(pmaasautomations.Warmer(fakeAtticThermometer, fakeOutsideThermometer, 10*time.Minute))),
			5*time.Minute),
		pmaasautomations.WithDescription("Checked every minute from 8:00 AM until an hour after sunset. On above 90°F, off below 88°F, and changed no more often than every five minutes."),
		pmaasautomations.WithCatchUp(0),
	))

	// And it's switched off when the day ends, since the automation above stops looking then.
	conf.Add(pmaasautomations.New(
		"attic-fan-off", "Turn off the attic fan an hour after sunset",
		trigger.Sun(home, sun.Sunset, time.Hour),
		pmaasautomations.TurnOff(fakeAtticFanSwitch),
	))

	// The garage door. If it is left open for five minutes there is an alert on the alert console: a warning in the
	// daylight, and critical once it is dark (from civil dusk to civil dawn). It is the same alert throughout, and
	// escalates or eases as the light changes. It clears when the door closes. Open and close the Fake Garage Door on the
	// Fakes page to watch it. While the sensor is out of touch the door alert is left as it is; the sensor being out of
	// touch is an alert of its own, which the fakes plugin raises (see addFakes).
	conf.Add(pmaasautomations.New(
		"garage-door-open", "Alert when the garage door is left open",
		trigger.Every(time.Minute),
		pmaasautomations.WatchAlert(conf.Alerts(), "garage-door-open", "Garage door open",
			pmaasautomations.SeverityBetween(
				trigger.Sun(home, sun.CivilDusk, 0), trigger.Sun(home, sun.CivilDawn, 0),
				alert.SeverityCritical, alert.SeverityWarning),
			pmaasautomations.DoorOpenFor(fakeGarageDoor, 5*time.Minute)),
		pmaasautomations.WithDescription("Warning in daylight, critical after dark, for a door open more than five minutes."),
		pmaasautomations.WithCatchUp(0),
	))

	coreConfig.AddPlugin(automations.NewPlugin(conf), config.PluginConfig{})
}
