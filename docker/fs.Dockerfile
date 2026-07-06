# syntax=docker/dockerfile:1.7
# ====== BUILD STAGE (temporary container) ======
FROM debian:bookworm AS builder

# Empty installs the newest version exposed by the SignalWire repository.
# Set this to an exact Debian package version for reproducible/manual builds.
ARG FREESWITCH_VERSION=""

# explicitly set user/group IDs
RUN groupadd -r freeswitch --gid=999 && useradd -r -g freeswitch --uid=999 freeswitch

# Install dependencies
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
RUN --mount=type=secret,id=signalwire_token \
    SIGNALWIRE_TOKEN="$(cat /run/secrets/signalwire_token)" && \
    wget --http-user=signalwire --http-password=${SIGNALWIRE_TOKEN} -O /usr/share/keyrings/signalwire-freeswitch-repo.gpg https://freeswitch.signalwire.com/repo/deb/debian-release/signalwire-freeswitch-repo.gpg

# Configure authentication
RUN --mount=type=secret,id=signalwire_token \
    SIGNALWIRE_TOKEN="$(cat /run/secrets/signalwire_token)" && \
    echo "machine freeswitch.signalwire.com login signalwire password ${SIGNALWIRE_TOKEN}" > /etc/apt/auth.conf
RUN chmod 600 /etc/apt/auth.conf

# Add the SignalWire FreeSWITCH repository to the sources.list
RUN echo "deb [signed-by=/usr/share/keyrings/signalwire-freeswitch-repo.gpg] https://freeswitch.signalwire.com/repo/deb/debian-release/ $(lsb_release -sc) main" > /etc/apt/sources.list.d/freeswitch.list
RUN echo "deb-src [signed-by=/usr/share/keyrings/signalwire-freeswitch-repo.gpg] https://freeswitch.signalwire.com/repo/deb/debian-release/ $(lsb_release -sc) main" >> /etc/apt/sources.list.d/freeswitch.list

# Install FreeSWITCH. Pinning the main package makes apt resolve matching module
# dependencies from the same repository version.
RUN apt-get update && apt-get install -y "freeswitch${FREESWITCH_VERSION:+=${FREESWITCH_VERSION}}" \
                                         freeswitch-init \
                                         freeswitch-lang \
                                         freeswitch-timezones \
                                         freeswitch-meta-codecs \
                                         freeswitch-meta-conf \
                                         freeswitch-music \
                                         freeswitch-mod-av \
                                         freeswitch-mod-callcenter \
                                         freeswitch-mod-commands \
                                         freeswitch-mod-conference \
                                         freeswitch-mod-curl \
                                         freeswitch-mod-db \
                                         freeswitch-mod-directory \
                                         freeswitch-mod-dptools \
                                         freeswitch-mod-esl \
                                         freeswitch-mod-expr \
                                         freeswitch-mod-fsv \
                                         freeswitch-mod-hash \
                                         freeswitch-mod-httapi \
                                         freeswitch-mod-http-cache \
                                         freeswitch-mod-pgsql \
                                         freeswitch-mod-png \
                                         freeswitch-mod-shout \
                                         freeswitch-mod-spandsp \
                                         freeswitch-mod-dialplan-directory \
                                         freeswitch-mod-dialplan-xml \
                                         freeswitch-mod-rtmp \
                                         freeswitch-mod-sofia \
                                         freeswitch-mod-verto \
                                         freeswitch-mod-cdr-csv \
                                         freeswitch-mod-event-socket \
                                         freeswitch-mod-snmp \
                                         freeswitch-mod-local-stream \
                                         freeswitch-mod-native-file \
                                         freeswitch-mod-sndfile \
                                         freeswitch-mod-tone-stream \
                                         freeswitch-mod-lua \
                                         freeswitch-mod-console \
                                         freeswitch-mod-logfile \
                                         freeswitch-mod-syslog \
                                         freeswitch-mod-posix-timer \
                                         freeswitch-mod-timerfd \
                                         freeswitch-mod-xml-cdr \
                                         freeswitch-mod-xml-curl \
                                         freeswitch-mod-fifo \
                                         freeswitch-mod-voicemail \
                                         freeswitch-mod-esf \
                                         freeswitch-mod-valet-parking \
                                         freeswitch-mod-rtc \
                                         freeswitch-mod-loopback \
                                         freeswitch-mod-enum \
                                         freeswitch-mod-amqp \
                                         freeswitch-mod-say-en \
                                         freeswitch-mod-cdr-pg-csv \
                                         freeswitch-mod-nibblebill \
                                         freeswitch-mod-cdr-mongodb \
                                         freeswitch-mod-perl \
                                         freeswitch-mod-distributor \
                                         freeswitch-mod-cdr-pg-csv \
                                         freeswitch-mod-alsa \
                                         freeswitch-mod-lcr \
                                         freeswitch-mod-memcache \
                                         freeswitch-mod-redis \
                                         freeswitch-mod-oreka \
                                         freeswitch-mod-pocketsphinx \
                                         freeswitch-mod-tts-commandline \
                                         freeswitch-mod-xml-rpc

# Clean up sensitive files
RUN rm -f /etc/apt/auth.conf /usr/share/keyrings/signalwire-freeswitch-repo.gpg

RUN if [ -f /etc/freeswitch/autoload_configs/event_socket.conf.xml ]; then \
    sed -i 's/<param name="listen-ip" value="::"\/>/<param name="listen-ip" value="freeswitch-host"\/>/g' /etc/freeswitch/autoload_configs/event_socket.conf.xml; \
    sed -i 's|<!--<param name="apply-inbound-acl" value="loopback.auto"/>-->|<param name="apply-inbound-acl" value="localnet.auto"/>|' /etc/freeswitch/autoload_configs/event_socket.conf.xml; \
  fi

# Modify the vars.xml file
RUN sed -i 's|<X-PRE-PROCESS cmd="stun-set" data="external_rtp_ip=stun:stun.freeswitch.org"/>|<X-PRE-PROCESS cmd="set" data="external_rtp_ip=127.0.0.1"/>|' /etc/freeswitch/vars.xml && \
    sed -i 's|<X-PRE-PROCESS cmd="stun-set" data="external_sip_ip=stun:stun.freeswitch.org"/>|<X-PRE-PROCESS cmd="set" data="external_sip_ip=127.0.0.1"/>|' /etc/freeswitch/vars.xml

RUN rm /etc/freeswitch/sip_profiles/internal-ipv6.xml
RUN rm /etc/freeswitch/sip_profiles/external-ipv6.xml

RUN sed -i 's/<param name="sip-capture" value="no"\/>/<param name="sip-capture" value="yes"\/>/g' /etc/freeswitch/sip_profiles/internal.xml
RUN sed -i 's/<param name="rtp-ip" value="\$\${local_ip_v4}"\/>/<param name="rtp-ip" value="freeswitch-host"\/>/g' /etc/freeswitch/sip_profiles/internal.xml
RUN sed -i 's/<param name="sip-ip" value="\$\${local_ip_v4}"\/>/<param name="sip-ip" value="freeswitch-host"\/>/g' /etc/freeswitch/sip_profiles/internal.xml
RUN sed -i 's/<param name="local-network-acl" value="localnet.auto"\/>/<param name="local-network-acl" value="loopback.auto"\/>/g' /etc/freeswitch/sip_profiles/internal.xml
RUN sed -i 's/<param name="apply-nat-acl" value="nat.auto"\/>/<param name="apply-nat-acl" value="rfc1918.auto"\/>/g' /etc/freeswitch/sip_profiles/internal.xml
RUN sed -i '/<param name="apply-nat-acl" value="rfc1918.auto"\/>/a\ <param name="apply-candidate-acl" value="rfc1918.auto"\/>' /etc/freeswitch/sip_profiles/internal.xml

COPY ./docker/fs_conf/sofia.conf.xml /etc/freeswitch/autoload_configs/
COPY ./docker/fs_conf/modules.conf.xml /etc/freeswitch/autoload_configs/
COPY ./docker/fs_conf/cdr_pg_csv.conf.xml /etc/freeswitch/autoload_configs/
COPY ./docker/fs_conf/switch.conf.xml /etc/freeswitch/autoload_configs/
COPY ./docker/fs_conf/xml_curl.conf.xml /etc/freeswitch/autoload_configs/

# ====== FINAL STAGE ======
FROM debian:bookworm

# Copy FreeSWITCH binaries from the builder stage
COPY --from=builder /usr/ /usr/
COPY --from=builder /etc/freeswitch /etc/freeswitch
COPY --from=builder /var/lib/freeswitch /var/lib/freeswitch

# Set up a dedicated user
RUN groupadd -r freeswitch --gid=999 && useradd -r -g freeswitch --uid=999 freeswitch

# Volumes
VOLUME ["/var/log/freeswitch/log"]
VOLUME ["/usr/local/freeswitch/db"]
## Tmp so we can get core dumps out
#VOLUME ["/tmp"]

# Expose the necessary ports
EXPOSE 5060/tcp 5060/udp 5080/tcp 5080/udp 8021/tcp 7443/tcp
EXPOSE 16384-16399/udp

# Start FreeSWITCH
CMD ["freeswitch", "-nonat"]
