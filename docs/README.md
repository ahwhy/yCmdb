## MySQL安装

```sh
$ docker run -p 3306:3306 -itd -e MARIADB_USER=yCmdb -e MARIADB_PASSWORD=123456 -e MARIADB_ROOT_PASSWORD=123456 --name mysql mariadb:latest
```

详细文档请参考: [mariadb docker hub](https://hub.docker.com/_/mariadb)

进入docker创建数据库和账号:
```
docker exec -it b272a099f6d9 mysql -uroot -p123456

create database yCmdb;
GRANT ALL PRIVILEGES ON yCmdb.* TO 'yCmdb'@'%' IDENTIFIED BY '123456';
flush privileges;
```
