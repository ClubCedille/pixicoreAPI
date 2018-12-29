FROM golang:1.11

env GOPATH /app
env WORKDIR  $GOPATH/src/github.com/ClubCedille/pixicoreAPI
RUN mkdir -p $WORKDIR
WORKDIR $WORKDIR

#RUN git clone --depth 1 -b master https://github.com/ClubCedille/pixicoreAPI .
COPY . $WORKDIR

ADD https://github.com/golang/dep/releases/download/v0.5.0/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep
RUN dep ensure

RUN go test ./... && go build ./cmd/pixicoreAPI &&  go build ./cmd/getFacts

CMD ./pixicoreAPI
