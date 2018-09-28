#!/bin/bash

source common.sh;

apt-get install shadowsocks;

SSCONFIG=~/.ss.conf;

echo "
{
    \"server\":\"0.0.0.0\",
    \"server_port\":8388,
    \"local_address\": \"127.0.0.1\",
    \"local_port\":1080,
    \"password\":\"helloworld\",
    \"timeout\":300,
    \"method\":\"aes-256-cfb\",
    \"fast_open\": false
}
" > $SSCONFIG;

echo -e "
[program:shadowsocks]
directory=/home/go/src/issues
command=ssserver -c $SSCONFIG start
autostart=true
stderr_logfile=~/ss.log
stdout_logfile=~/ss.log
" > /etc/supervisor/conf.d/ss.conf;

service supervisor restart;
