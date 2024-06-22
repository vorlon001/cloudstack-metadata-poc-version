package system

import (
	"fmt"
	"github.com/sirupsen/logrus"
        logs "gitlab.iblog.pro/cobra/metadata/internal/cobra/logs"
)

func VersionBuild(version string, service string) {
	logs.Log.WithFields(logrus.Fields{ "Version": version,}).Info("https://github.com/vorlon001, (C) Vorlon001")
	logs.Log.Info(fmt.Sprintf("HomeLabs CloudStack %s:service (Golang version)", service))
}

