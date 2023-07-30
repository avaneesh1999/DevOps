
#!/bin/bash


#################### Shell script to monitor AWS resources ##############

# EC2
# S3
# lambda

## AUTHOR :- AVANEESH SING
## Version:- V1

#command use to see executes commands in output
set -x 

echo "List s3 buckets"
aws s3 ls > output.txt


echo "List EC2 instance and copy to file mentioned "
aws ec2 describe-instances > output.txt

echo "List lambda functions and copy to file mentioned "
aws lambda list-functions > output.txt

