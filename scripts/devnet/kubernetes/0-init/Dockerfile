FROM rancher/kubectl:v1.23.7 AS kubectl
FROM us-central1-docker.pkg.dev/elesto/elesto/elestod:2.0.0-rc2-readiness AS elestod
FROM bash:latest
COPY --from=kubectl /bin/kubectl /usr/local/bin/kubectl
COPY --from=elestod /elestod /usr/local/bin/elestod
RUN wget -O /usr/local/bin/jq https://github.com/stedolan/jq/releases/download/jq-1.6/jq-linux64 && chmod +x /usr/local/bin/jq
COPY *.sh /usr/local/bin/
