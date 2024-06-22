package panicrecover

import (

        "encoding/json"
//        "os"
	"fmt"
        "strings"
        "runtime"
        "github.com/sirupsen/logrus"
        logs "gitlab.iblog.pro/cobra/metadata/internal/cobra/logs"
)


func PanicRecover() {
    if r := recover(); r != nil {
        logrus.Error("Internal error: %v", r)
        buf := make([]byte, 1<<16)
        stackSize := runtime.Stack(buf, true)
        logs.Log.Error("--------------------------------------------------------------------------------")
        logs.Log.Error(fmt.Sprintf("Internal error: %s\n", string(buf[0:stackSize])))
        logs.Log.Error("--------------------------------------------------------------------------------")

        }
}



func Herr(e error) {
        if e != nil {
                logs.Log.WithFields(logrus.Fields{ "err": strings.ReplaceAll(e.Error(), "\"", ""), }).Info("herr")
//                os.Exit(1)
        }
}

func Hok(message string) {
        logs.Log.WithFields(logrus.Fields{ "ok": strings.ReplaceAll(message, "\"", ""), }).Info("hok")
}

func Hret(i interface{}) {
        ret, err := json.Marshal(i)
        Herr(err)
        logs.Log.WithFields(logrus.Fields{ "ret": string(ret), }).Info("Inside hret Run")
}
