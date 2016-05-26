# set defaults
CMD_OPTS="--debug"

# check environment for additional options
if [ -n "$EZID_USER" ]; then
   CMD_OPTS="$CMD_OPTS --eziduser $EZID_USER"
fi

if [ -n "$EZID_PASSWD" ]; then
   CMD_OPTS="$CMD_OPTS --ezidpassword $EZID_PASSWD"
fi

echo $CMD_OPTS

bin/entity-id-ws.darwin $CMD_OPTS
