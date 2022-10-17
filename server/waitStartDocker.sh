#!/bin/bash

## Or whatever command is used for checking logstash availability
# until curl 'http://logstash:5045' 2> /dev/null; do
#   echo "Waiting for logtash..."
#   sleep 1; 
# done
echo "Задержка запуска..."
cd server
sleep 5

# Start your server
node index.js