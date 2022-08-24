DBDriver
A DB Driver for Go's database/sql package

## 初始化项目

安装完后，在命令行中启动如下 docker 容器：

```sh
# 创建 PostgreSQL 数据库
docker run --name postgres --restart always -d -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=testSQL -p 5432:5432  postgres
# 或者创建 MySQL 数据库（不推荐）
docker run --name mysql -d -e MYSQL_ROOT_PASSWORD=123456 -e MYSQL_DATABASE=testSQL -d -p 3306:3306 mysql

```
安装完可以创建张表
