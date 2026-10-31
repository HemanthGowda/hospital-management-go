# Register a patient chart

**Status:** Not built

## Who this is for

A doctor or an administrator needs to open a chart for a patient. Registering for an account creates a login only. A chart is a separate record.

## Requirement

A doctor or an administrator can open a chart for an existing patient account. The chart can be found later with the details that were saved.

## Information

- The patient account the chart belongs to (required)
- Full name (required)
- Primary doctor (required). Once doctors are part of the directory, the chart names a doctor the directory already knows.
- Medical record number (required)
- Date of birth (required)
- Gender (required): female, male, other, or unknown
- Phone (optional)
- Blood group (optional)
- Notes (optional)

## Who may do this

- Doctor and administrator: yes
- Patient: no
- Anyone who is not signed in: no

## Rules

- The chart belongs to a patient account that already exists. One patient account has at most one chart.
- Full name is required and is between 2 and 160 characters.
- Primary doctor is required. The displayed doctor name is between 1 and 160 characters.
- Medical record number is required and looks like `MRN-100001`: the letters MRN, a hyphen, and six digits. Two charts cannot share the same medical record number. Comparison ignores letter case and spaces.
- Date of birth is required and is not in the future.
- Gender is required and is one of female, male, other, or unknown.
- Phone, when present, is between 7 and 20 characters and uses digits, spaces, or the characters `+`, `(`, `)`, `.`, and `-`.
- Blood group, when present, is one of A+, A-, B+, B-, AB+, AB-, O+, or O-.
- Notes, when present, are at most 2,000 characters.
- A refused attempt does not open a chart.

## Done when

- A doctor can open a chart for an existing patient and find it later with the same name, doctor, medical record number, date of birth, gender, phone, blood group, and notes.
- An administrator can do the same.
- A patient, and anyone who is not signed in, is refused.
- A second chart for the same patient account, or a second chart with the same medical record number, is refused.
- A name, doctor, medical record number, date of birth, gender, phone, or blood group that breaks the rules above is refused, and the reason names the problem.
