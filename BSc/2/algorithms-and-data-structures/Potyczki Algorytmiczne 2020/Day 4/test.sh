for t in $2/*.in; do
echo $t
o=${t%.in}.out
echo $(./$1 < $t)
cat $o
diff -bsq $o <(./$1 < $t) || break
done