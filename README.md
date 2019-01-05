# pixicoreAPI

## How to run the API

- Install [Go](https://nats.io/documentation/tutorials/go-install/).

- Git clone the repo into `$GOPATH/src/github.com/ClubCedille/`.

- Enter the `pixicoreAPI` directory.

- Install [Dep](https://golang.github.io/dep/docs/installation.html), a program that will manage Go dependencies.

- Add `dep` to your PATH like this: `export PATH=$PATH:$GOPATH/bin`.

- Now you can install all the package dependencies with `dep ensure`.

- `go test ./... && go build ./cmd/pixicoreAPI && ./pixicoreAPI` will run the tests, build the program and run it.

## Run with vagrant 

- Run `dep ensure` and `go build ./cmd/pixicoreAPI`

- Install Vagrant

- Run the master vm with `vagrant up master`

    - SSH into the master VM with `vagrant ssh master`

    - Install pixiecore with `go get -v go.universe.tf/netboot/cmd/pixiecore`

    - Run the pixiecore client with `sudo ~/go/bin/pixiecore api http://localhost:3000 --dhcp-no-bind `

- In another terminal ssh into the master VM with `vagrant ssh master`

    - Run the pixiecoreAPI with `/home/cedille/pixicoreAPI`

- Start the child VM with `vagrant up child`

    - now if you open the child vm with virtualbox you can see the boot with pxe

#### Using Docker

- `docker build -t pixicoreapi .`

- `docker run -d -p 3000:3000 pixicoreapi`

## Usage

- Change the IP address for each server in the `servers-config.yaml` file

- You can run `curl -i http://localhost:3000/v1/install/SERVER_MAC_ADDRESS`: this will collect info and install coreOS for the server

- You can run `curl -i http://localhost:3000/v1/all/`: this will collect info  and install coreOS for each server

## API Endpoints

#### `GET v1/boot/:macAddress`

- Used by pixicore to get PXE config and boot each server (each server have a IP address assigned).

#### `GET v1/install/:macAddress`

- Gets information (cores, RAM, etc) from the server using its macAddress as ID and install coresOS.

#### `GET v1/all`

- Gets information (cores, RAM, etc) from each the server using its macAddress as ID and install coresOS for each one.

#### `GET v1/servers`

- Show information about all the registered servers.

