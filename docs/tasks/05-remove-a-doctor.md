# Remove a doctor

**Status:** Not built

## Who this is for

A doctor or an administrator needs to take a clinician out of the directory when that person should no longer be offered as a primary doctor.

## Requirement

A doctor or an administrator can remove a doctor who is not the primary doctor on any patient chart. After removal, that doctor cannot be found.

## Information

The hospital needs to know which doctor is being removed.

## Who may do this

- Doctor and administrator: yes
- Patient: no
- Anyone who is not signed in: no

## Rules

- The doctor must already exist. An unknown doctor is refused.
- Removal is refused while any patient chart still names that doctor as the primary doctor. Those charts stay on file.
- After a successful removal, a search and a direct view no longer find that doctor.
- A refused attempt leaves the doctor and the charts unchanged.

## Done when

- A doctor can remove a directory record that no chart names, and that doctor can no longer be found.
- An administrator can do the same.
- Removal is refused while a chart still names that doctor, and the doctor and the chart both remain.
- A patient, and anyone who is not signed in, is refused.
- An unknown doctor is refused.
