# AWS Lambda Function - COSEC Login API

This Lambda function calls the COSEC Login API endpoint.

## Files

- `lambda_function.py` - Main Lambda function code
- `requirements.txt` - Python dependencies

## Deployment

### Option 1: Using AWS CLI

1. Install dependencies:
```bash
pip install -r requirements.txt -t .
```

2. Create deployment package:
```bash
zip -r lambda-deployment.zip lambda_function.py requests/ urllib3/ certifi/ charset_normalizer/ idna/
```

3. Create/Update Lambda function:
```bash
aws lambda create-function \
  --function-name cosec-login-api \
  --runtime python3.11 \
  --role arn:aws:iam::YOUR_ACCOUNT:role/lambda-execution-role \
  --handler lambda_function.lambda_handler \
  --zip-file fileb://lambda-deployment.zip \
  --timeout 30 \
  --memory-size 256
```

### Option 2: Using AWS SAM

Create `template.yaml`:
```yaml
AWSTemplateFormatVersion: '2010-09-09'
Transform: AWS::Serverless-2016-10-31

Resources:
  CosecLoginFunction:
    Type: AWS::Serverless::Function
    Properties:
      Handler: lambda_function.lambda_handler
      Runtime: python3.11
      CodeUri: ./
      Timeout: 30
      MemorySize: 256
      Events:
        ApiEvent:
          Type: Api
          Properties:
            Path: /login
            Method: post
```

Then deploy:
```bash
sam build
sam deploy --guided
```

## Event Format

The Lambda expects an event with optional `body` and `headers`:

```json
{
  "body": {
    "username": "your_username",
    "password": "your_password"
  },
  "headers": {
    "Content-Type": "application/json"
  }
}
```

## Response Format

Success response:
```json
{
  "statusCode": 200,
  "body": {
    "success": true,
    "status_code": 200,
    "data": {...},
    "headers": {...}
  }
}
```

Error response:
```json
{
  "statusCode": 500,
  "body": {
    "success": false,
    "error": "Error message"
  }
}
```

## Local Testing

```bash
python lambda_function.py
```

## Notes

- The function uses a 30-second timeout for API calls
- Adjust CORS headers in the response as needed
- The API endpoint is hardcoded but can be moved to environment variables
- Ensure the Lambda has VPC access if the API is in a private network

