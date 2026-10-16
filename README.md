# PMAAS Demo Assembly

A PMAAS assembly that includes the core server and some plugins, to illustrate what is possible.

### Building

Just run `go build`

To build for another platform, say a 64-bit Raspberry Pi,
run `GOOS=linux GOARCH=arm64 go build -o pmaas-assembly-demo-arm64`

# What is in it

It runs without any hardware: the devices are fakes, so there is something to automate and something to alert about from
the first start. Open the **Fakes** page to press a switch, set a temperature, open the garage door or take the router
down, and watch the rest of the pages respond.

* **Fakes** (`pmaas-plugin-fakes`): holiday lights, two switches, a fan switch and thermometer, an attic fan with attic and
  outside thermometers, a garage door sensor and a router. They are registered as the entities the real devices are, so
  swapping one for a real device is a one-line change.
* **Automations** (`pmaas-plugin-automations`): a page that shows each automation, when it is next due and how it last
  went, with Run now, Pause/Resume and Re-enable buttons. The demo has:
  * the holiday lights come on 15 minutes before sunset and go off at 12:30 AM (set `home` in `automations.go` to your
    latitude and longitude);
  * the two fake switches are toggled together every five minutes;
  * the fan is turned off after 1:00 AM once the room is below 75°F, if it is on;
  * the attic fan runs from 8:00 AM until an hour after sunset while the attic is over 90°F and warmer than outside, and
    stops under 88°F, with no more than one change in five minutes;
  * an alert when the garage door is left open for five minutes: a warning in daylight, critical after dark. It clears when the
    door closes.
* **Alerts** (`pmaas-plugin-alertconsole`): the alert console, and an indicator in the header of every page that goes
  red for an unacknowledged critical alert. While one is showing, the page chimes about once a minute, and the speaker
  button beside the indicator silences it on that browser. There are two ways to see an alert:
  * leave the fake garage door open for five minutes; take the garage door sensor offline for five minutes, which is a
    setting on the device (`AlertWhenOffline`), not an automation; or take the fake router down for two minutes, which is a
    setting on the host (`AlertWhenUnreachable`).
* **Network monitor** (`pmaas-plugin-netmon`), **Environment**, **DbLog**, **Porkbun**, **HE Tunnel Broker** and (on
  Linux) **Bluetooth**, as before.

The Automations page can pause an automation, and the pause is kept across restarts.

# Web Access

The demo assembly's HTTP server port is 8090.  You can see status pages for plugins that have them, at their respective URLs.  For example the DbLog plugin is available at http://localhost:8090/plugins/dblog/
