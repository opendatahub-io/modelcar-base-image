in jupyter

```sh
VERSION="1.2.0"
curl -LO "https://github.com/oras-project/oras/releases/download/v${VERSION}/oras_${VERSION}_linux_amd64.tar.gz"
mkdir -p oras-install/
tar -zxf oras_${VERSION}_*.tar.gz -C oras-install/
mkdir .local/bin
mv oras-install/oras /usr/local/bin/
rm -rf oras_${VERSION}_*.tar.gz oras-install/
```

failed with

skopeo copy docker://quay.io/mmortari/demo20241108-base docker://modelregistry-poc.quaydev.org/repository/mmortari/demo20241108-base-private --multi-arch all

