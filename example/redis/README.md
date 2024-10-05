- start redis using docker

```
docker run -d --name redis_pub_sub -p 6379:6379 redis
```

- run example:

```
go run main.go
```
