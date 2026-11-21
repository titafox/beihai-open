#! /bin/bash
docker build . -t coltstail-server
docker-compose stop coltstail-server
docker-compose rm -f coltstail-server
docker-compose up -d coltstail-server
