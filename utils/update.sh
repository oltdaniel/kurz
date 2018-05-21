#!/bin/bash
while read package; do
  echo "Updating $package..."
  go get -u $package
done < deps
