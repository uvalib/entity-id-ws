#
# basic load test
#

if [ -z "$ENTITYID_URL" ]; then
   echo "ERROR: ENTITYID_URL is not defined"
   exit 1
fi

if [ -z "$ID_SERVICE_SHOULDER" ]; then
   echo "ERROR: ID_SERVICE_SHOULDER is not defined"
   exit 1
fi

if [ -z "$API_TOKEN" ]; then
   echo "ERROR: API_TOKEN is not defined"
   exit 1
fi

LT=../../bin/bombardier
if [ ! -f "$LT" ]; then
   echo "ERROR: Bombardier is not available"
   exit 1
fi

# generate the DOI
RES=$(curl -X POST --header "Content-Type: application/json" -d "{\"schema\":\"datacite\"}" $ENTITYID_URL/$ID_SERVICE_SHOULDER?auth=$API_TOKEN 2>/dev/null)
DOI=$(echo $RES | jq '.details.id'| tr -d \")
if [ -z "$DOI" ]; then
   echo "ERROR: unable to create a new DOI"
   exit 1
fi

# set the test parameters
endpoint=$ENTITYID_URL
concurrent=10
count=1000
url=$DOI?auth=$API_TOKEN

CMD="$LT -c $concurrent -n $count -l $endpoint/$url"
echo "Host = $endpoint, count = $count, concurrency = $concurrent"
echo $CMD
$CMD
exit $?

#
# end of file
#
