# PMAAS Demo Assembly

A PMAAS assembly that includes the core server and some plugins, to illustrate what is possible.

### Building

Just run `go build`

To build for another platform, say a 64-bit Raspberry Pi,
run `GOOS=linux GOARCH=arm64 go build -o pmaas-assembly-demo-arm64`

# Web Access

The demo assembly's HTTP server port is 8090.  You can see status pages for plugins that have them, at their respective URLs.  For example the DbLog plugin is available at http://localhost:8090/plugins/dblog/
