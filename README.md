# Bootstrap dependencies
docker-compose up -d

# Run API
cd to cmd/api
`go run ./...`

# Run processor
cd to cmd/processor
`go run ./...`

# Run report API
cd to cmd/reportapi
`export PORT=8081`
`go run ./...`

# Post message
```bash
curl -X POST http://localhost:8080/message \
   -H "Content-Type: application/json" \
   -d '{"sender": "foo", "receiver": "bar", "message": "test 82828288282"}'
```

# Fetch messages
`curl "http://localhost:8081/message/list?sender=foo&receiver=bar"`
