module github.com/DigitalVault/diva-cli

go 1.15
replace github.com/DigitalVault/diva => ../diva

require (
	github.com/DigitalVault/diva v0.0.5
	github.com/alecthomas/kong v0.2.12
	github.com/pkg/errors v0.9.1 // indirect
	github.com/sirupsen/logrus v1.7.0
)
