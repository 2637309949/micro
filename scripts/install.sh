#!/usr/bin/env bash

# GitHub Org and Repo to get archives from
GITHUB_ORG="2637309949"
GITHUB_REPO="micro"
GITHUB_SSH="git@github.com:${GITHUB_ORG}/${GITHUB_REPO}.git"
TMP_DIR=$(mktemp -d -t x-micro.XXXXXXXX)

getSystemInfo() {
    echo "Getting system information"
    case $ARCH in
        armv7*)
                ARCH="arm7";;
        aarch64)
                ARCH="arm64";;
        x86_64)
                ARCH="amd64";;
    esac

    # linux requires sudo permissions
    if [ "$OS" == "linux" ]; then
        SUDO="true"
    fi
    echo "Your machine is running ${OS} on ${ARCH} CPU architecture"
}

installFile() {
    cd ${TMP_DIR} && go install
    cd ${TMP_DIR}/cmd/protoc-gen-micro && go install
    cd ${TMP_DIR}/cmd/protoc-gen-openapi && go install
    cd ${TMP_DIR}/cmd/protoc-gen-client && go install
}

cleanup() {
    rm ${TMP_DIR} -rf
}

cloneRepo() {
    git clone ${GITHUB_SSH} --depth=1 ${TMP_DIR}
}

# execute installation
getSystemInfo
cloneRepo
installFile
cleanup
