# pixicoreAPI

## How to run the API

- Install [Go](https://nats.io/documentation/tutorials/go-install/).

- Git clone the repo into `$GOPATH/src/github.com/jalilbengoufa/`.

- Enter the `pixicoreAPI` directory. Run `go get ./...` which will install every dependencies.

- `go test ./... && go build ./cmd/pixicoreAPI && ./pixicoreAPI` will run the tests, build the program and run it.

#### Using Docker

- `docker build -t pixicoreapi . `

- `docker run -d -p 3000:3000 pixicoreapi `

## Usage

- Change the IP address for each server in the `servers-config.yaml` file

- You can run `curl -i http://localhost:3000/v1/install/SERVER_MAC_ADDRESS`  this will collect info  and install coreOS for the server

- You can run `curl -i http://localhost:3000/v1/all/`  this will collect info  and install coreOS for each server

## API Endpoints

#### `GET v1/boot/:macAddress` 

- Used by pixicore to get PXE config and boot each server (each server have a IP address assigned).

#### `GET v1/install/:macAddress` 

- get information (cores,ram,etc) from the server using her macAddress as ID and install coresOS. 

#### `GEt v1/all` 

- get information (cores,ram,etc) from each the server using her macAddress as ID and install coresOS for each one.

#### `GEt v1/servers`

- show information about all the registred servers 

### TO DO

- Unit tests
    - https://semaphoreci.com/community/tutorials/test-driven-development-of-go-web-applications-with-gin
    - https://medium.com/@craigchilds94/testing-gin-json-responses-1f258ce3b0b1
