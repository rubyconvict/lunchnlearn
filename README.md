# lunch-n-learn

## Lunch & Learn - Golang

The presentation slides are here:

[Lunch & Learn - Golang - An Intro to Go & gRPC](lunch-and-learn-golang.pdf)

## Workshop

Let's build a food ordering API.

Project is composed of a simple http API gateway service, communicating via gRPC (client),
with a backend microservice (server).

### Run the apps

Open 1st terminal window:

`cd client/ && go run main.go`

Open 2nd terminal window:

`cd server/ && go run main.go`

### See it in action

Visit this urls in a web browser:

Successfully order: 'fish':

`http://localhost:8080/new_order/fish`

#### Experiment with different errors:

Fail to order: 'pasta':

`http://localhost:8080/new_order/pasta`

Fail to order: 'pizza':

`http://localhost:8080/new_order/pizza`

Use invalid parameter: 'car':

`http://localhost:8080/new_order/car`

## Development

Update libs:

```bash
go get -v ./... && go mod vendor && go mod tidy
```

Regenerate protobuf file(s):

```bash
protoc --go_out=proto --go-grpc_out=proto ./proto/lunchnlearn/*.proto && \
mv proto/github.com/rubyconvict/lunchnlearn/proto/lunchnlearn/*.pb.* proto/lunchnlearn/ && \
rm -rf proto/github.com
```
