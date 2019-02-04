package dhcpserver

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
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

	return false
}
