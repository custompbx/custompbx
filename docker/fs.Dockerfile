# ====== BUILD STAGE (temporary container) ======
FROM debian:bookworm AS builder

# Set SIGNALWIRE_TOKEN build argument
ARG SIGNALWIRE_TOKEN

# Set XML_CURL_SERVER_HOST build argument
ARG XML_CURL_SERVER_HOST

# Set XML_CURL_SERVER_PORT build argument
ARG XML_CURL_SERVER_PORT

# Set XML_CURL_SERVER_ROUTE build argument
ARG XML_CURL_SERVER_ROUTE

# Set MEDIA_PORT_START build argument
ARG MEDIA_PORT_START

# Set MEDIA_PORT_END build argument
ARG MEDIA_PORT_END

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
RUN wget --http-user=signalwire --http-password=${SIGNALWIRE_TOKEN} -O /usr/share/keyrings/signalwire-freeswitch-repo.gpg https://freeswitch.signalwire.com/repo/deb/debian-release/signalwire-freeswitch-repo.gpg

# Configure authentication
RUN echo "machine freeswitch.signalwire.com login signalwire password ${SIGNALWIRE_TOKEN}" > /etc/apt/auth.conf
RUN chmod 600 /etc/apt/auth.conf

# Add the SignalWire FreeSWITCH repository to the sources.list
RUN echo "deb [signed-by=/usr/share/keyrings/signalwire-freeswitch-repo.gpg] https://freeswitch.signalwire.com/repo/deb/debian-release/ $(lsb_release -sc) main" > /etc/apt/sources.list.d/freeswitch.list
RUN echo "deb-src [signed-by=/usr/share/keyrings/signalwire-freeswitch-repo.gpg] https://freeswitch.signalwire.com/repo/deb/debian-release/ $(lsb_release -sc) main" >> /etc/apt/sources.list.d/freeswitch.list

# Install FreeSWITCH
RUN apt-get update && apt-get install -y freeswitch \
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
                                         freeswitch-mod-cdr-pg-csv

# Clean up sensitive files
RUN rm -f /etc/apt/auth.conf /usr/share/keyrings/signalwire-freeswitch-repo.gpg

RUN if [ -f /etc/freeswitch/autoload_configs/event_socket.conf.xml ]; then \
    sed -i 's/<param name="listen-ip" value="::"\/>/<param name="listen-ip" value="freeswitch-host"\/>/g' /etc/freeswitch/autoload_configs/event_socket.conf.xml; \
    sed -i 's|<!--<param name="apply-inbound-acl" value="loopback.auto"/>-->|<param name="apply-inbound-acl" value="localnet.auto"/>|' /etc/freeswitch/autoload_configs/event_socket.conf.xml; \
  fi

# Replace xml_curl.conf.xml with customized version
RUN echo '<configuration name="xml_curl.conf" description="cURL XML Gateway">' > /etc/freeswitch/autoload_configs/xml_curl.conf.xml && \
    echo '  <bindings>' >> /etc/freeswitch/autoload_configs/xml_curl.conf.xml && \
    echo '    <binding name="example">' >> /etc/freeswitch/autoload_configs/xml_curl.conf.xml && \
    echo '      <param name="gateway-url" value="https://'${XML_CURL_SERVER_HOST}':'${XML_CURL_SERVER_PORT}'/'${XML_CURL_SERVER_ROUTE}'" bindings="configuration|directory|dialplan"/>' >> /etc/freeswitch/autoload_configs/xml_curl.conf.xml && \
    echo '    </binding>' >> /etc/freeswitch/autoload_configs/xml_curl.conf.xml && \
    echo '  </bindings>' >> /etc/freeswitch/autoload_configs/xml_curl.conf.xml && \
    echo '</configuration>' >> /etc/freeswitch/autoload_configs/xml_curl.conf.xml

# Modify the vars.xml file
RUN sed -i 's|<X-PRE-PROCESS cmd="stun-set" data="external_rtp_ip=stun:stun.freeswitch.org"/>|<X-PRE-PROCESS cmd="set" data="external_rtp_ip=127.0.0.1"/>|' /etc/freeswitch/vars.xml && \
    sed -i 's|<X-PRE-PROCESS cmd="stun-set" data="external_sip_ip=stun:stun.freeswitch.org"/>|<X-PRE-PROCESS cmd="set" data="external_sip_ip=127.0.0.1"/>|' /etc/freeswitch/vars.xml

RUN rm /etc/freeswitch/sip_profiles/internal-ipv6.xml
RUN rm /etc/freeswitch/sip_profiles/external-ipv6.xml

# Left some mediaports for expose
RUN sed -i 's|<!-- <param name="rtp-start-port" value="16384"/> -->|<param name="rtp-start-port" value="${MEDIA_PORT_START}"/>|' /etc/freeswitch/autoload_configs/switch.conf.xml && \
    sed -i 's|<!-- <param name="rtp-end-port" value="32768"\/> -->|<param name="rtp-end-port" value="${MEDIA_PORT_END}"/>|' /etc/freeswitch/autoload_configs/switch.conf.xml

RUN sed -i 's/<param name="sip-capture" value="no"\/>/<param name="sip-capture" value="yes"\/>/g' /etc/freeswitch/sip_profiles/internal.xml
RUN sed -i 's/<param name="rtp-ip" value="\$\${local_ip_v4}"\/>/<param name="rtp-ip" value="freeswitch-host"\/>/g' /etc/freeswitch/sip_profiles/internal.xml
RUN sed -i 's/<param name="sip-ip" value="\$\${local_ip_v4}"\/>/<param name="sip-ip" value="freeswitch-host"\/>/g' /etc/freeswitch/sip_profiles/internal.xml
RUN sed -i 's/<param name="local-network-acl" value="localnet.auto"\/>/<param name="local-network-acl" value="loopback.auto"\/>/g' /etc/freeswitch/sip_profiles/internal.xml
RUN sed -i 's/<param name="apply-nat-acl" value="nat.auto"\/>/<param name="apply-nat-acl" value="rfc1918.auto"\/>/g' /etc/freeswitch/sip_profiles/internal.xml
RUN sed -i '/<param name="apply-nat-acl" value="rfc1918.auto"\/>/a\ <param name="apply-candidate-acl" value="rfc1918.auto"\/>' /etc/freeswitch/sip_profiles/internal.xml

COPY ./docker/fs_conf/sofia.conf.xml /etc/freeswitch/autoload_configs/
COPY ./docker/fs_conf/modules.conf.xml /etc/freeswitch/autoload_configs/
COPY ./docker/fs_conf/cdr_pg_csv.conf.xml /etc/freeswitch/autoload_configs/

# ====== FINAL STAGE (secure image) ======
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