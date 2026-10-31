# Remove a patient chart

**Status:** Not built

## Who this is for

A doctor or an administrator needs to take a chart off file when the hospital no longer keeps it.

## Requirement

A doctor or an administrator can remove a chart that has no upcoming appointment. After removal, that chart cannot be found.

## Information

The hospital needs to know which chart is being removed.

## Who may do this

- Doctor and administrator: yes
- Patient: no
- Anyone who is not signed in: no

## Rules

- The chart must already exist. An unknown chart is refused.
- An upcoming appointment is one that is still scheduled and whose time has not yet passed.
- Removal is refused while any upcoming appointment belongs to that patient. The chart and those appointments stay as they are.
- A chart whose appointments are all cancelled, or are all in the past, may be removed.
- After a successful removal, a search and a direct view no longer find that chart.
- A refused attempt leaves the chart unchanged.

## Done when

- A doctor can remove a chart with no upcoming appointment, and that chart can no longer be found.
- An administrator can do the same.
- Removal is refused while an upcoming appointment still belongs to that patient, and the chart remains.
- A patient, and anyone who is not signed in, is refused.
- An unknown chart is refused.
