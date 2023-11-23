Functionality:

"delete" command
- [ ] remove a branch from stack (and retarget children)

"down" command
- [ ] rename

"up" command
- [ ] rename

"pr" command
- [ ] Adds comment tracking all PR(s) in stack

Architectural:
- [ ] Update cache after any modifications
- [ ] Add helpful error messages (or at least surface git output)
- [ ] Add continue ability for sync operation (using rerere)
- [ ] Handle branches being deleted
- [ ] Potential idea - Rebuild stack based on local refs/heads with shared history (useful for codespaces)
- [ ] Be more thoughtful about how to initialize trunk when no config exists

----------------------------------------------------------------------------------------

plan to create a stack ${stack_name}:

create new hash-object
    ${parent branch name}
    ${parent branch revision}

update-ref refs/stacks/${stack_name} with resultant hash-object sha

update-ref refs/heads/${stack_name} with parent branch revision

checkout refs/heads/${stack_name}

----------------------------------------------------------------------------------------

Find location of .git
 - git rev-parse --show-toplevel

Simple status output:
- git status -sb

------------------------------------------------------------------------------------------
