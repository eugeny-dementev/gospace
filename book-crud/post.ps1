curl http://localhost:3007/book `
    --include `
    --header "Content-Type: application/json" `
    --request "POST" `
    --data '{"name":"Some book","author":"Some author","publication":"1989"}'
