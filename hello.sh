#! /bin/bash

 

echo 'hello'  $(curl  https://api.github.com/users/davidrobert99 | jq '.login')