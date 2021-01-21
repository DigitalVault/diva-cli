package main

import (
  "fmt"
  "github.com/alecthomas/kong"
  dCli "github.com/DigitalVault/diva/cli"
)

var (
  Version string
  Build string

  CLI struct {
    Scan struct {
      Path string `help:"Path to the digital vault location." default:"."`

    } `cmd:"" help:"Scan filesystem for files to be preserved."`

    Ingest struct {
      Path string `help:"Path to the digital vault location." default:"."`

    } `cmd:"" help:"Scan filesystem and ingest files to the digital vault."`

    Version struct {
    } `cmd:"" help:"Prints version info of the diva-cli"`
  }
)

func main() {
  ctx := kong.Parse(&CLI)
  switch ctx.Command() {
  case "scan":
    fmt.Println("Scan :", ctx.Args)
  case "ingest":
    fmt.Println("Ingest :", ctx.Args)
    case "version": {
      fmt.Println("\r\nVersion          :", Version)
      fmt.Println("Build date       :", Build)
      fmt.Println("\r\nPackage versions :")
      dCli.Version()
    }
  default:
    panic(ctx.Command())

  }
}
