
# git-stacks

Git utility for managing stacked branches - allowing developers to easily leverage smaller, dependent pull requests.


## Usage/Examples

#### Create new "stack" named "feature_branch"
```cli
gs stack feature_branch
```

#### View the current stack
```cli
gs show
```

#### Rebase all stacks
```cli
gs sync
```

#### Create pull request from current stack into its parent
```cli
gs pr
```

* requires gh cli to be installed and authenticated

#### Create pull requests for each stack
```cli
gs pr all
```

* requires gh cli to be installed and authenticated