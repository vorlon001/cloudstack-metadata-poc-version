package serverobject

type LinkServer struct {
        UUID            string   `json:"uuid"`
        EthernetMACAddress string `json:"ethernet_mac_address"`
}

type Server struct {
        Nodename         string      `json:"Nodename"`
        Flavor           string      `json:"Flavor"`
        BootDisk         string      `json:"boot_disk"`
        BootDevice       string      `json:"boot_device"`
        AvailabilityZone string      `json:"availability_zone"`
        UUID             string      `json:"uuid"`
        Hostname         string      `json:"hostname"`
        Name             string      `json:"name"`
        LaunchIndex      int64       `json:"launch_index"`
        RandomSeed       string      `json:"random_seed"`
        ProjectID        string      `json:"project_id"`
        NetworkData      string      `json:"network_data"`
        NetworkAddress   string      `json:"network_address"`

	ServerUserName	     	string      `json:"server_user_name"`
	ServerUserPassword   	string      `json:"server_user_password"`
	ServerRootPassword   	string      `json:"server_root_name"`
	TimeZone		string      `json:"timezone"`


        LinksServer      *LinkServer
}

