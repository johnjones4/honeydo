#!/bin/bash

bin/render > rendered.html
wkhtmltoimage --format bmp --width 825 --height 1200 rendered.html rendered.bmp
convert -rotate "90" rendered.bmp rendered.bmp
IT8951/IT8951 0 0 rendered.bmp
rm rendered.*
sleep 300
