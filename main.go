package main

import (
  "fmt"
  "github.com/alecthomas/kong"
)

var CLI struct {
  Scan struct {

  } `cmd help:"Scan filesystem for files to be preserved."`

  Ingest struct {

  } `cmd help:"Scan filesystem and ingest files to the digital vault."`
}

func main() {
  ctx := kong.Parse(&CLI)
  switch ctx.Command() {
  case "scan":
    fmt.Println("Scan")
  case "ingest":
    fmt.Println("Ingest")
  default:
    panic(ctx.Command())

  }
}
