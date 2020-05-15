FROM golang:latest

RUN GO111MODULE=on go get golang.org/x/tools/gopls@latest

RUN go get -u -v \
    github.com/stamblerre/gocode \
    github.com/ramya-rao-a/go-outline \
    github.com/newhook/go-symbols \
    github.com/uudashr/gopkgs/cmd/gopkgs \
    golang.org/x/tools/cmd/guru \
    golang.org/x/tools/cmd/gorename \
    golang.org/x/tools/cmd/goimports \
    github.com/rogpeppe/godef \
    github.com/zmb3/gogetdoc \
    golang.org/x/lint/golint \
    github.com/go-delve/delve/cmd/dlv \
    && go build -o gocode-gomod github.com/stamblerre/gocode \
    && mv gocode-gomod $GOPATH/bin/

VOLUME /src
WORKDIR /src

ENTRYPOINT sh
