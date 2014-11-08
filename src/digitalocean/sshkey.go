package digitalocean

type SshKeys struct {
	SshKeys []SshKey     `json:"ssh_keys"`
	Links   string       `json:"links"`
	Meta    DropletsMeta `json:"meta"`
}

type SshKey struct {
	Id          int    `json:"id"`
	FingerPrint string `json:"finger_print"`
	PublicKey   string `json:"public_key"`
	Name        string `json:"name"`
}
