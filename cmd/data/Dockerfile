FROM golang:1.22-alpine AS build
WORKDIR /src
COPY go.mod ./
COPY internal ./internal
COPY cmd/data ./cmd/data
RUN go build -o /out/service ./cmd/data

FROM alpine:3.20
WORKDIR /app
COPY --from=build /out/service ./service
ENV POLARIS_PORT=8803
EXPOSE 8803
CMD ["./service"]
