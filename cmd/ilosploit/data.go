package main

type Accounts struct {
	Description string    `json:"Description,omitempty"`
	Accounts    []Account `json:"Items,omitempty"`
}

type Account struct {
	Description string `json:"Description,omitempty"`
	Id          string `json:"Id,omitempty"`
	Name        string `json:"Name,omitempty"`
	Oem         Oem    `json:"Oem,omitempty"`
	Type        string `json:"Type,omitempty"`
	UserName    string `json:"UserName,omitempty"`
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
	iLOConfigPriv            bool `json:"iLOConfigPriv,omitempty"`
}
