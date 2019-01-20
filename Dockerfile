FROM golang:1.11

ENV GOPATH /app
ENV WORKDIR  $GOPATH/src/github.com/ClubCedille/pixicoreAPI
RUN mkdir -p $WORKDIR

RUN mkdir -p /home/cedille/
RUN wget -nv https://raw.githubusercontent.com/ClubCedille/infra-ng/master/pxe-config.ign?token=AargQBtmaDQI8Gtr20UWfP6STPWMFO8Pks5cTSUjwA%3D%3D -O /home/cedille/pxe-config.ign
RUN wget -nv https://stable.release.core-os.net/amd64-usr/current/coreos_production_pxe_image.cpio.gz -P /home/cedille/
RUN wget -nv https://stable.release.core-os.net/amd64-usr/current/coreos_production_pxe.vmlinuz -P /home/cedille/
RUN ssh-keygen -f /home/cedille/id_rsa -t rsa -N ''
RUN key=$(cat /home/cedille/id_rsa.pub) && sed -i "s|__KEY-SSH__|$key|g" /home/cedille/pxe-config.ign

WORKDIR $WORKDIR

COPY .  $WORKDIR

ADD https://github.com/golang/dep/releases/download/v0.5.0/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep
RUN dep ensure

RUN go test ./... && go build ./cmd/pixicoreAPI &&  go build ./cmd/getFacts && go get go.universe.tf/netboot/cmd/pixiecore 

CMD ./pixicoreAPI & /app/bin/pixiecore api http://localhost:3000 --dhcp-no-bind 