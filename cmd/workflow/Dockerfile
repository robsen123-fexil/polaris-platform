FROM golang:1.22-alpine AS build
WORKDIR /src
COPY go.mod ./
COPY internal ./internal
COPY cmd/workflow ./cmd/workflow
RUN go build -o /out/service ./cmd/workflow

FROM alpine:3.20
WORKDIR /app
COPY --from=build /out/service ./service
ENV POLARIS_PORT=8802
EXPOSE 8802
CMD ["./service"]
