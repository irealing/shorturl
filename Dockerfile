FROM alpine:3.14
LABEL name="github.com/irealing/shorturl" version="0.1"
ENV host="0.0.0.0" port=80 dbname="data.db"
ARG BRANCH=master
ARG GOVERSION=1.17.2
WORKDIR /data
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g'  /etc/apk/repositories && \
    apk update && \
    apk add curl git tzdata libc6-compat && \
    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone &&\
    export GOROOT=/data/go && \
    export GOPATH=/data/.go &&\
    export GOPROXY=https://goproxy.cn &&\
    export GO_PKG=go${GOVERSION}.linux-amd64.tar.gz && \
    curl -o $GO_PKG https://dl.google.com/go/$GO_PKG && \
    tar -xvf $GO_PKG  &&\
    export GOROOT=$PWD/go &&\
    export PATH=$PATH:$GOROOT/bin && \
    git clone -b ${BRANCH} https://github.com/irealing/shorturl.git --depth=1  &&\
    cd shorturl && \
    go mod vendor && \
    go build -o /usr/local/bin/shorturl shorturl/web && \
    echo -e "#!/bin/sh\nshorturl -host ${host} -port ${port} -data ${dbname} "> /usr/local/bin/run.sh &&\
    chmod +x /usr/local/bin/run.sh &&\
    cd / && \
    rm -rf /data && \
    rm -rf $GOROOT ~/.cache /var/cache/apk
VOLUME [ "/data"]
EXPOSE 80/tcp
CMD [ "run.sh" ] 
