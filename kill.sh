for((i=8090; i <= 8093; i++))
do
	kill -s SIGKILL "$(sudo netstat -ltp | \
		awk -v port="$i" '{
			regex = "^.+" port "$"
			where = match($4, regex)
			if (where != 0) {
				split($7, a, "/")
				print a[1]
			}
		}')"
done
echo "done"
