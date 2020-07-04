package main

import (
	"fmt"
	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/HewlettPackard/oneview-golang/utils"
	_ "os"
	"path/filepath"
)

func main() {
	var (
		ClientOV    *ov.OVClient
		testName    = "TestFCNetworkGOsdk"
		new_fc_name = "RenamedFCNetwork"
		falseVar    = false
	)
	config, err1 := ov.LoadConfigFile("oneview_config.json")
	if err1 != nil {
		fmt.Println(err1)
	}
	ovc := ClientOV.NewOVClient(
		config.UserName,
		config.Password,
		config.Domain,
		config.Endpoint,
		config.SSlVerify,
		config.ApiVersion,
		config.IfMatch)

	/*	apiversion, _ := strconv.Atoi(os.Getenv("ONEVIEW_APIVERSION"))
		ovc := ClientOV.NewOVClient(
			os.Getenv("ONEVIEW_OV_USER"),
			os.Getenv("ONEVIEW_OV_PASSWORD"),
			os.Getenv("ONEVIEW_OV_DOMAIN"),
			os.Getenv("ONEVIEW_OV_ENDPOINT"),
			false,
			apiversion,
			"*")*/
	initialScopeUris := &[]utils.Nstring{utils.NewNstring("/rest/scopes/7e4f76b0-bb2c-49d2-a641-d785475df423")}
	fcNetwork := ov.FCNetwork{
		AutoLoginRedistribution: falseVar,
		Description:             "Test FC Network",
		LinkStabilityTime:       30,
		FabricType:              "FabricAttach",
		Name:                    testName,
		Type:                    "fc-networkV4",    //The Type value is for API>500.
		InitialScopeUris:        *initialScopeUris, //added for API>500
	}
	fmt.Println(fcNetwork)
	err := ovc.CreateFCNetwork(fcNetwork)
	if err != nil {
		fmt.Println("Fc Network Creation Failed: ", err)
	} else {
		fmt.Println("Fc Network created successfully...")
	}

	sort := "name:desc"
	fcNetworks, err := ovc.GetFCNetworks("", sort, "", "")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("#---Get Fc Networks sorted by name in descending order----#")
		for i := range fcNetworks.Members {
			fmt.Println(fcNetworks.Members[i].Name)
		}
	}
	fcNetwork2, err := ovc.GetFCNetworkByName(testName)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("#-------------Get FCNetworks by name----------------#")
		fmt.Println(fcNetwork2)
	}
	fcNetwork2.Name = new_fc_name
	err = ovc.UpdateFcNetwork(fcNetwork2)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("FCNetwork has been updated with name: " + fcNetwork2.Name)
	}
	err = ovc.DeleteFCNetwork(new_fc_name)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Deleted FCNetworks successfully...")
	}

}
