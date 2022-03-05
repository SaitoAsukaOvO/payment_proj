docker search mysql
docker pull mysql
docker images
docker run --name mysql --privileged=true -p 3306:3306  -v /data/mysql/datadir:/var/lib/mysql -v /data/mysql/conf.d:/etc/mysql/conf.d -e  MYSQL_ROOT_PASSWORD=123456 -d  mysql:latest
docker exec -it mysql /bin/bash

# ALTER USER 'root'@'%' IDENTIFIED WITH mysql_native_password BY '123456';