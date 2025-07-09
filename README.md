# gokit

## 项目简介

gokit 是一个基于 Go 语言的服务端项目模板，集成了常用的开发组件，方便快速搭建后端服务。

## 环境配置

本项目的所有配置信息均存放于项目根目录下的 `.env` 文件中。**请务必在运行项目前，按照实际环境修改 `.env` 文件内容。**

### `.env` 配置项说明

| 配置项              | 说明                   | 示例值                |
|---------------------|------------------------|-----------------------|
| SERVER_PORT         | 服务监听端口           | 8080                  |
| DATABASE_HOST       | 数据库主机地址         | localhost             |
| DATABASE_PORT       | 数据库端口             | 5432                  |
| DATABASE_USER       | 数据库用户名           | postgres              |
| DATABASE_PASSWORD   | 数据库密码             | 123456                |
| DATABASE_NAME       | 数据库名称             | gokit_example_db      |
| REDIS_HOST          | Redis 主机地址         | localhost             |
| REDIS_PORT          | Redis 端口             | 6379                  |
| REDIS_PASSWORD      | Redis 密码（可为空）   |                       |
| REDIS_DB            | Redis 数据库编号       | 0                     |
| EMAIL_FROM          | Email发送账号          | XXXXXXXXXX@163.com    |
| EMAIL_SMTP_HOST     | 邮箱服务器             | smtp.163.com           |
| EMAIL_USER          | 发送者名称              | XXXXXXXXXX@163.com     |
| EMAIL_PASS          | 邮箱发送密码             | abc123456dxs           |
| EMAIL_SMTP_PORT     | 邮箱发送端口              |465                    | 
| JWT_SECRET_KEY      | jwt安全键                |jwt_sc_key          |

> 请参考项目根目录下的 `.env` 文件示例，填写正确的数据库和 Redis 连接信息。

## 使用方法

1. **克隆项目**

   ```bash
   git clone <仓库地址>
   cd gokit
   ```

2. **配置环境变量**

   编辑 `.env` 文件，填写正确的数据库和 Redis 连接信息。

3. **安装依赖**

   ```bash
   go mod tidy
   ```
4. **初始化gokit**
   ```
   // 初始化Gokit
	cmd.InitGokit()

	// 自动迁移数据库
	service.DB.AutoMigrate(&model.User{})
   ```

5. **运行项目**

   ```bash
   go run main.go
   ```

6. **访问服务**

   默认服务运行在 `http://localhost:8080`，可通过 `.env` 文件中的 `SERVER_PORT` 配置进行调整。

---

如有疑问请提交 issue 或联系维护者。

