ARG GO_VERSION=1
FROM golang:${GO_VERSION}-alpine as builder

WORKDIR /usr/src/app
RUN apk add --no-cache make 

COPY go.mod go.sum ./
COPY . .
RUN make install
RUN make build/api


FROM debian:bookworm

COPY --from=builder /bin/api /usr/local/bin/
CMD ["api", "-env", "production"]