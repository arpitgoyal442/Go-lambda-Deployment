# Go-lambda-Deployment
Deploying a simple Go lambda function --AWS

Step1: Create a trust policy giving the Lambda service permission to assume the execution role. Copy the following JSON and create a file named trust-policy.json in the current directory.
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Action": "sts:AssumeRole"
    }
  ]
}


Step2: aws iam create-role --role-name lambda-ex --assume-role-policy-document file://trust-policy.json
Step3 : aws iam attach-role-policy --role-name lambda-ex --policy-arn arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole
Step4: in Powershell Run the following commands one-byone
$env:GOOS = "linux"
$env:GOARCH = "amd64"
$env:CGO_ENABLED = "0"
go build -o main main.go
~\Go\Bin\build-lambda-zip.exe -o function.zip main

step5:  aws lambda create-function --function-name my-function5 --runtime go1.x --role arn:aws:iam::677339863582:role/lambda-ex --handler main --zip-file fileb://function.zip
step6:  aws lambda invoke --function-name my-function5 --cli-binary-format raw-in-base64-out  --payload file://my-input.json output.json

and Done


Reference : https://docs.aws.amazon.com/lambda/latest/dg/golang-package.html  AND
            https://github.com/aws/aws-lambda-go
