#!/bin/sh
set -e

echo Running startup script
./check-decrypt-file -fl secretFileList.txt 

echo Showing secret data from files
cat secretFile*.txt

#exec ./cmd
echo Waiting for terminate 
trap 'exit 143' SIGTERM
while true; do sleep 1; done
