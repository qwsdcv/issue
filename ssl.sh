#!/bin/bash

source common.sh

apt-get update
apt-get install software-properties-common
add-apt-repository ppa:certbot/certbot
apt-get update
apt-get install python-certbot-nginx 

certbot --nginx

check "certbot error"
