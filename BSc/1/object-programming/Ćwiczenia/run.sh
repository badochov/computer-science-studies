#!/bin/bash
name=$1
javac "$name.java"
java -ea "$name"
rm ./"$name.class"
