#!/bin/bash

echo '#########################'
echo '#    Surt Localstack    #'
echo '#########################'

echo 'Configuring S3...'

awslocal s3 mb s3://surt-bucket
awslocal s3 cp /testdata/test.txt s3://surt-bucket
awslocal s3api put-object-tagging \
  --bucket surt-bucket \
  --key test.txt \
  --tagging '{"TagSet": [{"Key": "SURT_LAST_SCAN", "Value": "2021-09-17"},{"Key": "SURT_SCAN_STATUS", "Value": "CLEAN"}]}'

