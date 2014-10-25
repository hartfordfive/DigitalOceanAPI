package digitalocean

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
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
			"action_history": {"GET", "actions"},
			"action_details": {"GET", "actions/[ID]", "int"},

			"domain_records":    {"GET", "domains/[ID]/records", "string"},
			"add_domain_record": {"POST", "domains/[DOMAIN]/records", "string"},

			"show_droplet":           {"GET", "droplets/[ID]", "int"},
			"list_droplets":          {"GET", "droplets"},
			"list_dropelet_kernels":  {"GET", "droplets/[ID]/kernels", "int"},
			"list_droplet_snapshots": {"GET", "droplets/[ID]/snapshots", "int"},
			"list_droplet_backups":   {"GET", "droplets/[ID]/backups", "int"},
			"list_droplet_actions":   {"GET", "droplets/[ID]/actions", "int"},
			"create_droplet":         {"POST", "droplets"},
			"delete_droplet":         {"DELETE", "droplets/[ID]", "int"},
			"transfer_droplet":       {"POST", "images/[ID]/actions", "int"},
		},
	}
}

func (c *DigitalOceanClient) doRequest(method string, request_headers map[string]string, params ...interface{}) (status string, resp_body []byte) {

	//var jsonStr = []byte(`{"title":"Buy cheese and bread for breakfast."}`)
	//jsonStr := []byte(`{}`)

	var req *http.Request

	num_params := len(params)
	url := ""

	switch {

	case num_params == 0:

		url = API_BASE + c.Methods[method][1]
		req, _ = http.NewRequest(c.Methods[method][0], url, bytes.NewBuffer([]byte(``)))

	case num_params == 1 && method == "action_details":

		url = API_BASE + strings.Replace(c.Methods[method][1], "[ID]", strconv.Itoa(params[0].(int)), -1)
		req, _ = http.NewRequest(c.Methods[method][0], url, bytes.NewBuffer([]byte(`{}`)))

	case num_params == 1 && method == "domain_records":

		url = API_BASE + strings.Replace(c.Methods[method][1], "[ID]", strconv.Itoa(params[0].(int)), -1)
		req, _ = http.NewRequest(c.Methods[method][0], url, bytes.NewBuffer([]byte(`{}`)))

	case num_params == 1 && method == "add_domain_record":

		jsonStr, err := json.Marshal(params[1])
		if err != nil {
			fmt.Println("JSON Encoding Error:", err)
			return "000", []byte(`{}`)
		}
		url = API_BASE + strings.Replace(c.Methods[method][1], "[DOMAIN]", strconv.Itoa(params[0].(int)), -1)
		req, _ = http.NewRequest(c.Methods[method][0], url, bytes.NewBuffer(jsonStr))

	default:
		fmt.Println("TODO")
	}

	/*
		if len(params) == 1 {

			fmt.Println("Requesting URL:")
			fmt.Println("\t", API_BASE+strings.Replace(c.Methods[method][1], "[ID]", params[0].(string), -1))
			req, _ = http.NewRequest(c.Methods[method][0], API_BASE+strings.Replace(c.Methods[method][1], "[ID]", strconv.Itoa(params[0].(int)), -1), bytes.NewBuffer(jsonStr))

		} else {

			fmt.Println("Requesting URL:")
			fmt.Println("\t", API_BASE+c.Methods[method][1])
			req, _ = http.NewRequest(c.Methods[method][0], API_BASE+c.Methods[method][1], bytes.NewBuffer(jsonStr))

		}
	*/

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

func (c *DigitalOceanClient) GetActionHistory() (status string, result Actions) {

	token := fmt.Sprintf("Bearer %s", c.ApiToken)
	headers := map[string]string{"Authorization": token}
	status, body := c.doRequest("action_history", headers)

	var json_resp Actions
	json.Unmarshal(body, &json_resp)

	return status, json_resp
}

func (c *DigitalOceanClient) GetAction(action_id int) (status string, result Action) {

	token := fmt.Sprintf("Bearer %s", c.ApiToken)
	headers := map[string]string{"Authorization": token}
	status, body := c.doRequest("action_details", headers, action_id)

	var json_resp Action
	json.Unmarshal(body, &json_resp)

	return status, json_resp
}

func (c *DigitalOceanClient) GetDomainRecords() (status string, result DomainRecords) {

	token := fmt.Sprintf("Bearer %s", c.ApiToken)
	headers := map[string]string{"Authorization": token}
	status, body := c.doRequest("domain_records", headers)

	var json_resp DomainRecords
	json.Unmarshal(body, &json_resp)

	return status, json_resp
}

func (c *DigitalOceanClient) CreateDomainRecord(r DomainRecord) (status string, result DomainRecords) {

	token := fmt.Sprintf("Bearer %s", c.ApiToken)
	headers := map[string]string{"Authorization": token}
	status, body := c.doRequest("domain_records", headers)

	var json_resp DomainRecords
	json.Unmarshal(body, &json_resp)

	return status, json_resp
}

// -------------------------------------------------------------

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
	status, body := c.doRequest("list_dropelet_kernels", headers, droplet_id)

	var json_resp Kernels
	json.Unmarshal(body, &json_resp)

	return status, json_resp
}
