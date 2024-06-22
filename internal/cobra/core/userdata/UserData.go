package userdata

///****************************************///

type UserDataObject struct {
        Hostname         string                         `json:"hostname"`
        ManageEtcHosts   bool                           `json:"manage_etc_hosts"`
        PreserveHostname bool                           `json:"preserve_hostname"`
        FQDN             string                         `json:"fqdn"`
        Users            []UserDataUser                 `json:"users"`
        SSHPwauth        bool                           `json:"ssh_pwauth"`
        DisableRoot      bool                           `json:"disable_root"`
        Apt              UserDataApt                    `json:"apt"`
        CACerts          UserDataCACerts                `json:"ca-certs"`
        Chpasswd         UserDataChpasswd               `json:"chpasswd"`
        WriteFiles       []UserDataWriteFile            `json:"write_files"`
        NTP              UserDataNTP                    `json:"ntp"`
        Growpart         UserDataGrowpart               `json:"growpart"`
        TimeZone         string                         `json:"timezone"`
        PackageUpdate    bool                           `json:"package_update"`
        PackageUpgrade   bool                           `json:"package_upgrade"`
        Packages         []string                       `json:"packages"`
        Output           UserDataOutput                 `json:"output"`
        Runcmd           []string                       `json:"runcmd"`
        FinalMessage     string                         `json:"final_message"`
        PowerState       UserDataPowerState             `json:"power_state"`
}

type UserDataApt struct {
        SourcesList string `json:"sources_list"`
}

type UserDataCACerts struct {
        Trusted []string `json:"trusted"`
}

type UserDataChpasswd struct {
        List   string `json:"list"`
        Expire bool   `json:"expire"`
}

type UserDataGrowpart struct {
        Mode    string   `json:"mode"`
        Devices []string `json:"devices"`
}

type UserDataNTP struct {
        Enabled   bool     `json:"enabled"`
        NTPClient string   `json:"ntp_client"`
        Servers   []string `json:"servers"`
}

type UserDataOutput struct {
        All string `json:"all"`
}

type UserDataPowerState struct {
        Delay   string `json:"delay"`
        Mode    string `json:"mode"`
        Message string `json:"message"`
        Timeout int64  `json:"timeout"`
}

type UserDataUser struct {
        Name              string   `json:"name"`
        Sudo              string   `json:"sudo"`
        Groups            string   `json:"groups"`
        Home              string   `json:"home"`
        Shell             string   `json:"shell"`
        LockPasswd        bool     `json:"lock_passwd"`
        SSHAuthorizedKeys []string `json:"ssh-authorized-keys"`
}


type UserDataWriteFile struct {
        Content string `json:"content"`
        Path    string `json:"path"`
}


