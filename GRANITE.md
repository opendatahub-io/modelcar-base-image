
```sh
cd granite
huggingface-cli download ibm-granite/granite-3b-code-instruct-2k
```

podman manifest create quay.io/mmortari/demo20241108-base:granite
podman build --platform linux/amd64,linux/arm64 -f Containerfile-granite --manifest quay.io/mmortari/demo20241108-base:granite .
podman manifest push --all --rm quay.io/mmortari/demo20241108-base:granite
skopeo inspect --raw docker://quay.io/mmortari/demo20241108-base:granite | jq
podman image rm quay.io/mmortari/demo20241108-base:latest


podman manifest create quay.io/mmortari/demo20241108-base:micro-granite
podman build --platform linux/amd64,linux/arm64 -f Containerfile-micro-granite --manifest quay.io/mmortari/demo20241108-base:micro-granite .
podman manifest push --all --rm quay.io/mmortari/demo20241108-base:micro-granite
skopeo inspect --raw docker://quay.io/mmortari/demo20241108-base:micro-granite | jq
podman image rm quay.io/mmortari/demo20241108-base:latest
