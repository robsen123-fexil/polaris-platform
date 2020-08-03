FROM golang:1.22-alpine AS build
WORKDIR /src
COPY go.mod ./
COPY internal ./internal
COPY cmd/gateway ./cmd/gateway
RUN go build -o /out/service ./cmd/gateway

FROM alpine:3.20
WORKDIR /app
COPY --from=build /out/service ./service
ENV POLARIS_PORT=8800
EXPOSE 8800
CMD ["./service"]
