# 安装指南

1. 克隆项目
   ```bash
   git clone https://github.com/your-repo/IntraTunnel.git

2. 生成 TLS 证书
   ```bash
   cd IntraTunnel/certs && openssl req -x509 -newkey rsa:4096 -keyout key.pem -out cert.pem -days 365 -nodes

3. 配置客户端和服务端
   编辑 config/client_config.json 和 config/server_config.json 以适应您的需求。

4. 运行服务端
   ```bash
   ./scripts/start_server.sh

5.	运行客户端
   ```bash
   ./scripts/start_client.sh