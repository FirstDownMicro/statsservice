FROM golang:alpine as builder
RUN mkdir -p $GOPATH/src/github.com/Nanixel/FirstDownMicro/statsservice
ADD . $GOPATH/src/github.com/Nanixel/FirstDownMicro/statsservice
RUN cd $GOPATH/src/github.com/Nanixel/FirstDownMicro/statsservice && go install

FROM alpine
WORKDIR /app
COPY --from=builder go/bin /app
ENTRYPOINT ./statsservice