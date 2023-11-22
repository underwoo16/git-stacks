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
