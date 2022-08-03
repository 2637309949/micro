#!/usr/bin/env bash

# GitHub Org and Repo to get archives from
GITHUB_ORG="2637309949"
GITHUB_REPO="micro"

git clone git@github.com:${GITHUB_ORG}/${GITHUB_REPO}.git --depth=1
cd micro && go install && cd ..
cd micro/cmd/protoc-gen-micro && go install && cd ../../../
cd micro/cmd/protoc-gen-openapi && go install && cd ../../../
cd micro/cmd/protoc-gen-client && go install && cd ../../../
rm micro -rf
