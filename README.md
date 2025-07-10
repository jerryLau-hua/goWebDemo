# Awesome Go Project - go web 基础框架搭建

基于Gin框架构建的RESTful API服务，提供基础用户管理功能

## 技术栈

- **语言**: Go 1.24.4
- **Web框架**: Gin v1.10.1
- **数据库**: MySQL
- **配置管理**: YAML
- **依赖管理**: Go Modules

## 项目结构

```text
/awesomeProject
. ├── cmd/ # 入口文件 
│ └── server/ 
│ └── main.go # 主程序入口 
├── configs/ # 配置文件 
│ └── config.dev.yaml 
├── internal/ # 内部模块 
│ ├── config/ # 配置加载 
│ ├── database/ # 数据库连接 
│ ├── models/ # 数据模型 
│ ├── repository/ # 数据访问层 
│ └── service/ # 业务逻辑层 
├── transport/ # 传输层 
│ └── http/ # HTTP处理 
└── go.mod # 依赖管理
```

## 快速开始

### 环境要求

- Go 1.24+
- MySQL 5.7+

### 安装步骤

```bash
 # 克隆项目 
 git clone https://github.com/yourusername/awesome-project.git 
 cd awesome-project 
 # 安装依赖 
 go mod download 
 # 启动MySQL服务（确保已安装） 
 mysqld --console 
 # 运行项目 
 go run cmd/server/main.go
```
### 配置说明
修改 configs/config.dev.yaml 配置数据库连接：
```yml
database:
  dsn: "root:yourpassword@tcp(127.0.0.1:3306)/go_test"
  max_open_conns: 25
  max_idle_conns: 25
```
### API文档
#### 用户注册
```text
POST /api/v1/users
Content-Type: application/json

{
  "name": "张三",
  "email": "zhangsan@example.com"
}
```
#### 用户查询
```text
GET /api/v1/users/{id}
```


### 常见问题
- Q: 数据库连接失败 
- A: 检查MySQL服务状态和config.dev.yaml中的账号密码配置
- Q: 端口8080被占用 
- A: 修改main.go中router.Run(":8080")的端口号