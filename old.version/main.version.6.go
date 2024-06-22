package main

import (
	"net/http"
	//"os"
	//"io"
	"strings"
	"fmt"
	"sync"
        "crypto/rand"
        b64 "encoding/base64"
	"github.com/gin-gonic/gin"
	"github.com/penglongli/gin-metrics/ginmetrics"
	"encoding/json"

        "io/ioutil"
        "os"
        "log/syslog"
        logrus_syslog "github.com/sirupsen/logrus/hooks/syslog"

        "github.com/sirupsen/logrus"
        "github.com/sirupsen/logrus/hooks/writer"


)

///****************************************///

func InitLogrus() *logrus.Logger {

        log := logrus.New()
        log.SetOutput(ioutil.Discard) // Send all logs to nowhere by default

        hook, err := logrus_syslog.NewSyslogHook("", "", syslog.LOG_INFO, "")
        if err == nil {
                log.Hooks.Add(hook)
        }

        log.AddHook(&writer.Hook{ // Send logs with level higher than warning to stderr
                Writer: os.Stderr,
                LogLevels: []logrus.Level{
                        logrus.PanicLevel,
                        logrus.FatalLevel,
                        logrus.ErrorLevel,
                        logrus.WarnLevel,
                },
        })
        log.AddHook(&writer.Hook{ // Send info and debug logs to stdout
                Writer: os.Stdout,
                LogLevels: []logrus.Level{
                        logrus.InfoLevel,
                        logrus.TraceLevel,
                        logrus.DebugLevel,
                },
        })

        log.SetReportCaller(true)

        log.SetFormatter(&logrus.TextFormatter{
                ForceColors:   true,
                DisableColors: false,
                FullTimestamp: true,
        })

        log.SetLevel(logrus.TraceLevel)

        return log
}


///****************************************///

func GetIP(ipport string) string {

        ipport_match := strings.Split(ipport, ":")
        return ipport_match[0]

}


func GetRandomSeed() (*string, error) {
        b := make([]byte, 512)
        _, err := rand.Read(b)
        if err !=nil {
                return nil, err
        }
        sEnc := b64.StdEncoding.EncodeToString([]byte(b))
        return &sEnc, nil
}

type LinkServer struct {
        UUID            string   `json:"uuid"`
        EthernetMACAddress string `json:"ethernet_mac_address"`
}

type Server struct {
        Nodename         string      `json:"Nodename"`
        Flavor           string      `json:"Flavor"`
        BootDisk         string      `json:"boot_disk"`
        BootDevice       string      `json:"boot_device"`
        AvailabilityZone string      `json:"availability_zone"`
        UUID             string      `json:"uuid"`
        Hostname         string      `json:"hostname"`
        Name             string      `json:"name"`
        LaunchIndex      int64       `json:"launch_index"`
        RandomSeed       string      `json:"random_seed"`
        ProjectID        string      `json:"project_id"`
        UserData         string      `json:"user_data"`
        NetworkData      string      `json:"network_data"`
        NetworkAddress   string      `json:"network_address"`
        LinksServer      *LinkServer
}

///****************************************///


type ServerMetaData struct {
	AvailabilityZone string      	`json:"availability_zone"`
	UUID             string      	`json:"uuid"`
	Hostname         string      	`json:"hostname"`
	Name             string      	`json:"name"`
	LaunchIndex      int64       	`json:"launch_index"`
	RandomSeed       string      	`json:"random_seed"`
	ProjectID        string    	`json:"project_id"`
	Devices		 []string	`json:"devices"`
	DedicatedCpus	 []string	`json:"dedicated_cpus"`
}

func CreateGetMetaData(server *Server) *ServerMetaData {
	return &ServerMetaData{
			AvailabilityZone: server.AvailabilityZone,
			UUID: server.UUID,
			Hostname: server.Hostname,
			Name: server.Name,
			LaunchIndex: server.LaunchIndex,
			RandomSeed: server.RandomSeed,
			ProjectID:  server.ProjectID,
			Devices: []string{},
			DedicatedCpus: []string{},
	}
}

func (serverMetaData *ServerMetaData) RenderGetMetaData() (*string, error) {
	u, err := json.Marshal(serverMetaData)
        if err != nil {
            return nil, err
        }

	metaData := string(u)
	return &metaData, nil
}

func GetUserData() string {

        return `#cloud-config

groups:
- admingroup:
  - root
  - sys
- cloud-users
hostname: node161-c56b4d
users:
  - name: vorlon
    sudo: ALL=(ALL) NOPASSWD:ALL
    groups: users, admin
    home: /home/vorlon
    shell: /bin/bash
    lock_passwd: false
    passwd: $6$rounds=4096$p2NsFgexlf7XdA5f$tEz3j6fbrzFvyaBfP4c6pwESz/cO3QH.AuGIu6bKRlkqurwHVzwz/Ke4kjVf3LjGWbAkd7WMKbCYKU4P/qztA1

ssh_pwauth: true
disable_root: false

chpasswd:
  list: |
    vorlon:123
    root:root
  expire: false

timezone: Asia/Yekaterinburg
package_update: true
package_upgrade: true

output:
  all: ">> /var/log/cloud-init.log"

runcmd:
  - systemctl disable systemd-udevd.service
  - ssh-keygen -A
  - ssh-keygen -t rsa -b 4096 -f /root/.ssh/id_rsa  -q -P ""
  - ssh-keygen -t rsa -b 4096 -f /home/vorlon/.ssh/id_rsa  -q -P ""
  - echo "LANGUAGE=ru_RU:ru" > /etc/default/locale
  - echo "LANG=ru_RU.UTF-8" >> /etc/default/locale
  - systemctl disable apt-news.service
  - systemctl disable esm-cache.service
  - systemctl stop apt-news.service
  - systemctl stop esm-cache.service
  - apt remove snapd -y
  - apt update

final_message: "The system is finally up, after $UPTIME seconds"



`
}


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



func GetServer() (*Server, error) {

        UserData := GetUserData()
        NetworkData := GetNetworkData()
        NetworkAddress := "192.168.100.161"
        RandomSeed, err := GetRandomSeed()

        if err != nil {
                return nil, err
        }

        server := Server{
                                UUID: "CLOUDSTACK-uuid",
                                UserData: UserData,
                                NetworkData: NetworkData,
                                NetworkAddress: NetworkAddress,
                                RandomSeed: *RandomSeed,
                                LaunchIndex: 0,
                                AvailabilityZone: "nova",
                                ProjectID: "CLOUDSTACK-project-uuid",
                                Flavor: "Flavor cloudstack-Flavor-uuid",
                                Nodename: "node170-c56b4d.cloud.local",
                                Name: "node170-c56b4d.cloud.local",
                                Hostname: "node170-c56b4d.cloud.local",
                                BootDisk: "sda",
                                BootDevice: "/dev/sda",
                        }

        fmt.Printf("%#v\n", server)
	return &server, nil
}

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

func (s *SingletonMetaData) GetIndex(c *gin.Context) string {
    return "SingletonMetaData:GetIndex"
}

func (s *SingletonMetaData) GetSubFolderIndex(c *gin.Context) string {
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

func (s *SingletonMetaData) GetAmiId(c *gin.Context) string {
	return "None"
}

func (s *SingletonMetaData) GetAmiLaunchIndex(c *gin.Context) string {
	return "0"
}

func (s *SingletonMetaData) GetAmiManifestPath(c *gin.Context) string {
	return "FIXME"
}

func (s *SingletonMetaData) GetHostname(c *gin.Context) string {
	return servers[GetIP(c.Request.RemoteAddr)].Hostname
}

func (s *SingletonMetaData) GetInstanceAction(c *gin.Context) string {
	return "none"
}

func (s *SingletonMetaData) GetInstanceType(c *gin.Context) string {
	return servers[GetIP(c.Request.RemoteAddr)].Flavor
}

func (s *SingletonMetaData) GetInstanceId(c *gin.Context) string {
	return servers[GetIP(c.Request.RemoteAddr)].UUID
}

func (s *SingletonMetaData) GetLocalHostname(c *gin.Context) string {
	return servers[GetIP(c.Request.RemoteAddr)].Hostname
}

func (s *SingletonMetaData) GetLocalIpv4(c *gin.Context) string {
	return servers[GetIP(c.Request.RemoteAddr)].NetworkAddress
}

func (s *SingletonMetaData) GetPublicHostname(c *gin.Context) string {
	return servers[GetIP(c.Request.RemoteAddr)].Hostname
}

func (s *SingletonMetaData) GetPublicIpv4(c *gin.Context) string {
    return ""
}

func (s *SingletonMetaData) GetReservationId(c *gin.Context) string {
	return servers[GetIP(c.Request.RemoteAddr)].UUID
}

func (s *SingletonMetaData) GetPlacementIndex(c *gin.Context) string {
    return `availability-zone`
}

func (s *SingletonMetaData) GetBlockDeviceMappingIndex(c *gin.Context) string {
    return `ami
ebs0
root`
}


func (s *SingletonMetaData) GetPlacementAvailabilityZone(c *gin.Context) string {
	return servers[GetIP(c.Request.RemoteAddr)].AvailabilityZone
}

func (s *SingletonMetaData) GetBlockDeviceMappingAmi(c *gin.Context) string {
	return servers[GetIP(c.Request.RemoteAddr)].BootDisk
}

func (s *SingletonMetaData) GetBlockDeviceMappingEbs0(c *gin.Context) string {
	return servers[GetIP(c.Request.RemoteAddr)].BootDevice
}

func (s *SingletonMetaData) GetBlockDeviceMappingRoot(c *gin.Context) string {
	return servers[GetIP(c.Request.RemoteAddr)].BootDevice
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

func (s *SingletonOpenstack) GetIndex(c *gin.Context) string {
    return `2018-08-27
latest`
}

func (s *SingletonOpenstack) GetSubFolderIndex(c *gin.Context) string {
    return `2018-08-27
latest`
}

func (s *SingletonOpenstack) GetVendorData2(c *gin.Context) string {
    return `{
  "static": {}
}`
}

func (s *SingletonOpenstack) GetVendorData(c *gin.Context) string {
    return "{}"
}

func (s *SingletonOpenstack) GetPassword(c *gin.Context) string {
    return ""
}
// fa:16:3e:6a:3c:d6 fa:16:3e:12:01:73
func (s *SingletonOpenstack) GetNetworkData(c *gin.Context) string {
	return servers[GetIP(c.Request.RemoteAddr)].NetworkData
}

func (s *SingletonOpenstack) GetMetaData(c *gin.Context) (*string,error) {

	server := servers[GetIP(c.Request.RemoteAddr)]
	serverMetaData := CreateGetMetaData(server)
	renderGetMetaData, err := serverMetaData.RenderGetMetaData()
	return renderGetMetaData, err
}

func (s *SingletonOpenstack) GetUserData(c *gin.Context) string {

	return servers[GetIP(c.Request.RemoteAddr)].UserData
}


///****************************************///

//var db = make(map[string]string)

func Ping(c *gin.Context) {

	fmt.Printf("%#v\n", c.Request)
	fmt.Printf("%v\n", c.Request.Method)
	fmt.Printf("%v\n", c.Request.Header["User-Agent"])
	fmt.Printf("%v\n", c.Request.Host)
	fmt.Printf("%v\n", c.Request.RemoteAddr)
	fmt.Printf("%v\n", c.Request.RequestURI)

	c.String(http.StatusOK, "pong")
}


func Method(c *gin.Context) {
	method := c.Params.ByName("method")
	// c.String(http.StatusOK, fmt.Sprintf("Method(): Run:%s", method))

        var response string
	responseCode := http.StatusOK

        switch method {
        case "openstack":
		response = instanceOpenstack.GetIndex(c)
	case "latest":
		response = instanceMetaData.GetIndex(c)
        default:
                response = fmt.Sprintf("Method not found\n")
        }

        c.String(responseCode, response)

}


func MethodSubmethod(c *gin.Context) {
	method := c.Params.ByName("method")
	submethod := c.Params.ByName("submethod")
	// c.String(http.StatusOK, fmt.Sprintf("MethodSubmethod(): Run:%s, %s", method, submethod))

	var response string
	responseCode := http.StatusOK

        switch method {
        case "openstack":
                switch submethod {
                case "2018-08-27":
                        response = instanceOpenstack.GetSubFolderIndex(c)
                case "latest":
                        response = instanceOpenstack.GetSubFolderIndex(c)
		default:
			response = fmt.Sprintf("SubMethod not found\n")
		}
	case "latest":
		switch submethod {
		case "meta-data":
			response = instanceMetaData.GetSubFolderIndex(c)
		default:
			response = fmt.Sprintf("SubMethod not found\n")
		}
	default:
                response = fmt.Sprintf("Method not found\n")
        }
	c.String(responseCode, response)

}


func MethodSubmethodFunction(c *gin.Context) {
	method := c.Params.ByName("method")
	submethod := c.Params.ByName("submethod")
	function := c.Params.ByName("function")

	//c.String(http.StatusOK, fmt.Sprintf("MethodSubmethodFunction(): Run:%s, %s, %s", method, submethod, function))


        fmt.Printf("%v\n", c.Request.Method)
        fmt.Printf("%v\n", c.Request.Header["User-Agent"])
        fmt.Printf("%v\n", c.Request.Host)
        fmt.Printf("%v\n", c.Request.RemoteAddr)
        fmt.Printf("%v\n", c.Request.RequestURI)

        var response string
	response = ""
	responseCode := http.StatusOK

        switch method {
        case "openstack":
                switch submethod {
                case "2018-08-27", "latest":
			switch function {
			case "vendor_data2.json":
				response = instanceOpenstack.GetVendorData2(c)
			case "vendor_data.json":
				response = instanceOpenstack.GetVendorData(c)
			case "password":
				response = instanceOpenstack.GetPassword(c)
			case "network_data.json":
				response = instanceOpenstack.GetNetworkData(c)
			case "meta_data.json":
				MetaData, err := instanceOpenstack.GetMetaData(c)
				if err != nil {
					fmt.Printf("InternalServerError: \n%v\n%v\n", c, err)
					responseCode = http.StatusInternalServerError
				}
				response = *MetaData
			case "user_data":
				response = instanceOpenstack.GetUserData(c)
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
				response = instanceMetaData.GetAmiId(c)
			case "ami-launch-index":
				response = instanceMetaData.GetAmiLaunchIndex(c)
			case "ami-manifest-path":
				response = instanceMetaData.GetAmiManifestPath(c)
			case "hostname":
				response = instanceMetaData.GetHostname(c)
			case "instance-action":
				response = instanceMetaData.GetInstanceAction(c)
			case "instance-type":
				response = instanceMetaData.GetInstanceType(c)
			case "instance-id":
				response = instanceMetaData.GetInstanceId(c)
			case "local-hostname":
				response = instanceMetaData.GetLocalHostname(c)
			case "local-ipv4":
				response = instanceMetaData.GetLocalIpv4(c)
			case "public-hostname":
				response = instanceMetaData.GetPublicHostname(c)
			case "public-ipv4":
				response = instanceMetaData.GetPublicIpv4(c)
			case "reservation-id":
				response = instanceMetaData.GetReservationId(c)
			case "placement":
				response = instanceMetaData.GetPlacementIndex(c)
			case "block-device-mapping":
				response = instanceMetaData.GetBlockDeviceMappingIndex(c)
			default:
				response = fmt.Sprintf("Function not found\n")
			}
                default:
                        response = fmt.Sprintf("SubMethod not found\n")
		}
        default:
                response = fmt.Sprintf("Method not found\n")
        }

        c.String(responseCode, response)


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
					response = instanceMetaData.GetPlacementAvailabilityZone(c)
				default:
					response = fmt.Sprintf("SubFunction not found %s\n",subfunction)
				}
                        case "block-device-mapping":
				switch subfunction {
				case "ami":
					response = instanceMetaData.GetBlockDeviceMappingAmi(c)
				case "ebs0":
					response = instanceMetaData.GetBlockDeviceMappingEbs0(c)
				case "root":
					response = instanceMetaData.GetBlockDeviceMappingRoot(c)
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


var Log *logrus.Logger
var servers map[string]*Server


func init () {
	Log = InitLogrus()
	servers = make(map[string]*Server)
}
func main() {

	_ = GetInstanceOpenstack()
	_ = GetInstanceMetaData()


	servers["192.168.100.161"], _ = GetServer()
        servers["192.168.100.1"], _ = GetServer()

	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":80")
}
