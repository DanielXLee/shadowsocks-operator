FROM registry.access.redhat.com/ubi7/ubi-minimal:latest

ENV OPERATOR=/usr/local/bin/shadowsocks-operator \
    USER_UID=1001 \
    USER_NAME=shadowsocks-operator

# install operator binary
COPY build/_output/bin/shadowsocks-operator ${OPERATOR}

COPY build/bin /usr/local/bin
RUN  /usr/local/bin/user_setup

ENTRYPOINT ["/usr/local/bin/entrypoint"]

USER ${USER_UID}
