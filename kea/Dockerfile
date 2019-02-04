FROM debian:buster-slim

ENV KEA_VERSION 1.5.0
ENV TINI_VERSION v0.18.0

EXPOSE 67/udp
EXPOSE 8080

RUN apt update && apt install -y git

RUN  cd /tmp \
	&& git clone -b Kea-${KEA_VERSION} --depth=1 https://github.com/isc-projects/kea.git \
	&& apt -y install automake libtool pkg-config build-essential \
	&& apt -y install libboost-dev libboost-system-dev liblog4cplus-dev libssl-dev  python3-minimal inotify-tools\
	&& cd /tmp/kea/ && ls  \
	&& autoreconf -iv \
	&& ./configure --enable-shell \
	&& make -j $(cat /proc/cpuinfo  | grep processor | wc -l) \
	&& make install \
	&& apt remove  automake libtool pkg-config build-essential -y && apt autoremove -y \
	&& apt-get clean && apt-get autoclean \
	&& rm -rf /tmp/* /var/tmp/* \
	&& rm -rf /var/lib/apt/lists/* 	&& rm -rf /var/lib/apt/lists/*

RUN echo "/usr/local/lib/hooks" > /etc/ld.so.conf.d/kea.conf && ldconfig -v | grep kea
RUN sed -i "s/dhcp6=yes/dhcp6=no/g" /usr/local/etc/kea/keactrl.conf

ADD https://github.com/krallin/tini/releases/download/${TINI_VERSION}/tini /usr/sbin/tini
RUN chmod +x /usr/sbin/tini
ENTRYPOINT ["/sbin/tini", "--"]

# start kea detached; sleep 10 seconds then wait until one of pid files are deleted
CMD  keactrl start && sleep 10 \
	&& inotifywait -e delete /usr/local/var/kea/kea-ctrl-agent.kea-ctrl-agent.pid  /usr/local/var/kea/kea-dhcp4.kea-dhcp4.pid
