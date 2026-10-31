# Hospital Management Backend (Go + PostgreSQL Intern Starter)

Half-baked Go and PostgreSQL hospital API used as a **reference implementation** for interns. It mirrors the Java/MongoDB starter: the same three reference slices, the same roles, and the same expectation that interns finish the product.

Only a few vertical slices are fully implemented. The **Doctor** resource is intentionally missing so interns can design the schema, API, tests, and docs from the ground up while copying the existing patterns.

## What is already implemented

| Method | Path | Auth | Purpose |
|--------|------|------|---------|
| `POST` | `/api/v1/users` | Public | Register a `PATIENT` account |
| `POST` | `/api/v1/auth/login` | Public | Issue a JWT access token |
| `PUT` | `/api/v1/patients/{id}` | `ADMIN` or `DOCTOR` | Update an existing patient chart |

Supporting pieces included:

- Layered architecture: handler → service → repository
- Request validation with field-level error responses
- JWT bearer authentication with role-based access
- SQL migrations and startup seed data
- Unit tests and Testcontainers integration tests
- OpenAPI document and Swagger UI
- Docker Compose runtime (API + PostgreSQL)

## What is intentionally incomplete

- **Doctor** resource: no table, model, repository, or API. Interns own this end-to-end. Patients currently store `primaryDoctorName` as free text.
- Patient create/list/get/delete HTTP endpoints (`Insert` exists only so startup seeding has a repository method to copy)
- Appointment booking workflows
- User admin APIs (list, role changes, disable)
- Refresh tokens, password reset, email verification

Each remaining capability is specified in [docs/tasks/README.md](docs/tasks/README.md). See [docs/intern-roadmap.md](docs/intern-roadmap.md) for the suggested pace.

## Architecture

```text
Client
  │
  ▼
patient.Handler / user.Handler / auth.Handler
  │
  ▼
patient.Service / user.Service / auth.Service + JWTService
  │
  ▼
PostgresRepository
  │
  ▼
PostgreSQL (users, user_roles, patients)  ← SQL migrations
```

Package layout:

- `cmd/api` — process entrypoint
- `internal/auth` — login, JWT, principal middleware
- `internal/user` — registration slice
- `internal/patient` — update slice
- `internal/db/migrations` — schema SQL
- `internal/seed` — admin user and sample patients
- `internal/httpserver` — routes, health, OpenAPI
- `internal/domain` / `internal/httpx` — shared errors and JSON responses
- `openapi/openapi.yaml` — API contract

## Prerequisites

- Docker Desktop (or Docker Engine + Compose v2)
- Optional for local IDE runs: Go 1.26+

## Quick start (Docker only)

```bash
cp .env.example .env
docker compose up --build
```

Services:

- API: [http://localhost:8082](http://localhost:8082)
- Health: [http://localhost:8082/actuator/health](http://localhost:8082/actuator/health)
- Swagger UI: [http://localhost:8082/swagger-ui.html](http://localhost:8082/swagger-ui.html)
- OpenAPI JSON: [http://localhost:8082/v3/api-docs](http://localhost:8082/v3/api-docs)
- PostgreSQL: `localhost:5433`, database `hospital_management`, user `hospital`

Seeded admin (override via `.env`):

- Email: `admin@hospital.local`
- Password: `ChangeMe123!`

Host ports are `8082` and `5433` so this stack can run beside the library starter. Inside the API container the process still listens on `8080`.

Swagger UI assets are loaded from a CDN. The OpenAPI document itself is served by the API and also stored at `openapi/openapi.yaml`.

### Example requests

Register:

```bash
curl -s -X POST http://localhost:8082/api/v1/users \
  -H 'Content-Type: application/json' \
  -d '{"email":"patient@example.com","password":"PatientPass1","fullName":"Ada Patient"}'
```

Login:

```bash
TOKEN=$(curl -s -X POST http://localhost:8082/api/v1/auth/login \
  -H 'Content-Type: application/json' \
  -d '{"email":"admin@hospital.local","password":"ChangeMe123!"}' \
  | jq -r .accessToken)
```

Update a patient (replace `{id}` with a seeded patient UUID):

```bash
curl -s -X PUT "http://localhost:8082/api/v1/patients/{id}" \
  -H "Authorization: Bearer $TOKEN" \
  -H 'Content-Type: application/json' \
  -d '{
    "fullName":"Ada Lovelace",
    "primaryDoctorName":"Dr. Grace Hopper",
    "medicalRecordNumber":"MRN-100001",
    "dateOfBirth":"1815-12-10",
    "gender":"FEMALE",
    "phone":"+15551230001",
    "bloodGroup":"O+",
    "notes":"Updated notes"
  }'
```

List seeded patient ids:

```bash
docker exec hospital-go-postgres psql -U hospital -d hospital_management -c 'SELECT id, full_name, medical_record_number FROM patients;'
```

## Configuration

| Variable | Default | Description |
|----------|---------|-------------|
| `DATABASE_URL` | `postgres://hospital:hospital@localhost:5433/hospital_management?sslmode=disable` | PostgreSQL connection string |
| `SERVER_PORT` | `8080` | HTTP port inside the container. Compose maps host port `8082` to it. |
| `JWT_SECRET` | long demo secret | HMAC signing secret (min 32 characters) |
| `JWT_EXPIRATION_MS` | `3600000` | Access token lifetime |
| `JWT_ISSUER` | `hospital-management` | JWT issuer claim |
| `SEED_ADMIN_EMAIL` | `admin@hospital.local` | Admin created at startup |
| `SEED_ADMIN_PASSWORD` | `ChangeMe123!` | Admin password (hashed before insert) |
| `SEED_ADMIN_FULL_NAME` | `Hospital Admin` | Admin display name |

The admin password is never stored in SQL. `internal/seed` hashes it with bcrypt when the process starts, and skips the insert if that email already exists.

## Tests

```bash
go test ./...
```

- Unit tests cover services and JWT helpers (no Docker required).
- Integration tests use Testcontainers PostgreSQL (Docker required).

## Local development without packaging the API image

```bash
docker compose up postgres -d
go run ./cmd/api
```

`go run` reads `DATABASE_URL`. If you use the Compose defaults, point it at `localhost:5433`.

## Patterns interns should copy

1. Keep handlers thin. Put business rules in services.
2. Validate input in the service and return `domain.ValidationFailed`.
3. Map database rows to response structs in the feature package.
4. Return `domain.Error` values and let `httpx.WriteError` shape the JSON body.
5. Protect endpoints by reading the principal set by `auth.Middleware`.
6. Add the next numbered SQL migration under `internal/db/migrations` with both `.up.sql` and `.down.sql`.
7. Add unit tests and an HTTP integration test for every new feature.
8. Update `openapi/openapi.yaml` when the contract changes.

Issued JWTs are trusted until they expire. Disabling a user does not revoke tokens that were already issued. Call that out if you add account-disable APIs.

## Troubleshooting

| Symptom | Likely fix |
|---------|------------|
| API unhealthy on first boot | Wait for Postgres health and migrations; check `docker compose logs api` |
| `401` on patient update | Login again; send `Authorization: Bearer <token>` |
| `403` on patient update | Use an `ADMIN` or `DOCTOR` account (the seeded admin works) |
| JWT secret error on startup | Set `JWT_SECRET` to at least 32 characters |
| Integration tests fail | Ensure Docker is running for Testcontainers |
| Port already in use | Host defaults are `8082` (API) and `5433` (Postgres) |

## License

Internal training material. Use and extend freely within your internship program.
