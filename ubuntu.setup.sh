#!/bin/bash

GOROOT=/home/go
GOVERSION=1.10

source ./common.sh

installSystem(){
    out "INSTALL git";
    apt-get install git;
    check "install git error";
}

installGolang(){
    rm -rf /usr/local/go;
    out "INSTALL golang";
    apt-get install golang-${GOVERSION}
#    mkdir -p /tmp/golang;
#    cd /tmp/golang;
#    wget https://dl.google.com/go/go1.10.3.linux-amd64.tar.gz;
#    check "Get go package from https://dl.google.com/go/go1.10.3.linux-amd64.tar.gz failed.";
#    tar -C /usr/local -xzf go1.10.3.linux-amd64.tar.gz;
#    rm -rf /tmp/golang;

    out "Setting environment variable";
    mkdir -p $GOROOT/src;
    echo -e "export GOPATH=$GOROOT\nexport PATH=\$PATH:${GOROOT}/bin" > /etc/profile.d/go.sh;
    . /etc/profile.d/go.sh;
}

buildIssues(){
    out "build Issues";
    . /etc/profile.d/go.sh;
    cd $GOROOT/src;
    go get -u github.com/golang/dep/cmd/dep;
    check "install go dep error";
    cd $GOROOT/src;
    if [ ! -d "issue" ];then
    git clone --depth=1 https://github.com/qwsdcv/issues.git;
    check "git clone error";
    fi
    cd issues;
    dep init;
    check "dep init error";
    dep status;
    check "dep status error";
    dep ensure;
    check "dep ensure error";

    go build;
    check "go build failed";

}

#installMariaDB(){
#    out "INSTALL mariaDB";
#    apt-get install  mariadb-server;
#    check "install mariadb failed";
#    cd $GOROOT/src/issues/models;
#    sh SYSTEM.sh;
#    check "CREATE DB failed";
#    sh SQL.sh;
#    check "CREATE TABLE failed";
#}

installSuperVisor(){
    out "INSTALL supervisor";
    apt-get install supervisor;
    check "install supervisor error";
    echo -e "[program:issues]\ndirectory=${GOROOT}/src/issues\ncommand=${GOROOT}/src/issues/issues\nautostart=true\nstderr_logfile=${GOROOT}/src/issues/out.log\nstdout_logfile=${GOROOT}/src/issues/out.log\n" > /etc/supervisor/conf.d/issues.conf;
    service supervisor restart;
}

installNginx(){
    out "INSTALL nginx";
    apt-get install nginx;
    check "install nginx error";
    out "config nginx";
    echo -e "server {\n    listen 80;\n    server_name gushijingcai.com;\n\n    location / {\n        root /home/go/src/issues/static;\n        index index.html;\n    }\n\n    location ^~ /issues/ {\n        proxy_pass http://127.0.0.1:8088;\n    }\n}\n" > /etc/nginx/conf.d/issues.conf;
    nginx -s reload;
    check "nginx reload error";
}



installSystem
installGolang
buildIssues
installSuperVisor
installNginx
