Functionality:

"stack" command
- [ ] remove a branch from stack (and retarget children)
- [ ] move directly to a given stack

"up" command
- [ ] move to child stack - select if many children

"show" command
- [ ] add commit info
- [ ] add date/time info

"pr submit" command
- [ ] Opens pull request from current stack into parent
- [ ] Optional flag to also open pull request for all stacks above current stack (whole tree)
- [ ] Checks for existing pull requests first
- [ ] Adds comment tracking all PR(s) in stack

"pr update" command
- [ ] Finds existing PR(s) and updates them
- [ ] Updates comment tracking all PR(s) in stack

Architectural:
- [ ] Update cache after any modifications
- [ ] Add helpful error messages (or at least surface git output)
- [ ] Add continue ability for sync operation
- [ ] Handle branches being deleted
