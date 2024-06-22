package ginrouter

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/penglongli/gin-metrics/ginmetrics"

	"github.com/sirupsen/logrus"
	logs "gitlab.iblog.pro/cobra/metadata/internal/cobra/logs"

	"gitlab.iblog.pro/cobra/metadata/internal/cobra/core/utils"
	"gitlab.iblog.pro/cobra/metadata/internal/cobra/core/storesingletonmetadata"
	"gitlab.iblog.pro/cobra/metadata/internal/cobra/core/storesingletonopenstack"
)


func Ping(c *gin.Context) {

	utils.RequestDump(c, "Ping")

	c.String(http.StatusOK, "pong")
}


func Method(c *gin.Context) {
	method := c.Params.ByName("method")
	// c.String(http.StatusOK, fmt.Sprintf("Method(): Run:%s", method))

        var response string
	responseCode := http.StatusOK

        switch method {
        case "openstack":
		response = storesingletonopenstack.InstanceOpenstack.GetIndex(c)
	case "latest":
		response = storesingletonmetadata.InstanceMetaData.GetIndex(c)
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
                        response = storesingletonopenstack.InstanceOpenstack.GetSubFolderIndex(c)
                case "latest":
                        response = storesingletonopenstack.InstanceOpenstack.GetSubFolderIndex(c)
		default:
			response = fmt.Sprintf("SubMethod not found\n")
		}
	case "latest":
		switch submethod {
		case "meta-data":
			response = storesingletonmetadata.InstanceMetaData.GetSubFolderIndex(c)
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

	utils.RequestDump(c, "MethodSubmethodFunction")

        var response string
	response = ""
	responseCode := http.StatusOK

        switch method {
        case "openstack":
                switch submethod {
                case "2018-08-27", "latest":
			switch function {
			case "vendor_data2.json":
				response = storesingletonopenstack.InstanceOpenstack.GetVendorData2(c)
			case "vendor_data.json":
				response = storesingletonopenstack.InstanceOpenstack.GetVendorData(c)
			case "password":
				response = storesingletonopenstack.InstanceOpenstack.GetPassword(c)
			case "network_data.json":
				response = storesingletonopenstack.InstanceOpenstack.GetNetworkData(c)
			case "meta_data.json":
				MetaData, err := storesingletonopenstack.InstanceOpenstack.GetMetaData(c)
				if err != nil {
					logs.Log.WithFields(logrus.Fields{ "Error": fmt.Sprintf("%v", err), }).Info("/openstack/[2018-08-27|latest]/meta_data.json")
					responseCode = http.StatusInternalServerError
				}
				response = *MetaData
			case "user_data":
				UserData, err := storesingletonopenstack.InstanceOpenstack.GetUserData(c)
				if err != nil {
					logs.Log.WithFields(logrus.Fields{ "Error": fmt.Sprintf("%v", err), }).Info("/openstack/[2018-08-27|latest]/meta_data.json")
					responseCode = http.StatusInternalServerError
				}
				fmt.Printf("%v\n", *UserData)
				response = *UserData
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
				response = storesingletonmetadata.InstanceMetaData.GetAmiId(c)
			case "ami-launch-index":
				response = storesingletonmetadata.InstanceMetaData.GetAmiLaunchIndex(c)
			case "ami-manifest-path":
				response = storesingletonmetadata.InstanceMetaData.GetAmiManifestPath(c)
			case "hostname":
				response = storesingletonmetadata.InstanceMetaData.GetHostname(c)
			case "instance-action":
				response = storesingletonmetadata.InstanceMetaData.GetInstanceAction(c)
			case "instance-type":
				response = storesingletonmetadata.InstanceMetaData.GetInstanceType(c)
			case "instance-id":
				response = storesingletonmetadata.InstanceMetaData.GetInstanceId(c)
			case "local-hostname":
				response = storesingletonmetadata.InstanceMetaData.GetLocalHostname(c)
			case "local-ipv4":
				response = storesingletonmetadata.InstanceMetaData.GetLocalIpv4(c)
			case "public-hostname":
				response = storesingletonmetadata.InstanceMetaData.GetPublicHostname(c)
			case "public-ipv4":
				response = storesingletonmetadata.InstanceMetaData.GetPublicIpv4(c)
			case "reservation-id":
				response = storesingletonmetadata.InstanceMetaData.GetReservationId(c)
			case "placement":
				response = storesingletonmetadata.InstanceMetaData.GetPlacementIndex(c)
			case "block-device-mapping":
				response = storesingletonmetadata.InstanceMetaData.GetBlockDeviceMappingIndex(c)
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
					response = storesingletonmetadata.InstanceMetaData.GetPlacementAvailabilityZone(c)
				default:
					response = fmt.Sprintf("SubFunction not found %s\n",subfunction)
				}
                        case "block-device-mapping":
				switch subfunction {
				case "ami":
					response = storesingletonmetadata.InstanceMetaData.GetBlockDeviceMappingAmi(c)
				case "ebs0":
					response = storesingletonmetadata.InstanceMetaData.GetBlockDeviceMappingEbs0(c)
				case "root":
					response = storesingletonmetadata.InstanceMetaData.GetBlockDeviceMappingRoot(c)
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

func SetupRouter() *gin.Engine {

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


