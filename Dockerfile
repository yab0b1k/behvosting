############################
# STEP 1 build executable binary
############################
FROM golang:1.10-alpine3.8 as builder

RUN apk add --update git vim bash && rm -rf /var/cache/apk/*
RUN go get github.com/golang/dep/cmd/dep

ENV APP_NAME "behvosting"
ENV BEEGO_RUNMODE prod
ENV APP_CONF_PATH conf

RUN mkdir -p /go/src/${APP_NAME}
WORKDIR /go/src/${APP_NAME}

ADD Gopkg.toml /go/src/${APP_NAME}
ADD Gopkg.lock /go/src/${APP_NAME}
RUN dep ensure -v --vendor-only

ADD . /go/src/${APP_NAME}
ADD /etc/rexad /go/src/${APP_NAME}/conf
WORKDIR /go/src/${APP_NAME}
RUN GOGC=off go build -v -i
############################
# STEP 2 build a small image
############################
FROM alpine:latest
RUN apk add --update bash tzdata ca-certificates less curl
WORKDIR /app/
ENV APP_NAME "behvosting"

COPY --from=builder /go/src/${APP_NAME}/ .

RUN echo 'alias ll="ls -la"' >> ~/.bashrc

CMD /app/${APP_NAME}
EXPOSE 80
EXPOSE 8080