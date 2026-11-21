#! /bin/bash
# 生产环境部署：交叉编译 → 上传 → 在服务器上执行 build/build-prod.sh
# 用法：DEPLOY_HOST=root@<生产服务器IP> ./build-prod.sh
set -e
: "${DEPLOY_HOST:?请设置 DEPLOY_HOST，例如 root@1.2.3.4}"
DEPLOY_DIR=${DEPLOY_DIR:-/root/coltstail-server}

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./build/coltstail-server main.go
scp ./build/* "$DEPLOY_HOST:$DEPLOY_DIR"
ssh "$DEPLOY_HOST" "cd $DEPLOY_DIR && ./build-prod.sh"
