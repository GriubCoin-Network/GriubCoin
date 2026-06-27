FROM golang:1.21-alpine AS build-env

RUN apk add --no-cache bash git make libc-dev gcc linux-headers eudev-dev jq curl

WORKDIR /root/grbc

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN make build

FROM alpine:3.20

RUN apk add --no-cache bash jq curl

COPY --from=build-env /root/grbc/out/linux/grbc /bin/grbc

CMD ["grbc"]