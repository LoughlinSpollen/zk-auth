FROM python:3.7.3-slim

RUN echo "deb http://archive.debian.org/debian stretch main" > /etc/apt/sources.list
RUN apt-get update && apt-get install -y bash

WORKDIR /usr/src/app

# version 24.0 required to install grpcio
RUN pip install --upgrade pip
# protobuf is from the grpc directory, not part of the requirements.txt
RUN pip --no-cache-dir install protobuf

COPY ./zk_client/requirements.txt ./
RUN pip --no-cache-dir install -r requirements.txt

COPY ./zk_client .

RUN groupadd -r zk_client_user && \
    useradd -r -g zk_client_user zk_client_user 

COPY --chown=zk_client_user . .

WORKDIR /usr/local/src
COPY ./scripts/zk-client-entrypoint.sh .
COPY --chown=zk_client_user . .

WORKDIR /usr/src/app

ENTRYPOINT ["sh", "/usr/local/src/zk-client-entrypoint.sh"]
