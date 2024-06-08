import argparse
import json
import time
from common import http_request

SUBMIT_URI = "/api/v1/qwen/submit"
RESULT_URI = "/api/v1/qwen/result"

def call(args):
    print(f"start interact with {args.model}")
    while True:
        req = input("you: ")
        if req == "quit":
            break
        # submit task
        data_map = {
            'user_id': args.user,
            'input_model': args.model,
            'input_role': "user",
            'input_content': req,
            'history_num': args.history
        }
        resp = http_request("POST", SUBMIT_URI, data_map)
        if resp.status_code != 200:
            print(f"submit http_request failed, code={resp.status_code}")
            continue
        resp_map = resp.json()
        task = json.loads(resp_map['status_msg'])
        id = task['id']
        # wait finish
        data_map = {
            'id': id
        }
        success = False
        start = time.time()
        while True:
            time.sleep(1)
            resp = http_request("GET", RESULT_URI, data_map)
            if resp.status_code != 200:
                print(f"result http_request failed, code={resp.status_code}")
                break
            resp_map = resp.json()
            if resp_map['status_code'] == 0:
                success = True
                break
            if resp_map['status_code'] != 202:
                print(f"result failed, err_msg={resp_map['status_msg']}")
                break
            if time.time() - start > 30:
                print(f"result timeout")
                break
        if not success:
            continue
        message = resp_map['status_msg']
        print(f"{args.model}: {message}")


if __name__ == "__main__":
    parser = argparse.ArgumentParser(description="qwen caller")

    parser.add_argument('-u', '--user', type=int, required=True, help='user id')
    parser.add_argument('-m', '--model', type=str, default="qwen-turbo", help='qwen model')
    parser.add_argument('--history', type=int, default=10, help='history num')

    args = parser.parse_args()
    call(args)
