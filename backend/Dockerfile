ARG GO_VERSION=1
FROM golang:${GO_VERSION}-alpine as builder

WORKDIR /usr/src/app
RUN apk add --no-cache make 

COPY go.mod go.sum ./
COPY . .
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN go install github.com/pressly/goose/v3/cmd/goose@latest
RUN make install
RUN make db/up
RUN make swagger
RUN make build/api

FROM debian:bookworm

ARG DB_DSN
ENV DB_DSN=${DB_DSN}
ARG API_URL
ENV API_URL=${API_URL}
ARG AUTH_SECRET
ENV AUTH_SECRET=${AUTH_SECRET}

RUN apt-get update && apt-get install -y \
  ca-certificates \
  && rm -rf /var/lib/apt/lists/*

COPY --from=builder /bin/api /usr/local/bin/
CMD ["api", "-env", "development"]
