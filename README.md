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
     -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySUQiOjF9.GVrvCL8mQdY1ubV1w3MEYbysYvQGAj2RzW_0GMYlvVs" \
     -d '{"answers": [{"question": "Whats the best programming language?", "answer": "Go"}]}'
```

Return codes:
- `200 OK` if signing is successfull
- `400 Bad Request` if body or userID is the wrong format
- `401 Unauthorized` if JWT isn't valid or userID doesn't match
- `500 Internal Server Error` for any server issues


### Verify Signature (GET /api/v1/users/{userID}/signatures/{signatureID})

Doesn't require a JWT, just send a GET request to fetch the signed test for a specific user:
```bash
curl -v "http://localhost:8000/api/v1/users/1/signatures/1"
```

If the signed test exists for that specific user, response:
```json
{
     "userId":1,
     "signatureId":1,
     "timestamp":"2024-02-24T19:45:59Z",
     "answers": [
          {
               "question":"Whats the best programming language?",
               "answer":"Go"
          }
     ]
}
```

Return codes:
- `200 OK` if signed test exists for that user
- `400 Bad Request` if signatureID or userID are invalid
- `404 Not Found` if the signature doesn't exist
- `500 Internal Server Error` for any server issues
