#!/bin/bash

if [ $# == 1 ]; then
	echo "execute checkout and pull!!"
	git reset --hard HEAD && git fetch && git checkout $1 && git pull origin $1
fi

# application build
(
cd webapp/go
make
)

# restart application services
# db, app, nginx, redis
echo 'systemctl are restarting...'
sudo systemctl restart isucari.golang.service
echo 'Finished to restart!!'
