# Description

An AI assistant based on LLM.

Currently supports dialogue service based on Qwen.

# Deployment

**Requires Docker environment.**

## Start service

```bash
docker-compose up --build
```

## Register

```bash
curl -X POST "http://127.0.0.1:8888/api/v1/user/register" \
    -H "Content-Type: application/json" \
    -d '{
    "name": "user_name",
    "qwen_api_key": "your_api_key"
}'
```

## Interact

```bash
cd python
python3 qwen_caller.py -u <user_id> [-m <model_name>] [--history <history_num>]
# Type 'quit' to exit.
```

# Devlopment

## Update IDL

```bash
make update
```

## Local test

```bash
make run
```
