#!/bin/bash
rm -r views
mkdir views

for f in $(find templates -name '*.tmpl'); do
  cat $f > views/$(echo ${f#*/} | tr '/' '.')
done
