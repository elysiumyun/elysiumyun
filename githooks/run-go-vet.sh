#!/usr/bin/env bash

set -e

for dir in $(go list ./...|sort -u); do
  go vet $dir
done