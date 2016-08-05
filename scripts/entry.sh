# set blank options variables
EZIDURL_OPT=""
EZIDUSER_OPT=""
EZIDPASSWD_OPT=""
TOKENURL_OPT=""
TIMEOUT_OPT=""
DEBUG_OPT=""

# EZID endpoint URL
if [ -n "$EZID_URL" ]; then
   EZIDURL_OPT="--ezidurl $EZID_URL"
fi

# EZID user name
if [ -n "$EZID_USER" ]; then
   EZIDUSER_OPT="--eziduser $EZID_USER"
fi

# EZID password
if [ -n "$EZID_PASSWD" ]; then
   EZIDPASSWD_OPT="--ezidpassword $EZID_PASSWD"
fi

# token authentication service URL
if [ -n "$TOKENAUTH_URL" ]; then
   TOKENURL_OPT="--tokenauth $TOKENAUTH_URL"
fi

# EZID service timeout
if [ -n "$EZID_TIMEOUT" ]; then
   TIMEOUT_OPT="--timeout $EZID_TIMEOUT"
fi

# service debugging
if [ -n "$ENTITYID_DEBUG" ]; then
   DEBUG_OPT="--debug"
fi

bin/entity-id-ws $EZIDURL_OPT $EZIDUSER_OPT $EZIDPASSWD_OPT $TOKENURL_OPT $TIMEOUT_OPT $DEBUG_OPT

#
# end of file
#
