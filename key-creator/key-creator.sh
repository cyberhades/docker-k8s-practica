#!/bin/sh

if [[ -z "${KEY}" ]]; then
  exit 1
else
  echo ${KEY} | sha512sum | head -c 128 > /usr/share/key/key.txt
fi
