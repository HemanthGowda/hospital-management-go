# View a patient chart

**Status:** Not built

## Who this is for

A patient needs to read their own chart. A doctor or an administrator needs to open any chart.

## Requirement

A patient can open only the chart that belongs to their account. A doctor or an administrator can open any chart that is already on file.

## Information

The view shows the patient name, primary doctor, medical record number, date of birth, gender, phone, blood group, and notes. Empty optional fields are shown as empty.

## Who may do this

- Patient: yes, for their own chart only
- Doctor and administrator: yes, for any chart
- Anyone who is not signed in: no

## Rules

- The chart must already exist. An unknown chart is refused.
- A patient who opens a chart that belongs to someone else is refused.
- A patient with no chart is told that no chart was found.
- Opening a chart does not change it.

## Done when

- A patient can open their own chart and see the saved details.
- A patient who asks for another person's chart is refused, and that chart is unchanged.
- A doctor and an administrator can each open any chart.
- Anyone who is not signed in is refused.
- An unknown chart is refused.
