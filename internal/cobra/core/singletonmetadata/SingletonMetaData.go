package singletonmetadata

import (
        "github.com/gin-gonic/gin"
	"gitlab.iblog.pro/cobra/metadata/internal/cobra/core/utils"
	"gitlab.iblog.pro/cobra/metadata/internal/cobra/core/store"
)

///****************************************///

type SingletonMetaData struct {
 	Name string
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
	return store.Servers[utils.GetIP(c.Request.RemoteAddr)].Hostname
}

func (s *SingletonMetaData) GetInstanceAction(c *gin.Context) string {
	return "none"
}

func (s *SingletonMetaData) GetInstanceType(c *gin.Context) string {
	return store.Servers[utils.GetIP(c.Request.RemoteAddr)].Flavor
}

func (s *SingletonMetaData) GetInstanceId(c *gin.Context) string {
	return store.Servers[utils.GetIP(c.Request.RemoteAddr)].UUID
}

func (s *SingletonMetaData) GetLocalHostname(c *gin.Context) string {
	return store.Servers[utils.GetIP(c.Request.RemoteAddr)].Hostname
}

func (s *SingletonMetaData) GetLocalIpv4(c *gin.Context) string {
	return store.Servers[utils.GetIP(c.Request.RemoteAddr)].NetworkAddress
}

func (s *SingletonMetaData) GetPublicHostname(c *gin.Context) string {
	return store.Servers[utils.GetIP(c.Request.RemoteAddr)].Hostname
}

func (s *SingletonMetaData) GetPublicIpv4(c *gin.Context) string {
    return ""
}

func (s *SingletonMetaData) GetReservationId(c *gin.Context) string {
	return store.Servers[utils.GetIP(c.Request.RemoteAddr)].UUID
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
	return store.Servers[utils.GetIP(c.Request.RemoteAddr)].AvailabilityZone
}

func (s *SingletonMetaData) GetBlockDeviceMappingAmi(c *gin.Context) string {
	return store.Servers[utils.GetIP(c.Request.RemoteAddr)].BootDisk
}

func (s *SingletonMetaData) GetBlockDeviceMappingEbs0(c *gin.Context) string {
	return store.Servers[utils.GetIP(c.Request.RemoteAddr)].BootDevice
}

func (s *SingletonMetaData) GetBlockDeviceMappingRoot(c *gin.Context) string {
	return store.Servers[utils.GetIP(c.Request.RemoteAddr)].BootDevice
}
