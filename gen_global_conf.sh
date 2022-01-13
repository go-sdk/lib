#!/usr/bin/env bash

lines=$(find . -type f -name '*.go' -print0 | xargs -0 grep 'conf.Get')

echo "$lines" | while IFS="" read -r p; do
  echo "${p}" | awk -v FS='(conf\\.Get\\(\\(|\\")' '{print $2}'
done
