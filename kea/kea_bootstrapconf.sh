#!/usr/bin/env bash
set -ueE
set -o pipefail


# Getting started. Read this script. As example, run me like this :
# $ KEA_INTERFACE="eth0" KEA_POOL="192.168.1.100 - 192.168.1.225" \
#    KEA_GATEWAY="192.168.1.1" KEA_NETWORK="192.168.1.0/24" \
# ./kea_bootstrapconf.sh


SPECIFIC_ERROR_MESSAGE=""
trap 'echo -e "\nThank you for using me. Read ./kea_bootstrapconf.sh on ./kea/ for more informations. EXIT code (rc: $?)"' EXIT


# Please, set KEA_INTERFACE variable with your host network interface like eth0.
export __INTERFACE__=${KEA_INTERFACE}

# Please, set KEA_POOL variable like "192.168.1.100 - 192.168.1.225".
export __POOL__=${KEA_POOL}

# Please, set KEA_GATEWAY variable like "192.168.1.1".
export __GATEWAY__=${KEA_GATEWAY}

# Please, set KEA_NETWORK variable like "192.168.1.0/24".
export __NETWORK__=${KEA_NETWORK}


MYVARS='$__INTERFACE__:$__POOL__:$__GATEWAY__:$__NETWORK__'

envsubst "$MYVARS" <${KEA_CONFIG_TEMPLATE:-kea-single-subnet.conf.tmpl} > kea-single-subnet.conf
