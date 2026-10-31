# API Overview

Interactive docs: `/swagger-ui.html`  
OpenAPI JSON: `/v3/api-docs`  
Source document: [`openapi/openapi.yaml`](../openapi/openapi.yaml)

## Security

- Public: registration, login, health, info, and Swagger
- Protected: every other `/api/v1/**` route requires `Authorization: Bearer <jwt>`
- Patient update additionally requires role `ADMIN` or `DOCTOR`

JWT claims used by the middleware:

- `sub` — user id
- `email`
- `roles` — array of role names
- `iss`, `iat`, `exp`

## Error shape

Handled errors return:

```json
{
  "timestamp": "2026-01-15T10:15:30.123Z",
  "status": 400,
  "error": "Validation Failed",
  "message": "Request validation failed",
  "path": "/api/v1/users",
  "violations": [
    { "field": "email", "message": "Email must be valid" }
  ]
}
```

## Persistence

| Table | Purpose |
|-------|---------|
| `users` | Accounts. Email uniqueness is case-insensitive. |
| `user_roles` | `ADMIN`, `DOCTOR`, or `PATIENT` |
| `patients` | Patient charts. Medical record number is unique. |

Identifiers are UUIDs. `primary_doctor_name` is free text until the Doctor resource exists.

## Implemented contracts

### `POST /api/v1/users`

Creates a `PATIENT` account.

- `201` user payload (password is never returned)
- `400` validation errors
- `409` duplicate email

### `POST /api/v1/auth/login`

Returns a bearer token.

- `200` with `accessToken`, `expiresInMs`, and roles
- `400` validation errors
- `401` invalid credentials or a disabled account

### `PUT /api/v1/patients/{id}`

Replaces mutable patient fields.

- `200` updated patient
- `400` validation failure, or a date of birth in the future
- `401` missing or invalid token
- `403` authenticated but wrong role
- `404` unknown id
- `409` medical record number belongs to another patient

## Not implemented (intern backlog)

- Doctors (entire resource)
- Patient create, list, get, and delete HTTP endpoints
- Appointments
- Admin user management
