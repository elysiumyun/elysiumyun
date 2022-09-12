FROM golang:1.19 AS builder

ARG GOPROXY https://goproxy.io,direct

WORKDIR /usr/src/

COPY . .

RUN go mod tidy
RUN make build

# final runner image
FROM alpine:latest as runner

LABEL maintaner="Mardan M"

WORKDIR /usr/bin

COPY --from=builder /usr/src/bin/elysium /usr/bin/elysium

EXPOSE 8080

ENTRYPOINT ["/usr/bin/elysium"]