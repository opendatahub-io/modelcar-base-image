FROM --platform=$BUILDPLATFORM registry.access.redhat.com/ubi8/go-toolset:1.22 AS build-stage
ARG TARGETOS
ARG TARGETARCH

USER root
WORKDIR /app

COPY go.mod go.sum ./
COPY *.go ./
RUN go mod download

RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} CGO_ENABLED=0 go build -o ./link-model-and-wait .


FROM scratch AS build-release-stage
ARG MODELCAR_LINK_FROM="/models"
ARG MODELCAR_LINK_TO="/mnt/models"

WORKDIR /bin
COPY --from=build-stage /app/link-model-and-wait ./link-model-and-wait

ENTRYPOINT ["/bin/link-model-and-wait", ${MODELCAR_LINK_FROM}, ${MODELCAR_LINK_TO}]
CMD ["dummy1", "dummy2"]