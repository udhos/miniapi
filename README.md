# miniapi
miniapi

```
./build.sh

miniapi

curl localhost:8080/v1/world
```

## Docker

Docker hub:

https://hub.docker.com/r/udhos/miniapi

Pull from docker hub:

```
docker pull udhos/miniapi:0.0.0
```

Build recipe:

```
./docker/build.sh

docker push -a udhos/miniapi
```

## Helm chart

### Create

```
mkdir charts
cd charts
helm create miniapi
```

Then edit files.

### Lint

```
helm lint ./charts/miniapi --values charts/miniapi/values.yaml
```

### Test redering chart templates locally

```
helm template miniapi ./charts/miniapi --values charts/miniapi/values.yaml
```

### Render templates at server

```
helm install miniapi ./charts/miniapi --values charts/miniapi/values.yaml --dry-run
```

### Package chart into file

```
helm package ./charts/miniapi
Successfully packaged chart and saved it to: miniapi-0.1.3.tgz
```

A chart repository is an HTTP server that houses one or more packaged charts.
A chart repository is an HTTP server that houses an index.yaml file and optionally some packaged charts.

### Install

```
helm install miniapi ./charts/miniapi --values charts/miniapi/values.yaml
```

### Upgrade

```
helm upgrade miniapi ./charts/miniapi --values charts/miniapi/values.yaml
```

### Uninstall

```
helm uninstall miniapi
```
