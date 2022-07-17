#!/bin/bash
curl "https://placement.iitk.ac.in/api/auth/login" \
  -H 'Content-Type: application/json' \
  -X POST \
  -d '{
  "user_id": "",
  "password": "",
  "remember_me": true
}'
