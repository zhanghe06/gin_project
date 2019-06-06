# 事务测试

iTerm 终端开启Broadcast Input
```
Shell >> Broadcast Input >> Broadcast Input to All Panes in Current Tab
```

并发测试

窗口1
```
➜  ~ curl -X POST \
  http://0.0.0.0:8080/v1/daily_sentence/transaction \
  -H 'Content-Type: application/json' \
  -H 'Postman-Token: c7c1497a-0d3b-45c8-b41e-8a17ef95520a' \
  -H 'cache-control: no-cache' \
  -d '{
    "id": "1",
    "Title": "this is a test",
    "Classification": "news"
}'
{"rows":1}
```

窗口2
```
➜  ~ curl -X POST \
  http://0.0.0.0:8080/v1/daily_sentence/transaction \
  -H 'Content-Type: application/json' \
  -H 'Postman-Token: c7c1497a-0d3b-45c8-b41e-8a17ef95520a' \
  -H 'cache-control: no-cache' \
  -d '{
    "id": "1",
    "Title": "this is a test",
    "Classification": "news"
}'
{"error":"Error 1213: Deadlock found when trying to get lock; try restarting transaction"}
```
