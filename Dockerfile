FROM golang:1.19 AS build-stage

WORKDIR /app

COPY go.mod go.sum cmd internal config.json ./
RUN GOPROXY=direct go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping cmd/main.go

FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /docker-gs-ping /docker-gs-ping

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/docker-gs-ping"]
