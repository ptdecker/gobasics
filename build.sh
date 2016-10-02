#!/bin/bash
SRC="./src"

for d in $SRC'/*'; do
    pd=$d'/*'
    for p in $pd; do
        f=$(echo $p | cut -d/ -f3-)
        go install -v $f 
        rc=$?; if [[ $rc != 0 ]]; then exit $rc; fi
    done
done

