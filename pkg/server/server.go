package server

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
	// "github.com/google/go-cmp/cmp/cmpopts"
	"github.com/jalilbengoufa/pixicoreAPI/pkg/helper"
	log "github.com/sirupsen/logrus"
	"net"
)

//Servers represent a list of Server type
type Servers map[string]*Server

//Server represent a config of a server
type Server struct {
	MacAddress       net.HardwareAddr `yaml:"macAddress"`
	IPAddress        string           `yaml:"ipAddress"`
	Installed        bool             `yaml:"installed"`
	Kernel           string           `yaml:"kernel"`
	SecondMacAddress string           `yaml:"secondmacAddress"`
}

type EmptyServerListError struct {
}

func (e *EmptyServerListError) Error() string {
	return fmt.Sprintf("%v: server error", "Server list are empty")
}

type NilServerListError struct {
}

func (e *NilServerListError) Error() string {
	return fmt.Sprintf("%v: server error", "Server list is nil and maps in Golang are useless when they are nil. Reference : https://blog.golang.org/go-maps-in-action")
}

type UnreconizeServerError struct {
	serverList   *Servers
	wantedServer string
}

func (e *UnreconizeServerError) Error() string {
	return fmt.Sprintf("Server named %v+ is not found in this server list: %v+ ", e.wantedServer, e.serverList)
}

//AddServer Add server using mac Address.
func (servers *Servers) AddServer(macAddressStr string) error {
	var macAddress net.HardwareAddr
	macAddress, err := net.ParseMAC(macAddressStr)
	if err != nil {
		return err
	}

	_, err = servers.GetServer(macAddress.String())
	server := Server{
		MacAddress: macAddress, IPAddress: "change me", Installed: false, Kernel: "linux", SecondMacAddress: "find me"}
	if err == nil {
		log.Warnln("The server already exist in the list. Overwrite it.")
	} else {
		switch err.(type) {
		case *EmptyServerListError:
			log.Infoln(err)
			break
		case *UnreconizeServerError:
			break
		case *NilServerListError:

			// A map should not be nil
			// Refence : https://blog.golang.org/go-maps-in-action
			return err

		default:
			return err
		}
	}

	(*servers)[macAddress.String()] = &server

	return nil
}

//GetServer Get server from a list of servers using context as input
func (servers *Servers) IsExist(macAddressStr string) bool {

	server, _ := servers.GetServer(macAddressStr)

	if server != nil {
		return true
	} else {
		return false
	}

}

//GetServer Get server from a list of servers using context as input
func (servers *Servers) GetServer(macAddressStr string) (*Server, error) {

	if servers == nil {
		return nil, new(NilServerListError)
	} else if cmp.Equal(*servers, make(Servers)) {
		return nil, new(EmptyServerListError)

	} else if server, ok := (*servers)[macAddressStr]; ok {
		return server, nil
	} else {
		err := UnreconizeServerError{serverList: servers, wantedServer: macAddressStr}
		return nil, &err
	}

}

//Boot Boot server specified in gin Context
func (server Server) Boot() (bootLog helper.PxeSpec) {

	pxeSpec := helper.PixicoreInit(server.MacAddress.String())
	return pxeSpec
}
