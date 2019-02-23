#!/usr/bin/env bash

set -e

export KEA_INTERFACE="eth1"
export KEA_POOL="10.1.1.100 - 10.1.1.254"
export KEA_GATEWAY="10.1.1.1"
export KEA_NETWORK="10.1.1.3/24"
export KEA_CONFIG_TEMPLATE="kea-single-subnet.conf.tmpl"
export KEA_DOCKER_IMAGE="clubcedille/kea:latest"

# Generate kea-single-subnet.conf.tmpl from kea-single-subnet.conf
./kea_bootstrapconf.sh

sudo docker run -ti --rm --cap-add=NET_ADMIN --net=host -ti -v $PWD/kea-single-subnet.conf:/usr/local/etc/kea/kea-dhcp4.conf $KEA_DOCKER_IMAGE
