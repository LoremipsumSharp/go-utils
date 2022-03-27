#!/bin/bash

if [ -n "$(git status --short)" ]; then
    echo "Error: There are untracked/modified changes, commit or discard them before the release."
    exit 1
fi


RELEASE_VERSION=$1
PUSH_CHANGES=$2
CURRENT_VERSION=$3


if [ -z "${RELEASE_VERSION}" ]; then
    if [ -z "${FROM_MAKEFILE}" ]; then
        echo "Error: VERSION is missing. e.g. ./release.sh <version>"
    else
        echo "Error: missing value for 'version'. e.g. 'make release VERSION=x.y.z'"
    fi
    exit 1
fi

if [ -z "${CURRENT_VERSION}" ]; then
    CURRENT_VERSION=$(git describe --tags --exact-match 2>/dev/null || git describe --tags 2>/dev/null || echo "v0.0.1-$(COMMIT_HASH)")
fi

if [ "v${RELEASE_VERSION}" == "${CURRENT_VERSION}" ]; then
    echo "Error: provided version (v${RELEASE_VERSION}) already exists."
    exit 1
fi

if [ "$(git describe --tags "v${RELEASE_VERSION}" 2>/dev/null)" ]; then
    echo "Error: provided version (v${RELEASE_VERSION}) already exists."
    exit 1
fi

git tag --annotate --message "v${RELEASE_VERSION} Release" "v${RELEASE_VERSION}"
if [ "${PUSH_CHANGES}" == "true" ]; then
    git push origin "v${RELEASE_VERSION}"
fi