export GOPATH=$(pwd)
RUN=""
if [ $# -ge 1 ]; then
   RUN="-run $*"
fi

go test -v entityidws $RUN
