#### Golang爬虫项目

###### 项目简介

- 爬去链家二手房相关信息

### 新建数据表

```sql
CREATE TABLE `houses`
(
    `id`          int unsigned NOT NULL AUTO_INCREMENT,
    `city`        varchar(100) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
    `name`        varchar(100) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
    `total_price` varchar(50) COLLATE utf8mb4_general_ci  NOT NULL DEFAULT '',
    `address`     varchar(100) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
    `dan_price`   varchar(50) COLLATE utf8mb4_general_ci  NOT NULL DEFAULT '',
    `info`        varchar(200) COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
    `created_at`  datetime                                         DEFAULT NULL,
    `updated_at`  datetime                                         DEFAULT NULL,
    `url`         varchar(200) COLLATE utf8mb4_general_ci NOT NULL,
    `area`        varchar(100) COLLATE utf8mb4_general_ci NOT NULL,
    PRIMARY KEY (`id`),
    KEY           `city` (`city`)
) ENGINE=InnoDB AUTO_INCREMENT=361 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
```

- 修改 .env 相关数据库配置


- 添加爬去城市信息

```azure
修改文件 helpers/const.go 新增爬取城市
```

- 运行项目

```golang
// cd:表示成都 all:暂时只支持成都、北京、上海,如果想支持更多城市，在配置里面加一下就可以了
go run main.go spider -c cd | all
```