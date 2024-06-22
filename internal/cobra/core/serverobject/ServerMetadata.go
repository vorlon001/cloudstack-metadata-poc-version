package serverobject

import (
	"encoding/json"
)

type ServerMetaData struct {
	AvailabilityZone string      	`json:"availability_zone"`
	UUID             string      	`json:"uuid"`
	Hostname         string      	`json:"hostname"`
	Name             string      	`json:"name"`
	LaunchIndex      int64       	`json:"launch_index"`
	RandomSeed       string      	`json:"random_seed"`
	ProjectID        string    	`json:"project_id"`
	Devices		 []string	`json:"devices"`
	DedicatedCpus	 []string	`json:"dedicated_cpus"`
}

func CreateGetMetaData(server *Server) *ServerMetaData {
	return &ServerMetaData{
			AvailabilityZone: server.AvailabilityZone,
			UUID: server.UUID,
			Hostname: server.Hostname,
			Name: server.Name,
			LaunchIndex: server.LaunchIndex,
			RandomSeed: server.RandomSeed,
			ProjectID:  server.ProjectID,
			Devices: []string{},
			DedicatedCpus: []string{},
	}
}

func (serverMetaData *ServerMetaData) RenderGetMetaData() (*string, error) {
	u, err := json.Marshal(serverMetaData)
        if err != nil {
            return nil, err
        }

	metaData := string(u)
	return &metaData, nil
}

