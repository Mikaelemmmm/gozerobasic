FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOOS linux


WORKDIR /build/zero

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
COPY app/gateway/api/v1/captcha/etc /app/etc
RUN go build -ldflags="-s -w" -o /app/captcha app/gateway/api/v1/captcha/captcha.go


FROM alpine

RUN apk update --no-cache
RUN apk add --no-cache ca-certificates
RUN apk add --no-cache tzdata
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/captcha /app/captcha
COPY --from=builder /app/etc /app/etc

CMD ["./captcha"]
