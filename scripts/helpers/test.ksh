if [ -z "$GOPATH" ]; then
   echo "ERROR: GOPATH is not defined"
   exit 1
fi

RUN=""
if [ $# -ge 1 ]; then
   RUN="-run $*"
fi

go test -v entityidws $RUN
