# gokit

## 项目简介

gokit 是一个基于 Go 语言的项目模板，集成了常用的服务端开发组件。

## 配置说明

本项目的配置信息存放于 `.env` (该文件需要与main.go同级)文件中。请根据实际环境修改相关配置后再启动项目。

### `.env` 配置项说明

- **SERVER_PORT**：服务监听端口（如：8080）
- **DATABASE_HOST**：数据库主机地址（如：localhost）
- **DATABASE_PORT**：数据库端口（如：5432）
- **DATABASE_USER**：数据库用户名（如：postgres）
- **DATABASE_PASSWORD**：数据库密码（如：123456）
- **DATABASE_NAME**：数据库名称（如：gokit_example_db）
- **REDIS_HOST**：Redis 主机地址（如：localhost）
- **REDIS_PORT**：Redis 端口（如：6379）
- **REDIS_PASSWORD**：Redis 密码（如为空则表示无密码）
- **REDIS_DB**：Redis 数据库编号（如：0）

请参考项目根目录下的 `.env` 文件示例，填写正确的数据库和 Redis 连接信息。

## 使用方法

1. **克隆项目**

   ```bash
   git clone <仓库地址>
   cd gokit
   ```

2. **修改配置**

   根据实际需求编辑 `.env` 文件，填写正确的数据库和 Redis 连接信息。

3. **安装依赖**

   ```bash
   go mod tidy
   ```

4. **运行项目**

   ```bash
   go run main.go
   ```

5. **访问服务**

   默认服务运行在 `http://localhost:8080`，可根据 `.env` 文件中的 `SERVER_PORT` 配置进行调整。

---

如有问题请提交 issue 或联系维护者。

