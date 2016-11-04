#!/bin/bash

set -vex

# Setup GOPATH and repodir
export GOPATH=${WORKSPACE}/.jenkins-tmp
export PATH=${GOPATH}/bin:${PATH}
REPODIR="${GOPATH}/src/code.uber.internal/infra/mesos-go"

# Create go workspace
mkdir -p ${REPODIR}

# rsync everything to go workspace
rsync -avzu ${WORKSPACE}/ ${REPODIR}/ --exclude=${REPODIR} --exclude=vendor --exclude=.git

# cd into workspace
cd ${REPODIR}

# Install test dependencies
go get -u github.com/jstemmer/go-junit-report
go get -u github.com/axw/gocov/
go get -u github.com/AlekSi/gocov-xml

# Install dependencies with glide
glide install

# Run test
go test -v -coverprofile=coverage.out | tee test.log

# Generate test report
go-junit-report < test.log > ${WORKSPACE}/junit.xml

# Generate coverage report
go tool cover -html=coverage.out
gocov convert coverage.out | gocov-xml | sed 's|filename=".*code.uber.internal/infra/mesos-go/|filename="|' > ${WORKSPACE}/coverage.xml
