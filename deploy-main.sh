#!/bin/bash

if [ $# == 1 ]; then
	echo "execute checkout and pull!!"
	git reset --hard HEAD && git fetch && git checkout $1 && git pull origin $1
fi

# Database data reset
#sudo ./isucari/db/init.sh

# application build
(
cd webapp/go
make
)

# replaces log data 
LOGPATH=/var/log
NOW=`date +'%H-%M-%S'`

sudo cp $LOGPATH/nginx/access.log $LOGPATH/nginx/access-$NOW.log
sudo sh -c 'echo "" > /var/log/nginx/access.log'

sudo cp $LOGPATH/mysql/slow.log $LOGPATH/mysql/slow-$NOW.log
sudo sh -c 'echo "" > /var/log/mysql/slow.log'

# replace mysql conf
sudo cp /home/isucon/isucari/conf/mysqld.cnf /etc/mysql/mysql.conf.d/mysqld.cnf

# replace redis conf    
sudo cp /home/isucon/isucari/conf/redis.conf /etc/redis/redis.conf    
sudo chown redis:redis /etc/redis/redis.conf

# restart application services
# db, app, nginx, redis
echo 'systemctl are restarting...'
sudo systemctl restart mysql.service
sudo systemctl restart isucari.golang.service
sudo systemctl restart nginx.service
sudo systemctl restart redis.service
echo 'Finished to restart!!'
