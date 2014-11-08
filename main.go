package main

import (
	"digitalocean"
	"fmt"
	"github.com/kr/pretty"
	"utils"
)

func main() {

	conf := &utils.Config{}
	utils.LoadConfig("settings.conf", conf)
	do := digitalocean.NewClient(conf.ApiToken)

	/******************************************
		Get Action History
	*******************************************/
	status_code, headers, actions := do.GetActionHistory()
	displayMethodInfo("Action History", map[string]interface{}{
		"status code:":      status_code,
		"response headers:": headers,
		"history:":          actions,
	})

	/******************************************
		Get details for a specific action ID
	*******************************************/
	status_code, headers, action := do.GetAction(actions.Actions[0].Id)
	displayMethodInfo("Action Details", map[string]interface{}{
		"status code:":      status_code,
		"response headers:": headers,
		"action details":    action,
	})

	/******************************************
		Get list of domain records
	*******************************************/
	_, _, domain_records := do.GetDomainRecords("whatsmydevice.mobi")
	displayMethodInfo("Domain Records List", map[string]interface{}{"domain records": domain_records})

	/******************************************
		Create a new domain record
	*******************************************/
	/*
		dr := &digitalocean.DomainRecord{
			Name: "devlocal",
			Data: "127.0.0.1",
			Type: "A",
		}
		_, _, new_domain_record := do.CreateDomainRecord(dr)
		displayMethodInfo("New Domain Record", map[string]interface{}{"domain records": new_domain_record})
	*/

	/******************************************
		Delete a domain record
	*******************************************/
	/*
		_, _, deleted_record := do.DeleteDomainRecord("dev", 2456134)
		displayMethodInfo("Deleted Domain Record", map[string]interface{}{"deleted domain record": deleted_record})
	*/

	/******************************************
		Update a domain record
	*******************************************/
	/*
		updated_dr := &digitalocean.DomainRecord{
			Name: "devlocal",
			Data: "127.0.0.10",
			Type: "A",
		}
			_, _, updated_record := do.UpdateDomainRecord(updated_dr)
			displayMethodInfo("Updated Domain Record", map[string]interface{}{"updated domain record": updated_record})
	*/

	/******************************************
		Get list of domains
	*******************************************/
	status_code, headers, domains := do.GetDomains()
	displayMethodInfo("List Domains", map[string]interface{}{
		"status code:":      status_code,
		"response headers:": headers,
		"domains":           domains,
	})

	/******************************************
		Create a domain
	*******************************************/
	/*
		domain_to_create := &digitalocean.NewDomain{
			Name:      "yourdomain.com",
			IpAddress: "192.168.1.2",
		}

		status_code, headers, new_domain := do.CreateDomain(*domain_to_create)
		displayMethodInfo("Create Domain", map[string]interface{}{
			"status code:":      status_code,
			"response headers:": headers,
			"domain creation":   new_domain,
		})
	*/

	/******************************************
		Get domain details
	*******************************************/
	status_code, headers, domain := do.GetDomain("yourdomain.com")
	displayMethodInfo("Domain Details", map[string]interface{}{
		"status code:":      status_code,
		"response headers:": headers,
		"domain":            domain,
	})

	/******************************************
		Delete domain
	*******************************************/
	/*
		status_code, headers = do.DeleteDomain("yourdomain.com")
		displayMethodInfo("Delete Domain", map[string]interface{}{
			"status code:":      status_code,
			"response headers:": headers,
		})
	*/

	/******************************************
		Get a list of active droplets
	*******************************************/
	status_code, _, droplets := do.GetDroplets()
	displayMethodInfo("Active Droplet List", map[string]interface{}{
		"Response Status Code": status_code,
		"Total Droplets":       len(droplets.DropletList),
		"Droplet #1 ID":        droplets.DropletList[0].Id,
		"Droplet #1 Name":      droplets.DropletList[0].Name,
		"Droplet #1 Memory":    droplets.DropletList[0].Memory,
		"Droplet #1 VCPUs":     droplets.DropletList[0].Vcpus,
	},
	)

	/******************************************
		Get a list of available kernels
	*******************************************/

	_, _, kernels := do.GetKernels(int(droplets.DropletList[0].Id))
	displayMethodInfo("Available Kernels", map[string]interface{}{"kernels": kernels})

	/******************************************
		Get a list of available kernels
	*******************************************/

	_, _, action_result, err := do.PerformDropletAction(int(droplets.DropletList[0].Id), "power_cycle")
	if err != nil {
		fmt.Println("Error performing droplet action:", err)
	} else {
		displayMethodInfo("Droplet Action - Powercycle", map[string]interface{}{"action power cycle": action_result})
	}
}

func displayMethodInfo(title string, data map[string]interface{}) {
	fmt.Println("\n----------------------------------------")
	fmt.Printf("           %s \n", title)
	fmt.Println("----------------------------------------")
	for k, v := range data {
		fmt.Printf("%s:\n", k)
		fmt.Printf("%# v\n", pretty.Formatter(v))
	}
	fmt.Println("----------------------------------------\n")
}
