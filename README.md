# Kubernetes-Cleaner

[![Go Report Card](https://goreportcard.com/badge/github.com/Cytram/kubernetes-cleaner)](https://goreportcard.com/report/github.com/Cytram/kubernetes-cleaner)

CronJob that remove pods from cluster.


# Usage

This cleaner will look in an given namespace and search for a given pod or a regex can be given.

Set Enviroment variables
```
export CLEANER_NAMESPACE=dev
export CLEANER_SEARCH=telepresence.*
```

Run application

```
$ ./cleaner
```

# Build

The cleaner can be build using the standard Go tool chain if you have it available.

```
go build
```

You can build inside a docker imagee as well.
This produces a `cleaner` image that can run with the binary as entry point.

```
docker build -t cleaner .
```

# Deployment

To deploy the Cleaner in Kubernetes, you can find a simple Kubernetes Cronjob and RBAC yaml in the `examples` folder. You have to change the cronjob and RBAC namespace so it has `LIST` and `DELETE` rights in Kubernetes.


To deploy it to your kubernetes cluster run the following commands:

```
kubectl apply -f examples/cronjob.yaml
kubectl apply -f examples/rbac.yaml
```
