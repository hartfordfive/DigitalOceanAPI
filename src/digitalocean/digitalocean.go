package digitalocean

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	//"strconv"
	//"strings"
	"time"
)

const (
	API_BASE = "https://api.digitalocean.com/v2/"
)

type DigitalOceanClient struct {
	ApiToken           string
	Methods            map[string][]string
	RateLimitMax       int
	RateLimitRemaining int
	RateLimitResetTime time.Time
}

func NewClient(api_token string) (c *DigitalOceanClient) {
	return &DigitalOceanClient{
		ApiToken: api_token,
		Methods: map[string][]string{
			"action_history": {"GET", "actions"},
			"action_details": {"GET", "actions/%d"},

			"domain_records":       {"GET", "domains/%s/records"},
			"add_domain_record":    {"POST", "domains/%s/records"},
			"get_domain_record":    {"GET", "domains/%s/records/%d"},
			"delete_domain_record": {"DELETE", "domains/%s/records/%d"},
			"update_domain_record": {"PUT", "domains/%s/records/%d"},

			"list_domains":  {"GET", "domains"},
			"create_domain": {"POST", "domains"},
			"get_domain":    {"GET", "domains/%s"},
			"delete_domain": {"DELETE", "domains/%s"},

			"show_droplet":           {"GET", "droplets/%d"},
			"list_droplets":          {"GET", "droplets"},
			"list_droplet_kernels":   {"GET", "droplets/%d/kernels"},
			"list_droplet_snapshots": {"GET", "droplets/%d/snapshots"},
			"list_droplet_backups":   {"GET", "droplets/%d/backups"},
			"list_droplet_actions":   {"GET", "droplets/%d/actions"},
			"create_droplet":         {"POST", "droplets"},
			"delete_droplet":         {"DELETE", "droplets/%d"},
			"transfer_droplet":       {"POST", "images/%d/actions"},

			"perform_droplet_action": {"POST", "droplets/%d/actions"},
		},
	}
}

func (c *DigitalOceanClient) doRequest(method string, request_headers map[string]string, params ...interface{}) (status string, headers http.Header, resp_body []byte) {

	var req *http.Request

	fmt.Println("DEBUG METHOD: ", method)

	num_params := len(params)
	url := ""

	switch {

	case num_params == 0:

		url = API_BASE + c.Methods[method][1]
		req, _ = http.NewRequest(c.Methods[method][0], url, bytes.NewBuffer([]byte(``)))

	case num_params == 1 && method == "action_details":

		url = API_BASE + fmt.Sprintf(c.Methods[method][1], params[0].(int))
		req, _ = http.NewRequest(c.Methods[method][0], url, bytes.NewBuffer([]byte(``)))

	case num_params == 1 && method == "domain_records":

		url = API_BASE + fmt.Sprintf(c.Methods[method][1], params[0].(string))
		req, _ = http.NewRequest(c.Methods[method][0], url, bytes.NewBuffer([]byte(``)))

	case num_params == 1 && method == "add_domain_record":

		jsonStr, err := json.Marshal(params[1])
		if err != nil {
			fmt.Println("JSON Encoding Error:", err)
			return "000", nil, []byte(`{}`)
		}

		url = API_BASE + fmt.Sprintf(c.Methods[method][1], params[0].(string))
		req, _ = http.NewRequest(c.Methods[method][0], url, bytes.NewBuffer(jsonStr))

	case num_params == 2 && method == "get_domain_record":

		url = API_BASE + fmt.Sprintf(c.Methods[method][1], params[0].(string), params[1].(int))
		req, _ = http.NewRequest(c.Methods[method][0], url, bytes.NewBuffer([]byte(``)))

	case num_params == 2 && method == "delete_domain_record":

		url = API_BASE + fmt.Sprintf(c.Methods[method][1], params[0].(string), params[1].(int))
		req, _ = http.NewRequest(c.Methods[method][0], url, bytes.NewBuffer([]byte(``)))

	case num_params == 1 && method == "update_domain_record":

		jsonStr, err := json.Marshal(params[2])
		if err != nil {
			fmt.Println("JSON Encoding Error:", err)
			return "000", nil, []byte(`{}`)
		}
		url = API_BASE + fmt.Sprintf(c.Methods[method][1], params[0].(string), params[1].(int))
		req, _ = http.NewRequest(c.Methods[method][0], url, bytes.NewBuffer(jsonStr))

	case num_params == 1 && method == "list_domains":

		url = API_BASE + fmt.Sprintf(c.Methods[method][1], params[0].(string))
		req, _ = http.NewRequest(c.Methods[method][0], url, bytes.NewBuffer([]byte(``)))

	case num_params == 1 && method == "create_domain":

		jsonStr, err := json.Marshal(params[0])
		if err != nil {
			fmt.Println("JSON Encoding Error:", err)
			return "000", nil, []byte(`{}`)
		}
		url = API_BASE + c.Methods[method][1]
		req, _ = http.NewRequest(c.Methods[method][0], url, bytes.NewBuffer(jsonStr))

	case num_params == 1 && method == "get_domain":

		url = API_BASE + fmt.Sprintf(c.Methods[method][1], params[0].(string))
		req, _ = http.NewRequest(c.Methods[method][0], url, bytes.NewBuffer([]byte(``)))

	case num_params == 1 && method == "list_droplet_kernels":

		url = API_BASE + fmt.Sprintf(c.Methods[method][1], params[0].(int))
		req, _ = http.NewRequest(c.Methods[method][0], url, bytes.NewBuffer([]byte(``)))

	case num_params == 2 && method == "perform_droplet_action":

		jsonStr, err := json.Marshal(map[string]string{"type": params[1].(string)})
		if err != nil {
			fmt.Println("JSON Encoding Error:", err)
			return "000", nil, []byte(`{}`)
		}

		url = API_BASE + fmt.Sprintf(c.Methods[method][1], params[0].(int))
		req, _ = http.NewRequest(c.Methods[method][0], url, bytes.NewBuffer(jsonStr))

	default:
		fmt.Println("\t*** Method unimplemented***")
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
	return resp.Status, resp.Header, body
}

func (c *DigitalOceanClient) GetActionHistory() (status string, resp_headers http.Header, result Actions) {

	token := fmt.Sprintf("Bearer %s", c.ApiToken)
	headers := map[string]string{"Authorization": token}
	status, resp_headers, body := c.doRequest("action_history", headers)

	var json_resp Actions
	json.Unmarshal(body, &json_resp)

	return status, resp_headers, json_resp
}

func (c *DigitalOceanClient) GetAction(action_id int) (status string, resp_headers http.Header, result Action) {

	token := fmt.Sprintf("Bearer %s", c.ApiToken)
	headers := map[string]string{"Authorization": token}
	status, resp_headers, body := c.doRequest("action_details", headers, action_id)

	var json_resp Action
	json.Unmarshal(body, &json_resp)

	return status, resp_headers, json_resp
}

// -------------------------------------------------------------

func (c *DigitalOceanClient) GetDomainRecords(domain string) (status string, resp_headers http.Header, result DomainRecords) {

	token := fmt.Sprintf("Bearer %s", c.ApiToken)
	headers := map[string]string{"Authorization": token}
	status, resp_headers, body := c.doRequest("domain_records", headers, domain)

	var json_resp DomainRecords
	json.Unmarshal(body, &json_resp)

	return status, resp_headers, json_resp
}

func (c *DigitalOceanClient) CreateDomainRecord(r DomainRecord) (status string, resp_headers http.Header, result DomainRecords) {

	token := fmt.Sprintf("Bearer %s", c.ApiToken)
	headers := map[string]string{"Authorization": token, "Content-Type": "application/json"}
	status, resp_headers, body := c.doRequest("domain_records", headers)

	var json_resp DomainRecords
	json.Unmarshal(body, &json_resp)

	return status, resp_headers, json_resp
}

/*
func (c *DigitalOceanClient) GetDomainRecord(domain string, id int) (status string, resp_headers http.Header, result DomainRecords) {

	token := fmt.Sprintf("Bearer %s", c.ApiToken)
	headers := map[string]string{"Authorization": token}
	status, resp_headers, body := c.doRequest("get_domain_record", headers, domain, id)

	var json_resp DomainRecord
	json.Unmarshal(body, &json_resp)

	return status, resp_headers, json_resp
}

func (c *DigitalOceanClient) DeleteDomainRecord(domain string, id int) () {

	token := fmt.Sprintf("Bearer %s", c.ApiToken)
	headers := map[string]string{"Authorization": token, "Content-Type": "application/x-www-form-urlencoded"}
	status, resp_headers, body := c.doRequest("get_domain_record", headers, domain, id)

	var json_resp DomainRecord
	json.Unmarshal(body, &json_resp)

	return status, resp_headers, json_resp
}
*/

// -------------------------------------------------------------

func (c *DigitalOceanClient) GetDomains() (status string, resp_headers http.Header, result Domains) {

	token := fmt.Sprintf("Bearer %s", c.ApiToken)
	headers := map[string]string{"Authorization": token}
	status, resp_headers, body := c.doRequest("list_domains", headers)

	var json_resp Domains
	json.Unmarshal(body, &json_resp)

	return status, resp_headers, json_resp
}

func (c *DigitalOceanClient) CreateDomain(domain NewDomain) (status string, resp_headers http.Header, result Domain) {

	token := fmt.Sprintf("Bearer %s", c.ApiToken)
	headers := map[string]string{"Authorization": token, "Content-Type": "application/json"}
	status, resp_headers, body := c.doRequest("create_domain", headers, domain)

	var json_resp Domain
	json.Unmarshal(body, &json_resp)

	return status, resp_headers, json_resp
}

func (c *DigitalOceanClient) GetDomain(domain string) (status string, resp_headers http.Header, result Domain) {

	token := fmt.Sprintf("Bearer %s", c.ApiToken)
	headers := map[string]string{"Authorization": token}
	status, resp_headers, body := c.doRequest("get_domain", headers, domain)

	var json_resp Domain
	json.Unmarshal(body, &json_resp)

	return status, resp_headers, json_resp
}

func (c *DigitalOceanClient) DeleteDomain(domain string) (status string, resp_headers http.Header) {

	token := fmt.Sprintf("Bearer %s", c.ApiToken)
	headers := map[string]string{"Authorization": token, "Content-Type": "application/x-www-form-urlencoded"}
	status, resp_headers, _ = c.doRequest("delete_domain", headers, domain)

	return status, resp_headers
}

// -------------------------- DROPLET ACTIONS -----------------------------------

func (c *DigitalOceanClient) PerformDropletAction(droplet_id int, action string) (status string, resp_headers http.Header, result DropletAction) {

	valid_actions := []string{
		"disable_babckups", "reboot", "powercycle",
		"shutdown", "poweroff", "poweron",
		"", "",
	}

	token := fmt.Sprintf("Bearer %s", c.ApiToken)
	headers := map[string]string{"Authorization": token, "Content-Type": "application/json"}
	status, resp_headers, body := c.doRequest("perform_droplet_action", headers, droplet_id, action)

	var json_resp DropletAction
	json.Unmarshal(body, &json_resp)

	return status, resp_headers, json_resp
}

// -------------------------------------------------------------

func (c *DigitalOceanClient) GetDroplets() (status string, resp_headers http.Header, result Droplets) {

	token := fmt.Sprintf("Bearer %s", c.ApiToken)
	headers := map[string]string{"Authorization": token}
	status, resp_headers, body := c.doRequest("list_droplets", headers)

	var json_resp Droplets
	json.Unmarshal(body, &json_resp)

	return status, resp_headers, json_resp
}

// -------------------------------------------------------------

func (c *DigitalOceanClient) GetKernels(droplet_id int) (status string, resp_headers http.Header, result Kernels) {

	token := fmt.Sprintf("Bearer %s", c.ApiToken)
	headers := map[string]string{"Authorization": token}
	status, resp_headers, body := c.doRequest("list_droplet_kernels", headers, droplet_id)

	var json_resp Kernels
	json.Unmarshal(body, &json_resp)

	return status, resp_headers, json_resp
}
