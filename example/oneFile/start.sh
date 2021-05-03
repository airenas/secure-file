#!/bin/sh
set -e

echo Running startup script
./check-decrypt-file -f secretFile.txt 

echo Showing secret data from file
cat secretFile.txt

#exec ./cmd
echo Waiting for terminate 
trap 'exit 143' SIGTERM
while true; do sleep 1; done
