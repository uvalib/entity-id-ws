
export GOPATH=$(pwd)

res=0
if [ $res -eq 0 ]; then
  env GOOS=darwin go build -o bin/entity-id-ws.darwin entityidws
  res=$?
fi

if [ $res -eq 0 ]; then
  env GOOS=linux go build -o bin/entity-id-ws.linux entityidws
  res=$?
fi

exit $res
