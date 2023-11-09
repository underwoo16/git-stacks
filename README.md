Functionality:

"stack" command
- [ ] start a new stacked workflow (ie git checkout -b <branch>-<stack>-1)
- [ ] add a new layer to stack (ie git checkout -b <branch>-<stack>-<n+1>)
- [ ] remove a layer from the stack (ie git branch -D <branch>-<stack>-<n>)

"write" command
- [ ] add current diff to the stack (ie git commit --amend)
- [ ] add subset of current diff to the stack

"show" command
- [ ] visualize the stack (???)

"move" command
- [ ] move to a different layer in the stack (ie git checkout <branch>-<stack>-<n>)

"sync" command
- [ ] sync the whole stack (ie git rebase --update-refs ???)