#!/bin/bash
while read package; do
  echo "Installing $package..."
  go get $package
done < deps
