package helper

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"strings"

	"github.com/matishsiao/goInfo"
	log "github.com/sirupsen/logrus"
)

type PxeSpec struct {
	K   string   `json:"kernel"`
	I   []string `json:"initrd"`
	CMD string   `json:"cmdline"`
}

//// *************************** HELPER FUNCTIONS ****************************
func PixicoreInit(IPAddress string) PxeSpec {
	cmd := "coreos.autologin coreos.first_boot=1 coreos.config.url={{ URL \"file:///home/cedille/pxe-config.ign\" }}"
	ip := "ip="
	ip = strings.Join([]string{ip, IPAddress}, "")
	cmd = strings.Join([]string{cmd, ip}, " ")

	pxeSpec := PxeSpec{

		K: "file:///home/cedille/coreos_production_pxe.vmlinuz",
		I: []string{
			"file:///home/cedille/coreos_production_pxe_image.cpio.gz",
		},
		CMD: "coreos.autologin coreos.first_boot=1 coreos.config.url={{ URL \"file:///home/cedille/pxe-config.ign\" }}"}

	return pxeSpec
}

// CollectPhysicalsIfaces Used to collect physicals interfaces by excluding virtuals interfaces from all interfaces
func CollectPhysicalsIfaces() ([]*net.Interface, error) {
	// Use system path containing all interfaces
	// Since everything is a file on *Nix systems we can only use /sys to discover nic.
	basePath := "/sys/class/net"

	files, err := ioutil.ReadDir(basePath)
	if err != nil {
		return nil, err
	}

	// This variable will contain the return for this function
	var phyIfaceList []*net.Interface

	for _, file := range files {
		ifacePath, err := os.Readlink(fmt.Sprint(basePath, "/", file.Name()))
		if err != nil {
			log.Println(err)
		}

		// If the nic symlink doesn't contain "devices/virtual/net" then we got a physical device.
		if !strings.Contains(ifacePath, "devices/virtual/net") {
			phyIface, err := net.InterfaceByName(file.Name())
			if err != nil {
				log.Println(err)
			}
			phyIfaceList = append(phyIfaceList, phyIface)
		}
	}
	return phyIfaceList, err
}

//GetServerInfo scan this info from server: Kernel version, CPU architecture, Hostname and CPU number
func GetServerInfo() *goInfo.GoInfoObject {
	gi := goInfo.GetInfo()
	return gi
}
