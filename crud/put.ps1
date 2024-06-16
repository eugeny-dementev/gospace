curl http://localhost:3005/movies/someid `
    --include `
    --header "Content-Type: application/json" `
    --request "PUT" `
    --data '{"title":"Some movie","director":{"first":"Another"},"year":1999}'
