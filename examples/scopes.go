package main

import (
	"fmt"
	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/HewlettPackard/oneview-golang/utils"
	"os"
	"strconv"
)

func main() {
	var (
		ClientOV    *ov.OVClient
		scp_name    = "ScopeTest"
		scp_name2   = "Auto-Scope"
		new_scope   = "new-scope"
		upd_scope   = "update-scope"
		eth_network = "Auto-ethernet_network"
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

	scope_test := ov.Scope{Name: scp_name, Description: "Test from script", Type: "ScopeV3"}
	scope_test_2 := ov.Scope{Name: scp_name2, Description: "Test from script", Type: "ScopeV3"}
	er_test := ovc.CreateScope(scope_test)
	er_test = ovc.CreateScope(scope_test_2)
	if er_test != nil {
		fmt.Println("Error Creating Scope: ", er_test)
	}

	fmt.Println("#................... Scope by Name ...............#")
	scp, scperr := ovc.GetScopeByName(scp_name)
	if scperr != nil {
		fmt.Println(scperr)
	}
	fmt.Println(scp)

	sort := "name:desc"
	scp_list, err := ovc.GetScopes("", "", "", "", sort)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("# ................... Scopes List .................#")
	for i := 0; i < len(scp_list.Members); i++ {
		fmt.Println(scp_list.Members[i].Name)
	}
	eth_uri, err := ovc.GetEthernetNetworkByName(eth_network)
	if err != nil {
		fmt.Println(err)
	}

	initialScopeUris := &[]utils.Nstring{(scp.URI)}
	addedResourceUris := &[]utils.Nstring{(eth_uri.URI)}
	scope := ov.Scope{Name: new_scope, Description: "Test from script", Type: "ScopeV3", InitialScopeUris: *initialScopeUris, AddedResourceUris: *addedResourceUris}

	er := ovc.CreateScope(scope)
	if er != nil {
		fmt.Println("............... Scope Creation Failed:", er)
	} else {
		fmt.Println("# ................... Scope Created Successfully.................#")
	}

	new_scp, err := ovc.GetScopeByName(new_scope)
	if err != nil {
		fmt.Println(err)
	} else {
		new_scp.Name = upd_scope
		err = ovc.UpdateScope(new_scp)

		if err != nil {
			fmt.Println("#.................... Scope Updation failed ...........#")
			panic(err)
		} else {
			fmt.Println("#.................... Scope after Updating ...........#")
		}
	}
	up_list, err := ovc.GetScopes("", "", "", "", sort)
	for i := 0; i < len(up_list.Members); i++ {
		fmt.Println(up_list.Members[i].Name)
	}

	scp_list, err = ovc.GetScopes("", "", "", "", sort)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("# ................... Scopes List .................#")
	for i := 0; i < len(scp_list.Members); i++ {
		fmt.Println(scp_list.Members[i].Name)
	}

	scopesInResource, err := ovc.GetScopeFromResource(eth_uri.URI.String())
	if err == nil {
		fmt.Println("#.................Scopes assigned to a resource ..............#")
		fmt.Println(scopesInResource)
	}

	scopeByUri, err := ovc.GetScopeByUri(up_list.Members[0].URI.String())
	if err == nil {
		fmt.Println("#.................Scope by Uri ..............#")
		fmt.Println(scopeByUri)
	}

	scopesInResource.ScopeUris = []string{up_list.Members[0].URI.String()}
	err = ovc.UpdateScopeForResource(scopesInResource)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("resource %s updated\n", scopesInResource.ResourceUri)
	}

	err = ovc.DeleteScope(upd_scope)
	if err != nil {
		panic(err)
	}
}
