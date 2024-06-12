curl http://localhost:8888/albums `
    --include `
    --header "Content-Type: application/json" `
    --request "POST" `
    --data '{"title": "Some albume","artist": "Some artist","price": 49.99}'
