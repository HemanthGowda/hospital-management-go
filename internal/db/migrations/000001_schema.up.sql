CREATE TABLE users (
    id UUID PRIMARY KEY,
    email TEXT NOT NULL,
    password_hash TEXT NOT NULL,
    full_name TEXT NOT NULL,
    enabled BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE UNIQUE INDEX uk_users_email ON users (lower(email));

CREATE TABLE user_roles (
    user_id UUID NOT NULL REFERENCES users (id) ON DELETE CASCADE,
    role TEXT NOT NULL CHECK (role IN ('ADMIN', 'DOCTOR', 'PATIENT')),
    PRIMARY KEY (user_id, role)
);

CREATE TABLE patients (
    id UUID PRIMARY KEY,
    full_name TEXT NOT NULL,
    primary_doctor_name TEXT NOT NULL,
    medical_record_number TEXT NOT NULL,
    date_of_birth DATE NOT NULL,
    gender TEXT NOT NULL CHECK (gender IN ('FEMALE', 'MALE', 'OTHER', 'UNKNOWN')),
    phone TEXT,
    blood_group TEXT,
    notes TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE UNIQUE INDEX uk_patients_mrn ON patients (medical_record_number);
CREATE INDEX idx_patients_full_name ON patients (full_name);
