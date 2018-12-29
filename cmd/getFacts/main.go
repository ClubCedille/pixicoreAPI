package main

import (
	"github.com/ClubCedille/pixicoreAPI/pkg/helper"
	log "github.com/sirupsen/logrus"
)


func main() {

	physIfaces, err := helper.CollectPhysicalsIfaces()
	if err != nil {
		log.Error(err)
	}

	for _, iface := range physIfaces{
		log.Printf("Iface Index: %v  IfaceName: %v Iface Mac: %s ", iface.Index , iface.Name, iface.HardwareAddr.String())
	}

	serverInfo := helper.GetServerInfo()

	log.Println(serverInfo)
}
