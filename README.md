Functionality:
- [ ] start a new stacked workflow (ie git checkout -b <branch>-<stack>-1)
- [ ] add current diff to the stack (ie git commit --amend)
- [ ] add subset of current diff to the stack
- [ ] add a new layer to stack (ie git checkout -b <branch>-<stack>-<n+1>)
- [ ] move to a different layer in the stack (ie git checkout <branch>-<stack>-<n>)
- [ ] remove a layer from the stack (ie git branch -D <branch>-<stack>-<n>)
- [ ] sync the whole stack (ie git rebase --update-refs ???)