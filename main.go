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
    return "SingletonMetaData:GetSubFolderIndex"
}

func (s *SingletonMetaData) GetAmiId() string {
    return "SingletonMetaData:GetAmiId"
}

func (s *SingletonMetaData) GetAmiLaunchIndex() string {
    return "SingletonMetaData:GetAmiLaunchIndex"
}

func (s *SingletonMetaData) GetAmiManifestPath() string {
    return "SingletonMetaData:GetAmiManifestPath"
}

func (s *SingletonMetaData) GetHostname() string {
    return "SingletonMetaData:GetHostname"
}

func (s *SingletonMetaData) GetInstanceAction() string {
    return "SingletonMetaData:GetInstanceAction"
}

func (s *SingletonMetaData) GetInstanceType() string {
    return "SingletonMetaData:GetInstanceType"
}

func (s *SingletonMetaData) GetInstanceId() string {
    return "SingletonMetaData:GetInstanceId"
}

func (s *SingletonMetaData) GetLocalHostname() string {
    return "SingletonMetaData:GetLocalHostname"
}

func (s *SingletonMetaData) GetLocalIpv4() string {
    return "SingletonMetaData:GetLocalIpv4"
}

func (s *SingletonMetaData) GetPublicHostname() string {
    return "SingletonMetaData:GetPublicHostname"
}

func (s *SingletonMetaData) GetPublicIpv4() string {
    return "SingletonMetaData:GetPublicIpv4"
}

func (s *SingletonMetaData) GetReservationId() string {
    return "SingletonMetaData:GetReservationId"
}

func (s *SingletonMetaData) GetPlacementIndex() string {
    return "SingletonMetaData:GetPlacementIndex"
}

func (s *SingletonMetaData) GetBlockDeviceMappingIndex() string {
    return "SingletonMetaData:GetBlockDeviceMappingIndex"
}


func (s *SingletonMetaData) GetPlacementAvailabilityZone() string {
    return "SingletonMetaData:GetPlacementAvailabilityZone"
}

func (s *SingletonMetaData) GetBlockDeviceMappingAmi() string {
    return "SingletonMetaData:GetBlockDeviceMappingAmi"
}

func (s *SingletonMetaData) GetBlockDeviceMappingEbs0() string {
    return "SingletonMetaData:GetBlockDeviceMappingEbs0"
}

func (s *SingletonMetaData) GetBlockDeviceMappingRoot() string {
    return "SingletonMetaData:GetBlockDeviceMappingRoot"
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
    return "SingletonOpenstack:GetSubFolderIndex"
}

func (s *SingletonOpenstack) GetVendorData2() string {
    return "SingletonOpenstack:GetVendorData2"
}

func (s *SingletonOpenstack) GetVendorData() string {
    return "SingletonOpenstack:GetVendorData"
}

func (s *SingletonOpenstack) GetPassword() string {
    return "SingletonOpenstack:GetPassword"
}

func (s *SingletonOpenstack) GetNetworkData() string {
    return "SingletonOpenstack:GetNetworkData"
}

func (s *SingletonOpenstack) GetMetaData() string {
    return "SingletonOpenstack:GetMetaData"
}

func (s *SingletonOpenstack) GetUserData() string {
    return "SingletonOpenstack:GetUseraData"
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
	r.Run(":18080")
}
