DBDriver
A DB Driver for Go's database/sql package

## 初始化项目

1、安装完后，在命令行中启动如下 docker 容器：

```sh
# 创建 PostgreSQL 数据库
docker run --name postgres --restart always -d -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=testSQL -p 5432:5432  postgres
# 或者创建 MySQL 数据库（不推荐）
docker run --name mysql -d -e MYSQL_ROOT_PASSWORD=123456 -e MYSQL_DATABASE=testSQL -d -p 3306:3306 mysql

```
2、安装完可以创建张表
```markdown

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
`id` int NOT NULL AUTO_INCREMENT,
`name` varchar(255) NOT NULL,
`age` int DEFAULT NULL,
PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of users
-- ----------------------------
BEGIN;
INSERT INTO `users` (`id`, `name`, `age`) VALUES (1, 'hell0', 12);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
```

3、执行测试
go test -v ./client_test.go