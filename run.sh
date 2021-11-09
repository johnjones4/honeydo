#!/bin/bash

while true
do
  /usr/local/bin/go/bin/go run .
  # convert -rotate "-90" rendered.bmp rendered.bmp
  timeout 60 sudo IT8951/IT8951 0 0 rendered.bmp
  rm rendered.bmp
  sleep 300
done
