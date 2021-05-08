#!/bin/zsh

curl http://localhost:9000/api/posts -v | jq # need jq installed to pretty print json or remove the | jq

