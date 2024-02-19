#!/usr/bin/env bash

export GONOSUMDB=github.com/Azure/azure-sdk-for-go,github.com/go-kratos/kratos/v2

_flags=" --verbose --pagesize 1000"

if [[ "$USE_PXY" -eq 1 ]]; then
  export http_proxy=http://127.0.0.1:7890
  export https_proxy=http://127.0.0.1:7890
  export all_proxy=socks5://127.0.0.1:7890
  echo Using proxy...
fi

if [[ "$USE_MRR" -eq 1 ]]; then
  export GOPROXY=https://goproxy.io,https://goproxy.cn,direct
  echo Using mirrors...
else
  export GOPROXY=https://proxy.golang.org,direct
fi

echo "Using GOPROXY:     ${GOPROXY:-'Unset'}"
echo "Using GOSUMDB:     ${GOSUMDB:-'Unset'}"
echo "Using GONOSUMDB:   ${GONOSUMDB:-'Unset'}"

echo "Using http_proxy:  ${http_proxy:-'Unset'}"
echo "Using https_proxy: ${https_proxy:-'Unset'}"


if [[ "$USE_FF" -eq 1 ]]; then
 _flags=" --force $_flags "
fi

# go list -u -f '{{if (and (not (or .Main .In direct)) .Update)}}{{.Path}}: {{.Version}} -> {{.Update.Version}}{{end}}' -m all 2> /dev/null
go-mod-upgrade "$_flags"
