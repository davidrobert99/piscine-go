#! /bin/bash

echo "$(ls -t *.sh)" | cut -d. -f1 

