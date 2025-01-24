# k8s-go-rest
simple restful api written in go for testing it on k8s

### Build Image From Dockerfile(alptht is dockerhub username)
docker build --tag k8s-go-rest .
docker build -t alptht/k8s-go-rest:multistage -f Dockerfile.multistage .

### Check Docker Images
docker image ls

### Tag Images
docker image tag k8s-go-rest:latest k8s-go-rest:v1.0

### Remove Tag
docker image rm k8s-go-rest:v1.0

### Run Image inside Container(-detached --name and --published on ports)
docker run -d -p 8080:8080 --name k8s-go-rest k8s-go-rest:multistage

### List Containers
docker ps -a

### Stop Container
docker stop k8s-go-rest

### Remove all non working container.
docker container prune

### Restart Container
docker restart k8s-go-rest

### Login to Docker
docker login

### Push to DockerHub(alptht is dockerhub username)
docker push alptht/k8s-go-rest:multistage

---

###### K8s IMPERATIVE

### Create New Deployment
kubectl create deployment k8s-go-rest-deployment --image=alptht/k8s-go-rest:multistage

### Show deployments
kubectl get deploy

### Create Service
kubectl expose deployment k8s-go-rest-deployment --port=8080
kubectl expose deployment k8s-go-rest-deployment --type=NodePort --port=8080
kubectl expose deployment k8s-go-rest-deployment --type=LoadBalancer --port=8080

### List Services
kubectl get services

### Delete Service
kubectl delete svc k8s-go-rest-deployment

### Scale Deployment
kubectl scale deployment k8s-go-rest-deployment --replicas=4

### Get Pods
kubectl get pods -o wide

### Minikube problem fix for sending request
minikube service k8s-go-rest

### Describe deployment
kubectl describe deployment 

### Delete everything(in Minikube kubernetes service will also be deleted, but automaticly recreated)
kubectl delete all --all

---

##### K8s DECLERATIVE

### Apply deployment and service
kubectl apply -f deployment.yaml
kubectl apply -f service.yaml

### Delete deployment and service
kubectl delete -f deployment.yaml
kubectl delete -f service.yaml
kubectl delete -f deployment.yaml -f service.yaml

### Try to resolve service name inside pod
kubectl exec <podName> -- nslookup <other-service-name>