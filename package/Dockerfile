#
# build the target application
#
FROM golang:1.14 as builder
RUN apt-get -y update

WORKDIR /build
COPY go.mod go.sum Makefile ./
COPY entityidws ./entityidws
RUN make build-linux

#
# build the target container
#
FROM alpine:3.11

# update the packages
RUN apk update && apk upgrade && apk add bash tzdata ca-certificates && rm -fr /var/cache/apk/*

# Create the run user and group
RUN addgroup --gid 18570 sse && adduser --uid 1984 docker -G sse -D

# set the timezone appropriatly
ENV TZ=UTC
RUN cp /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

# Specify home 
ENV APP_HOME /entity-id-ws
WORKDIR $APP_HOME

# Create necessary directories
RUN mkdir -p $APP_HOME/scripts $APP_HOME/bin $APP_HOME/data $APP_HOME/assets
RUN chown -R docker $APP_HOME && chgrp -R sse $APP_HOME

# port and run command
EXPOSE 8080
CMD scripts/entry.sh

# Move in necessary assets
COPY data/container_bash_profile /home/docker/.profile
COPY scripts/entry.sh $APP_HOME/scripts/entry.sh
COPY data/crossref-template.xml $APP_HOME/data/crossref-template.xml
COPY data/datacite-template.xml $APP_HOME/data/datacite-template.xml
COPY assets/* $APP_HOME/assets/
COPY --from=builder /build/bin/entity-id-ws.linux $APP_HOME/bin/entity-id-ws

# Ensure permissions are correct
RUN chown docker:sse /home/docker/.profile $APP_HOME/scripts/entry.sh $APP_HOME/bin/entity-id-ws $APP_HOME/data/*.xml && chmod 755 /home/docker/.profile $APP_HOME/scripts/entry.sh $APP_HOME/bin/entity-id-ws

# Add the build tag
ARG BUILD_TAG
RUN test -n "$BUILD_TAG" && touch $APP_HOME/buildtag.build-$BUILD_TAG || touch $APP_HOME/buildtag.build-0

# Specify the user
USER docker

#
# end of file
#
