import requests

SERVER = "http://127.0.0.1:8888"

def http_request(method: str, uri: str, data: str):
    headers = {
        'Content-Type': "application/json"
    }
    print(f"request: {SERVER+uri}")
    if method == "POST":
        return requests.post(SERVER+uri, json=data, headers=headers)
    elif method == "GET":
        return requests.get(SERVER+uri, json=data, headers=headers)
