package main

import (
	"digitalocean"
	"fmt"
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
	fmt.Println("Action History: \n", actions)
	fmt.Println("-------------------------------------------")

	/******************************************
		Get details for a specific action ID
	*******************************************/
	_, action := do.GetAction(actions.Actions[0].Id)
	fmt.Println("Action History: \n", action)
	fmt.Println("-------------------------------------------")

	/******************************************
		Get a list of active droplets
	*******************************************/
	status_code, droplets := do.GetDroplets()
	fmt.Println("Status Code: ", status_code)
	fmt.Println("Total Droplets: \n", len(droplets.DropletList))
	fmt.Println("Droplet #1 ID: ", droplets.DropletList[0].Id)
	fmt.Println("Droplet #1 Name: ", droplets.DropletList[0].Name)
	fmt.Println("Droplet #1 Memory: ", droplets.DropletList[0].Memory)
	fmt.Println("Droplet #1 VCPUs: ", droplets.DropletList[0].Vcpus)
	fmt.Println("-------------------------------------------")

	/******************************************
		Get a list of available kernels
	*******************************************/
	_, kernels := do.GetKernels(int(droplets.DropletList[0].Id))
	fmt.Println("Kernels: \n", kernels)
	fmt.Println("-------------------------------------------")

}
