package digitalocean

type Droplets struct {
	DropletList []*Droplet    `json:"droplets"`
	Meta        *DropletsMeta `json:"meta"`
}

type DropletsMeta struct {
	Total int `json:"total"`
}

type Droplet struct {
	Id          uint32              `json:"id"`
	Name        string              `json:"name"`
	Memory      int                 `json:"memory"`
	Vcpus       int                 `json:"vcpus"`
	Disk        int                 `json:"disk"`
	Region      *DropletRegion      `json:"region"`
	Image       *DropletImage       `json:"image"`
	Size        *DropletSize        `json:"meta"`
	Locked      bool                `json:"locked"`
	Status      string              `json:"status"`
	Networks    *DropletNetworks    `json:"networks"`
	Kernel      *DropletKernel      `json:"kernel"`
	CreatedAt   string              `json:"created_at"`
	Features    *DropletFeatures    `json:"features"`
	BackupIds   *DropletBabckupIds  `json:"backup_ids"`
	SnapshotIds *DropletSnapshotIds `json:"snapshot_ids"`
}

type DropletRegion struct {
	Slug      string   `json:"slug"`
	Name      string   `json:"name"`
	Sizes     []string `json:"sizes"`
	Available bool     `json:"available"`
	Features  []string `json:"features"`
}

type DropletImage struct {
	Id           uint32   `json:"id"`
	Name         string   `json:"name"`
	Distribution string   `json:"distribution"`
	Slug         string   `json:"slug"`
	Public       bool     `json:"public"`
	Regions      []string `json:"regions"`
	CreatedAt    string   `json:"created_at"`
}

type DropletSize struct {
	Slug         string  `json:"slug"`
	Transfer     int     `json:"transfer"`
	PriceMonthly float64 `json:"price_monthly"`
	PriceHourly  float64 `json:"price_hourly"`
}

type DropletNetworks struct {
	IPv4 []*DropletIPv4Network
	IPv6 []*DropletIPv6Network
}

type DropletIPv4Network struct {
	IpAddress string `json:"ip_address"`
	Netmask   string `json:"netmask"`
	Gateway   string `json:"gateway"`
	Type      string `json:"type"`
}

type DropletIPv6Network struct {
	IpAddress string `json:"ip_address"`
	Cidr      int    `json:"cidr"`
	Gateway   string `json:"gateway"`
	Type      string `json:"type"`
}

type DropletKernel struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Version string `json:"version"`
}

type DropletFeatures struct {
	Features []string `json:"features"`
}

type DropletBabckupIds struct {
	BackupIds []int `json:"backup_ids"`
}

type DropletSnapshotIds struct {
	SnapshotIds []int `json:"snapshot_ids"`
}
