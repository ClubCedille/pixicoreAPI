#!/bin/bash

set -e

curl -X POST -H "Content-Type: application/json" -d '{ "command": "lease4-add", "service": [ "dhcp4" ], "arguments": {"ip-address": "192.168.1.202", "hw-address": "1a:1b:1c:1d:1e:1f"} }' http://127.0.0.1:8080

curl -X POST -H "Content-Type: application/json" -d '{ "command": "lease4-get-all", "service": [ "dhcp4" ] }' http://127.0.0.1:8080 | jq
