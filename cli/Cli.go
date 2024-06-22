package main

import (

	"gitlab.iblog.pro/cobra/metadata/internal/cobra/core/serverobject"

	"gitlab.iblog.pro/cobra/metadata/internal/cobra/core/store"
	"gitlab.iblog.pro/cobra/metadata/internal/cobra/core/storesingletonmetadata"
	"gitlab.iblog.pro/cobra/metadata/internal/cobra/core/storesingletonopenstack"
	"gitlab.iblog.pro/cobra/metadata/internal/cobra/core/ginrouter"
	"gitlab.iblog.pro/cobra/metadata/internal/cobra/core/utils"

        PanicRecover "gitlab.iblog.pro/cobra/metadata/internal/cobra/core/panicrecover"
        "gitlab.iblog.pro/cobra/metadata/internal/cobra/system"

)

///****************************************///

func GetNetworkData() string {
    return `
{
  "links": [
    {
      "id": "node161-16",
      "vif_id": "cloudstack-uuid-vif_id-1",
      "type": "ovs",
      "mtu": null,
      "ethernet_mac_address": "fa:16:3e:6a:3c:d6"
    },{
      "id": "node161-17",
      "vif_id": "cloudstack-uuid-vif_id-2",
      "type": "ovs",
      "mtu": null,
      "ethernet_mac_address": "fa:16:3e:12:01:73"
    },{
        "id": "vlan200",
        "type": "vlan",
        "vlan_link": "node161-16",
        "vlan_id": 200,
        "vlan_mac_address": "fa:16:3e:6a:3c:d6",
        "vif_id": "cloudstack-uuid-vif_id-200"
    },{
        "id": "vlan400",
        "type": "vlan",
        "vlan_link": "node161-16",
        "vlan_id": 400,
        "vlan_mac_address": "fa:16:3e:6a:3c:d6",
        "vif_id": "cloudstack-uuid-vif_id-400"
    },{
        "id": "vlan600",
        "type": "vlan",
        "vlan_link": "node161-16",
        "vlan_id": 600,
        "vlan_mac_address": "fa:16:3e:6a:3c:d6",
        "vif_id": "cloudstack-uuid-vif_id-600"
    },{
        "id": "vlan800",
        "type": "vlan",
        "vlan_link": "node161-16",
        "vlan_id": 800,
        "vlan_mac_address": "fa:16:3e:6a:3c:d6",
        "vif_id": "cloudstack-uuid-vif_id-800"
    }
  ],
  "networks": [
    {
      "id": "network0",
      "type": "ipv4_dhcp",
      "link": "node161-16",
      "network_id": "cloudstack-uuid-network_id",
      "network_name": "Network cloudstack-uuid-network_id"
    },{
      "id": "network1",
      "type": "ipv4_dhcp",
      "link": "node161-17",
      "network_id": "cloudstack-uuid-network_id",
      "network_name": "Network cloudstack-uuid-network_id"
    },{
        "id": "publicnet-ipv4",
        "type": "ipv4_dhcp",
        "link": "vlan200",
        "network_id": "cloudstack-uuid-network_id",
        "network_name": "Network cloudstack-uuid-network_id"
    },{
        "id": "publicnet-ipv4",
        "type": "ipv4_dhcp",
        "link": "vlan400",
        "network_id": "cloudstack-uuid-network_id",
        "network_name": "Network cloudstack-uuid-network_id"
    },{
        "id": "publicnet-ipv4",
        "type": "ipv4_dhcp",
        "link": "vlan600",
        "network_id": "cloudstack-uuid-network_id",
        "network_name": "Network cloudstack-uuid-network_id"
    },{
        "id": "publicnet-ipv4",
        "type": "ipv4_dhcp",
        "link": "vlan800",
        "network_id": "cloudstack-uuid-network_id",
        "network_name": "Network cloudstack-uuid-network_id"
    }
  ],
  "services": [
    {
      "type": "dns",
      "address": "8.8.8.8"
    }
  ]
}`

}



func GetServer() (*serverobject.Server, error) {

        NetworkData := GetNetworkData()
        NetworkAddress := "192.168.100.161"
        RandomSeed, err := utils.GetRandomSeed()

        if err != nil {
                return nil, err
        }

        server := serverobject.Server{
                                UUID: "CLOUDSTACK-uuid",
                                NetworkData: NetworkData,
                                NetworkAddress: NetworkAddress,
                                RandomSeed: *RandomSeed,
                                LaunchIndex: 0,
                                AvailabilityZone: "nova",
                                ProjectID: "CLOUDSTACK-project-uuid",
                                Flavor: "Flavor cloudstack-Flavor-uuid",
                                Nodename: "node161-c56b4d.cloud.local",
                                Name: "node161-c56b4d.cloud.local",
                                Hostname: "node161-c56b4d",
                                BootDisk: "sda",
                                BootDevice: "/dev/sda",
				ServerUserName: "vorlon",
			        ServerUserPassword: "123",
        			ServerRootPassword: "root",
				TimeZone: "Asia/Yekaterinburg",
                        }

	return &server, nil
}

///****************************************///

func main() {

        defer PanicRecover.PanicRecover()

        system.VersionBuild(Version,"Metadata")


	_ = storesingletonopenstack.GetInstanceOpenstack()
	_ = storesingletonmetadata.GetInstanceMetaData()


	store.Servers["192.168.100.161"], _ = GetServer()
        store.Servers["192.168.100.1"], _ = GetServer()

	r := ginrouter.SetupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":80")
}
