FROM alpine AS builder

RUN mkdir -p /opt
RUN apk add --update git alpine-sdk ncurses-dev bash linux-headers
ADD toybox /opt/toybox
WORKDIR /opt/toybox
RUN make CFLAGS=-static

WORKDIR /opt
RUN git clone https://github.com/mkj/dropbear.git
WORKDIR /opt/dropbear
RUN ./configure --disable-zlib --disable-syslog --enable-static --disable-wtmp --disable-lastlog
ADD localoptions.h .
RUN make PROGRAMS="dropbear"

RUN mkdir -p /opt/newfs/opt/ /opt/newfs/etc/
RUN cp /opt/toybox/toybox /opt/newfs/opt/toybox
RUN cp /opt/dropbear/dropbear /opt/newfs/opt/dropbear
RUN adduser -h /home/ractf -s /bin/sh -D ractf
RUN echo "ractf:ractf" | chpasswd
RUN echo "root:x:0:0:root:/root:" > /opt/newfs/etc/passwd
RUN echo "root:!::0:::::" > /opt/newfs/etc/shadow
RUN tail -n 1 /etc/passwd >> /opt/newfs/etc/passwd
RUN tail -n 1 /etc/shadow >> /opt/newfs/etc/shadow
ADD motd /opt/newfs/etc/motd

WORKDIR /opt/newfs
RUN mkdir -p usr/bin bin sbin lib tmp run var srv media home/ractf root etc/dropbear
RUN for cmd in $(opt/toybox --long); do \
        ln -s /opt/toybox $cmd; \
    done

ADD reee.sh /opt/newfs/opt/reee.sh
RUN chmod +x /opt/newfs/opt/reee.sh
WORKDIR /opt/newfs/bin
RUN for cmd in 'ls' 'cat' 'tail' 'head' 'dir' 'less' 'more' 'vi' 'vim' 'grep'; do \
        ln -s /opt/reee.sh $cmd; \
    done
ADD flag.txt /opt/newfs/home/ractf/flag.txt
RUN chown -R ractf:ractf /opt/newfs/home/ractf/
RUN chmod -R +x /opt/newfs/home/ractf/
# Throw away this builder container now

FROM scratch
COPY --from=0 /opt/newfs/ /
EXPOSE 22
CMD ["/opt/dropbear", "-R", "-F"]