FROM ubuntu

# switch apt mirror
RUN sed -i 's/archive.ubuntu.com/mirrors.tuna.tsinghua.edu.cn/g' \ 
    /etc/apt/sources.list &&\
    sed -i 's/archive.ubuntu.com/mirrors.tuna.tsinghua.edu.cn/g' \
    /etc/apt/sources.list.d/ubuntu.sources && \
    sed -i 's/security.ubuntu.com/mirrors.tuna.tsinghua.edu.cn/g' \ 
    /etc/apt/sources.list && \
    sed -i 's/security.ubuntu.com/mirrors.tuna.tsinghua.edu.cn/g' \
    /etc/apt/sources.list.d/ubuntu.sources && \
    apt update

# install wget
RUN apt install -y wget

# install go
RUN wget https://go.dev/dl/go1.25.3.linux-amd64.tar.gz && tar -C /usr/local/ -xzf go1.25.3.linux-amd64.tar.gz

ENV GOPATH=/usr/local/go 
ENV GOBIN=$GOPATH/bin/
ENV PATH=$PATH:$GOBIN

CMD [ "sleep","365d" ]