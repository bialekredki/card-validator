# syntax=docker/dockerfile:1

FROM golang:1.21.5 AS build

WORKDIR /app

COPY go.mod  ./

RUN go mod download

COPY internal ./internal
COPY main.go ./

RUN CGO_ENABLED=0 go build -o /card-validator


FROM build AS test

RUN go test -v ./...


FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build /card-validator /card-validator

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/card-validator"]


