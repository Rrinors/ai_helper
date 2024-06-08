# Description

An AI assistant based on LLM.

Currently supports dialogue service based on Qwen.

# Deployment

**Requires Docker environment.**

## Init MySQL

1.Run MySQL in Docker.

```bash
docker run --name mysql \
    -e MYSQL_ROOT_PASSWORD=mysql \
    -v $HOME/mnt/mysql:/var/lib/mysql \
    -p 3306:3306 -d mysql
```

2.Enter MySQL terminal.

```bash
docker exec -it mysql mysql -uroot -pmysql
```

3.Create database `ai_helper`.

Since Gorm cannot automatically create databases, it is necessary to manually create them through SQL.

```sql
CREATE DATABASE ai_helper;
```

## Init MinIO

Run MinIO in Docker

```bash
docker run -p 9000:9000 -p 9001:9001 --name minio \
    -e "MINIO_ROOT_USER=admin" \
    -e "MINIO_ROOT_PASSWORD=minio-key" \
    -v $HOME/mnt/minio:/data \
    minio/minio server /data --console-address ":9001"
```

## Start service

```bash
make run
```

## Interaction

1.Register user with API-KEY.

```bash
curl -X POST "http://127.0.0.1:8888/api/v1/user/register" \
    -H "Content-Type: application/json" \
    -d '{
    "name": "user_name",
    "qwen_api_key": "your_api_key"
}'
```

2.Interact

```bash
cd python
python3 qwen_caller.py -u <user_id> [-m <model_name> --history <history_num>]
```

# Devlopment

## DB

Just modify the definition in `biz/db/db.go`, and Gorm will automatically update the database, **unless you want to delete something in the database**.

## OSS

You can modify the `ModuleConcurrencyMap` in `package/config/config.go` to configure bucket information. **Similarly, deletion operations must be performed manually**.

## IDL

After updating the IDL, execute the following commands.

```bash
make update
```