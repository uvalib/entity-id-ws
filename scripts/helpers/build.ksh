export GOPATH=$(pwd)

res=0
if [ $res -eq 0 ]; then
  GOOS=darwin go build -o bin/entity-id-ws.darwin entityidws
  res=$?
fi

if [ $res -eq 0 ]; then
  CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/entity-id-ws.linux entityidws
  res=$?
fi

exit $res
