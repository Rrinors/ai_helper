import requests

SERVER = "http://127.0.0.1:8888"

def http_request(method: str, uri: str, data: dict):
    headers = {
        'Content-Type': "application/json"
    }
    url = SERVER + uri
    if method == "POST":
        response = requests.post(url, json=data, headers=headers)
    elif method == "GET":
        response = requests.get(url, json=data, headers=headers)
    return response
