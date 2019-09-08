#!/bin/sh

sudo /usr/local/bin/alp -f /var/log/nginx/access.log -r --sum | head -n 30
#sudo mysqldumpslow -s t /var/log/mariadb/slow.log | head -n 30
