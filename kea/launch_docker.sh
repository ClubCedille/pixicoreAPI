#!/usr/bin/env bash

set -e

export KEA_INTERFACE="eth1"
export KEA_POOL="10.1.1.100 - 10.1.1.254"
export KEA_GATEWAY="10.1.1.1"
export KEA_NETWORK="10.1.1.3/24"
#export KEA_CONFIG_TEMPLATE="kea-single-subnet.conf.tmpl"
export KEA_DOCKER_IMAGE="clubcedille/kea:latest"

# Generate kea-single-subnet.conf.tmpl from kea-single-subnet.conf

#SPECIFIC_ERROR_MESSAGE=""
#trap 'echo -e "\nThank you for using me. Read ./kea_bootstrapconf.sh on ./kea/ for more informations. EXIT code (rc: $?)"' EXIT


# Please, set KEA_INTERFACE variable with your host network interface like eth0.
export __INTERFACE__=${KEA_INTERFACE}

# Please, set KEA_POOL variable like "192.168.1.100 - 192.168.1.225".
export __POOL__=${KEA_POOL}

# Please, set KEA_GATEWAY variable like "192.168.1.1".
export __GATEWAY__=${KEA_GATEWAY}

# Please, set KEA_NETWORK variable like "192.168.1.0/24".
export __NETWORK__=${KEA_NETWORK}


MYVARS='$__INTERFACE__:$__POOL__:$__GATEWAY__:$__NETWORK__'

cd /vagrant/kea && envsubst "$MYVARS" <${KEA_CONFIG_TEMPLATE:-kea-single-subnet.conf.tmpl} > kea-single-subnet.conf

#sudo docker run -ti --rm --cap-add=NET_ADMIN --net=host -ti -v $PWD/kea-single-subnet.conf:/usr/local/etc/kea/kea-dhcp4.conf $KEA_DOCKER_IMAGE
