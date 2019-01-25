package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/ClubCedille/pixicoreAPI/pkg/config"
	"github.com/ClubCedille/pixicoreAPI/pkg/server"
	"github.com/ClubCedille/pixicoreAPI/pkg/sshclient"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh"
)

type Controller struct {
	currentConfig *config.ConfigFile
}

//Getlocal pixicore demands
func (ctrl *Controller) Getlocal(c *gin.Context) {
	c.JSON(200, "success")

}

//BootServers called by pixicore client to register a new server
func (ctrl *Controller) BootServer(c *gin.Context) {

	servers, err := ctrl.currentConfig.GetServers()
	if err != nil {
		log.Warn(err)
	}

	macAddr := c.Param("macAddress")

	err = servers.AddServer(macAddr)
	if err != nil {
		log.Warn(err)
	}

	server, err := servers.GetServer(macAddr)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusNotFound, gin.H{"status": err})
	} else {

		pxeSpec := server.Boot()
		c.JSON(200, pxeSpec)
	}

}

//InstallServer Install a single server
func (ctrl *Controller) InstallServer(c *gin.Context) {
	macAddr := c.Param("macAddress")
	servers := ctrl.currentConfig.Servers
	if _, err := ctrl.currentConfig.Servers.GetServer(macAddr); err == nil {

		err := fmt.Sprint("This Requested server doesn't exist : ", macAddr)
		c.JSON(http.StatusNotFound, gin.H{"status": err})
	}

	ctrl.CollectServerInfo((*servers)[c.Param("macAddress")])

	ctrl.currentConfig.WriteYamlConfig()

	c.JSON(200, (*servers)[c.Param("macAddress")])
}

//InstallAll install all the servers available
func (ctrl *Controller) InstallAll(c *gin.Context) {
	servers := ctrl.currentConfig.Servers
	for svr := range *servers {
		ctrl.CollectServerInfo((*servers)[svr])
	}

	ctrl.currentConfig.WriteYamlConfig()

	c.JSON(200, &servers)
}

//CollectServerInfo collect information about a server with ssh
func (ctrl *Controller) CollectServerInfo(currentServer *server.Server) {

	sshConfig := ssh.ClientConfig{
		User: "core",
		Auth: []ssh.AuthMethod{
			sshclient.PublicKeyFile("~/.ssh/id_rsa"),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	clientSSH := sshclient.SSHClient{
		Config: &sshConfig,
		Host:   currentServer.IPAddress,
		Port:   22,
	}

	// run command with ssh
	kernel, err := clientSSH.RunCommand("uname -r")
	if err != nil {
		log.Errorf("command run error: %s", err)
	}

	macAddressFirst, err := clientSSH.RunCommand("cat /sys/class/net/enp4s0/address")
	if err != nil {
		log.Errorf("command run error: %s\n", err)
	}
	macAddressSecond, err := clientSSH.RunCommand("cat /sys/class/net/enp5s0/address")
	if err != nil {
		log.Errorf("command run error: %s\n", err)
	}

	var machineId = ""
	machineId, err = clientSSH.RunCommand("cat /var/lib/dbus/machine-id")
	if err != nil {
		// It failed, so try again with another path.
		machineId, err = clientSSH.RunCommand("cat /etc/machine-id")

		if err != nil {
			log.Errorf("Failed to get the machine Id. Command run error: %s\n", err)
		}
	}

	if currentServer.MacAddress.String() == strings.TrimSuffix(macAddressFirst, "\r\n") {
		currentServer.SecondMacAddress = strings.TrimSuffix(macAddressSecond, "\r\n")
	} else {
		currentServer.SecondMacAddress = strings.TrimSuffix(macAddressFirst, "\r\n")
	}
	currentServer.Kernel = strings.TrimSuffix(kernel, "\r\n")

	// Set machine id as the server's id. This is an unique identifier that will help differentiate the servers.
	currentServer.Id = strings.TrimSuffix(machineId, "\r\n")

	_, err = clientSSH.RunCommand("sudo coreos-install -d /dev/sda -i /run/ignition.json -C stable")
	if err != nil {
		log.Errorf("command run error: %s\n", err)
	}
	currentServer.Installed = true

}

// GetServers return config of the all the servers
func (ctrl *Controller) GetServers(c *gin.Context) {
	servers, err := ctrl.currentConfig.GetServers()

	if err != nil {
		log.Error(err)
		c.JSON(403, gin.H{"status": err})
	}

	c.JSON(200, gin.H{"success": servers})
}
