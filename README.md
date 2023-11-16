Functionality:

"stack" command
- [x] start a new stack (ie git checkout -b <branch>-<stack>-1)
- [x] add a new branch to stack (ie git checkout -b <branch>-<stack>-<n+1>)
- [ ] remove a branch from stack (and retarget children)
- [ ] move directly to a given stack

"up" command
- [x] move to child stack

"down" command
- [x] move to parent stack

"write" command
- [ ] commit to current branch
- [ ] sync from current branch 
- [ ] commit subset of current diff to the branch

"show" command
- [x] visualize the stack
- [ ] add commit info
- [ ] add date/time info
- [ ] add sync info

"sync" command
- [x] rebase everything in current stack

Architectural:
- [x] Cache metadata locally (trunk, branches + children, parent branch + ref)
- [x] Allow for multiple children per stack
- [ ] Update cache after any modifications
- [ ] Add helpful error messages (or at least surface git output)
- [ ] Handle branches being deleted
