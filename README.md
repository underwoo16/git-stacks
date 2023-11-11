Functionality:

"stack" command
- [ x ] start a new stack (ie git checkout -b <branch>-<stack>-1)
- [ x ] add a new branch to stack (ie git checkout -b <branch>-<stack>-<n+1>)
- [ ] remove a branch from stack (ie git branch -D <branch>-<stack>-<n>)
- [ ] move to a given branch (ie git checkout <branch>-<stack>-<n>)
- [ ] offer to commit working tree before stacking

"up" command
- [ ] move to child stack

"down" command
- [ ] move to parent stack

"write" command
- [ ] commit to current branch (ie git commit [--amend])
- [ ] commit subset of current diff to the branch

"show" command
- [ ] visualize the stack

"sync" command
- [ ] sync the whole stack (ie git rebase --update-refs ???)