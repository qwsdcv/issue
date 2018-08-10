#!/bin/sh

GOROOT=/home/go

check(){
if [ $? -eq 0 ];then
echo "OK"
else
if [ -n "$1" ];then
echo $1
else
echo "Error"
fi
exit -1
fi
}

out(){
echo "$(tput setab 7)$1$(tput sgr 0)"
}

installSystem(){
out "INSTALL git"
apt-get install git
check "install git error"
}

installGolang(){
out "INSTALL golang"
#apt-get install golang-go
mkdir -p /tmp/golang
cd /tmp/golang
wget https://dl.google.com/go/go1.10.3.linux-amd64.tar.gz
check "Get go package from https://dl.google.com/go/go1.10.3.linux-amd64.tar.gz failed."
tar -C /usr/local -xzf go1.10.3.linux-amd64.tar.gz
rm -rf /tmp/golang

out "Setting environment variable"
mkdir -p $GOROOT/src
echo "export GOPATH=$GOROOT\nexport PATH=\$PATH:${GOROOT}/bin:/usr/local/go/bin" > /etc/profile.d/go.sh
. /etc/profile.d/go.sh
}

buildIssues(){
out "build Issues"
. /etc/profile.d/go.sh
cd $GOROOT/src
go get -u github.com/golang/dep/cmd/dep
check "install go dep error"
cd $GOROOT/src
git clone https://github.com/qwsdcv/issues.git
check "git clone error"
cd issues
dep init
check "dep init error"
dep status
check "dep status error"
dep ensure
check "dep ensure error"

go build
check "go build failed"

}

installSuperVisor(){
out "INSTALL supervisor"
apt-get install supervisor
check "install supervisor error"
echo "[program:issues]\ndirectory=${GOROOT}/src/issues\ncommand=${GOROOT}/src/issues/issues\nautostart=true\nstderr_logfile=${GOROOT}/src/issues/out.log\nstdout_logfile=${GOROOT}/src/issues/out.log\n" > /etc/supervisor/conf.d/issues.conf
service supervisor restart
}

installSystem
installGolang
buildIssues
installSuperVisor
