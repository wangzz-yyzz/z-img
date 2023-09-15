FROM golang:1.20

ENV GOPROXY https://goproxy.cn,direct
WORKDIR /go/src/wzz/z-img
COPY . /go/src/wzz/z-img
RUN go build .
EXPOSE 8086

ENTRYPOINT ["./z-img"]