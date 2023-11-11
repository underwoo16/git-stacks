plan to create a stack ${stack_name}:

create new hash-object
    ${parent branch name}
    ${parent branch revision}

update-ref refs/stacks/${stack_name} with resultant hash-object sha

update-ref refs/heads/${stack_name} with parent branch revision

checkout refs/heads/${stack_name}

--------------------------------------------

Find location of .git
 - git rev-parse --show-toplevel

----------------------------------------------
plan to show stack
 - each refs/stacks points to its parent
 - create nodes and then point them at eachother
 - trunk is node that points to nil
 - tip is node at other end
 - start at tip and print each node
 - mark current node when it is seen