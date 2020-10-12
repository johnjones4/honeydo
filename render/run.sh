#!/bin/bash

while true
do
  bin/render > rendered.html
  wkhtmltoimage --format bmp --width 825 --height 1200 rendered.html rendered.bmp
  convert -rotate "-90" rendered.bmp rendered.bmp
  sudo IT8951/IT8951 0 0 rendered.bmp
  rm rendered.*
  sleep 300
done
