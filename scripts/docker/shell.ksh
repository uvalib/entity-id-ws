if [ -z "$DOCKER_HOST" ]; then
   echo "ERROR: no DOCKER_HOST defined"
   exit 1
fi

# set the definitions
INSTANCE=entity-id-ws
NAMESPACE=uvadave

docker run -ti -p 8180:8080 $NAMESPACE/$INSTANCE /bin/bash
