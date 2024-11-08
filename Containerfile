FROM registry.access.redhat.com/ubi8/go-toolset:1.22 AS build-stage

USER root

WORKDIR /app

COPY go.mod go.sum ./
COPY *.go ./
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o ./tailf


FROM scratch AS build-release-stage

WORKDIR /app
COPY --from=build-stage /app/tailf ./tailf

ENTRYPOINT ["/app/tailf"]
