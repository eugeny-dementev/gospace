curl http://localhost:3007/book/2 `
    --include `
    --header "Content-Type: application/json" `
    --request "POST" `
    --data '{"name":"Better book","author":"Another author","publication":"1992"}'
