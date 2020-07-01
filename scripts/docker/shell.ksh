#if [ -z "$DOCKER_HOST" ]; then
#   echo "ERROR: no DOCKER_HOST defined"
#   exit 1
#fi

if [ -z "$DOCKER_HOST" ]; then
   DOCKER_TOOL=docker
else
   DOCKER_TOOL=docker-17.04.0
fi

# set the definitions
INSTANCE=entity-id-ws
NAMESPACE=uvadave

$DOCKER_TOOL run -ti -p 8180:8080 $NAMESPACE/$INSTANCE /bin/bash -l
