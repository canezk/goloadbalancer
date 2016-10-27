#!bin/sh

if [ ! -f install ]; then
    echo 'install must be run within its container folder' 1>&2
    exit 1
fi

CURDIR=`pwd`
OLDGOPATH="$GOPATH"
export GOPATH="$CURDIR"

gofmt -w src

go install golangloadbalancer

export GOPATH="$OLDGOPATH"

echo 'finished'
