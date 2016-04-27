FROM centos:7

# Create the run user and group
RUN groupadd -r webservice && useradd -r -g webservice webservice

# Specify home 
ENV APP_HOME /entity-id-ws
WORKDIR $APP_HOME

# Create necessary directories
RUN mkdir -p $APP_HOME/scripts $APP_HOME/bin
RUN chown -R webservice $APP_HOME && chgrp -R webservice $APP_HOME

# Specify the user
USER webservice

# port and run command
EXPOSE 8080
CMD scripts/entry.sh

# Move in necessary assets
COPY scripts/entry.sh $APP_HOME/scripts/entry.sh
COPY bin/entity-id-ws.linux $APP_HOME/bin/entity-id-ws
