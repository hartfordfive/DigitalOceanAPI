package digitalocean

type Regions struct {
	Regions []Region     `json:"region"`
	Meta    DropletsMeta `json:"meta"`
}

type Region struct {
	Id           int      `json:"id"`
	Name         string   `json:"name"`
	Distribution string   `json:"distribution"`
	Slug         string   `json:"slug"`
	Public       bool     `json:"public"`
	Regions      []Region `json:"regions"`
}
