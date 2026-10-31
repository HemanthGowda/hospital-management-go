# Tie patients to doctors

**Status:** Not built

## Who this is for

Clinicians need a patient chart to name a real doctor from the directory, so a rename or a correction shows up on every chart that uses that doctor.

## Requirement

A chart names a doctor the directory already knows. When that doctor's name changes, every chart that names them shows the new name. Charts already on file remain readable.

## Information

Each chart remembers which directory doctor is the primary doctor, and shows that doctor's current name.

## Who may do this

This changes the rules for registering a chart and for updating a chart.

- Doctor and administrator: may register or update a chart so that it names a known doctor
- Patient: no
- Anyone who is not signed in: no

## Rules

- Registering a chart, or changing a chart, requires a primary doctor that can already be found in the directory.
- An unknown doctor is refused, and the chart is left unchanged. A new chart is not created.
- When a doctor's name is corrected, every chart that names that doctor shows the corrected name.
- Charts already on file before this capability exists stay available. Each of them still shows a primary doctor name a reader can recognize.
- Removing a doctor is still refused while any chart names that doctor.

## Done when

- A doctor can save a chart only when it names a doctor already in the directory.
- Saving a chart with an unknown doctor is refused.
- After a doctor's name is corrected, a reader sees the new name on every chart that names that doctor.
- Charts that were already on file can still be opened, and each still shows a primary doctor name.
