FROM golang:1.10.3-stretch

LABEL version="0.1"
LABEL mantainer="Pietro Bongiovanni"

# Install clang-5.0 for go memory_sanitizer
RUN wget -O - https://apt.llvm.org/llvm-snapshot.gpg.key| apt-key add - && \
    echo "deb http://apt.llvm.org/stretch/ llvm-toolchain-stretch-5.0 main" > /etc/apt/sources.list.d/llvm.list && \
    echo "deb-src http://apt.llvm.org/stretch/ llvm-toolchain-stretch-5.0 main" >> /etc/apt/sources.list.d/llvm.list && \
    apt-get update && \
    apt-get install -y --no-install-recommends clang-5.0

# Clean apt
RUN apt-get clean && \
    rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

# Install golint
ENV GOPATH /go
ENV PATH ${GOPATH}/bin:$PATH
RUN go get -u github.com/golang/lint/golint

WORKDIR /go/src
