# Cancel an appointment

**Status:** Not built

## Who this is for

The patient, the doctor on that appointment, or an administrator needs to call off an appointment so the doctor's time is free again.

## Requirement

The patient who booked the appointment, the doctor it is with, or an administrator can cancel it. The appointment is marked cancelled, and that doctor can be booked at that time again.

## Information

The cancelled appointment remembers that it is cancelled and no longer holds the doctor's time.

## Who may do this

- The patient who booked that appointment: yes
- The doctor named on that appointment: yes
- Administrator: yes, for any scheduled appointment
- A different patient, or a doctor who is not on that appointment: no
- Anyone who is not signed in: no

## Rules

- The appointment must already exist and still be scheduled.
- Only an upcoming appointment can be cancelled. The scheduled time has not yet passed.
- Cancelling an appointment that is already cancelled is refused.
- Cancelling an appointment whose time has already passed is refused, and it stays as it was.
- A successful cancellation marks the appointment cancelled and frees that minute for the same doctor.
- A refused attempt leaves the appointment scheduled.

## Done when

- The patient who booked an upcoming appointment can cancel it. It shows as cancelled, and another patient can book that doctor at that minute.
- The doctor named on the appointment, and an administrator, can cancel it with the same result.
- A different patient, and a doctor who is not on the appointment, are refused, and the appointment stays scheduled.
- Cancelling an appointment that is already cancelled, or whose time has passed, is refused.
- Anyone who is not signed in is refused.
