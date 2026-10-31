# Intern Roadmap

This project ships with **partial** functionality on purpose. Use the implemented user registration, login, and patient-update slices as the pattern to copy. The **Doctor** resource has no table and no code—design it from scratch.

Product requirements for each remaining capability are in [tasks/README.md](tasks/README.md), one file per capability. This roadmap is the suggested pace. Follow the task file for what the product must do.

Suggested pace: 1–2 milestones per week. Adjust freely.

## Roles reminder

| Role | Intended capabilities |
|------|------------------------|
| `PATIENT` | Register, view own chart, book appointments (future) |
| `DOCTOR` | Update patient charts, manage appointments |
| `ADMIN` | Everything doctors can do, plus user administration |

---

## Phase 0 — Orientation (Day 1)

**Goal:** Run the stack and understand the reference slices.

Acceptance criteria:

- [ ] `docker compose up --build` succeeds
- [ ] Swagger UI loads and login works with the seeded admin
- [ ] You can identify handler → service → repository for users and patients
- [ ] You can explain how JWT middleware and role checks protect `PUT /api/v1/patients/{id}`

---

## Phase 1 — Doctor domain design (Week 1)

**Goal:** Propose the Doctor resource with **no production code yet**.

Deliverables:

- [ ] Short design note (`docs/design/doctor.md`) covering:
  - Columns (for example full name, specialty, license number, department)
  - Primary key, indexes, and uniqueness rules
  - Relationship to `patients` (free-text `primary_doctor_name` today versus `primary_doctor_id` foreign key)
  - API sketch (paths, methods, roles)
- [ ] Review and approval from a mentor before coding

Acceptance criteria:

- [ ] Design answers create, read, update, delete, and list/search
- [ ] Migration plan described (up and down SQL, plus optional seed doctors)
- [ ] Authorization matrix defined per endpoint

---

## Phase 2 — Doctor persistence + migration (Week 1–2)

**Goal:** Create the PostgreSQL table the same way `users` and `patients` exist.

Deliverables:

- [ ] `000002_doctors.up.sql` and `.down.sql`
- [ ] `Doctor` model and `Repository`
- [ ] Unit tests for any normalization helpers

Acceptance criteria:

- [ ] App starts with the new migration applied
- [ ] Unique constraints behave as designed
- [ ] No Doctor HTTP API yet

---

## Phase 3 — Doctor API vertical slice (Week 2)

**Goal:** Implement at least one full Doctor operation end-to-end, then complete CRUD.

Recommended order (mirror the starter style):

1. `POST /api/v1/doctors`
2. `GET /api/v1/doctors/{id}`
3. `GET /api/v1/doctors`
4. `PUT /api/v1/doctors/{id}`
5. `DELETE /api/v1/doctors/{id}`

Acceptance criteria for each endpoint:

- [ ] Validation and `domain.Error` responses
- [ ] Correct role checks
- [ ] Service unit tests
- [ ] Integration tests (auth success and failure, plus validation)
- [ ] `openapi/openapi.yaml` updated

---

## Phase 4 — Wire Doctors into Patients (Week 3)

**Goal:** Replace or complement the free-text `primary_doctor_name` on patients.

Options (pick one with a mentor):

- Store `primary_doctor_id` on `patients` and resolve doctor details in responses
- Keep `primary_doctor_name` denormalized and sync it when a doctor is renamed

Acceptance criteria:

- [ ] Migration strategy for existing seeded patients
- [ ] Patient update validation rejects unknown doctors if you use a foreign key
- [ ] Tests cover the relationship rules

Also complete the remaining Patient APIs using the update slice as a template:

- [ ] `POST /api/v1/patients`
- [ ] `GET /api/v1/patients/{id}`
- [ ] `GET /api/v1/patients`
- [ ] `DELETE /api/v1/patients/{id}`

---

## Phase 5 — Appointments (Week 3–4)

**Goal:** Introduce a new aggregate from scratch (similar ambition to Doctors).

Suggested table: `appointments` with `patient_id`, `doctor_id`, `scheduled_at`, `status`, and `reason`.

Acceptance criteria:

- [ ] A patient can book an appointment with an existing doctor
- [ ] Double-booking the same doctor at the same time fails
- [ ] Cancel marks the appointment cancelled
- [ ] Doctors can list upcoming appointments
- [ ] Full test and OpenAPI coverage

---

## Phase 6 — User administration and hardening (Week 4+)

Stretch goals:

- [ ] Admin endpoints to list users, change roles, and enable or disable accounts
- [ ] Refresh tokens or shorter-lived access tokens
- [ ] Request logging and audit fields
- [ ] CI workflow running `go test ./...`
- [ ] Login throttling

Remember: today's JWT middleware trusts the token until expiry. A disable-user endpoint should define whether existing tokens keep working.

---

## Definition of Done (every feature)

1. Code follows existing package and naming patterns
2. Validation and authorization are explicit
3. A reversible SQL migration is added when the schema changes
4. Unit and integration tests pass locally
5. OpenAPI documents the new surface area
6. Roadmap checkboxes in this file are updated

## Reference endpoints (already done)

Use these as the gold standard:

- `POST /api/v1/users`
- `POST /api/v1/auth/login`
- `PUT /api/v1/patients/{id}`
