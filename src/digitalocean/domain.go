package digitalocean

type Domains struct {
	Domains []Domain     `json:"domains"`
	Meta    DropletsMeta `json:"meta"`
}

type Domain struct {
	Name     string `json:"name"`
	Ttl      int    `json:"ttl"`
	ZoneFile string `json:"zone_file"`
}

type NewDomain struct {
	Name      string `json:"name"`
	IpAddress int    `json:"ip_address"`
}
