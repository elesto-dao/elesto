ARG ELESTOD_IMAGE_TAG
FROM $ELESTOD_IMAGE_TAG AS elestod
FROM rancher/kubectl:v1.23.7 AS kubectl
FROM bash:latest
COPY --from=kubectl /bin/kubectl /usr/local/bin/kubectl
COPY --from=elestod /elestod /usr/local/bin/elestod
RUN wget -O /usr/local/bin/jq https://github.com/stedolan/jq/releases/download/jq-1.6/jq-linux64 && chmod +x /usr/local/bin/jq
COPY *.sh /usr/local/bin/
