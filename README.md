# ranked-distributor

Given ranked preferences for multiple users, allocate items in accordance with their
preferences.  A random initial order is chosen for the customers.  Each round the
order is reversed (the user who selected last selects first in the next round) in
a snake draft.

That is, given 4 users, one repeating selection pattern might be (0,2,3,1,1,3,2,0).
Each user will select the item they want the most during each of their turns until
no items remain.  This requires the users to rank the items beforehand.