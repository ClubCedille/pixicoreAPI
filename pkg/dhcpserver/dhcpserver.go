package dhcpserver

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"syscall"
	"time"
)

type DhcpdServer struct {
	Mac string
	Ip  string
}

func getServerIP(macAddress string) string {

	var servers []DhcpdServer

	var httpClient = &http.Client{Timeout: 10 * time.Second}
	res, _ := httpClient.Get("http://localhost:8008")
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	json.Unmarshal(body, &servers)

	for _, element := range servers {

		if element.Mac == macAddress {
			if verifyServerConnection(element.Ip) {
				return element.Ip
			} else {
				return "false"
			}
		}
	}

	return "false"
}

func verifyServerConnection(ipAddress string) bool {
	var exitCode int
	for i := 0; i < 3; i++ {

		time.Sleep(6 * time.Second)

		cmd := exec.Command("nc -vz 142.137.247.120 22")
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}

		ws := cmd.ProcessState.Sys().(syscall.WaitStatus)
		exitCode = ws.ExitStatus()

		if exitCode == 0 {
			return true
		}
	}
	return false
}
