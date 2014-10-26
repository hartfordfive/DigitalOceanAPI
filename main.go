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
	_, actions := do.GetActionHistory()
	displayMethodInfo("Action History", map[string]interface{}{"history": actions})

	/******************************************
		Get details for a specific action ID
	*******************************************/
	//fmt.Println("\tAction ID:", actions.Actions[0].Id)
	_, action := do.GetAction(actions.Actions[0].Id)
	displayMethodInfo("Action Details", map[string]interface{}{"history": action})

	/******************************************
		Get list of domain records
	*******************************************/
	_, domain_records := do.GetDomainRecords("whatsmydevice.mobi")
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
		_, new_domain_record := do.CreateDomainRecord(dr)
		displayMethodInfo("New Domain Record", map[string]interface{}{"domain records": new_domain_record})
	*/

	/******************************************
		Get a list of active droplets
	*******************************************/
	status_code, droplets := do.GetDroplets()
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

	_, kernels := do.GetKernels(int(droplets.DropletList[0].Id))
	displayMethodInfo("Available Kernels", map[string]interface{}{"kernels": kernels})

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
