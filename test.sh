#!/bin/bash

# http://localhost:4869/upload


curl -F "blob=@1.jpeg;type=image/jpeg" "http://127.0.0.1:4869/upload"

curl -H "Content-Type:jpeg" --data-binary @1.jpeg "http://127.0.0.1:4869/upload"
