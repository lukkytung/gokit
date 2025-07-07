# gokit

## 项目简介

gokit 是一个基于 Go 语言的项目模板，集成了常用的服务端开发组件。

## 配置说明

本项目的配置信息存放于 `config.yaml` 文件中。请根据实际环境修改相关配置后再启动项目。

### 配置项说明

- **server**
  - `port`：服务监听端口（如：8080）
- **database**
  - `host`：数据库主机地址（如：localhost）
  - `port`：数据库端口（如：5432）
  - `user`：数据库用户名（如：postgres）
  - `password`：数据库密码（如：123456）
  - `name`：数据库名称（如：gokit_example_db）
- **redis**
  - `host`：Redis 主机地址（如：localhost）
  - `port`：Redis 端口（如：6379）
  - `password`：Redis 密码（如：空字符串表示无密码）
  - `db`：Redis 数据库编号（如：0）

## 使用方法

1. **克隆项目**

   ```bash
   git clone <仓库地址>
   cd gokit
   ```

2. **修改配置**

   根据实际需求编辑 `config.yaml` 文件，填写正确的数据库和 Redis 连接信息。

3. **安装依赖**

   ```bash
   go mod tidy
   ```

4. **运行项目**

   ```bash
   go run main.go
   ```

5. **访问服务**

   默认服务运行在 `http://localhost:8080`，可根据 `config.yaml` 中的 `server.port` 配置进行调整。

---

如有问题请提交 issue 或联系维护者。

