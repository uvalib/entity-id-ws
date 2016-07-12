#
#
#

# ensure we have an endpoint
if [ -z "$ENTITYID_URL" ]; then
   echo "ERROR: ENTITYID_URL is not defined"
   exit 1
fi

# issue the command
echo "$ENTITYID_URL"
curl $ENTITYID_URL/version

exit 0

#
# end of file
#
