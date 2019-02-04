#!/usr/bin/env bash

set -e

export KEA_INTERFACE="wlp113s0"
export KEA_POOL="192.168.1.100 - 192.168.1.225"
export KEA_GATEWAY="192.168.1.1"
export KEA_NETWORK="192.168.1.0/24"
export KEA_CONFIG_TEMPLATE="kea-single-subnet.conf.tmpl"
export KEA_DOCKER_IMAGE="clubcedille/kea:latest"

# Generate kea-single-subnet.conf.tmpl from kea-single-subnet.conf
./kea_bootstrapconf.sh

sudo docker run -ti --rm --cap-add=NET_ADMIN --net=host -ti -v $PWD/kea-single-subnet.conf:/usr/local/etc/kea/kea-dhcp4.conf $KEA_DOCKER_IMAGE
