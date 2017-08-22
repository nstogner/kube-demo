# Kubernetes Demo

```sh
minikube start --kubernetes-version v1.6.4
eval $(minikube docker-env)
docker build -t kube-demo-img:v1 .
kubectl apply -f kubernetes.yaml
kubectl get deployment
kubectl get pods
kubectl get services
minikube service kube-demo --url
curl -v $(minikube service kube-demo --url)
```
