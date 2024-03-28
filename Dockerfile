FROM ubuntu:latest

ENV TZ=Etc/UTC
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
RUN sed -i "s@http://.*archive.ubuntu.com@http://mirrors.aliyun.com@g" /etc/apt/sources.list && \
    sed -i "s@http://.*security.ubuntu.com@http://mirrors.aliyun.com@g" /etc/apt/sources.list
RUN apt-get -y update && apt-get -y install vim curl gcc g++ python3 python3-pip

WORKDIR /opt
ENTRYPOINT [ "/opt/go-judge" ]
COPY go-judge mount.yaml polaris.yaml /opt/
