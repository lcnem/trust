FROM centos:latest

RUN curl -O https://dl.google.com/go/go1.13.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.13.linux-amd64.tar.gz

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && \
    chmod -R 777 "$GOPATH"
WORKDIR $GOPATH

RUN yum update -y && \
    yum install git -y

RUN cd /usr/local/src && \
    git clone https://github.com/lcnem/lcnemint && \
    cd lcnemint && \
    git checkout master && \
    go mod download && \
    go install ./cmd/lcnemintd && \
    go install ./cmd/lcnemintcli && \
    cp scripts/lcnemintd.service /etc/systemd/system/lcnemintd.service && \
    systemctl enable lcnemintd && \
    firewall-cmd --add-port=26656/tcp --zone=public --permanent && \
    firewall-cmd --add-port=26657/tcp --zone=public --permanent && \
    firewall-cmd --reload

EXPOSE 26656
EXPOSE 26657
