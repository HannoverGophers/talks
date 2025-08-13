BASE_URL="https://multiple-routes-643277839112.europe-west1.run.app"

curl -s "$BASE_URL/info"
# service=cloudrun-example host=<pod> time=...

curl -s -X POST "$BASE_URL/echo" \
  -H "Content-Type: application/json" \
  -d '{"message":"hello cloud run"}' | jq .
