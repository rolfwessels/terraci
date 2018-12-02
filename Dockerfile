FROM golang:1.10-alpine3.7


RUN apk update
RUN apk upgrade
RUN apk add ca-certificates wget && update-ca-certificates
RUN apk add --no-cache --update \
    curl \
    git \
    zip \
    ncurses \
    busybox

RUN go get github.com/ahmetb/go-linq
RUN go get github.com/gorilla/handlers
RUN go get github.com/gorilla/mux
RUN go get github.com/tkanos/gonfig
RUN go get github.com/urfave/cli
RUN go get github.com/smartystreets/assertions


RUN rm /var/cache/apk/*

#install terraform
RUN wget https://releases.hashicorp.com/terraform/0.11.10/terraform_0.11.10_linux_amd64.zip
RUN unzip terra*
RUN mv terraform /usr/local/bin/
RUN terraform version
RUN rm -f terra*

WORKDIR /go/src/github.com/rolfwessels/continues-terraforming
#COPY ./ ./
COPY ./readme.md ./


#RUN go build main.go terrastate.go cli.go terraform.go


CMD ["top"]
