package storesingletonmetadata

import (
	"sync"
	"gitlab.iblog.pro/cobra/metadata/internal/cobra/core/singletonmetadata"
)

var (
        InstanceMetaData *singletonmetadata.SingletonMetaData
        Once     sync.Once
)



func GetInstanceMetaData() *singletonmetadata.SingletonMetaData {
    Once.Do(func() {
        instanceMetaData := singletonmetadata.SingletonMetaData{Name: "Safe Golang Singleton"}
	InstanceMetaData = &instanceMetaData
    })
    return InstanceMetaData
}
