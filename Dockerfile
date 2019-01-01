FROM golang
ENV DEBIAN_NONINTERACTIVE 1
RUN apt update && apt-get install -y build-essential wget git autoconf automake libtool musl-tools
WORKDIR /tmp
RUN wget http://ftp.gnu.org/gnu/libtasn1/libtasn1-4.13.tar.gz && tar -xzf libtasn1-4.13.tar.gz
RUN wget https://ftp.gnu.org/pub/gnu/libiconv/libiconv-1.15.tar.gz && tar -xzf libiconv-1.15.tar.gz
RUN git clone https://github.com/videolabs/libdsm.git
RUN git clone https://github.com/sahlberg/libsmb2.git
WORKDIR /tmp/libsmb2
RUN ./bootstrap
RUN CFLAGS="-O2 -fPIC -Wno-everything -DHAVE_SOCKADDR_LEN=1 -DHAVE_SOCKADDR_STORAGE=1" ./configure --enable-static=libsmb2 --without-libkrb5 && make && make install
WORKDIR /tmp/libiconv-1.15
RUN ./configure --enable-static=libiconv && make && make install
ENV CC=musl-gcc
ENV CFLAGS=-fPIC
ENV LDFLAGS=-static
WORKDIR /tmp/libtasn1-4.13
RUN ./configure --enable-static=libtasn1 && make && make install
WORKDIR /tmp/libdsm
ENV TASN1_CFLAGS="-I/usr/local/include"
ENV TASN1_LIBS="-L/usr/local/lib -ltasn1 -liconv"
RUN ./bootstrap
RUN ./configure --enable-static=libdsm && make && make install