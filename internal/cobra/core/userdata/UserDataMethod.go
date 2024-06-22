package userdata

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"gitlab.iblog.pro/cobra/metadata/internal/cobra/core/serverobject"
)

func CreateUserDataObject(server *serverobject.Server) *UserDataObject {

        userDataObject := UserDataObject{
  Hostname: string(server.Hostname),
  ManageEtcHosts: bool(false),
  PreserveHostname: bool(false),
  FQDN: string(""),
  Users: []UserDataUser{
    UserDataUser{
      Name: server.ServerUserName,
      Sudo: string("ALL=(ALL) NOPASSWD:ALL"),
      Groups: string("users, admin"),
      Home: fmt.Sprintf("/home/%s",server.ServerUserName),
      Shell: string("/bin/bash"),
      LockPasswd: bool(false),
      SSHAuthorizedKeys: []string{},
    },
  },
  SSHPwauth: bool(false),
  DisableRoot: bool(false),
  Apt: UserDataApt{
    SourcesList: string(""),
  },
  CACerts: UserDataCACerts{
    Trusted: []string{},
  },
  Chpasswd: UserDataChpasswd{
    List: fmt.Sprintf("%s:%s\nroot:%s\n", server.ServerUserName, server.ServerUserPassword, server.ServerRootPassword),
    Expire: bool(false),
  },
  WriteFiles: []UserDataWriteFile{},
  NTP: UserDataNTP{
    Enabled: bool(false),
    NTPClient: string(""),
    Servers: []string{},
  },
  Growpart: UserDataGrowpart{
    Mode: string(""),
    Devices: []string{},
  },
  TimeZone: server.TimeZone,
  PackageUpdate: bool(false),
  PackageUpgrade: bool(false),
  Packages: []string{},
  Output: UserDataOutput{
    All: string(">> /var/log/cloud-init.log"),
  },
  Runcmd: []string{
    string("systemctl disable systemd-udevd.service"),
    string("ssh-keygen -A"),
    string("ssh-keygen -t rsa -b 4096 -f /root/.ssh/id_rsa  -q -P \"\""),
    string("ssh-keygen -t rsa -b 4096 -f /home/vorlon/.ssh/id_rsa  -q -P \"\""),
    string("echo \"LANGUAGE=ru_RU:ru\" > /etc/default/locale"),
    string("echo \"LANG=ru_RU.UTF-8\" >> /etc/default/locale"),
    string("systemctl disable apt-news.service"),
    string("systemctl disable esm-cache.service"),
    string("systemctl stop apt-news.service"),
    string("systemctl stop esm-cache.service"),
    string("apt remove snapd -y"),
    string("apt update"),
  },
  FinalMessage: string(""),
  PowerState: UserDataPowerState{
    Delay: string(""),
    Mode: string(""),
    Message: string(""),
    Timeout: int64(0),
  },
}

        return &userDataObject

}


func RenderUserDataObject(userDataObject *UserDataObject) (*string, error) {

        yamlObject, err := yaml.Marshal(&userDataObject)
        if err != nil {
                return nil, err
        }

        userDataString := string(yamlObject)
        userDataString = fmt.Sprintf("#cloud-config\n\n%s\n", userDataString)
        return &userDataString, nil

}
