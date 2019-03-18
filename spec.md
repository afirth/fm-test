1. Auth
https://github.com/youngpm/gbdx/blob/master/api.go
1. Listen for Post
  1. Validate valid geojson
  https://github.com/paulmach/go.geojson
  1. extract AOI
1. Search catalogue
  1. Return results as json

## design

### golang server

handlers
middleware (inject auth)
tests
  - transcoding
  - http request (outgoing) is reasonable based on sample input
  - http response (return to client) is reasonable with mocked gdbx response
  - auth throws error if invalid
  - http request errors as expected
geojson->wkt transcoding
healthcheck

makefile for build of package
build package in container, then copy in scratch or minimal

add travis.yml, build badges
https://github.com/twpayne/go-geom/blob/master/.travis.yml

### docker-compose

straightforward. expose 80:8080

### kube

Use local docker registry
on OSX
```
minikube start --vm-driver=hyperkit
eval $(minikube docker-env)
```

create secret
create deployment
and service

apply the manifests

```
$ kubectl expose deployment <name> --type=NodePort #optional - only if no service manifest
$ curl $(minikube service <name> --url)
```

do this with say, `make up`

## Secrets
```
kubectl create secret generic gbdx-secret \
--from-literal=username=${USERNAME} \
--from-literal=password=${PASSWORD}
```

## Out of scope / not implemented

- Pagination of catalog responses [because this is broken](https://gbdxdocs.digitalglobe.com/docs/catalog-v2-course#section-searching-the-catalog)
- Searching with anything besides a single polygon [because this is also broken](https://gbdxdocs.digitalglobe.com/docs/catalog-v2-course#section-spatial-area-search-format)
- API doco (e.g. swagger) - for a larger app I would consider using go-swagger generation
- Testing for configuration and server setup
- Testing of the 
- Decent logging

## Known deviations from the spec

- The spec says to "return GeoJSON to the client" however no geometry is requested. I chose to return plain json.

- The spec asks for "all catalog entries". As pagination is currently not implemented by the `catalog/v2/search` endpoint, the limit is 1000 results

## Notes

My first choice of implementation language would be python to leverage the existing GDBX library and because it's not in the user latency path. However, I understand you're probably more interested in my go :)
