package singletonopenstack

import (
	"github.com/gin-gonic/gin"
	"gitlab.iblog.pro/cobra/metadata/internal/cobra/core/utils"
	"gitlab.iblog.pro/cobra/metadata/internal/cobra/core/store"
	"gitlab.iblog.pro/cobra/metadata/internal/cobra/core/userdata"
	"gitlab.iblog.pro/cobra/metadata/internal/cobra/core/serverobject"
)

///****************************************///

type SingletonOpenstack struct {
	Name string
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
	return store.Servers[utils.GetIP(c.Request.RemoteAddr)].NetworkData
}

func (s *SingletonOpenstack) GetMetaData(c *gin.Context) (*string,error) {

	server := store.Servers[utils.GetIP(c.Request.RemoteAddr)]
	serverMetaData := serverobject.CreateGetMetaData(server)
	renderGetMetaData, err := serverMetaData.RenderGetMetaData()
	return renderGetMetaData, err
}

func (s *SingletonOpenstack) GetUserData(c *gin.Context) (*string, error) {
	server := store.Servers[utils.GetIP(c.Request.RemoteAddr)]
	userDataObject := userdata.CreateUserDataObject(server)
	userDataString, err := userdata.RenderUserDataObject(userDataObject)
	return userDataString, err
}


