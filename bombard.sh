#!/bin/bash
for ((i=1;i<=200;i++)); 
do 
    echo "\nSending request number "$i
    echo '{"number":'$i'}'
    curl  -X POST "http://localhost:8080/" -d '{"number":'$i'}' -H "Content-Type: application/json"
done