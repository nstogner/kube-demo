# Kubernetes Demo

```sh
# Start a kubernetes cluster running on a local VM
minikube start --kubernetes-version v1.6.4

# Setup docker client to point to the docker instance running inside VM
eval $(minikube docker-env)

# Build the docker image
docker build -t kube-demo-img:v1 .

# Create the kubernetes resources
kubectl apply -f kubernetes.yaml

# Inspect things
kubectl get deployments
kubectl get pods
kubectl get services

# Inspect the exposed service url
minikube service kube-demo --url

# Hit the service
curl -v $(minikube service kube-demo --url)
```
