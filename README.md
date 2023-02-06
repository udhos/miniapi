[![license](http://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/udhos/gateboard/blob/main/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/udhos/miniapi)](https://goreportcard.com/report/github.com/udhos/miniapi)
[![Go Reference](https://pkg.go.dev/badge/github.com/udhos/miniapi.svg)](https://pkg.go.dev/github.com/udhos/miniapi)
[![Artifact Hub](https://img.shields.io/endpoint?url=https://artifacthub.io/badge/repository/miniapi)](https://artifacthub.io/packages/search?repo=miniapi)
[![Docker Pulls](https://img.shields.io/docker/pulls/udhos/miniapi)](https://hub.docker.com/r/udhos/miniapi)

# miniapi
miniapi

```
./build.sh

miniapi
```

# test

```
curl localhost:8080/v1/world

# JSON
curl -d '{"hello":"world"}' -H 'account: 4321'  'localhost:8080/v1/hello?a=b' | jq

# multipart/form-data
curl -F param1=value1 -F param2=value2 'localhost:8080/v1/world?a=b' | jq

# application/x-www-form-urlencoded
curl -H "Content-Type: application/x-www-form-urlencoded" -d "param1=value1&param2=value2" 'localhost:8080/v1/world?a=b' | jq
```

## Docker

Docker hub:

https://hub.docker.com/r/udhos/miniapi

Run from docker hub:

```
docker run -p 8080:8080 --rm udhos/miniapi:0.0.1
```

Build recipe:

```
./docker/build.sh

docker push -a udhos/miniapi
```

## Helm chart

### Using the repository

See <https://udhos.github.io/miniapi/>.

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

### Test rendering chart templates locally

```
helm template miniapi ./charts/miniapi --values charts/miniapi/values.yaml
```

### Render templates at server

```
helm install miniapi ./charts/miniapi --values charts/miniapi/values.yaml --dry-run
```

### Generate files for a chart repository

A chart repository is an HTTP server that houses one or more packaged charts.
A chart repository is an HTTP server that houses an index.yaml file and optionally (*) some packaged charts.

(*) Optionally since the package charts could be hosted elsewhere and referenced by the index.yaml file.

    docs
    ├── index.yaml
    └── miniapi-0.1.3.tgz

See script [update-charts.sh](update-charts.sh):

    # generate chart package from source
    helm package ./charts/miniapi -d ./docs

    # regenerate the index from existing chart packages
    helm repo index ./docs --url https://udhos.github.io/miniapi/

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
