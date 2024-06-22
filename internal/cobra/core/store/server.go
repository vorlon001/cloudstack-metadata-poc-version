package store

import (
        "gitlab.iblog.pro/cobra/metadata/internal/cobra/core/serverobject"
)
var Servers map[string]*serverobject.Server

func init () {
        Servers = make(map[string]*serverobject.Server)
}
