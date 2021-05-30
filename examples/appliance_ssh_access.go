package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/HewlettPackard/oneview-golang/ov"
)

func main() {
	var (
		ClientOV *ov.OVClient
	)
	apiversion, _ := strconv.Atoi(os.Getenv("ONEVIEW_APIVERSION"))
	ovc := ClientOV.NewOVClient(
		os.Getenv("ONEVIEW_OV_USER"),
		os.Getenv("ONEVIEW_OV_PASSWORD"),
		os.Getenv("ONEVIEW_OV_DOMAIN"),
		os.Getenv("ONEVIEW_OV_ENDPOINT"),
		false,
		apiversion,
		"*")
	getsshaccess, err := ovc.GetSshAccess()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("#--- Got the Appliance SSH access ---#")
		fmt.Println(getsshaccess)
	}

	sshaccess := ov.ApplianceSshAccess{
		allowSshAccess: false
	}
	fmt.Println(sshAccess)
	err := ovc.SetSshAccess(sshaccess)
	if err != nil {
		fmt.Println("Appliance SSH access set failed: ", err)
	} else {
		fmt.Println("Appliance SSH access set successfully...")
	}

}
