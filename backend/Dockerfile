FROM golang:1.22-alpine as builder

WORKDIR /app
RUN apk add --no-cache make 

COPY . ./
RUN make install
RUN make build/api
RUN > /app/.env

FROM alpine:latest

COPY --from=builder /bin/api /api
COPY --from=builder /app/.env .env

EXPOSE 8080
ENTRYPOINT [ "./api" ]