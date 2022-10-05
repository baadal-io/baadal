package main

import (
	"baadal-server/common"
	"log"
	"net"
	"time"

	"github.com/digitalocean/go-libvirt"
	libvirtxml "libvirt.org/libvirt-go-xml"
)

func main() {
	conn, err := net.DialTimeout("unix", "/var/run/libvirt/libvirt-sock", time.Second*2)
	common.CheckFatalError(err, "Could not connect to libvirt socket")
	defer conn.Close()

	l := libvirt.New(conn)
	err = l.Connect()
	common.CheckFatalError(err, "Could not initialize the connection to libvirt service")
	defer l.Disconnect()

	version, err := l.ConnectGetLibVersion()
	common.CheckError(err, "Could not fetch version")
	log.Println(version)

	testDom := libvirtxml.Domain{
		Name: "Test",
		Memory: &libvirtxml.DomainMemory{
			Value: 512,
			Unit:  "M",
		},
		Type: "qemu",
		OS: &libvirtxml.DomainOS{
			Type: &libvirtxml.DomainOSType{
				Type: "hvm",
			},
		},
	}

	testDomXml, err := testDom.Marshal()
	common.CheckFatalError(err, "XML could not be generated")

	dom, err := l.DomainDefineXML(testDomXml)
	common.CheckFatalError(err, "Could not define domain")

	err = l.DomainCreate(dom)
	common.CheckFatalError(err, "Could not create domain")
}
