#!/bin/sh

ACTION="build"

if [ "$1" == "clean" ]; then
	ACTION="clean"
fi

for x in "`pwd`"/src/cmd/*; do
	if [ -d $x ]; then
		echo "${ACTION}ing `basename $x` ..."
		cd $x
		go $ACTION
	fi
done
