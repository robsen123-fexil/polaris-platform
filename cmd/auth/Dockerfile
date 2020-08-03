FROM golang:1.22-alpine AS build
WORKDIR /src
COPY go.mod ./
COPY internal ./internal
COPY cmd/auth ./cmd/auth
RUN go build -o /out/service ./cmd/auth

FROM alpine:3.20
WORKDIR /app
COPY --from=build /out/service ./service
ENV POLARIS_PORT=8801
EXPOSE 8801
CMD ["./service"]
