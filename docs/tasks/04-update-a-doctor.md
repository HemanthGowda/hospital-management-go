# Update a doctor

**Status:** Not built

## Who this is for

A doctor or an administrator needs to correct a clinician's name, specialty, license number, or department after the record was first saved.

## Requirement

A doctor or an administrator can replace the saved details of one doctor. The updated doctor can still be found later.

## Information

The same information as adding a doctor: full name, specialty, license number, and department. All four are required.

## Who may do this

- Doctor and administrator: yes
- Patient: no
- Anyone who is not signed in: no

## Rules

- The doctor must already exist. An unknown doctor is refused.
- Full name, specialty, license number, and department follow the same limits as when a doctor is added.
- The license number stays unique regardless of letter case and extra spaces.
- The doctor may keep their own current license number.
- A refused attempt leaves that doctor unchanged.

## Done when

- A doctor can change a directory record, and a later view shows the new values.
- An administrator can do the same.
- Changing a license number to one another doctor already has is refused, and both records stay as they were.
- A patient, and anyone who is not signed in, is refused.
- An unknown doctor is refused.
