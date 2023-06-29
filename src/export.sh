#!/bin/sh

for i in {0..14}
do
  output_file="output$i.png"
  openscad --o $output_file --export-format png --camera 100,-180,100,40,40,40 --autocenter --viewall --D i=$i case.scad
done
