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