# Base image
# docker build --build-arg SIGNALWIRE_TOKEN=pat_S4dEV6fPvGiUwgVxe15zbeGc --build-arg XML_CURL_SERVER_HOST=192.168.0.29 --build-arg XML_CURL_SERVER_PORT=8081 --build-arg XML_CURL_SERVER_ROUTE=/conf/config -t freeswitch-image -f Dockerfile-fs .
# docker run -d --name freeswitch-container -p 5060:5060/tcp -p 5060:5060/udp -p 5080:5080/tcp -p 5080:5080/udp -p 8021:8021/tcp -p 7443:7443/tcp freeswitch-image
FROM debian:buster

# Set SIGNALWIRE_TOKEN build argument
ARG SIGNALWIRE_TOKEN

# Set XML_CURL_SERVER_HOST build argument
ARG XML_CURL_SERVER_HOST

# Set XML_CURL_SERVER_PORT build argument
ARG XML_CURL_SERVER_PORT

# Set XML_CURL_SERVER_ROUTE build argument
ARG XML_CURL_SERVER_ROUTE

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
RUN wget --http-user=signalwire --http-password=$SIGNALWIRE_TOKEN -O /usr/share/keyrings/signalwire-freeswitch-repo.gpg https://freeswitch.signalwire.com/repo/deb/debian-release/signalwire-freeswitch-repo.gpg

# Configure authentication
RUN echo "machine freeswitch.signalwire.com login signalwire password $SIGNALWIRE_TOKEN" > /etc/apt/auth.conf
RUN chmod 600 /etc/apt/auth.conf

# Add the SignalWire FreeSWITCH repository to the sources.list
RUN echo "deb [signed-by=/usr/share/keyrings/signalwire-freeswitch-repo.gpg] https://freeswitch.signalwire.com/repo/deb/debian-release/ $(lsb_release -sc) main" > /etc/apt/sources.list.d/freeswitch.list
RUN echo "deb-src [signed-by=/usr/share/keyrings/signalwire-freeswitch-repo.gpg] https://freeswitch.signalwire.com/repo/deb/debian-release/ $(lsb_release -sc) main" >> /etc/apt/sources.list.d/freeswitch.list

# Install FreeSWITCH
RUN apt-get update && apt-get install -y --no-install-recommends freeswitch-meta-all

# Create default modules.conf.xml file if it doesn't exist
RUN if [ ! -f /etc/freeswitch/autoload_configs/modules.conf.xml ]; then \
    echo '<include>' > /etc/freeswitch/autoload_configs/modules.conf.xml && \
    echo '  <extension name="modules.conf" dialplan="XML" include-subdirs="true" reload="true"/>' >> /etc/freeswitch/autoload_configs/modules.conf.xml && \
    echo '</include>' >> /etc/freeswitch/autoload_configs/modules.conf.xml; \
  fi

# Modify the modules.conf.xml file to load mod_xml_curl
RUN if [ -f /etc/freeswitch/autoload_configs/modules.conf.xml ]; then \
    sed -i '/<!-- <load module="mod_xml_curl"\/> -->/ s/<!-- \(<load module="mod_xml_curl"\/>\) -->/\1/' /etc/freeswitch/autoload_configs/modules.conf.xml; \
  fi

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
RUN sed -i '/<!-- <param name="rtp-start-port" value="16384"\/> -->/{n;s/<!-- <param name="rtp-start-port" value="16384"\/> -->/<param name="rtp-start-port" value="16384"\/>/}; /<!-- <param name="rtp-end-port" value="32768"\/> -->/{n;s/<!-- <param name="rtp-end-port" value="32768"\/> -->/<param name="rtp-end-port" value="16399"\/>/}' /etc/freeswitch/autoload_configs/switch.conf.xml

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