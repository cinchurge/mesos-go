#!/bin/bash

# Setup constants
export GOPATH=${WORKSPACE}/.jenkins-tmp
REPODIR="${GOPATH}/src/code.uber.internal/infra/mesos-go"

# Create go workspace
mkdir -p ${REPODIR} &&

# rsync everything to go workspace
rsync -avzu ${WORKSPACE}/ ${REPODIR}/ --exclude=${REPODIR} --exclude=vendor --exclude=.git &&

# cd into workspace
cd ${REPODIR} &&

# Install dependencies with glide
glide install &&

# Run test
go test
