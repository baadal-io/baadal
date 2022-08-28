package main

import (
	"baadal-server/common"
	"log"
	"net"
	"time"

	"github.com/digitalocean/go-libvirt"
)

func main() {
	conn, err := net.DialTimeout("unix", "/var/run/libvirt/libvirt-sock", time.Second*2)
	common.CheckFatalError(err, "Could not connect to libvirt socket")
	defer conn.Close()

	l := libvirt.New(conn)
	version, err := l.ConnectGetLibVersion()
	common.CheckError(err, "Could not fetch version")
	log.Println(version)

	err = l.Connect()
	common.CheckFatalError(err, "Could not initialize the connection to libvirt service")
	defer l.Disconnect()
}
