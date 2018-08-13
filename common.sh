#!/bin/bash


out(){
echo "$(tput setaf 2)$(tput setab 7)$1$(tput sgr 0)"
}

err(){
echo "$(tput setaf 1)$(tput setab 7)$1$(tput sgr 0)"
}

check(){
if [ $? -eq 0 ];then
out "OK"
else
if [ -n "$1" ];then
err $1
else
err "Error"
fi
exit -1
fi
}


