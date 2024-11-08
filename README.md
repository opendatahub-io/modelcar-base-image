# A minimal base image for sidecar puposes that does "nothing"

## Local dev

```sh
podman build -f Containerfile -t mio .
podman run -it mio
```

## Publishing

```sh
podman build -f Containerfile -t quay.io/mmortari/demo20241108-base .
podman push quay.io/mmortari/demo20241108-base 
skopeo inspect --raw docker://quay.io/mmortari/demo20241108-base | jq
```

## Credits

Many thanks Jason for the idea and to Daniele
