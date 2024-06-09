# Description

An AI assistant based on LLM.

Currently supports dialogue service based on Qwen.

# Deployment

**Requires Docker environment.**

1.Start service.

```bash
docker-compose up --build
```

2.Register user with API-KEY.

```bash
curl -X POST "http://127.0.0.1:8888/api/v1/user/register" \
    -H "Content-Type: application/json" \
    -d '{
    "name": "user_name",
    "qwen_api_key": "your_api_key"
}'
```

3.Interact.

```bash
cd python
python3 qwen_caller.py -u <user_id> [-m <model_name> --history <history_num>]

# Type quit to exit.
```

# Devlopment

After updating the IDL, execute the following commands.

```bash
make update
```

Start service locally.

```bash
make run
```
