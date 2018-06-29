mkdir abcd
mkdir test1
mkdir test2

touch abcd/test.dta
touch abcd/test.gz
touch abcd/test.txt

touch test1/a.a
touch test1/a.b
touch test1/a.c

for i in {1..100}
do
touch test2/$i.txt
done
