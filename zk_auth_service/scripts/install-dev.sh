#! /usr/bin/env sh

if [[ "$OSTYPE" == "darwin"* ]]; then
    echo "installing dev env on MacOS"
else
    echo "command install-dev only supports MacOS"
    exit
fi   

eval xcodebuild -version 2> /dev/null 2>&1
if [ ! $? -eq 0 ] ; then
    echo "Please Install XCode first - bye"
    exit
fi


is_installed()
{
  command -v "$1" >/dev/null 2>&1
}

install() 
{
  if is_installed $1; then
    echo "$1 already installed"
  else
    echo "installing $"
    brew install $1
  fi
}

add_path()
{
    if [[ ":$PATH:" != *":$1:"* ]]; then
        echo "adding $1 to PATH"
        echo 'export PATH="$PATH:$1"' >> ~/.zshrc
    else
        echo "$1 already in PATH"
    fi
}

add_export()
{    
    echo "export $1=$2" >> ~/.zshrc
}

if is_installed go; then
    echo "go already installed"
else
    echo "installing go"
    brew install go
    add_export "GOPATH" "$(go env GOPATH)"
    add_export "GOROOT" "$(go env GOROOT)"
    add_export "GOBIN" "$(go env GOBIN)"
    add_path "$(go env GOPATH)/bin"
    add_path "$(go env GOROOT)/bin"

    LINT_VERSION="v1.56.2"
    install golangci-lint

    # reload shell
    exec $SHELL -l
fi

go mod init zk_auth_service

echo "Installing ginkgo for testing"
go install -mod=mod github.com/onsi/ginkgo/v2/ginkgo@latest

go mod tidy
go mod download
