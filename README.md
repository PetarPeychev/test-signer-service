# test-signer-service

## Running locally

```
docker compose build
docker compose up
```

## Endpoints

### Sign Answers (POST /api/v1/users/{userID}/signatures)
Requires a JWT with a claim containing a userID:
```json
{
  "userID": "1"
}
```
HS256-signed with this key `b139b70d-3c7d-4aab-9244-c7c6da093b9e`.

Which can then be used to submit the POST request:
```bash
curl -v -X POST "http://localhost:8000/api/v1/users/1/signatures" \
     -H "Content-Type: application/json" \
     -H "Authorization: Bearer eyJhbGciOiJIUzI1NiJ9.eyJ1c2VySUQiOiIxIn0.8cyk4-zbs3h1u7RGKG4kmM_zdbHdx-4EXGGoQklgoyc" \
     -d '{"answers": [{"question": "Whats the best programming language?", "answer": "Go"}]}'
```

### Verify Signature (GET /api/v1/users/{userID}/signatures/{signatureID})

Doesn't require a JWT, just send a GET request to fetch the signed test for a specific user:
```bash
curl -v "http://localhost:8000/api/v1/users/1/signatures/1"
```

If the signed test exists for that specific user, response is `200 OK`:
```json
{"userId":1,"signatureId":1,"timestamp":"2024-02-24T19:45:59Z"}
```

If the signature doesn't exist, it's a `404 Not Found`.