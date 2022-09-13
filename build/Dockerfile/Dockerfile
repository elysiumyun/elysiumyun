FROM golang:1.19 AS builder

ENV GOPROXY https://proxy.golang.com.cn,direct

WORKDIR /usr/src/

COPY . .

RUN make tidy \
 && make build

# final runner image
FROM alpine:latest as runner

LABEL maintaner="Mardan M"

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories \
 && apk add tzdata

WORKDIR /usr/bin

COPY --from=builder /usr/src/bin/elysium /usr/bin/elysium

EXPOSE 8080

ENTRYPOINT ["/usr/bin/elysium"]