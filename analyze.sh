#!/bin/sh

sudo /usr/local/bin/alp -f /var/log/nginx/access.log -r --sum | head -n 40
sudo mysqldumpslow -s -t /var/log/mysql/slow.log | head -n 20
