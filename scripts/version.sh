#!/bin/bash

# this script updates version information

# version details
export BUILD_DATE=`date -u +'%Y-%m-%dT%H:%M:%S%:z'`
export GIT_BRANCH=`git rev-parse --abbrev-ref HEAD`
export GIT_SHA=`git rev-parse HEAD`
export GIT_SHA_SHORT=`git rev-parse --short HEAD`
export VERSION_PKG="github.com/mycontroller-org/esphome_api/cli/version"

# update tag, if available (detached HEAD on a release tag, e.g. v1.4.0)
if [ "${GIT_BRANCH}" = "HEAD" ]; then
  export GIT_BRANCH=`git describe --abbrev=0 --tags`
fi

# Prefer full semver X.Y.Z (e.g. v1.4.0 -> 1.4.0).
# Do not use X.Y at end of string: that turns v1.4.0 into 4.0.
export VERSION=`echo "${GIT_BRANCH}" | awk 'match($0, /[0-9]+\.[0-9]+\.[0-9]+/) { print substr($0, RSTART, RLENGTH); exit }'`
if [ -z "$VERSION" ]; then
  # fallback: major.minor only (e.g. branch release/1.4)
  export VERSION=`echo "${GIT_BRANCH}" | awk 'match($0, /[0-9]+\.[0-9]+/) { print substr($0, RSTART, RLENGTH); exit }'`
fi
if [ -z "$VERSION" ]; then
  # takes version from versions file and adds devel suffix with that
  STATIC_VERSION=`grep esphomectl= versions.txt | awk -F= '{print $2}'`
  export VERSION="${STATIC_VERSION}-devel"
fi

export LD_FLAGS="-X $VERSION_PKG.version=$VERSION -X $VERSION_PKG.buildDate=$BUILD_DATE -X $VERSION_PKG.gitCommit=$GIT_SHA"
