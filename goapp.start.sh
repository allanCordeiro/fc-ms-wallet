#!/bin/sh

sleep 5s && migrate -path=sql/migrations -database "mysql://root:root@tcp(mysql:3306)/wallet" -verbose up  
tail -f /dev/null