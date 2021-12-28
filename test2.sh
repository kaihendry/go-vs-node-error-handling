#!/bin/bash

cat <<EOJSON >/tmp/test.json
{ "human": { "age": 12 } }
EOJSON

curlie -X POST http://localhost:3000 \
	-H "Content-Type: application/json" \
	-d @/tmp/test.json
