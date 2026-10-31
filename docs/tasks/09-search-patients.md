# Search patients

**Status:** Not built

## Who this is for

A doctor or an administrator needs to find charts without already knowing which record to open.

## Requirement

A doctor or an administrator can page through charts and narrow the list by patient name or medical record number.

## Information

Each entry shows the same details as opening that chart: patient name, primary doctor, medical record number, date of birth, gender, phone, blood group, and notes.

## Who may do this

- Doctor and administrator: yes
- Patient: no
- Anyone who is not signed in: no

## Rules

- Results come back one page at a time. The reader can ask for a later page.
- A page holds at most 20 charts.
- A name filter matches charts whose patient name contains the given text, ignoring letter case.
- A medical record number filter matches charts whose number contains the given text. Letter case and spaces are ignored.
- Filters that are used together all have to match.
- With no filters, the page lists charts in patient name order.
- When nothing matches, the result is an empty page.
- Searching does not change any chart.

## Done when

- A doctor can move through pages and see charts in patient name order.
- An administrator can do the same.
- Name and medical record number filters return only the charts that match, including when both filters are used together.
- An empty match is an empty page.
- A patient, and anyone who is not signed in, is refused.
