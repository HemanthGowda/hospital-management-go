# Add a doctor

**Status:** Not built

## Who this is for

A doctor or an administrator needs one directory record per clinician so patient charts can name that clinician consistently.

## Requirement

The hospital keeps a doctor and can find that same doctor later.

## Information

- Full name (required)
- Specialty (required)
- License number (required)
- Department (required)

## Who may do this

- Doctor and administrator: yes
- Patient: no
- Anyone who is not signed in: no

## Rules

- Full name is required. After trimming extra spaces, it is between 2 and 160 characters.
- Specialty is required and is between 1 and 80 characters.
- License number is required and is between 1 and 40 characters.
- Two doctors cannot share the same license number. Comparison ignores letter case and extra spaces.
- Department is required and is between 1 and 80 characters.
- A refused attempt leaves the directory unchanged.

## Done when

- A doctor can add a directory record and find that doctor again with the name, specialty, license number, and department that were saved.
- An administrator can do the same.
- A patient, and anyone who is not signed in, is refused.
- A second doctor with the same license number is refused, and the first record is unchanged.
- A missing name, specialty, license number, or department, or a value that is too long, is refused, and the reason names the problem.
