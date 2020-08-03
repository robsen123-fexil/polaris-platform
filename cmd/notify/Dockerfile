FROM golang:1.22-alpine AS build
WORKDIR /src
COPY go.mod ./
COPY internal ./internal
COPY cmd/notify ./cmd/notify
RUN go build -o /out/service ./cmd/notify

FROM alpine:3.20
WORKDIR /app
COPY --from=build /out/service ./service
ENV POLARIS_PORT=8804
EXPOSE 8804
CMD ["./service"]
