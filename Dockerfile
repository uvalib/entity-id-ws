FROM alpine:3.4

# update the packages
RUN apk update && apk upgrade && apk add bash tzdata ca-certificates

# Create the run user and group
RUN addgroup webservice && adduser webservice -G webservice -D

# set the timezone appropriatly
ENV TZ=EST5EDT
RUN cp /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

# Specify home 
ENV APP_HOME /entity-id-ws
WORKDIR $APP_HOME

# Create necessary directories
RUN mkdir -p $APP_HOME/scripts $APP_HOME/bin $APP_HOME/data
RUN chown -R webservice $APP_HOME && chgrp -R webservice $APP_HOME

# Specify the user
USER webservice

# port and run command
EXPOSE 8080
CMD scripts/entry.sh

# Move in necessary assets
COPY scripts/entry.sh $APP_HOME/scripts/entry.sh
COPY data/container_bash_profile /home/webservice/.profile
COPY bin/entity-id-ws.linux $APP_HOME/bin/entity-id-ws
COPY data/crossref-template.xml $APP_HOME/data/crossref-template.xml

# Add the build tag
COPY buildtag.* $APP_HOME/
