package main

import (
	"fmt"
	"io/ioutil"
//	"os"
	"runtime"

	dCli "github.com/DigitalVault/diva/cli"
	"github.com/DigitalVault/diva/core/info"
	"github.com/alecthomas/kong"
	log "github.com/sirupsen/logrus"
)

var (
	Version string
	Build   string

	CLI struct {
		Scan struct {
			Path string `help:"Path to the digital vault location." default:"."`
		} `cmd:"" help:"Scan filesystem for files to be preserved."`

		Ingest struct {
			Path string `help:"Path to the digital vault location." default:"."`
		} `cmd:"" help:"Scan filesystem and ingest files to the digital vault."`

		Version struct {
		} `cmd:"" help:"Prints version info of the diva-cli"`

		LogLevel string `help:"Setting log level [Off | Debug | Info | Warning | Error]"`
	}
)

func init() {
	log.SetOutput(ioutil.Discard)
	log.SetReportCaller(true)
	log.SetLevel(log.WarnLevel)

}

func main() {
	ctx := kong.Parse(&CLI)
	switch ctx.Command() {
	case "scan":
		fmt.Println("Scan :", ctx.Args)
	case "ingest":
		fmt.Println("Ingest :", ctx.Args)
	case "version":
		{
			fmt.Printf("diva-cli (%s %s)\r\n\r\nVersion          : %s\r\n", runtime.GOOS, runtime.GOARCH, Version)
			fmt.Println("Build date       :", Build)
			fmt.Println("\r\nPackage versions :")
			dCli.Version()
//			log.SetOutput(os.Stdout)
//			log.SetLevel(log.DebugLevel)
//			log.Info("SysInfo")
			info.Info.SysInfo()
		}
	default:
		panic(ctx.Command())

	}
}
