# Run in docker container

```
docker run -p 1234:8088 -d diptadas/grpc-example
docker ps
docker logs -f <container_id>
curl --data "{\"name\": \"dipta\"}" http://localhost:1234/echo
```

# Deploy in kubernetes (minikube)

```
minikube start
kubectl cluster-info
kubectl create -f k8s_grpc_example.yaml
kubectl get pods
kubectl logs -f <pod_name>
```

**Access using service**

```
kubectl expose deployment grpc-example --type=LoadBalancer --name=grpc-example-svc
minikube service --url grpc-example-svc
curl --data "{\"name\": \"dipta\"}" <url>/echo
```

**Access using port-forwarding**

```
kubectl port-forward <pod_name> 8088
curl --data "{\"name\": \"dipta\"}" http://127.0.0.1:8088/echo
```
