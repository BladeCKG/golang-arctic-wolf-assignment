# Notes

- **State Validation**:

  There are _4_ possible states: **_open_**, **_closed_**, **_accepted_**, **_investigating_**<br>
  In the assignment, no mention to validate this in create risk process<br>
  so i added validation of the created risk' state in create risk handler<br>
  if the state is not matching one of these 4, then respond error: Invalid Risk State<br>
