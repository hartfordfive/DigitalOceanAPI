package digitalocean

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	API_BASE = "https://api.digitalocean.com/v2/"
)

type DigitalOceanClient struct {
	ApiToken string
	Methods  map[string][]string
}

func NewClient(api_token string) (c *DigitalOceanClient) {
	return &DigitalOceanClient{
		ApiToken: api_token,
		Methods: map[string][]string{
			"show_droplet":           {"GET", "droplets/[ID]"},
			"list_droplets":          {"GET", "droplets"},
			"list_dropelet_kernels":  {"GET", "droplets/[ID]/kernels"},
			"list_droplet_snapshots": {"GET", "droplets/[ID]/snapshots"},
			"list_droplet_backups":   {"GET", "droplets/[ID]/backups"},
			"list_droplet_actions":   {"GET", "droplets/[ID]/actions"},
			"create_droplet":         {"POST", "droplets"},
			"delete_droplet":         {"DELETE", "droplets/[ID]"},
			"transfer_droplet":       {"POST", "images/[ID]/actions"},
		},
	}
}

func (c *DigitalOceanClient) doRequest(method string, request_headers map[string]string, params ...string) (status string, resp_body []byte) {

	//var jsonStr = []byte(`{"title":"Buy cheese and bread for breakfast."}`)
	jsonStr := []byte(`{}`)
	if len(params) == 1 {
		req, err := http.NewRequest(c.Methods[method][0], API_BASE+strings.Replace(c.Methods[method][1], "[ID]", params[0]), bytes.NewBuffer(jsonStr))
	} else {
		req, err := http.NewRequest(c.Methods[method][0], API_BASE+c.Methods[method][1], bytes.NewBuffer(jsonStr))
	}
	for k, v := range request_headers {
		req.Header.Set(k, v)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	return resp.Status, body
}

func (c *DigitalOceanClient) GetDroplets() (status string, result Droplets) {

	token := fmt.Sprintf("Bearer %s", c.ApiToken)
	headers := map[string]string{"Authorization": token}
	status, body := c.doRequest("list_droplets", headers)

	var json_resp Droplets
	json.Unmarshal(body, &json_resp)

	return status, json_resp
}

func (c *DigitalOceanClient) GetKernels(droplet_id int) (status string, result Kernels) {

	token := fmt.Sprintf("Bearer %s", c.ApiToken)
	headers := map[string]string{"Authorization": token}
	status, body := c.doRequest("list_dropelet_kernels", headers)

	var json_resp Kernels
	json.Unmarshal(body, &json_resp)

	return status, json_resp
}

func (c *DigitalOceanClient) Test() bool {
	return true
}
