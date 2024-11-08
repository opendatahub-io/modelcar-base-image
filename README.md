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

## Using

```sh
podman build -f Containerfile.modelcar -t quay.io/mmortari/demo20241108-base:modelcar .
podman push quay.io/mmortari/demo20241108-base:modelcar
skopeo inspect --raw docker://quay.io/mmortari/demo20241108-base:modelcar | jq
```

follow tutorial from https://kserve.github.io/website/latest/admin/kubernetes_deployment/#3-install-kserve

then:

```sh
./enable-modelcar.sh
kubectl apply -f isvc-modelcar.yaml
```

The problem seems to be here:

```yaml
    - args:
        - sh
        - -c
        - ln -s /proc/$$$$/root/models /mnt/models && sleep infinity
      image: quay.io/mmortari/demo20241108-base:modelcar
      imagePullPolicy: IfNotPresent
      name: modelcar
```

## Credits

Many thanks Jason for the idea and to Daniele
