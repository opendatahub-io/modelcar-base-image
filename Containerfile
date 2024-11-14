FROM registry.access.redhat.com/ubi8/go-toolset:1.22 AS build-stage

USER root

WORKDIR /app

COPY go.mod go.sum ./
COPY *.go ./
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o ./link-model-and-wait


FROM scratch AS build-release-stage

WORKDIR /bin
COPY --from=build-stage /app/link-model-and-wait ./link-model-and-wait

ENTRYPOINT ["/bin/link-model-and-wait"]
