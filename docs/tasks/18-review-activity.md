# Review activity

**Status:** Not built

## Who this is for

An administrator needs to see who changed the doctors, the charts, the appointments, and the accounts, and when they did it.

## Requirement

The hospital records accountable activity, and an administrator can read that history. Passwords are never part of the history.

## Information

Each entry records:

- When it happened
- Who did it (name and email), or that the person was not signed in when the action allows that
- What happened
- Which doctor, chart, appointment, or account was affected

The history covers:

- Adding, updating, and removing a doctor
- Registering, updating, and removing a patient chart
- Booking and cancelling an appointment
- Changing a role
- Suspending and restoring an account

## Who may do this

- Administrator: yes
- Patient and doctor: no
- Anyone who is not signed in: no

## Rules

- An entry is recorded when the action succeeds. A refused attempt is not recorded.
- Passwords, renewals, and anything else that could be used to sign in are never stored in the history.
- The list can be paged, at most 20 entries at a time, newest first.
- The reader can narrow the list to one kind of record (doctors, charts, appointments, or accounts).
- Reading the history does not change it.

## Done when

- After a doctor opens a chart, an administrator can see that doctor, the chart, and the time.
- After a patient books or cancels an appointment, an administrator can see that action in the history.
- After an administrator changes a role or suspends an account, that action appears in the history.
- No history entry contains a password.
- A patient, a doctor, and anyone who is not signed in are refused.
