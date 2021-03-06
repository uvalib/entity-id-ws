# set blank options variables
IDSERVICE_URL_OPT=""
IDSERVICE_USER_OPT=""
IDSERVICE_PASSWD_OPT=""
SECRET_OPT=""
TIMEOUT_OPT=""
DEBUG_OPT=""

# ID service endpoint URL
if [ -n "$ID_SERVICE_URL" ]; then
   IDSERVICE_URL_OPT="--idserviceurl $ID_SERVICE_URL"
fi

# ID service user name
if [ -n "$ID_SERVICE_USER" ]; then
   IDSERVICE_USER_OPT="--idserviceuser $ID_SERVICE_USER"
fi

# ID service password
if [ -n "$ID_SERVICE_PASSWD" ]; then
   IDSERVICE_PASSWD_OPT="--idservicepasswd $ID_SERVICE_PASSWD"
fi

# shared secret
if [ -n "$AUTH_SHARED_SECRET" ]; then
   SECRET_OPT="--secret $AUTH_SHARED_SECRET"
fi

# service timeout
if [ -n "$SERVICE_TIMEOUT" ]; then
   TIMEOUT_OPT="--timeout $SERVICE_TIMEOUT"
fi

# service debugging
if [ -n "$ENTITYID_DEBUG" ]; then
   DEBUG_OPT="--debug"
fi

bin/entity-id-ws $IDSERVICE_URL_OPT $IDSERVICE_USER_OPT $IDSERVICE_PASSWD_OPT $SECRET_OPT $TIMEOUT_OPT $DEBUG_OPT

#
# end of file
#
