FROM centos
WORKDIR /srv
COPY coltstail-server /srv/coltstail-server
RUN cd /srv
EXPOSE 8000
ENV GOPATH /root/go
ENV ENV prod
ENTRYPOINT ["./coltstail-server"]