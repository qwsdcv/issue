#!/bin/bash

source ./common.sh

setupFirewall(){
    out "Setup Firewall";
    apt-get install ufw;
    check "INSTALL ufw failed";
    ufw default deny;
    ufw allow 22;
    ufw allow 80;
    ufw allow 443;
    ufw allow 8388;
    ufw enable;
    check "ufw setup error";
}

setupFirewall
