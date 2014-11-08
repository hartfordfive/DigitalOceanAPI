package digitalocean

type Images struct {
	Images []Image      `json:"images"`
	Meta   DropletsMeta `json:"meta"`
}

type Image struct {
	Id           int      `json:"id"`
	Name         string   `json:"name"`
	Distribution string   `json:"distribution"`
	Slug         string   `json:"slug"`
	Public       bool     `json:"public"`
	Regions      []Region `json:"regions"`
}
