# prod-rdy-go-microservice
Sample for production ready microservice in golang

# TODO
- Add logger (ex:zap)
- Add metrics endpoint (prometheus)
- Add health check endpoint
- Add Graceful shutdown mechanic
- Add Config reader
- Add server and http client timeouts

docker build -t prod-rdy-go-microservice:multistage -f Dockerfile.multistage .

docker run -d -p 8080:8080 --name rest-server prod-rdy-go-microservice:multistage