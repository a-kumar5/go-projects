import json
import requests
from typing import Dict, Any

def makeApiCall(method, url, payload=None):
        headers = {
            "Accept": "application/json",
            "Content-Type": "application/json"
        }
        ###This function makes an API call to BitBucket using the specified HTTP method (GET, POST, etc.) 
        ###URL, and optional payload. It handles both GET requests (when no payload is provided) and POST/PUT requests (when a payload is included). 
        ###The function then returns the API response (in Json).
        # Handle GET and POST requests
        try:
            response = requests.request(
                method,
                url,
                headers=headers,
                data=payload
            )
            response.raise_for_status()
            return response
        except requests.exceptions.HTTPError as http_err:
            print(f"HTTP error occurred: {response.status_code} - {http_err}")
        except Exception as err:
            print(f"An error occurred: {err}")
        return None

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
    res= makeApiCall('GET', api_url)
    print(res)
    print(res.status_code)


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

