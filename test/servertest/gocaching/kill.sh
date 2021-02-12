va=`ps -aux | grep "main -host" | cut -d" " -f 3`
for k in $va
do
echo "killed" $k
kill $k
done