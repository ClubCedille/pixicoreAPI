FROM golang:1.11

ENV GOPATH /app
ENV WORKDIR  $GOPATH/src/github.com/ClubCedille/pixicoreAPI
RUN mkdir -p $WORKDIR

RUN mkdir -p /app/coreosPxeConfig/


RUN wget -nv https://stable.release.core-os.net/amd64-usr/current/coreos_production_pxe_image.cpio.gz -P /app/coreosPxeConfig/
RUN wget -nv https://stable.release.core-os.net/amd64-usr/current/coreos_production_pxe.vmlinuz -P /app/coreosPxeConfig/
WORKDIR $WORKDIR

COPY . $WORKDIR



ADD https://github.com/golang/dep/releases/download/v0.5.0/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep
RUN dep ensure

RUN go test ./... && go build ./cmd/pixicoreAPI &&  go build ./cmd/getFacts && go get go.universe.tf/netboot/cmd/pixiecore 

CMD ./pixicoreAPI & /app/bin/pixiecore api http://localhost:3000 --dhcp-no-bind 