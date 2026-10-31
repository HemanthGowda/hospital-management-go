# Review upcoming appointments

**Status:** Not built

## Who this is for

A doctor needs to see the appointments still ahead on their own schedule. An administrator needs the same view across every doctor.

## Requirement

A doctor can see their own upcoming appointments. An administrator can see upcoming appointments for every doctor.

## Information

Each entry shows the patient, the doctor, the scheduled time, the reason, and that the appointment is still scheduled.

## Who may do this

- Doctor: yes, for their own appointments
- Administrator: yes, for every doctor
- Patient: no
- Anyone who is not signed in: no

A patient seeing the appointments they booked is covered by booking an appointment.

## Rules

- An upcoming appointment is still scheduled, and its time has not yet passed.
- A cancelled appointment does not appear.
- An appointment whose time has passed does not appear.
- A doctor sees only appointments that name them.
- The list can be paged, at most 20 appointments at a time, soonest first.
- When nothing is upcoming, the result is an empty page.
- Reading the list does not change any appointment.

## Done when

- A doctor can see a future scheduled appointment that names them, including the patient, the time, and the reason.
- That doctor does not see another doctor's appointments.
- An administrator can see upcoming appointments across doctors.
- A cancelled appointment, and an appointment whose time has passed, are absent from the list.
- A patient, and anyone who is not signed in, is refused.
