Functionality:

"delete" command
- [ ] remove a branch from stack (and retarget children)

"down" command
- [ ] rename

"up" command
- [x] move to child stack - select if many children
- [ ] rename

"show" command
- [x] add commit info
- [ ] add date/time info

"pr submit" command
- [x] Opens pull request from current stack into parent
- [x] Opens pull requests for all children into current stack (and so on and so forth)
- [ ] Checks for existing pull requests first
- [ ] Adds comment tracking all PR(s) in stack
- [ ] Updates PRs if they already exist (resync and force push and update comments)

Architectural:
- [ ] Update cache after any modifications
- [ ] Add helpful error messages (or at least surface git output)
- [ ] Add continue ability for sync operation (use rerere)
- [ ] Handle branches being deleted
- [ ] Rebuild stack based on local refs/heads with shared history (useful for codespaces)
- [ ] Be more thoughtful about how to initialize trunk when no config exists
