# time=$( TIMEFORMAT="%R"; { time ls; } 2>&1 )
# AMOUNT=100000
# echo $[AMOUNT] > test.txt
# ./rand.sh $[AMOUNT] 100000 >> test.txt

gcc C/QSort.c
echo $(time(./a.out))
rm a.out

javac Java/QSort.java -d .
echo $(time(java QSort))
echo $(time(java -Xint QSort))
rm QSort.class
