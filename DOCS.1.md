#### STEP ZERO
Need two or more servers with network  mtu more than 4096 
urls:
- https://docs.openstack.org/nova/latest/user/metadata.html
- https://github.com/canonical/cloud-init/blob/main/cloudinit/sources/DataSourceOpenStack.py
- https://cloudinit.readthedocs.io/en/latest/reference/datasources/openstack.html
- https://specs.openstack.org/openstack/nova-specs/specs/liberty/implemented/metadata-service-network-info.html
- https://github.com/canonical/cloud-init/issues/5389


### STEP 1
Config Libvirt XML VM
```
<domain type='kvm'>
  <name>node170</name>
  <uuid>19302da1-49f7-4bf4-babb-e7d64b7f0d55</uuid>
  <metadata>
    <libosinfo:libosinfo xmlns:libosinfo="http://libosinfo.org/xmlns/libvirt/domain/1.0">
      <libosinfo:os id="http://ubuntu.com/ubuntu/20.04"/>
    </libosinfo:libosinfo>
  </metadata>
  <memory unit='KiB'>8388608</memory>
  <currentMemory unit='KiB'>8388608</currentMemory>
  <vcpu placement='static'>5</vcpu>
  <sysinfo type='smbios'>
    <system>
      <entry name='manufacturer'>OpenStack Foundation</entry>
      <entry name='product'>OpenStack Nova</entry>
      <entry name='version'>26.2.1</entry>
    </system>
  </sysinfo>
  <os>
    <type arch='x86_64' machine='pc-q35-6.0'>hvm</type>
    <boot dev='hd'/>
    <smbios mode='sysinfo'/>
  </os>
  <features>
    <acpi/>
    <apic/>
    <vmport state='off'/>
  </features>
  <cpu mode='host-model' check='partial'/>
  <clock offset='utc'>
    <timer name='rtc' tickpolicy='catchup'/>
    <timer name='pit' tickpolicy='delay'/>
    <timer name='hpet' present='no'/>
  </clock>
  <on_poweroff>destroy</on_poweroff>
  <on_reboot>restart</on_reboot>
  <on_crash>destroy</on_crash>
  <pm>
    <suspend-to-mem enabled='no'/>
    <suspend-to-disk enabled='no'/>
  </pm>
.........
```

### STEP 2
Create DHCP server, CoreDHCP
url: 
- https://github.com/vorlon001/cloudstack-coredhcp-poc-version


```

cat <<EOF>config.yaml
config:
    sw1:
        server4:
            listen:
                - "%sw1"
            plugins:
                - lease_time: 3600s
                - server_id: 192.168.1.10
                - file: file_leases_sw1.txt
                - router: 192.168.100.1
                - netmask: 255.255.255.255
                - range: leases_sw1.txt 192.168.100.10 192.168.100.50 60s file_leases_sw1.txt
                - staticroute: 192.168.100.0/24,192.168.100.1 169.254.169.254/32,192.168.100.9
    vlan200:
        server4:
            listen:
                - "%vlan200"
            plugins:
                - lease_time: 3600s
                - server_id: 192.168.1.10
                - file: file_leases_vlan200.txt
                - dns: 192.168.200.1
                - mtu: 1500
                - searchdomains: cloud.local
                - router: 192.168.200.1
                - netmask: 255.255.255.255
                - range: leases_vlan200.txt 192.168.200.10 192.168.200.50 60s file_leases_vlan200.txt
                - staticroute: 10.20.20.0/24,192.168.200.1 0.0.0.0/0,192.168.200.1 169.254.169.254/32,192.168.100.9
    vlan400:
        server4:
            listen:
                - "%vlan400"
            plugins:
                - lease_time: 3600s
                - server_id: 192.168.1.10
                - file: file_leases_vlan400.txt
                - mtu: 1500
                - router: 192.168.201.1
                - netmask: 255.255.255.255
                - range: leases_vlan400.txt 192.168.201.10 192.168.201.50 60s file_leases_vlan400.txt
                - staticroute: 10.20.21.0/24,192.168.201.1 192.168.201.0/24,192.168.201.1
    vlan600:
        server4:
            listen:
                - "%vlan600"
            plugins:
                - lease_time: 3600s
                - server_id: 192.168.1.10
                - file: file_leases_vlan600.txt
                - mtu: 1500
                - router: 192.168.202.1
                - netmask: 255.255.255.255
                - range: leases_vlan600.txt 192.168.202.10 192.168.202.50 60s file_leases_vlan600.txt
                - staticroute: 10.20.22.0/24,192.168.202.1 192.168.202.0/24,192.168.202.1

    vlan800:
        server4:
            listen:
                - "%vlan800"
            plugins:
                - lease_time: 3600s
                - server_id: 192.168.1.10
                - file: file_leases_vlan800.txt
                - mtu: 1500
                - router: 192.168.203.1
                - netmask: 255.255.255.255
                - range: leases_vlan800.txt 192.168.203.10 192.168.203.50 60s file_leases_vlan800.txt
                - staticroute: 10.20.21.0/24,192.168.203.1 192.168.203.0/24,192.168.203.1
EOF

```
### STEP 3
Create OVS Cloud
```

apt-get install openvswitch-switch -y

ovs-vsctl del-br sw1
systemctl restart ovs-vswitchd.service
systemctl restart ovsdb-server.service
ovs-vsctl add-br sw1

ovs-vsctl \
  -- add-port sw1 patch_sw1_to_sw2 \
  -- set interface patch_sw1_to_sw2 type=patch options:peer=patch_sw2_to_sw1 \
  -- add-port sw2 patch_sw2_to_sw1 \
  -- set interface patch_sw2_to_sw1 type=patch options:peer=patch_sw1_to_sw2

ovs-vsctl add-port sw2 tun_node140tonode141 -- set interface tun_<nodeOne>to<nodeTwo> type=geneve options:remote_ip=a.b.c.d   options:key=<secret> mtu_request=<mtu>
.......
ovs-vsctl show


ovs-vsctl add-port sw1 vlan200 tag=200 --\
                set interface vlan200 type=internal

ip addr add 192.168.93.10/24 dev vlan200
ip link set vlan200 up

ovs-vsctl add-port sw1 vlan400 tag=400 --\
                set interface vlan400 type=internal

ovs-vsctl set Bridge sw1 rstp_enable=true
ovs-vsctl set Bridge sw1 stp_enable=false

ovs-vsctl del-port sw1 vlan600
ovs-vsctl add-port sw1 vlan600 tag=600 --\
                set interface vlan600 type=internal mtu_request=4000


ovs-vsctl del-port sw1 vlan800
ovs-vsctl add-port sw1 vlan800 tag=800 --\
                set interface vlan800 type=internal mtu_request=4000

ip link set vlan200 up
ip link set vlan400 up
ip link set vlan600 up
ip link set vlan800 up

```
### STEP 4
Create Metadata Server
```
apt install -y nginx

cat <<EOF>/etc/netplan/50-cloud-init.yaml
network:
    ethernets:
        enp1s0:
            dhcp4: false
            dhcp6: false
            addresses:
            - 192.168.100.130/24
            - 192.168.100.9/24
            match:
                macaddress: fa:16:3e:f3:1d:5c
                name: enp*s0
            set-name: enp1s0
        enp2s0:
            dhcp4: false
            dhcp6: false
            match:
                macaddress: fa:16:3e:2f:78:e0
                name: enp*s0
            set-name: enp2s0
        lo:
            match:
                name: lo
            addresses:
                - 169.254.169.254/32
    version: 2
    vlans:
        enp1s0.200:
            addresses:
            - 192.168.200.130/24
            - 192.168.200.9/24
            dhcp4: false
            dhcp6: false
            gateway4: 192.168.200.1
            id: 200
            link: enp1s0
            nameservers:
                addresses:
                - 192.168.1.10
                search:
                - cloud.local
        enp1s0.400:
            addresses:
            - 192.168.201.130/24
            dhcp4: false
            dhcp6: false
            id: 400
            link: enp1s0
        enp1s0.600:
            addresses:
            - 192.168.202.130/24
            dhcp4: false
            dhcp6: false
            id: 600
            link: enp1s0
        enp1s0.800:
            addresses:
            - 192.168.203.130/24
            dhcp4: false
            dhcp6: false
            id: 800
            link: enp1s0
EOF




mkdir -p /var/www/html/openstack/2018-08-27/
mkdir -p /var/www/html/openstack/latest/
mkdir -p /var/www/html/latest/meta-data/

cat <<EOF>/var/www/html/openstack/index.html
2018-08-27
latest
EOF


cat <<EOF>/var/www/html/openstack/2018-08-27/vendor_data2.json
{
  "static": {}
}
EOF

cat <<EOF>/var/www/html/openstack/2018-08-27/vendor_data.json
{}
EOF

touch /var/www/html/openstack/2018-08-27/password


cat <<EOF>/var/www/html/openstack/2018-08-27/network_data.json
{
  "links": [
    {
      "id": "node170-16",
      "vif_id": "cloudstack-uuid-vif_id-1",
      "type": "ovs",
      "mtu": null,
      "ethernet_mac_address": "fa:16:3e:86:09:64"
    },{
      "id": "node170-17",
      "vif_id": "cloudstack-uuid-vif_id-2",
      "type": "ovs",
      "mtu": null,
      "ethernet_mac_address": "fa:16:3e:23:4e:24"
    },{
        "id": "vlan200",
        "type": "vlan",
        "vlan_link": "node170-16",
        "vlan_id": 200,
        "vlan_mac_address": "fa:16:3e:86:09:64",
        "vif_id": "cloudstack-uuid-vif_id-200"
    },{
        "id": "vlan400",
        "type": "vlan",
        "vlan_link": "node170-16",
        "vlan_id": 400,
        "vlan_mac_address": "fa:16:3e:86:09:64",
        "vif_id": "cloudstack-uuid-vif_id-400"
    },{
        "id": "vlan600",
        "type": "vlan",
        "vlan_link": "node170-16",
        "vlan_id": 600,
        "vlan_mac_address": "fa:16:3e:86:09:64",
        "vif_id": "cloudstack-uuid-vif_id-600"
    },{
        "id": "vlan800",
        "type": "vlan",
        "vlan_link": "node170-16",
        "vlan_id": 800,
        "vlan_mac_address": "fa:16:3e:86:09:64",
        "vif_id": "cloudstack-uuid-vif_id-800"
    }
  ],
  "networks": [
    {
      "id": "network0",
      "type": "ipv4_dhcp",
      "link": "node170-16",
      "network_id": "cloudstack-uuid-network_id",
      "network_name": "Network cloudstack-uuid-network_id"
    },{
      "id": "network1",
      "type": "ipv4_dhcp",
      "link": "node170-17",
      "network_id": "cloudstack-uuid-network_id",
      "network_name": "Network cloudstack-uuid-network_id"
    },{
        "id": "publicnet-ipv4",
        "type": "ipv4_dhcp",
        "link": "vlan200",
        "network_id": "cloudstack-uuid-network_id",
        "network_name": "Network cloudstack-uuid-network_id"
    },{
        "id": "publicnet-ipv4",
        "type": "ipv4_dhcp",
        "link": "vlan400",
        "network_id": "cloudstack-uuid-network_id",
        "network_name": "Network cloudstack-uuid-network_id"
    },{
        "id": "publicnet-ipv4",
        "type": "ipv4_dhcp",
        "link": "vlan600",
        "network_id": "cloudstack-uuid-network_id",
        "network_name": "Network cloudstack-uuid-network_id"
    },{
        "id": "publicnet-ipv4",
        "type": "ipv4_dhcp",
        "link": "vlan800",
        "network_id": "cloudstack-uuid-network_id",
        "network_name": "Network cloudstack-uuid-network_id"
    }
  ],
  "services": [
    {
      "type": "dns",
      "address": "8.8.8.8"
    }
  ]
}
EOF


cat <<EOF>/var/www/html/openstack/2018-08-27/meta_data.json
{
  "uuid": "CLOUDSTACK-uuid",
  "hostname": "node170.cloud.local",
  "name": "Vm Node170",
  "launch_index": 0,
  "availability_zone": "nova",
  "random_seed": "..........",
  "project_id": "CLOUDSTACK-project-uuid",
  "devices": [],
  "dedicated_cpus": []
}
EOF
```

```

cat <<EOF>/var/www/html/openstack/2018-08-27/user_data
#cloud-config

groups:
- admingroup:
  - root
  - sys
- cloud-users
hostname: node170-<.....>
users:
- groups: users
  lock_passwd: false
  name: vorlon
  passwd: <.............>
  primary_group: <........>
  sudo: ALL=(ALL) NOPASSWD:ALL
EOF


cat <<EOF>/var/www/html/latest/meta-data/index.html
ami-id
ami-launch-index
ami-manifest-path
block-device-mapping/
hostname
instance-action
instance-id
instance-type
local-hostname
local-ipv4
placement/
public-hostname
public-ipv4
reservation-id
EOF

cat <<EOF>/var/www/html/latest/meta-data/ami-id
None
EOF

cat <<EOF>/var/www/html/latest/meta-data/ami-launch-index
0
EOF

cat <<EOF>/var/www/html/latest/meta-data/ami-manifest-path
FIXME
EOF

cat <<EOF>/var/www/html/latest/meta-data/hostname
node170-c56b4d.cloud.local
EOF

cat <<EOF>/var/www/html/latest/meta-data/instance-action
none
EOF

cat <<EOF>/var/www/html/latest/meta-data/instance-type
Flavor cloudstack-Flavor-uuid
EOF


cat <<EOF>/var/www/html/latest/meta-data/instance-id
cloudstack-instance-uuid
EOF

cat <<EOF>/var/www/html/latest/meta-data/local-hostname
node170-c56b4d.cloud.local
EOF

cat <<EOF>/var/www/html/latest/meta-data/local-ipv4
192.168.200.170
EOF

cat <<EOF>/var/www/html/latest/meta-data/public-hostname
node170-c56b4d.cloud.local
EOF

cat <<EOF>/var/www/html/latest/meta-data/public-ipv4
EOF

cat <<EOF>/var/www/html/latest/meta-data/reservation-id
cloudstatck-uuid-reservation
EOF


mkdir -p /var/www/html/latest/meta-data/placement/

cat <<EOF>/var/www/html/latest/meta-data/placement/index.html
availability-zone
EOF

cat <<EOF>/var/www/html/latest/meta-data/placement/availability-zone
nova
EOF

mkdir -p /var/www/html/latest/meta-data/block-device-mapping

cat <<EOF>/var/www/html/latest/meta-data/block-device-mapping/index.html
ami
ebs0
root
EOF

cat <<EOF>/var/www/html/latest/meta-data/block-device-mapping/ami
sda
EOF

cat <<EOF>/var/www/html/latest/meta-data/block-device-mapping/ebs0
/dev/sda
EOF

cat <<EOF>/var/www/html/latest/meta-data/block-device-mapping/root
/dev/sda
EOF


```

### STEP 5
First run VM configs
```

cat <<EOF>/etc/netplan/50-cloud-init.yaml
# This file is generated from information provided by the datasource.  Changes
# to it will not persist across an instance reboot.  To disable cloud-init's
# network configuration capabilities, write a file
# /etc/cloud/cloud.cfg.d/99-disable-network-config.cfg with the following:
# network: {config: disabled}
network:
    version: 2
    ethernets:
        enp1s0f0:
            dhcp4: true
EOF


rm -R /var/lib/cloud/*
rm /var/log/cloud-init*
reboot or cloud-init init

```

### STEP 6 ADDON
nginx log
```
192.168.100.10 - - [16/Jun/2024:23:54:07 +0500] "GET /openstack HTTP/1.1" 301 178 "-" "Cloud-Init/24.1.3-0ubuntu3.3"
192.168.100.10 - - [16/Jun/2024:23:54:07 +0500] "GET /openstack/ HTTP/1.1" 200 100 "-" "Cloud-Init/24.1.3-0ubuntu3.3"
192.168.100.10 - - [16/Jun/2024:23:54:07 +0500] "GET /openstack HTTP/1.1" 301 178 "-" "Cloud-Init/24.1.3-0ubuntu3.3"
192.168.100.10 - - [16/Jun/2024:23:54:07 +0500] "GET /openstack/ HTTP/1.1" 200 100 "-" "Cloud-Init/24.1.3-0ubuntu3.3"
192.168.100.10 - - [16/Jun/2024:23:54:07 +0500] "GET /openstack/2018-08-27/meta_data.json HTTP/1.1" 200 934 "-" "Cloud-Init/24.1.3-0ubuntu3.3"
192.168.100.10 - - [16/Jun/2024:23:54:07 +0500] "GET /openstack/2018-08-27/user_data HTTP/1.1" 200 336 "-" "Cloud-Init/24.1.3-0ubuntu3.3"
192.168.100.10 - - [16/Jun/2024:23:54:07 +0500] "GET /openstack/2018-08-27/vendor_data.json HTTP/1.1" 200 3 "-" "Cloud-Init/24.1.3-0ubuntu3.3"
192.168.100.10 - - [16/Jun/2024:23:54:07 +0500] "GET /openstack/2018-08-27/vendor_data2.json HTTP/1.1" 200 19 "-" "Cloud-Init/24.1.3-0ubuntu3.3"
192.168.100.10 - - [16/Jun/2024:23:54:07 +0500] "GET /openstack/2018-08-27/network_data.json HTTP/1.1" 200 2543 "-" "Cloud-Init/24.1.3-0ubuntu3.3"
192.168.100.10 - - [16/Jun/2024:23:54:07 +0500] "GET /latest/meta-data/ HTTP/1.1" 200 149 "-" "Cloud-Init/24.1.3-0ubuntu3.3"
192.168.100.10 - - [16/Jun/2024:23:54:07 +0500] "GET /latest/meta-data/block-device-mapping/ HTTP/1.1" 200 14 "-" "Cloud-Init/24.1.3-0ubuntu3.3"
192.168.100.10 - - [16/Jun/2024:23:54:07 +0500] "GET /latest/meta-data/block-device-mapping/ami HTTP/1.1" 200 4 "-" "Cloud-Init/24.1.3-0ubuntu3.3"
192.168.100.10 - - [16/Jun/2024:23:54:07 +0500] "GET /latest/meta-data/block-device-mapping/ebs0 HTTP/1.1" 200 9 "-" "Cloud-Init/24.1.3-0ubuntu3.3"
192.168.100.10 - - [16/Jun/2024:23:54:07 +0500] "GET /latest/meta-data/block-device-mapping/root HTTP/1.1" 200 9 "-" "Cloud-Init/24.1.3-0ubuntu3.3"
192.168.100.10 - - [16/Jun/2024:23:54:07 +0500] "GET /latest/meta-data/placement/ HTTP/1.1" 200 18 "-" "Cloud-Init/24.1.3-0ubuntu3.3"
192.168.100.10 - - [16/Jun/2024:23:54:07 +0500] "GET /latest/meta-data/placement/availability-zone HTTP/1.1" 200 5 "-" "Cloud-Init/24.1.3-0ubuntu3.3"
192.168.100.10 - - [16/Jun/2024:23:54:07 +0500] "GET /latest/meta-data/ami-id HTTP/1.1" 200 5 "-" "Cloud-Init/24.1.3-0ubuntu3.3"
192.168.100.10 - - [16/Jun/2024:23:54:07 +0500] "GET /latest/meta-data/ami-launch-index HTTP/1.1" 200 2 "-" "Cloud-Init/24.1.3-0ubuntu3.3"
192.168.100.10 - - [16/Jun/2024:23:54:07 +0500] "GET /latest/meta-data/ami-manifest-path HTTP/1.1" 200 6 "-" "Cloud-Init/24.1.3-0ubuntu3.3"
192.168.100.10 - - [16/Jun/2024:23:54:07 +0500] "GET /latest/meta-data/hostname HTTP/1.1" 200 27 "-" "Cloud-Init/24.1.3-0ubuntu3.3"
192.168.100.10 - - [16/Jun/2024:23:54:07 +0500] "GET /latest/meta-data/instance-action HTTP/1.1" 200 5 "-" "Cloud-Init/24.1.3-0ubuntu3.3"
192.168.100.10 - - [16/Jun/2024:23:54:07 +0500] "GET /latest/meta-data/instance-id HTTP/1.1" 200 25 "-" "Cloud-Init/24.1.3-0ubuntu3.3"
192.168.100.10 - - [16/Jun/2024:23:54:07 +0500] "GET /latest/meta-data/instance-type HTTP/1.1" 200 30 "-" "Cloud-Init/24.1.3-0ubuntu3.3"
192.168.100.10 - - [16/Jun/2024:23:54:07 +0500] "GET /latest/meta-data/local-hostname HTTP/1.1" 200 27 "-" "Cloud-Init/24.1.3-0ubuntu3.3"
192.168.100.10 - - [16/Jun/2024:23:54:07 +0500] "GET /latest/meta-data/local-ipv4 HTTP/1.1" 200 16 "-" "Cloud-Init/24.1.3-0ubuntu3.3"
192.168.100.10 - - [16/Jun/2024:23:54:07 +0500] "GET /latest/meta-data/public-hostname HTTP/1.1" 200 27 "-" "Cloud-Init/24.1.3-0ubuntu3.3"
192.168.100.10 - - [16/Jun/2024:23:54:07 +0500] "GET /latest/meta-data/public-ipv4 HTTP/1.1" 200 0 "-" "Cloud-Init/24.1.3-0ubuntu3.3"
192.168.100.10 - - [16/Jun/2024:23:54:07 +0500] "GET /latest/meta-data/reservation-id HTTP/1.1" 200 29 "-" "Cloud-Init/24.1.3-0ubuntu3.3"
```

url: https://bugzilla.redhat.com/show_bug.cgi?id=1746627
cloud init log
```
2024-06-16 18:54:07,136 - util.py[DEBUG]: Read 4 bytes from /run/dhcpcd/enp1s0f0-4.pid
2024-06-16 18:54:07,137 - util.py[DEBUG]: Reading from /proc/756/stat (quiet=True)
2024-06-16 18:54:07,137 - util.py[DEBUG]: Read 299 bytes from /proc/756/stat
2024-06-16 18:54:07,137 - dhcp.py[DEBUG]: killing dhcpcd with pid=756 gid=755
2024-06-16 18:54:07,137 - ephemeral.py[DEBUG]: Received dhcp lease on enp1s0f0 for 192.168.100.10/255.255.255.255
2024-06-16 18:54:07,140 - util.py[DEBUG]: Resolving URL: http://[fe80::a9fe:a9fe] took 0.002 seconds
2024-06-16 18:54:07,141 - util.py[DEBUG]: Resolving URL: http://169.254.169.254 took 0.000 seconds
2024-06-16 18:54:07,145 - url_helper.py[DEBUG]: [0/1] open 'http://[fe80::a9fe:a9fe]/openstack' with {'url': 'http://[fe80::a9fe:a9fe]/openstack', 'stream': False, 'allow_redirects': True,>
2024-06-16 18:54:07,296 - url_helper.py[DEBUG]: [0/1] open 'http://169.254.169.254/openstack' with {'url': 'http://169.254.169.254/openstack', 'stream': False, 'allow_redirects': True, 'me>
2024-06-16 18:54:07,298 - url_helper.py[WARNING]: Exception(s) [UrlError("HTTPConnectionPool(host='fe80::a9fe:a9fe', port=80): Max retries exceeded with url: /openstack (Caused by NewConne>
2024-06-16 18:54:07,299 - url_helper.py[DEBUG]: Calling 'http://169.254.169.254/openstack' failed [0/-1s]: request error [HTTPConnectionPool(host='169.254.169.254', port=80): Max retries e>
2024-06-16 18:54:07,299 - url_helper.py[ERROR]: Timed out, no response from urls: ['http://[fe80::a9fe:a9fe]/openstack', 'http://169.254.169.254/openstack']
2024-06-16 18:54:07,299 - DataSourceOpenStack.py[DEBUG]: Giving up on OpenStack md from ['http://[fe80::a9fe:a9fe]/openstack', 'http://169.254.169.254/openstack'] after 0 seconds
2024-06-16 18:54:07,299 - util.py[DEBUG]: Crawl of metadata service took 0.161 seconds
2024-06-16 18:54:07,300 - subp.py[DEBUG]: Running command ['ip', '-family', 'inet', 'link', 'set', 'dev', 'enp1s0f0', 'down'] with allowed return codes [0] (shell=False, capture=True)
2024-06-16 18:54:07,310 - subp.py[DEBUG]: Running command ['ip', '-family', 'inet', 'addr', 'del', '192.168.100.10/32', 'dev', 'enp1s0f0'] with allowed return codes [0] (shell=False, captu>
2024-06-16 18:54:07,313 - util.py[WARNING]: No active metadata service found
2024-06-16 18:54:07,313 - util.py[DEBUG]: No active metadata service found
Traceback (most recent call last):
  File "/usr/lib/python3/dist-packages/cloudinit/sources/DataSourceOpenStack.py", line 159, in _get_data
    results = util.log_time(
              ^^^^^^^^^^^^^^
  File "/usr/lib/python3/dist-packages/cloudinit/util.py", line 2827, in log_time
    ret = func(*args, **kwargs)
          ^^^^^^^^^^^^^^^^^^^^^
  File "/usr/lib/python3/dist-packages/cloudinit/sources/DataSourceOpenStack.py", line 213, in _crawl_metadata
    raise sources.InvalidMetaDataException(
cloudinit.sources.InvalidMetaDataException: No active metadata service found
2024-06-16 18:54:07,315 - sources[DEBUG]: Datasource DataSourceOpenStackLocal [net,ver=None] not updated for events: boot-new-instance
2024-06-16 18:54:07,315 - handlers.py[DEBUG]: finish: init-local/search-OpenStackLocal: SUCCESS: no local data found from DataSourceOpenStackLocal
2024-06-16 18:54:07,315 - main.py[DEBUG]: No local datasource found
2024-06-16 18:54:07,316 - util.py[DEBUG]: Reading from /sys/class/net/lo/address (quiet=False)
2024-06-16 18:54:07,316 - util.py[DEBUG]: Read 18 bytes from /sys/class/net/lo/address
2024-06-16 18:54:07,316 - util.py[DEBUG]: Reading from /sys/class/net/ens1f0/address (quiet=False)
2024-06-16 18:54:07,316 - util.py[DEBUG]: Read 18 bytes from /sys/class/net/ens1f0/address
2024-06-16 18:54:07,316 - util.py[DEBUG]: Reading from /sys/class/net/enp1s0f0/address (quiet=False)
2024-06-16 18:54:07,316 - util.py[DEBUG]: Read 18 bytes from /sys/class/net/enp1s0f0/address
2024-06-16 18:54:07,317 - util.py[DEBUG]: Reading from /sys/class/net/ens1f0/name_assign_type (quiet=False)
2024-06-16 18:54:07,317 - util.py[DEBUG]: Read 2 bytes from /sys/class/net/ens1f0/name_assign_type
2024-06-16 18:54:07,317 - util.py[DEBUG]: Reading from /sys/class/net/enp1s0f0/name_assign_type (quiet=False)
2024-06-16 18:54:07,317 - util.py[DEBUG]: Read 2 bytes from /sys/class/net/enp1s0f0/name_assign_type
2024-06-16 18:54:07,317 - util.py[DEBUG]: Reading from /sys/class/net/lo/address (quiet=False)

root@node170-c56b4d:~# ip a s
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
    inet6 ::1/128 scope host noprefixroute
       valid_lft forever preferred_lft forever
2: enp1s0f0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc mq state UP group default qlen 1000
    link/ether fa:16:3e:86:09:64 brd ff:ff:ff:ff:ff:ff
    inet 192.168.100.10/32 metric 100 scope global dynamic enp1s0f0
       valid_lft 36sec preferred_lft 36sec
    inet6 fe80::f816:3eff:fe86:964/64 scope link
       valid_lft forever preferred_lft forever
3: ens1f0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc pfifo_fast state UP group default qlen 1000
    link/ether fa:16:3e:23:4e:24 brd ff:ff:ff:ff:ff:ff
    altname enp3s1f0
    inet 192.168.100.11/32 metric 100 scope global dynamic ens1f0
       valid_lft 34sec preferred_lft 34sec
    inet6 fe80::f816:3eff:fe23:4e24/64 scope link
       valid_lft forever preferred_lft forever
4: enp1s0f0.800@enp1s0f0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default qlen 1000
    link/ether fa:16:3e:86:09:64 brd ff:ff:ff:ff:ff:ff
    inet 192.168.203.12/32 metric 100 scope global dynamic enp1s0f0.800
       valid_lft 60sec preferred_lft 60sec
    inet6 fe80::f816:3eff:fe86:964/64 scope link
       valid_lft forever preferred_lft forever
5: enp1s0f0.600@enp1s0f0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default qlen 1000
    link/ether fa:16:3e:86:09:64 brd ff:ff:ff:ff:ff:ff
    inet 192.168.202.12/32 metric 100 scope global dynamic enp1s0f0.600
       valid_lft 35sec preferred_lft 35sec
    inet6 fe80::f816:3eff:fe86:964/64 scope link
       valid_lft forever preferred_lft forever
6: enp1s0f0.400@enp1s0f0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default qlen 1000
    link/ether fa:16:3e:86:09:64 brd ff:ff:ff:ff:ff:ff
    inet 192.168.201.12/32 metric 100 scope global dynamic enp1s0f0.400
       valid_lft 34sec preferred_lft 34sec
    inet6 fe80::f816:3eff:fe86:964/64 scope link
       valid_lft forever preferred_lft forever
7: enp1s0f0.200@enp1s0f0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default qlen 1000
    link/ether fa:16:3e:86:09:64 brd ff:ff:ff:ff:ff:ff
    inet 192.168.200.170/32 metric 100 scope global dynamic enp1s0f0.200
       valid_lft 2375sec preferred_lft 2375sec
    inet6 fe80::f816:3eff:fe86:964/64 scope link
       valid_lft forever preferred_lft forever

root@node170-c56b4d:~# ip r s
default via 192.168.200.1 dev enp1s0f0.200 proto dhcp src 192.168.200.170 metric 100
10.20.20.0/24 via 192.168.200.1 dev enp1s0f0.200 proto dhcp src 192.168.200.170 metric 100
10.20.21.0/24 via 192.168.203.1 dev enp1s0f0.800 proto dhcp src 192.168.203.12 metric 100
10.20.21.0/24 via 192.168.201.1 dev enp1s0f0.400 proto dhcp src 192.168.201.12 metric 100
10.20.22.0/24 via 192.168.202.1 dev enp1s0f0.600 proto dhcp src 192.168.202.12 metric 100
169.254.169.254 via 192.168.100.9 dev enp1s0f0 proto dhcp src 192.168.100.10 metric 100
169.254.169.254 via 192.168.100.9 dev enp1s0f0.200 proto dhcp src 192.168.200.170 metric 100
169.254.169.254 via 192.168.100.9 dev ens1f0 proto dhcp src 192.168.100.11 metric 100
192.168.100.0/24 via 192.168.100.1 dev enp1s0f0 proto dhcp src 192.168.100.10 metric 100
192.168.100.0/24 via 192.168.100.1 dev ens1f0 proto dhcp src 192.168.100.11 metric 100
192.168.100.1 dev enp1s0f0 proto dhcp scope link src 192.168.100.10 metric 100
192.168.100.1 dev ens1f0 proto dhcp scope link src 192.168.100.11 metric 100
192.168.100.9 dev enp1s0f0 proto dhcp scope link src 192.168.100.10 metric 100
192.168.100.9 dev enp1s0f0.200 proto dhcp scope link src 192.168.200.170 metric 100
192.168.100.9 dev ens1f0 proto dhcp scope link src 192.168.100.11 metric 100
192.168.200.1 dev enp1s0f0.200 proto dhcp scope link src 192.168.200.170 metric 100
192.168.200.1 via 192.168.200.1 dev enp1s0f0.200 proto dhcp src 192.168.200.170 metric 100
192.168.201.0/24 via 192.168.201.1 dev enp1s0f0.400 proto dhcp src 192.168.201.12 metric 100
192.168.201.1 dev enp1s0f0.400 proto dhcp scope link src 192.168.201.12 metric 100
192.168.202.0/24 via 192.168.202.1 dev enp1s0f0.600 proto dhcp src 192.168.202.12 metric 100
192.168.202.1 dev enp1s0f0.600 proto dhcp scope link src 192.168.202.12 metric 100
192.168.203.0/24 via 192.168.203.1 dev enp1s0f0.800 proto dhcp src 192.168.203.12 metric 100
192.168.203.1 dev enp1s0f0.800 proto dhcp scope link src 192.168.203.12 metric 100


```

source code from openstack cloud init.
```
# Versions and names taken from nova source nova/api/metadata/base.py
OS_LATEST = "latest"
OS_FOLSOM = "2012-08-10"
OS_GRIZZLY = "2013-04-04"
OS_HAVANA = "2013-10-17"
OS_LIBERTY = "2015-10-15"
# NEWTON_ONE adds 'devices' to md (sriov-pf-passthrough-neutron-port-vlan)
OS_NEWTON_ONE = "2016-06-30"
# NEWTON_TWO adds vendor_data2.json (vendordata-reboot)
OS_NEWTON_TWO = "2016-10-06"
# OS_OCATA adds 'vif' field to devices (sriov-pf-passthrough-neutron-port-vlan)
OS_OCATA = "2017-02-22"
# OS_ROCKY adds a vf_trusted field to devices (sriov-trusted-vfs)
OS_ROCKY = "2018-08-27"


# Various defaults/constants...
DEF_MD_URLS = ["http://[fe80::a9fe:a9fe]", "http://169.254.169.254"]
DEFAULT_IID = "iid-dsopenstack"
DEFAULT_METADATA = {
    "instance-id": DEFAULT_IID,
}

# OpenStack DMI constants
DMI_PRODUCT_NOVA = "OpenStack Nova"
DMI_PRODUCT_COMPUTE = "OpenStack Compute"
VALID_DMI_PRODUCT_NAMES = [DMI_PRODUCT_NOVA, DMI_PRODUCT_COMPUTE]
DMI_ASSET_TAG_OPENTELEKOM = "OpenTelekomCloud"
# See github.com/sapcc/helm-charts/blob/master/openstack/nova/values.yaml
# -> compute.defaults.vmware.smbios_asset_tag for this value
DMI_ASSET_TAG_SAPCCLOUD = "SAP CCloud VM"
DMI_ASSET_TAG_HUAWEICLOUD = "HUAWEICLOUD"
VALID_DMI_ASSET_TAGS = VALID_DMI_PRODUCT_NAMES
VALID_DMI_ASSET_TAGS += [
    DMI_ASSET_TAG_HUAWEICLOUD,
    DMI_ASSET_TAG_OPENTELEKOM,
    DMI_ASSET_TAG_SAPCCLOUD,
]


```


### STEP END
custom cloud.cfg
```
cat <<EOF>/etc/cloud/cloud.cfg
# The top level settings are used as module
# and base configuration.

# A set of users which may be applied and/or used by various modules
# when a 'default' entry is found it will reference the 'default_user'
# from the distro configuration specified below
users:
  - default

# If this is set, 'root' will not be able to ssh in and they
# will get a message to login instead as the default $user
disable_root: true

# This will cause the set+update hostname module to not operate (if true)
preserve_hostname: false

# If you use datasource_list array, keep array items in a single line.
# If you use multi line array, ds-identify script won't read array items.
# Example datasource config
#datasource:
#  NoCloudNet:
#    seedfrom: http://169.254.169.254/cloud-init/configs/
#  NoCloud:
#    seedfrom: http://169.254.169.254/cloud-init/configs/
#
#datasource_list: ["openstack"]
#datasource:
#   Ec2:
#     metadata_urls: [ 'blah.com' ]
#     timeout: 5 # (defaults to 50 seconds)
#     max_wait: 10 # (defaults to 120 seconds)

# The modules that run in the 'init' stage
cloud_init_modules:
  - seed_random
  - bootcmd
  - write_files
  - growpart
  - resizefs
  - disk_setup
  - mounts
  - set_hostname
  - update_hostname
  - update_etc_hosts
  - ca_certs
  - rsyslog
  - users_groups
  - ssh

# The modules that run in the 'config' stage
cloud_config_modules:
  - wireguard
  - snap
  - ubuntu_autoinstall
  - ssh_import_id
  - keyboard
  - locale
  - set_passwords
  - grub_dpkg
  - apt_pipelining
  - apt_configure
  - ubuntu_pro
  - ntp
  - timezone
  - disable_ec2_metadata
  - runcmd
  - byobu

# The modules that run in the 'final' stage
cloud_final_modules:
  - package_update_upgrade_install
  - fan
  - landscape
  - lxd
  - ubuntu_drivers
  - write_files_deferred
  - puppet
  - chef
  - ansible
  - mcollective
  - salt_minion
  - reset_rmc
  - scripts_vendor
  - scripts_per_once
  - scripts_per_boot
  - scripts_per_instance
  - scripts_user
  - ssh_authkey_fingerprints
  - keys_to_console
  - install_hotplug
  - phone_home
  - final_message
  - power_state_change
# System and/or distro specific settings
# (not accessible to handlers/transforms)
system_info:
  # This will affect which distro class gets used
  distro: ubuntu
  # Default user name + that default users groups (if added/used)
  default_user:
    name: ubuntu
    lock_passwd: True
    gecos: Ubuntu
    groups: [adm, cdrom, dip, lxd, sudo]
    sudo: ["ALL=(ALL) NOPASSWD:ALL"]
    shell: /bin/bash
  network:
    dhcp_client_priority: [dhcpcd, dhclient, udhcpc]
    renderers: ['netplan', 'eni', 'sysconfig']
    activators: ['netplan', 'eni', 'network-manager', 'networkd']
  # Automatically discover the best ntp_client
  ntp_client: auto
  # Other config here will be given to the distro class and/or path classes
  paths:
    cloud_dir: /var/lib/cloud/
    templates_dir: /etc/cloud/templates/
  package_mirrors:
    - arches: [i386, amd64]
      failsafe:
        primary: http://archive.ubuntu.com/ubuntu
        security: http://security.ubuntu.com/ubuntu
      search:
        primary:
          - http://%(ec2_region)s.ec2.archive.ubuntu.com/ubuntu/
          - http://%(availability_zone)s.clouds.archive.ubuntu.com/ubuntu/
          - http://%(region)s.clouds.archive.ubuntu.com/ubuntu/
        security: []
    - arches: [arm64, armel, armhf]
      failsafe:
        primary: http://ports.ubuntu.com/ubuntu-ports
        security: http://ports.ubuntu.com/ubuntu-ports
      search:
        primary:
          - http://%(ec2_region)s.ec2.ports.ubuntu.com/ubuntu-ports/
          - http://%(availability_zone)s.clouds.ports.ubuntu.com/ubuntu-ports/
          - http://%(region)s.clouds.ports.ubuntu.com/ubuntu-ports/
        security: []
    - arches: [default]
      failsafe:
        primary: http://ports.ubuntu.com/ubuntu-ports
        security: http://ports.ubuntu.com/ubuntu-ports
  ssh_svcname: ssh
EOF

```
