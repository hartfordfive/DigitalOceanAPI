package digitalocean

type DomainRecords struct {
	DomainRecords []DomainRecord `json:"domain_records"`
	Meta          DropletsMeta   `json:"meta"`
}

type DomainRecord struct {
	Id       int    `json:"id"`
	Type     string `json:"type"`
	Name     string `json:"name"`
	Data     string `json:"data"`
	Priority int    `json:"priority"`
	Port     int    `json:"port"`
	Weight   int    `json:"weight"`
}
