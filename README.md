# Note Sync

一个支持将笔记同步到 GitHub 的笔记管理系统。支持文章和问答两种形式，文章会同步为 Markdown 文件，问答会同步为 GitHub Issues。

> [!NOTE]
> 该项目是玩具项目，使用cursor进行开发，后续可能会优化重构，请勿用于生产环境

## 功能特点

- 🔐 GitHub OAuth2 认证
- 📝 Markdown 编辑器
- 💾 自动同步到 GitHub
- 🏷️ 文章和问答管理
- 🌓 自动跟随系统的暗色模式
- 🚀 前后端分离架构

## 技术栈

### 前端
- Vue 3
- TypeScript
- Naive UI
- Vue Router
- Pinia
- Axios

### 后端
- Go
- Gin
- GORM
- GitHub API

## 项目配置

### 前端配置
1. 进入前端目录
```bash
cd frontend
```

2. 安装依赖
```bash
npm install
```

3. 开发环境运行
```bash
npm run dev
```

4. 生产环境构建
```bash
npm run build
```

### 后端配置
1. 进入后端目录
```bash
cd backend
```

2. 安装依赖
```bash
go mod tidy
```

3. 配置环境变量
复制 `.env.example` 到 `.env` 并填写配置：
```env
# Server Configuration
SERVER_PORT=8080

# Database Configuration
DB_HOST=127.0.0.1
DB_PORT=3306
DB_USER=note_sync
DB_PASSWORD=your_password
DB_NAME=note_sync

# GitHub Configuration
GITHUB_CLIENT_ID=your_client_id
GITHUB_CLIENT_SECRET=your_client_secret
GITHUB_REPO_OWNER=your_username
GITHUB_REPO_NAME=your_repo

# Proxy Configuration (可选,用于本地开发时连接GitHub API)
PROXY_ENABLE=false
PROXY_URL=socks5://127.0.0.1:7890
```

4. 运行服务
```bash
go run main.go
```

## 数据库配置

1. 创建 MySQL 数据库
```sql
CREATE DATABASE note_sync CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

2. 创建用户并授权
```sql
CREATE USER 'note_sync'@'localhost' IDENTIFIED BY 'your_password';
GRANT ALL PRIVILEGES ON note_sync.* TO 'note_sync'@'localhost';
FLUSH PRIVILEGES;
```

## GitHub 配置

1. 创建 GitHub OAuth 应用
- 访问 https://github.com/settings/developers
- 点击 "New OAuth App"
- 填写应用信息：
  - Application name: Note Sync
  - Homepage URL: http://localhost:5173
  - Authorization callback URL: http://localhost:5173/login/callback

2. 创建存储仓库
- 创建一个新的 GitHub 仓库用于存储笔记
- 更新 `.env` 中的 `GITHUB_REPO_OWNER` 和 `GITHUB_REPO_NAME`

## 使用说明

1. 访问 http://localhost:5173
2. 使用 GitHub 账号登录
3. 开始创建和管理你的笔记
4. 笔记会自动同步到 GitHub 仓库

## 注意事项

- 确保 GitHub OAuth 应用的回调地址配置正确
- 确保有正确的 GitHub 仓库访问权限
- 如果在中国大陆使用，可能需要配置代理

## 许可证

MIT License
