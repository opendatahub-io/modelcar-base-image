FROM --platform=$BUILDPLATFORM registry.access.redhat.com/ubi8/go-toolset:1.22@sha256:25f2884bcf8ba92eca7613c8dbcbdfb3e4951db2875732b01486628d57623745 AS build-stage
ARG TARGETOS
ARG TARGETARCH

USER root
WORKDIR /app

COPY go.mod go.sum ./
COPY *.go ./
RUN go mod download

RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} CGO_ENABLED=0 go build -o ./link-model-and-wait .


FROM scratch AS build-release-stage

WORKDIR /bin
COPY --from=build-stage /app/link-model-and-wait ./link-model-and-wait

ENTRYPOINT ["/bin/link-model-and-wait"]
