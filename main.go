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
		Get a list of active droplets
	*******************************************/
	status_code, droplets := do.GetDroplets()
	fmt.Println("Status Code: ", status_code)
	fmt.Println("Droplets: \n", droplets)
	fmt.Println("Droplet #1 ID: ", droplets.DropletList[0].Id)
	fmt.Println("-------------------------------------------")

	/******************************************
		Get a list of available kernels
	*******************************************/
	/*
		_, kernels := do.GetKernels(1495765)
		fmt.Println("Kernels: \n", kernels)
		fmt.Println("-------------------------------------------")
	*/

	/******************************************
		Get
	*******************************************/

	_, test := do.Test()
	fmt.Println("Test: \n", test)
	fmt.Println("-------------------------------------------")

}
