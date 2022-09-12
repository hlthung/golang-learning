#!/bin/sh

BIN_DIR=$1
VERSION=1.49.0

if [ -z "$BIN_DIR" ]; then
  echo "1 arg is required: destination directory"
  exit 1
fi

if [ -x "$BIN_DIR"/golangci-lint ]; then
  CURRENT_VERSION=$($BIN_DIR/golangci-lint --version | grep -Eo '\d+\.\d+\.\d+')
  if [[ $CURRENT_VERSION != $VERSION ]]; then
    echo "Version ${CURRENT_VERSION} does not match ${VERSION}.. Updating..."
  else
    echo "golangci-lint is already installed with version $CURRENT_VERSION"
    exit 0
  fi
fi

# https://golangci-lint.run/usage/install/
wget -O- -nv https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | BINDIR="$BIN_DIR" sh -s v$VERSION
"$BIN_DIR"/golangci-lint --version
