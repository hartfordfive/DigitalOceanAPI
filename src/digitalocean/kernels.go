package digitalocean

type Kernels struct {
	Kernels []DropletKernel `json:"kernels"`
	Meta    DropletsMeta    `json:"meta"`
}
