Functionality:

"stack" command
- [x] start a new stack (ie git checkout -b <branch>-<stack>-1)
- [x] add a new branch to stack (ie git checkout -b <branch>-<stack>-<n+1>)
- [ ] remove a branch from stack (and retarget children)
- [ ] move directly to a given stack
- [ ] offer to commit working tree before stacking

"up" command
- [x] move to child stack

"down" command
- [x] move to parent stack

"write" command
- [ ] commit to current branch (ie git commit [--amend])
- [ ] commit subset of current diff to the branch

"show" command
- [x] visualize the stack
- [ ] add commit info
- [ ] add date/time info
- [ ] add sync info

"restack" command
- [ ] rebase everything in current stack

"sync" command
- [ ] pull trunk and restack everything

Architectural:
- [ ] Allow for multiple children per stack
