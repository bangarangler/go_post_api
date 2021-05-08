#!/bin/zsh

curl -X POST http://localhost:9000/api/posts \
  -H "Content-Type: application/json" \
  -d "{\"title\": \"Test Post\", \"content\": \"This is content\", \"author\":
  \"Jon P\"}" \
  -v | jq # need jq installed to pretty print json or remove the | jq

