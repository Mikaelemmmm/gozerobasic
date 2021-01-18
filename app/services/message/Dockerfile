FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOOS linux


WORKDIR /build/zero

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
COPY app/services/message/etc /app/etc
RUN go build -ldflags="-s -w" -o /app/message app/services/message/message.go


FROM alpine

RUN apk update --no-cache
RUN apk add --no-cache ca-certificates
RUN apk add --no-cache tzdata
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/message /app/message
COPY --from=builder /app/etc /app/etc

CMD ["./message", "-f", "etc/message.yaml"]
