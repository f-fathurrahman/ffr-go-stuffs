#!/bin/bash
basn=`basename $1 .go`

rm -vf $basn.x
go build $1
mv $basn $basn.x
echo "Executable: $basn.x"
