# View a doctor

**Status:** Not built

## Who this is for

A signed-in person needs to open one doctor and read the directory record.

## Requirement

Anyone who is signed in can open a single doctor that is already on file.

## Information

The view shows the full name, specialty, license number, and department saved for that doctor.

## Who may do this

- Patient, doctor, and administrator: yes
- Anyone who is not signed in: no

## Rules

- The doctor must already exist.
- An unknown doctor is refused, and the reason says the doctor was not found.
- Opening a doctor does not change the record.

## Done when

- A patient, a doctor, and an administrator can each open a doctor and see the saved details.
- Anyone who is not signed in is refused.
- An unknown doctor is refused, and no doctor record changes.
