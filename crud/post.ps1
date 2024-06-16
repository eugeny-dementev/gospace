curl http://localhost:3005/movies `
    --include `
    --header "Content-Type: application/json" `
    --request "POST" `
    --data '{"title":"Some movie","director":{"first":"Some","last":"artist"},"year":1989,"isbn":"239847"}'
