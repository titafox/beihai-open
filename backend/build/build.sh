#! /bin/bash
#docker run运行方式
docker build . -t coltstail-server

#rm -r ./*
docker stop coltstail-server
docker rm coltstail-server

# 配置通过 env 文件注入（参考 backend/.env.example）
docker run --name="coltstail-server"  --restart=always --network main -d  \
                --env-file ${ENV_FILE:-./.env}      \
                -p 8000:8000             \
                -v /data/docker/nfts:/srv/nfts      \
                -v /data/docker/tmp:/srv/tmp      \
                -v /data/docker/files:/srv/files      \
                -v /data/docker/static:/srv/static      \
                -v /data/docker/images:/srv/images      \
                coltstail-server
exit