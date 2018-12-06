package main

import (
	"fmt"
)

type Accounts struct {
	Description string    `json:"Description,omitempty"`
	Accounts    []Account `json:"Items,omitempty"`
}

type Account struct {
	UserName    string `json:"UserName,omitempty"`
	Password    string `json:"Password,omitempty"`
	Description string `json:"Description,omitempty"`
	Id          string `json:"Id,omitempty"`
	Name        string `json:"Name,omitempty"`
	Oem         Oem    `json:"Oem,omitempty"`
	Type        string `json:"Type,omitempty"`
}

type Oem struct {
	Hp   Hp     `json:"Hp,omitempty"`
	Type string `json:"Type,omitempty"`
}

type Hp struct {
	LoginName  string     `json:"LoginName,omitempty"`
	Privileges Privileges `json:"Privileges,omitempty"`
}

type Privileges struct {
	LoginPriv                bool `json:"LoginPriv,omitempty"`
	RemoteConsolePriv        bool `json:"RemoteConsolePriv,omitempty"`
	UserConfigPriv           bool `json:"UserConfigPriv,omitempty"`
	VirtualMediaPriv         bool `json:"VirtualMediaPriv,omitempty"`
	VirtualPowerAndResetPriv bool `json:"VirtualPowerAndResetPriv,omitempty"`
	ILOConfigPriv            bool `json:"iLOConfigPriv,omitempty"`
}

func (a Accounts) Print() {
	for _, account := range a.Accounts {
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
}
