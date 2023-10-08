#!/bin/zsh

# By Wojciech Przytula

idx=hb417666

if (($# == 0)); then
    echo "Usage: provide solution.rs code; as a result you get zip of proper format"
elif (($# != 1)); then
    echo "Wrong number of args."
else
   mkdir $idx && cp $1 $idx && zip $idx.zip -r $idx && rm -r $idx
fi
