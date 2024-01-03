#!/bin/bash

export GOLANG="go"

$GOLANG env -w GO111MODULE=on

export GOPROXY=https://proxy.golang.org,direct

echo "GOPROXY => ${GOPROXY}"

GREEN="\e[32m"

tools=(
  'github.com/blink-io/x/kratos/v2/cmd/protoc-gen-x-go-http@latest'
  'github.com/GaijinEntertainment/go-exhaustruct/cmd/exhaustruct@latest'
  'connectrpc.com/connect/cmd/protoc-gen-connect-go@latest'
  'github.com/bufbuild/buf/cmd/buf@latest'
  'github.com/fullstorydev/grpcurl/cmd/grpcurl@latest'
  'google.golang.org/protobuf/cmd/protoc-gen-go@latest'
  'google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest'
  'github.com/go-kratos/kratos/cmd/kratos/v2@latest'
  'github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest'
  'github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2@latest'
  'github.com/google/wire/cmd/wire@latest'
  'github.com/oligot/go-mod-upgrade@latest'
  'github.com/sqlc-dev/sqlc/cmd/sqlc@latest'
  'github.com/golangci/golangci-lint/cmd/golangci-lint@latest'
  'github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest'
  'github.com/rubenv/sql-migrate/...@latest'
  'github.com/golang-migrate/migrate/v4/cmd/migrate@latest'
  'github.com/swaggo/swag/cmd/swag@latest'
  'github.com/go-swagger/go-swagger/cmd/swagger@latest'
  'github.com/alta/protopatch/cmd/protoc-gen-go-patch@latest'
  'github.com/automation-co/husky@latest'
  'mvdan.cc/garble@latest'
#  'golang.org/x/vuln/cmd/govulncheck@latest'
  # EXP
  'github.com/ServiceWeaver/weaver/cmd/weaver@latest'
  'github.com/rvflash/goup@latest'
)

echo "Begin to fetch tools..."

for t in "${tools[@]}"; do
  echo "Installing ${t} ..."
  $GOLANG install "$t" || (echo -e "\033[31m Unable to install ${t} \033[0m")
done

echo "----------------------------------"
echo "----------------------------------"

echo "------ Done ------"
#!/bin/bash

export GOLANG="go"

$GOLANG env -w GO111MODULE=on

export GOPROXY=https://proxy.golang.org,direct

echo "GOPROXY => ${GOPROXY}"

GREEN="\e[32m"

tools=(
  'google.golang.org/protobuf/cmd/protoc-gen-go@latest'
  'google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest'
  'github.com/go-kratos/kratos/cmd/kratos/v2@latest'
  'github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest'
  'github.com/go-kratos/kratos/cmd/protoc-gen-go-errors/v2@latest'
  'github.com/google/wire/cmd/wire@latest'
  'github.com/oligot/go-mod-upgrade@latest'
  'github.com/sqlc-dev/sqlc/cmd/sqlc@latest'
  'github.com/golangci/golangci-lint/cmd/golangci-lint@latest'
  'github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest'
  'github.com/rubenv/sql-migrate/...@latest'
  'github.com/golang-migrate/migrate/v4/cmd/migrate@latest'
  'github.com/swaggo/swag/cmd/swag@latest'
  'github.com/go-swagger/go-swagger/cmd/swagger@latest'
  'github.com/alta/protopatch/cmd/protoc-gen-go-patch@latest'
  'github.com/automation-co/husky@latest'
  'mvdan.cc/garble@latest'
#  'golang.org/x/vuln/cmd/govulncheck@latest'
  # EXP
  'github.com/ServiceWeaver/weaver/cmd/weaver@latest'
  'github.com/rvflash/goup@latest'
)

#failed=()

echo "Begin to fetch tools..."

for t in "${tools[@]}"; do
  echo "Installing ${t} ..."
  $GOLANG install "$t" || (echo -e "\033[31m Unable to install ${t} \033[0m")
done

echo "----------------------------------"
echo "----------------------------------"

echo "Installation is completed"
