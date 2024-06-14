FROM debian:buster

# Set environment variables
#ENV FREESWITCH_PATH=/usr/local/freeswitch
ENV CMAKE_VERSION=3.18.6
ENV CMAKE_DIR=/opt/cmake
ENV SW_TOKEN=<TOKEN>

RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates \
    gnupg2 \
    wget \
    lsb-release

# Download the certificate bundle from Mozilla
RUN wget --no-check-certificate -O /usr/local/share/ca-certificates/mozilla.crt https://curl.se/ca/cacert.pem

# Update the CA certificates
RUN update-ca-certificates

# Download and import the SignalWire FreeSWITCH repository key
RUN wget --http-user=signalwire --http-password=$SW_TOKEN -O /usr/share/keyrings/signalwire-freeswitch-repo.gpg https://freeswitch.signalwire.com/repo/deb/debian-release/signalwire-freeswitch-repo.gpg

# Configure authentication
RUN echo "machine freeswitch.signalwire.com login signalwire password $SW_TOKEN" > /etc/apt/auth.conf
RUN chmod 600 /etc/apt/auth.conf

# Add the SignalWire FreeSWITCH repository to the sources.list
RUN echo "deb [signed-by=/usr/share/keyrings/signalwire-freeswitch-repo.gpg] https://freeswitch.signalwire.com/repo/deb/debian-release/ $(lsb_release -sc) main" > /etc/apt/sources.list.d/freeswitch.list
RUN echo "deb-src [signed-by=/usr/share/keyrings/signalwire-freeswitch-repo.gpg] https://freeswitch.signalwire.com/repo/deb/debian-release/ $(lsb_release -sc) main" >> /etc/apt/sources.list.d/freeswitch.list

# Install dependencies
RUN apt-get update
RUN apt-get install -y apt-transport-https
RUN apt-get install -y git
RUN apt-get install -y build-essential
RUN apt-get install -y libfreeswitch-dev
RUN apt-get install -y libssl-dev
RUN apt-get install -y zlib1g-dev
RUN apt-get install -y libspeexdsp-dev
RUN apt-get install -y pkg-config

# Install CMake 3.18
RUN wget https://github.com/Kitware/CMake/releases/download/v${CMAKE_VERSION}/cmake-${CMAKE_VERSION}-Linux-x86_64.sh -O /tmp/cmake.sh && \
    mkdir -p ${CMAKE_DIR} && \
    sh /tmp/cmake.sh --skip-license --prefix=${CMAKE_DIR} && \
    ln -s ${CMAKE_DIR}/bin/cmake /usr/local/bin/cmake

# Clone the repository and initialize submodules
RUN git clone https://github.com/amigniter/mod_audio_stream.git /mod_audio_stream && \
    cd /mod_audio_stream && \
    git submodule init && \
    git submodule update

# Set the working directory
WORKDIR /mod_audio_stream

# Build the project
RUN mkdir build
RUN cd build && cmake -DCMAKE_BUILD_TYPE=Release ..
RUN cd build && make

# Optional: if you want to install the module directly
RUN cd build && make install

# Clean up APT when done.
RUN apt-get clean && rm -rf /var/lib/apt/lists/*
