package utils

import (

	"fmt"
        "strings"
	"crypto/rand"
	b64 "encoding/base64"

        "github.com/sirupsen/logrus"
        logs "gitlab.iblog.pro/cobra/metadata/internal/cobra/logs"

	"github.com/gin-gonic/gin"
)

///****************************************///


func GetIP(ipport string) string {

        ipport_match := strings.Split(ipport, ":")
        return ipport_match[0]

}


func RequestDump(c *gin.Context, method string) {

        logs.Log.WithFields(logrus.Fields{ "Request.Method": fmt.Sprintf("%v",c.Request.Method), }).Info(method)
        logs.Log.WithFields(logrus.Fields{ "Request.Header[User-Agent]": fmt.Sprintf("%v",c.Request.Header["User-Agent"]), }).Info(method)
        logs.Log.WithFields(logrus.Fields{ "Request.Host": fmt.Sprintf("%v",c.Request.Host), }).Info(method)
        logs.Log.WithFields(logrus.Fields{ "Request.RemoteAddr": fmt.Sprintf("%v",c.Request.RemoteAddr), }).Info(method)
        logs.Log.WithFields(logrus.Fields{ "Request.RequestURI": fmt.Sprintf("%v",c.Request.RequestURI), }).Info(method)

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
