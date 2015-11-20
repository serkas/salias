#! /bin/bash

DOC='{"id":"demo_doc", "context":"1", "text": "Nice red apple"}'
curl -XGET -H "Content-Type: application/json" -d "$DOC"  http://localhost:8080/classify