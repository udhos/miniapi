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
mkdir chart
cd chart
helm create miniapi
```

### Install

```
helm install miniapi chart/miniapi/ --values chart/miniapi/values.yaml
```

### Upgrade

```
helm upgrade miniapi chart/miniapi/ --values chart/miniapi/values.yaml
```

### Uninstall

```
helm uninstall miniapi
```
