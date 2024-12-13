#u/bin/bash

#Simple test case

cat <<EOF > ./TestFile
Hello World!
This is a test!
We should insert a line at the end of this file!


test
EOF

time go run lineinfile.go ./TestFile 'I added a line'

grep -q 'added' ./TestFile
if [ $? -ne 0 ]
then
	echo "Test failed! Line is not present"
	exit 1;
else
	echo "Test passed. Line is present"
	exit 0;
fi
