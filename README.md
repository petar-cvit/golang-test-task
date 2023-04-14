docker-compose up -d

cd to cmd/api
`go run ./...`

```bash
curl -X POST http://localhost:8080/message \
   -H "Content-Type: application/json" \
   -d '{"sender": "foo", "receiver": "bar", "message": "test 82828288282"}'
```

cd to cmd/reportapi
`export PORT=8081`
`go run ./...`

`curl "http://localhost:8081/message/list?sender=foo&receiver=bar"`
