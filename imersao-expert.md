`go build`

`./pokemon-api`

*OUTPUT:* Welcome to POKEMON API from 192.168.0.15

*Test the application:*

`curl http://localhost:8080 `

*OUTPUT:* Welcome to POKEMON API from 192.168.0.15

Test the other endpoints of our POKEMON API:

`curl http://localhost:8080/pikachu`

`curl http://localhost:8080/bulbasaur`

`curl http://localhost:8080/squirtle`

`curl http://localhost:8080/charmander`


----------

Dockerfile

```
# Start from the latest golang base image
FROM golang:alpine

# Add Maintainer Info
LABEL maintainer="Andre Almar <andre@y7mail.com>"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o pokemon-api .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./pokemon-api"]
```
----------

*Build the image:*

```
docker build -t pokemon-api-v1 .
```
*Run the Docker Image:*

```
docker run -d -p 8080:8080 pokemon-api-v1
```

See the container running: 

```
docker container ls
```

Interact with the app inside the running container:

`curl http://localhost:8080`

*OUTPUT:* Welcome to POKEMON API from 172.17.0.2

Test the other endpoints of our POKEMON API running now INSIDE the container:

`curl http://localhost:8080/pikachu`

`curl http://localhost:8080/bulbasaur`

`curl http://localhost:8080/squirtle`

`curl http://localhost:8080/charmander`

----------

**MULTISTAGE BUILDING**

Dockerfile Multistage

```
# build stage
FROM golang:alpine AS build-env
RUN apk --no-cache add build-base git mercurial gcc
ADD . /src
RUN cd /src && go build -ldflags '-w -s' -a -o pokemon-api

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /src/pokemon-api /app/
EXPOSE 8080
ENTRYPOINT ./pokemon-api

```

Build the optimized Docker image:

```
docker build -t go-docker-optimized -f Dockerfile.multistage .
```

Check if the image is created:

```
docker images

REPOSITORY            TAG                 IMAGE ID            CREATED             SIZE
go-docker-optimized   latest              4ba2c8b851fe        7 seconds ago       14MB
```

Comparing the 2 images:

```
REPOSITORY            TAG                 IMAGE ID            CREATED             SIZE
go-docker-optimized   latest              4ba2c8b851fe        7 seconds ago       14MB
go-docker             latest              8be37dd67655        10 minutes ago      819MB
```

Now interact with the Optimized Docker image:

`curl http://localhost:8080 `

*OUTPUT:* Hello, Guest

`curl http://localhost:8080\?name=Andre`

*OUTPUT:* Hello, Andre


----------

Create your Kubernetes cluster (AKS)

*Launch Cloud Shell*

Then type the command below to create a resource group:

```
az group create --name myResourceGroup --location eastus
```

The output must be something like this:

```
{
  "id": "/subscriptions/<guid>/resourceGroups/myResourceGroup",
  "location": "eastus",
  "managedBy": null,
  "name": "myResourceGroup",
  "properties": {
    "provisioningState": "Succeeded"
  },
  "tags": null
}
```

Don't forget to register the following Subscription:

`az provider register --namespace Microsoft.Network`

Type the following command to create the AKS cluster:

```
az aks create --resource-group myResourceGroup --name myAKSCluster --node-count 1 --enable-addons monitoring --generate-ssh-keys
```

The following example creates a cluster named myAKSCluster with one node. Azure Monitor for containers is also enabled using the --enable-addons monitoring parameter. This will take several minutes to complete.

After a few minutes, the command completes and returns JSON-formatted information about the cluster.

**Deploying our app on AKS**

Type the following:

To create a namespace: 

```
apiVersion: v1
kind: Namespace
metadata:
  name: imersao
  labels:
    name: imersao
```

and then

`kubectl apply -f namespace.yaml`

To create a deployment:

```
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  name: go-docker-app-deployment
  namespace: imersao
spec:
  replicas: 2
  template:
    metadata:
      labels:
        app: go-docker
    spec:
      containers:
        - name: go-docker
          image: andrealmar/go-docker-optimized:latest
          imagePullPolicy: Always
          resources:
            limits:
              cpu: 400m
              memory: 300Mi
            requests:
              cpu: 200m
              memory: 150Mi
```

and then

`kubectl apply -f deployment.yaml`

To create a service (to expose our deployment to external world)

```
apiVersion: v1
kind: Service
metadata:
  name: go-docker-service
  namespace: imersao
spec:
  type: LoadBalancer
  ports:
    - name: http
      protocol: TCP
      port: 8080
      targetPort: 8080
  selector:
    app: go-docker

```

and then

`kubectl apply -f service.yaml`

Check if the service was created:

`kubectl get service go-docker-service -n imersao --watch`

Wait a few minutes to get the EXTERNAL IP:

```
➜  kubernetes kubectl get service go-docker-service -n imersao --watch
NAME                TYPE           CLUSTER-IP    EXTERNAL-IP    PORT(S)          AGE
go-docker-service   LoadBalancer   10.0.67.193   20.42.32.175   8080:31702/TCP   2m24s
```

Now try to access the app (already deployed on AKS) from your browser:

`http://20.42.32.175:8080/?name=Andre`

*OUTPUT:* Hello, Andre

Now DELETE the cluster so you don't get charged ;)

`az group delete --name myResourceGroup --yes --no-wait`

references: https://docs.microsoft.com/en-us/azure/aks/kubernetes-walkthrough
from zero to a bank in one year: https://www.starlingbank.com
