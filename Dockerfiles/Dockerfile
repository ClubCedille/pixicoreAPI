FROM golang:1.11

ENV GOPATH /app
ENV WORKDIR  $GOPATH/src/github.com/ClubCedille/pixicoreAPI
RUN mkdir -p $WORKDIR

RUN mkdir  /home/cedille/
WORKDIR /home/cedille/
RUN wget -nv https://raw.githubusercontent.com/ClubCedille/infra-ng/master/pxe-config.ign?token=AargQFYElAzZ1bBN-4UvultnJm8rtaR5ks5cO3yvwA%3D%3D -O pxe-config.ign
RUN wget -nv https://stable.release.core-os.net/amd64-usr/current/coreos_production_pxe_image.cpio.gz 
RUN wget -nv https://stable.release.core-os.net/amd64-usr/current/coreos_production_pxe.vmlinuz 

WORKDIR $WORKDIR

COPY . $WORKDIR

ADD https://github.com/golang/dep/releases/download/v0.5.0/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep
RUN dep ensure

RUN go test ./... && go build ./cmd/pixicoreAPI &&  go build ./cmd/getFacts

CMD ./pixicoreAPI
