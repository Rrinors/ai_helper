import argparse
import json
import time
from common import http_request

SUBMIT_URI = "/api/v1/qwen/submit"
RESULT_URI = "/api/v1/qwen/result"

def call(args):
    print(f"start interact with {args.model}")
    while True:
        req = input("You: ")
        if req == "quit":
            break
        data_map = {
            'user_id': args.user,
            'input_model': args.model,
            'input_role': "user",
            'input_content': req,
            'history_num': args.history
        }
        print(f"req: {json.dumps(data_map)}")
        resp = http_request("POST", SUBMIT_URI, json.dumps(data_map))
        if resp.status_code != 200:
            print(f"http request failed, code={resp.status_code}")
            continue
        try:
            task = resp.json()
        except Exception:
            print(f"response format error")
            continue
        task_id = task['id']
        data_map = {
            'id': task_id
        }
        # wait finish
        success = False
        start = time.time()
        while True:
            time.sleep(0.5)
            resp = http_request("GET", RESULT_URI, json.dumps(data_map))
            if resp.status_code != 200:
                break
            try:
                resp_data = resp.json()
            except Exception:
                break
            if resp_data['status_code'] == 0:
                success = True
                break
            if resp_data['status_code'] != 202 or time.time() - start > 30:
                break
        if not success:
            print(f"get qwen response failed")
            continue
        message = resp['status_msg']
        print(f"Qwen: {message}")


if __name__ == "__main__":
    parser = argparse.ArgumentParser(description="qwen caller")

    parser.add_argument('-u', '--user', type=int, required=True, help='user id')
    parser.add_argument('-m', '--model', type=str, default="qwen-turbo", help='qwen model')
    parser.add_argument('--history', type=int, default=10, help='history num')

    args = parser.parse_args()
    call(args)
