package storesingletonopenstack

import (
	"gitlab.iblog.pro/cobra/metadata/internal/cobra/core/singletonopenstack"
	"gitlab.iblog.pro/cobra/metadata/internal/cobra/core/storesingletonmetadata"
)

///****************************************///

var (
        InstanceOpenstack *singletonopenstack.SingletonOpenstack
)

func GetInstanceOpenstack() *singletonopenstack.SingletonOpenstack {
    storesingletonmetadata.Once.Do(func() {
        instanceOpenstack := singletonopenstack.SingletonOpenstack{Name: "Safe Golang Singleton"}
	InstanceOpenstack = &instanceOpenstack
    })
    return InstanceOpenstack
}
