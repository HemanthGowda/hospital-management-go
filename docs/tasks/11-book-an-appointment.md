# Book an appointment

**Status:** Not built

## Who this is for

A patient needs to book a time with a doctor who is already in the directory, and later see the appointments they booked.

## Requirement

A patient can book an appointment with an existing doctor. The same doctor cannot have two appointments at the same time. The patient can see their own appointments.

## Information

An appointment remembers:

- Which patient booked it
- Which doctor it is with
- The scheduled date and time, to the minute
- The reason
- That it is scheduled

## Who may do this

- Patient: yes, for themselves
- Doctor and administrator: no
- Anyone who is not signed in: no

A patient can list the appointments they booked. A doctor listing their own upcoming appointments is a separate capability.

## Rules

- The doctor must already exist in the directory. An unknown doctor is refused, and no appointment is created.
- The scheduled time is required and is in the future.
- The reason is required and is between 1 and 500 characters.
- Two appointments for the same doctor at the same minute are refused. The first appointment stays scheduled, and that minute stays taken.
- A cancelled appointment no longer holds the minute. A new appointment may use it.
- Two patients may book different doctors at the same minute.
- A refused attempt creates no appointment.

## Done when

- A patient can book a future time with an existing doctor and later see that appointment, including the doctor, the time, the reason, and that it is scheduled.
- Booking the same doctor at that same minute again is refused, and the first appointment is unchanged.
- After the first appointment is cancelled, that same minute can be booked again.
- An unknown doctor, a time in the past, or a missing reason is refused.
- A doctor, an administrator, and anyone who is not signed in are refused when they try to book.
- A patient sees only the appointments they booked.
