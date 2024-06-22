package main

import (
	"net/http"
	//"os"
	//"io"
	"fmt"
	"sync"
	"github.com/gin-gonic/gin"
	"github.com/penglongli/gin-metrics/ginmetrics"
)

///****************************************///

type SingletonMetaData struct {
 	name string
}

var (
	instanceMetaData *SingletonMetaData
	once     sync.Once
)

func GetInstanceMetaData() *SingletonMetaData {
    once.Do(func() {
        instanceMetaData = &SingletonMetaData{name: "Safe Golang Singleton"}
    })
    return instanceMetaData
}

func (s *SingletonMetaData) GetIndex() string {
    return "SingletonMetaData:GetIndex"
}

func (s *SingletonMetaData) GetSubFolderIndex() string {
    return `
ami-id
ami-launch-index
ami-manifest-path
block-device-mapping/
hostname
instance-action
instance-id
instance-type
local-hostname
local-ipv4
placement/
public-hostname
public-ipv4
reservation-id`

}

func (s *SingletonMetaData) GetAmiId() string {
    return "None"
}

func (s *SingletonMetaData) GetAmiLaunchIndex() string {
    return "0"
}

func (s *SingletonMetaData) GetAmiManifestPath() string {
    return "FIXME"
}

func (s *SingletonMetaData) GetHostname() string {
    return "node170-c56b4d.cloud.local"
}

func (s *SingletonMetaData) GetInstanceAction() string {
    return "none"
}

func (s *SingletonMetaData) GetInstanceType() string {
    return "Flavor cloudstack-Flavor-uuid"
}

func (s *SingletonMetaData) GetInstanceId() string {
    return "cloudstack-instance-uuid"
}

func (s *SingletonMetaData) GetLocalHostname() string {
    return "node170-c56b4d.cloud.local"
}

func (s *SingletonMetaData) GetLocalIpv4() string {
    return "192.168.200.170"
}

func (s *SingletonMetaData) GetPublicHostname() string {
    return "node170-c56b4d.cloud.local"
}

func (s *SingletonMetaData) GetPublicIpv4() string {
    return ""
}

func (s *SingletonMetaData) GetReservationId() string {
    return "cloudstatck-uuid-reservation"
}

func (s *SingletonMetaData) GetPlacementIndex() string {
    return `availability-zone`
}

func (s *SingletonMetaData) GetBlockDeviceMappingIndex() string {
    return `ami
ebs0
root`
}


func (s *SingletonMetaData) GetPlacementAvailabilityZone() string {
    return "nova"
}

func (s *SingletonMetaData) GetBlockDeviceMappingAmi() string {
    return "sda"
}

func (s *SingletonMetaData) GetBlockDeviceMappingEbs0() string {
    return "/dev/sda"
}

func (s *SingletonMetaData) GetBlockDeviceMappingRoot() string {
    return "/dev/sda"
}


///****************************************///

type SingletonOpenstack struct {
	name string
}

var (
	instanceOpenstack *SingletonOpenstack
)

func GetInstanceOpenstack() *SingletonOpenstack {
    once.Do(func() {
        instanceOpenstack = &SingletonOpenstack{name: "Safe Golang Singleton"}
    })
    return instanceOpenstack
}

func (s *SingletonOpenstack) GetIndex() string {
    return "SingletonOpenstack:GetIndex"
}

func (s *SingletonOpenstack) GetSubFolderIndex() string {
    return `2018-08-27
latest`
}

func (s *SingletonOpenstack) GetVendorData2() string {
    return `{
  "static": {}
}`
}

func (s *SingletonOpenstack) GetVendorData() string {
    return "{}"
}

func (s *SingletonOpenstack) GetPassword() string {
    return ""
}

func (s *SingletonOpenstack) GetNetworkData() string {
    return `
{
  "links": [
    {
      "id": "node170-16",
      "vif_id": "cloudstack-uuid-vif_id-1",
      "type": "ovs",
      "mtu": null,
      "ethernet_mac_address": "fa:16:3e:86:09:64"
    },{
      "id": "node170-17",
      "vif_id": "cloudstack-uuid-vif_id-2",
      "type": "ovs",
      "mtu": null,
      "ethernet_mac_address": "fa:16:3e:23:4e:24"
    },{
        "id": "vlan200",
        "type": "vlan",
        "vlan_link": "node170-16",
        "vlan_id": 200,
        "vlan_mac_address": "fa:16:3e:86:09:64",
        "vif_id": "cloudstack-uuid-vif_id-200"
    },{
        "id": "vlan400",
        "type": "vlan",
        "vlan_link": "node170-16",
        "vlan_id": 400,
        "vlan_mac_address": "fa:16:3e:86:09:64",
        "vif_id": "cloudstack-uuid-vif_id-400"
    },{
        "id": "vlan600",
        "type": "vlan",
        "vlan_link": "node170-16",
        "vlan_id": 600,
        "vlan_mac_address": "fa:16:3e:86:09:64",
        "vif_id": "cloudstack-uuid-vif_id-600"
    },{
        "id": "vlan800",
        "type": "vlan",
        "vlan_link": "node170-16",
        "vlan_id": 800,
        "vlan_mac_address": "fa:16:3e:86:09:64",
        "vif_id": "cloudstack-uuid-vif_id-800"
    }
  ],
  "networks": [
    {
      "id": "network0",
      "type": "ipv4_dhcp",
      "link": "node170-16",
      "network_id": "cloudstack-uuid-network_id",
      "network_name": "Network cloudstack-uuid-network_id"
    },{
      "id": "network1",
      "type": "ipv4_dhcp",
      "link": "node170-17",
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

func (s *SingletonOpenstack) GetMetaData() string {
    return `{
  "uuid": "CLOUDSTACK-uuid",
  "hostname": "node170.cloud.local",
  "name": "Vm Node170",
  "launch_index": 0,
  "availability_zone": "nova",
  "random_seed": "<..................>",
  "project_id": "CLOUDSTACK-project-uuid",
  "devices": [],
  "dedicated_cpus": []
}`

}

func (s *SingletonOpenstack) GetUserData() string {
    return `#cloud-config

groups:
- admingroup:
  - root
  - sys
- cloud-users
hostname: node170-c56b4d
users:
- groups: users
  lock_passwd: false
  name: vorlon
  passwd: <................>
  primary_group: vorlon
  sudo: ALL=(ALL) NOPASSWD:ALL
`
}


///****************************************///

var db = make(map[string]string)

func Ping (c *gin.Context) {
	fmt.Printf("%#v",c.Request)
	c.String(http.StatusOK, "pong")
}


func Method(c *gin.Context) {
	method := c.Params.ByName("method")
	// c.String(http.StatusOK, fmt.Sprintf("Method(): Run:%s", method))

        var response string

        switch method {
        case "openstack":
		response = instanceOpenstack.GetIndex()
	case "latest":
		response = instanceMetaData.GetIndex()
        default:
                response = fmt.Sprintf("Method not found\n")
        }
        c.String(http.StatusOK, response)

}


func MethodSubmethod(c *gin.Context) {
	method := c.Params.ByName("method")
	submethod := c.Params.ByName("submethod")
	// c.String(http.StatusOK, fmt.Sprintf("MethodSubmethod(): Run:%s, %s", method, submethod))

	var response string

        switch method {
        case "openstack":
                switch submethod {
                case "2018-08-27":
                        response = instanceOpenstack.GetSubFolderIndex()
                case "latest":
                        response = instanceOpenstack.GetSubFolderIndex()
		default:
			response = fmt.Sprintf("SubMethod not found\n")
		}
	case "latest":
		switch submethod {
		case "meta-data":
			response = instanceMetaData.GetSubFolderIndex()
		default:
			response = fmt.Sprintf("SubMethod not found\n")
		}
	default:
                response = fmt.Sprintf("Method not found\n")
        }
	c.String(http.StatusOK, response)

}


func MethodSubmethodFunction(c *gin.Context) {
	method := c.Params.ByName("method")
	submethod := c.Params.ByName("submethod")
	function := c.Params.ByName("function")

	//c.String(http.StatusOK, fmt.Sprintf("MethodSubmethodFunction(): Run:%s, %s, %s", method, submethod, function))

        var response string
	response = ""

        switch method {
        case "openstack":
                switch submethod {
                case "2018-08-27", "latest":
			switch function {
			case "vendor_data2.json":
				response = instanceOpenstack.GetVendorData2()
			case "vendor_data.json":
				response = instanceOpenstack.GetVendorData()
			case "password":
				response = instanceOpenstack.GetPassword()
			case "network_data.json":
				response = instanceOpenstack.GetNetworkData()
			case "meta_data.json":
				response = instanceOpenstack.GetMetaData()
			case "user_data":
				response = instanceOpenstack.GetUserData()
			default:
				response = fmt.Sprintf("Function not found\n")
			}
                default:
                        response = fmt.Sprintf("SubMethod not found\n")
                }

        case "latest":
                switch submethod {
                case "meta-data":
			switch function {
			case "ami-id":
				response = instanceMetaData.GetAmiId()
			case "ami-launch-index":
				response = instanceMetaData.GetAmiLaunchIndex()
			case "ami-manifest-path":
				response = instanceMetaData.GetAmiManifestPath()
			case "hostname":
				response = instanceMetaData.GetHostname()
			case "instance-action":
				response = instanceMetaData.GetInstanceAction()
			case "instance-type":
				response = instanceMetaData.GetInstanceType()
			case "instance-id":
				response = instanceMetaData.GetInstanceId()
			case "local-hostname":
				response = instanceMetaData.GetLocalHostname()
			case "local-ipv4":
				response = instanceMetaData.GetLocalIpv4()
			case "public-hostname":
				response = instanceMetaData.GetPublicHostname()
			case "public-ipv4":
				response = instanceMetaData.GetPublicIpv4()
			case "reservation-id":
				response = instanceMetaData.GetReservationId()
			case "placement":
				response = instanceMetaData.GetPlacementIndex()
			case "block-device-mapping":
				response = instanceMetaData.GetBlockDeviceMappingIndex()
			default:
				response = fmt.Sprintf("Function not found\n")
			}
                default:
                        response = fmt.Sprintf("SubMethod not found\n")
		}
        default:
                response = fmt.Sprintf("Method not found\n")
        }

        c.String(http.StatusOK, response)


}


func MethodSubmethodFunctionSubFunction(c *gin.Context) {
	method := c.Params.ByName("method")
	submethod := c.Params.ByName("submethod")
	function := c.Params.ByName("function")
	subfunction := c.Params.ByName("subfunction")
	//c.String(http.StatusOK, fmt.Sprintf("MethodSubmethodFunctionSubFunction(): Run:%s, %s, %s, %s", method, submethod, function, subfunction))

        var response string
        response = ""

        switch method {
        case "latest":
                switch submethod {
                case "meta-data":
                        switch function {
                        case "placement":
				switch subfunction {
				case "availability-zone":
					response = instanceMetaData.GetPlacementAvailabilityZone()
				default:
					response = fmt.Sprintf("SubFunction not found %s\n",subfunction)
				}
                        case "block-device-mapping":
				switch subfunction {
				case "ami":
					response = instanceMetaData.GetBlockDeviceMappingAmi()
				case "ebs0":
					response = instanceMetaData.GetBlockDeviceMappingEbs0()
				case "root":
					response = instanceMetaData.GetBlockDeviceMappingRoot()
				default:
					response = fmt.Sprintf("SubFunction not found %s\n",subfunction)
				}
                        default:
                                response = fmt.Sprintf("Function not found\n")
                	}
        default:
                response = fmt.Sprintf("Method not found\n")
        }
        c.String(http.StatusOK, response)
	}
}

func setupRouter() *gin.Engine {

	// Logging to a file.

	//f, _ := os.Create("gin2.log")
	//gin.DisableConsoleColor()
	//gin.DefaultWriter = io.MultiWriter(f)

	r := gin.Default()


	m := ginmetrics.GetMonitor()

	// +optional set metric path, default /debug/metrics
	m.SetMetricPath("/metrics")
	// +optional set slow time, default 5s
	m.SetSlowTime(10)
	// +optional set request duration, default {0.1, 0.3, 1.2, 5, 10}
	// used to p95, p99
	m.SetDuration([]float64{0.1, 0.3, 1.2, 5, 10})

	// set middleware for gin
	m.Use(r)


	// Ping test
	r.GET("/ping", Ping)

	r.GET("/:method", Method)
        r.GET("/:method/", Method)

	r.GET("/:method/:submethod", MethodSubmethod)
	r.GET("/:method/:submethod/", MethodSubmethod)
	r.GET("/:method/:submethod/:function", MethodSubmethodFunction)
        r.GET("/:method/:submethod/:function/", MethodSubmethodFunction)
	r.GET("/:method/:submethod/:function/:subfunction", MethodSubmethodFunctionSubFunction)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	return r
}

func main() {

	_ = GetInstanceOpenstack()
	_ = GetInstanceMetaData()

	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":80")
}
