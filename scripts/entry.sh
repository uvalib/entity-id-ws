# set blank options variables
EZIDURL_OPT=""
EZIDUSER_OPT=""
EZIDPASSWD_OPT=""
TOKENURL_OPT=""

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

bin/entity-id-ws $EZIDURL_OPT $EZIDUSER_OPT $EZIDPASSWD_OPT $TOKENURL_OPT

#
# end of file
#
