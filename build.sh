#!/bin/bash
SRC="./src"

export GOPATH=$HOME/Code/gobasics/
export GOBIN=$HOME/Code/gobasics/bin

echo Building...
for d in $SRC'/ch*' $SRC'/extras'; do
    pd=$d'/*'
    for p in $pd; do
        f=$(echo $p | cut -d/ -f3-)
        go install -v $f 
        rc=$?; if [[ $rc != 0 ]]; then exit $rc; fi
    done
done

