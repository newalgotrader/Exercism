#!/usr/bin/env bash

input=$1
output=""

length=${#input}

for ((i=$length-1;i>=0;i--))
do
  output=$output${input:$i:1}
done

echo "$output"

