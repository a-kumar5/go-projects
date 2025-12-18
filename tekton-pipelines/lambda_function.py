import json
import requests
from typing import Dict, Any


def lambda_handler(event: Dict[str, Any], context: Any) -> Dict[str, Any]:
    """
    AWS Lambda function to call the COSEC Login API
    
    Args:
        event: Lambda event object (can contain request body, headers, etc.)
        context: Lambda context object
    
    Returns:
        Dictionary with statusCode and body
    """
    api_url = "http://10.15.4.230/COSEC/Login"
    
    try:
        # Extract request body from event if present
        request_body = event.get('body', {})
        if isinstance(request_body, str):
            request_body = json.loads(request_body)
        
        # Extract headers from event if present
        headers = event.get('headers', {})
        if not isinstance(headers, dict):
            headers = {}
        
        # Set default headers
        default_headers = {
            'Content-Type': 'application/json',
            'Accept': 'application/json'
        }
        default_headers.update(headers)
        
        # Make the API call
        response = requests.post(
            api_url,
            json=request_body,
            headers=default_headers,
            timeout=30  # 30 second timeout
        )
        
        # Return successful response
        return {
            'statusCode': response.status_code,
            'headers': {
                'Content-Type': 'application/json',
                'Access-Control-Allow-Origin': '*'  # Adjust as needed
            },
            'body': json.dumps({
                'success': response.status_code < 400,
                'status_code': response.status_code,
                'data': response.json() if response.headers.get('content-type', '').startswith('application/json') else response.text,
                'headers': dict(response.headers)
            })
        }
    
    except requests.exceptions.Timeout:
        return {
            'statusCode': 504,
            'body': json.dumps({
                'success': False,
                'error': 'Request timeout - API did not respond within 30 seconds'
            })
        }
    
    except requests.exceptions.ConnectionError:
        return {
            'statusCode': 503,
            'body': json.dumps({
                'success': False,
                'error': 'Connection error - Could not reach the API endpoint'
            })
        }
    
    except requests.exceptions.RequestException as e:
        return {
            'statusCode': 500,
            'body': json.dumps({
                'success': False,
                'error': f'Request failed: {str(e)}'
            })
        }
    
    except json.JSONDecodeError:
        return {
            'statusCode': 500,
            'body': json.dumps({
                'success': False,
                'error': 'Invalid JSON in response'
            })
        }
    
    except Exception as e:
        return {
            'statusCode': 500,
            'body': json.dumps({
                'success': False,
                'error': f'Unexpected error: {str(e)}'
            })
        }


# Example usage for local testing
if __name__ == "__main__":
    # Test event
    test_event = {
        'body': {
            'username': 'test_user',
            'password': 'test_password'
        },
        'headers': {}
    }
    
    result = lambda_handler(test_event, None)
    print(json.dumps(result, indent=2))

