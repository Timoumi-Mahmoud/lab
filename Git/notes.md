##
- Git at its core, is like a key value store.
    Value = Data
    Key = Hash of the Data
I can use the key to retrive the content.

SHA1: Is a cryptographic hash function
Data----> produce 40-digit hexadecimal
The value should always be the same if the given input is the same.

### BLOB
Git store the compressed data in a blob, along with metadata in a header:
[Blob|size ]
[\n
[ hello world]
Bolbs are stored under .git/objects/2CharsOfTheHash/theRestOfHash
Blobs missing the information: filenames and directory structures ==> Git stores this information in a tree

### TREE
Tree contains pointer (using SHA1) :
    to blobs
    to other trees
    and metadata: 
        type of pointer(tree or blob)
        filename or directory name
        mode (executable file, symbolic link, ...)

[tree|size ]
[\n
[blob | 8a55 | go.mod]
[tree | a5df | cmd]

### COMMIT OBJECT
a commit points to:
    a tree
and contains metadata:
    author and committer
    date
    message
    parent commit (one or more)


[commit|size ]
[tree | z0abb ]
[parent | 589ao ]
[author | mahmoud ]
[message |"added readme" ]


```bash
echo "Hello, World!" | git hash-object --stdin
echo "blob 14\0Hello, World!" | openssl sha1

git cat-file -p 8ab686eaf   # contents
git cat-file -t 8ab686eaf   # type: tree, blob, commit
cat .git/HEAD # HEAD a pointer to the current commit

```


```bash
git log --oneline
git --no-pager log --oneline
```

### Objects
Git object are compressed 
As files change, their contents remain mostly similar
Git optimizes for this by compressing these files together, into a Packfile
Packfile stores the object and "deltas", or the difference between one version of the file and the next
Packfile generated when :1. gc 2. Push to remote


-------------------------------------------------------
### Area

1. Working area:
    The files in my working area that are also not in the staging are not handled by git. Intracked files

2. Staging area:
    What files are going to be part of the next commit.
    is how git knows what will change between the current commit and the next commit.

```bash
git ls-files -s ## mode sha1 filename
### Moving files In & out of the staging area
git add <file> ## add a file to the next commit
git rm <file> ## delete a file in the next commit
git mv <file> ## rename a file in the next commit


git add -p
```
Unstage Files from the staging area: Not removing the files , but I replace them with a copy that's currently in the repository. (git checkout)


3. Repo
    The files git know about
    contains all of my commits

---------------------------------------------
### Stash
Save un-committed work, is safe from destructive operations.


```bash
git stash
git stash list
git stash show stash@{0} # show the contents
git stash  apply # apply the last stash 
git stash apply stash@{0} ## apply specific stash

git stash drop ## remove the last stash
git stash drop stash@{0} ## remove specific stash

# By default git stash only tracked files 
git stash --include-untracked 
git statsh --all

git stash save "WIP: making progress on feature X" # name stashes
git stash branch <optional branch name> # start a new branch from a stash
git checkout <stash name> -- <filename> # Grab a single firl from a stash

git stash pop # remove the last stash and applying changes
# doesn't remove if there's a merge conflict
git stash clear

git stash -p # try it
```
-------------------------------------------------------------

### REFERENCES
Three Types of Git References:
    Tags & Annotated Tags
    Branches
    HEAD


* Branch: is just a pointer to a particular commit, will change when new commit will be made
* HEAD: is how git knows what branch I am currently on, and what the next parent will be. 
It's a pointer:
    point at the name of the current branch 
    but it can point at a commit too (detached HEAD)
it moves when:
    make a commit in the currently active branch 
    when I checkhout a new branch


* Tags:
Lightweight tags: simple pointer to a commit
When I create a tag with no arguments, it captures the value HEAD

Annotated tags: point to a commit, but store additional information
    Author, message, date.
```bash
git tag -a v1.0 -m "Version 1.0"

$ git show v1.0
tag v1.0
Tagger: Timoumi_Mahmoud <mahmoud.timoumi@espirt.tn>
Date:   Sat Jun 21 14:03:49 2025 +0100

Version 1.0

commit e704c552dbb5b626f80442aae30101a35ffea215 (HEAD -> Exercice-2, tag: v1.0, tag: my-first-tag, Exercice-3)
Author: Timoumi_Mahmoud <mahmoud.timoumi@espirt.tn>
Date:   Sat Jun 21 13:57:03 2025 +0100

    Update hello.txt

    diff --git a/hello.txt b/hello.txt
    index 980a0d5..d9e24a9 100644
    --- a/hello.txt
    +++ b/hello.txt
    @@ -1 +1,2 @@

     Hello World!
     +new update

```

```bash
git tag # list all the tags in a repo
git show-ref --tags # list all tags, and what commit they're pointing to
git tag --points-at <commit> # list all the tags pointing at a commit
git show <tag-name>

```


### Tags & branches
Branch 
    The current branch pointer moves with every commit to the repository

Tag
    The commit that a tag points doesn't change
    It's a snapshot


### HEAD-LESS / DETACHED HEAD
Checkout to a specific commit(or tag) instead of a branch
git moves the HEAD pointer to that commit

Save my work: create a new branch that points to the last commit I made
```bash
git branch <new-branch-name><commit>
or git switch -s 

lenovo@DESKTOP-U5PPG5V MINGW64 ~/Desktop/GIT/ww (Exercice-2)
$ git checkout e704
Note: switching to 'e704'.

You are in 'detached HEAD' state. You can look around, make experimental
changes and commit them, and you can discard any commits you make in this
state without impacting any branches by switching back to a branch.

If you want to create a new branch to retain commits you create, you may
do so (now or later) by using -c with the switch command. Example:

  git switch -c <new-branch-name>

  Or undo this operation with:

    git switch -

    Turn off this advice by setting config variable advice.detachedHead to false

    HEAD is now at e704c55 Update hello.txt

    lenovo@DESKTOP-U5PPG5V MINGW64 ~/Desktop/GIT/ww ((v1.0))
    $ vim main.go

lenovo@DESKTOP-U5PPG5V MINGW64 ~/Desktop/GIT/ww ((v1.0))
$ git commit -m "added main.go"
[detached HEAD 51a2d85] added main.go
 1 file changed, 7 insertions(+)
  create mode 100644 main.go

  lenovo@DESKTOP-U5PPG5V MINGW64 ~/Desktop/GIT/ww ((51a2d85...))
  $ git status
  HEAD detached from e704c55
  nothing to commit, working tree clean

  lenovo@DESKTOP-U5PPG5V MINGW64 ~/Desktop/GIT/ww ((51a2d85...))
  $ git log
  commit 51a2d852dc1b27609ee0c4441911f6f3711e759e (HEAD)
  Author: Timoumi_Mahmoud <mahmoud.timoumi@espirt.tn>
  Date:   Sat Jun 21 14:25:19 2025 +0100

      added main.go

enovo@DESKTOP-U5PPG5V MINGW64 ~/Desktop/GIT/ww ((51a2d85...))
$ git branch Exercice-4 51a2
```

---------------------------------------------------------------------------------------------
### MERGING

Under the Hood , merge commits are just commits
```bash

 git log
 commit 4c44769cd3cadf495902fac215e994cf7b590ce9 (HEAD -> feature-2)
 Merge: cef8535 7554fc2
 Author: Timoumi_Mahmoud <mahmoud.timoumi@espirt.tn>
 Date:   Sun Jun 22 01:43:23 2025 +0100

     Merge branch 'feature-1' into feature-2

     commit cef853562d55494e17b70c3379fe4977018a3bc7
     Author: Timoumi_Mahmoud <mahmoud.timoumi@espirt.tn>
     Date:   Sun Jun 22 01:41:59 2025 +0100

         Update hello.txt, feature-2

         commit 7554fc2cdc83000e649fc6400edd82d18331cc7d (feature-1)
         Author: Timoumi_Mahmoud <mahmoud.timoumi@espirt.tn>
         Date:   Sun Jun 22 01:39:56 2025 +0100

             Updated hello.txt, feature-1
             
git cat-file -p  4c44
tree 02eb4bcadcd6da5726a7cb20fa0d316c8ed1b016
parent cef853562d55494e17b70c3379fe4977018a3bc7
parent 7554fc2cdc83000e649fc6400edd82d18331cc7d
author Timoumi_Mahmoud <mahmoud.timoumi@espirt.tn> 1750553003 +0100
committer Timoumi_Mahmoud <mahmoud.timoumi@espirt.tn> 1750553003 +0100

Merge branch 'feature-1' into feature-2
```


Fast Forward: happens when there are no commits on the base branch that occurred after the feature branch was created 
```bash
 git merge feature-2
 Updating 4eb494e..4c44769
 Fast-forward
  hello.txt | 2 ++
   1 file changed, 2 insertions(+)

   lenovo@DESKTOP-U5PPG5V MINGW64 ~/Desktop/GIT/ww (master)
$ git log ## Nothing changes in the commit history
   commit 4c44769cd3cadf495902fac215e994cf7b590ce9 (HEAD -> master, feature-2)
   Merge: cef8535 7554fc2
   Author: Timoumi_Mahmoud <mahmoud.timoumi@espirt.tn>
   Date:   Sun Jun 22 01:43:23 2025 +0100

       Merge branch 'feature-1' into feature-2

       commit cef853562d55494e17b70c3379fe4977018a3bc7
       Author: Timoumi_Mahmoud <mahmoud.timoumi@espirt.tn>
       Date:   Sun Jun 22 01:41:59 2025 +0100

           Update hello.txt, feature-2

           commit 7554fc2cdc83000e649fc6400edd82d18331cc7d (feature-1)
           Author: Timoumi_Mahmoud <mahmoud.timoumi@espirt.tn>
           Date:   Sun Jun 22 01:39:56 2025 +0100

               Updated hello.txt, feature-1

```
- Git REUSE RECORDED RESOLUTION git rerere (advance) : git saves how I resolved a conflict , so in next conflict reuse the same resolutionuseful for long lived feature branch(like a refactor, rebasing)

```bash
git config rerere.enable true # --global flag to enable for all projects

```

### HISTORY & DIFFS
Good commit :
```
commit message: in futre tense. 'Fix' vs 'Fixed'
Short subject, followeed by a blank line.

A description of the current behavior, a short summary of why the fix is needed. Mention side effects. (The description is broken into 72 char lines)

```

- Git log 
```bash
git log --since="yesterday"
git log --since="2 weeks ago"

git log --name-status --follow -- <file> ## log files that have been moved or renamer
```


### FIXING MISTAKES
#### CHECKOUT

What happens when I git checkout a branch:
1. Change HEAD to point to the new branch
2. copy the commit snapshot to the staging area
3. Update the working area with the branch contents

What happens when I git checkout a file:
Overwrite the wokring area file with staging area version from the last commit

What happens when I git checkout a <commit>--file:
1. Update the staging area to match the commit 
2. Update the working area to match the staging area
```bash
git checkout <commit> -- <file_path>
```

#### Git clean
Remove untracked files the working tree
```bash
git clean ## -f:force, --dry-run: don't actually remove anything, jsut show me what would be done, -d: directory

git clean --dry-run
git clean -d --dry-run
git clean -d -f
```

#### Git RESET
Git Reset is a command that performs different actions depending on the arguments:
    with a path
    without a path
    by default, git performs a git reset -mixed

for commits:
    movies the HEAD pointer, optionally modifies files
for file paths:
    does not move the HEAD pointer, modifies files
    

Three options:
1. Soft: not frequently 
```bash
git reset --sort HEAD~  #moves the HEAD pointer
```
2. Hard
```bash
git reset  --hard HEAD~ 
# moves the HEAD pointer
# Copy files from the Head to the staging Area
# Copy files from the Head to the working Area
```
3. mixed 

```bash
git reset --mixed HEAD~ r
git reset  HEAD~ 
# moves the HEAD pointer
# Copy files from the Head to the staging Area

```

1. Move HEAD and current branch
2. Reset the staging area
3. Reset the working area

--soft =(1)
--mixed= (1) & (2) (default)
--hard= (1) & (2) & (3)

```bash
git reset  <commit> --<file> ## (2)
git reset <file> ## (2)
```
undo a git reset with ORIG_HEAD
```bash
git reset ORIG_HEAD
# () -->(HEAD)---->(ORIG_HEAD)
```

#### GIT REVERT 
- The safe reset
- git revert create a new commit that introduces the opposite changes from hte specified commit.
- The original commit stays in the repository

usefull when:
    if I am undoing a commit that has already been shared
    Revert does not change history


#### GIT AMEND

- make changes to previous commit message (but the sha1 will change) because commits can't be edited so it creates a new commit.

```bash
lenovo@DESKTOP-U5PPG5V MINGW64 ~/Desktop/GIT/lab4 (master)
$ git log --oneline
b9811fa (HEAD -> master) add main test bla
3413eff Revert "Delete main.go"
c226f51 Delete main.go

git commit --amend -m "add main_test.go file"
 [master e89fc28] add main_test.go file
  Date: Mon Jun 23 16:45:37 2025 +0100
   1 file changed, 7 insertions(+)
    create mode 100644 main_test.go

lenovo@DESKTOP-U5PPG5V MINGW64 ~/Desktop/GIT/lab4 (master)
$ git log --oneline
    e89fc28 (HEAD -> master) add main_test.go file
    3413eff Revert "Delete main.go"
    c226f51 Delete main.go


```
## REBASE
























### GitHub
git : open source version control software

GitHub: Repository hosting, browse code, Issues, Pull Requests, Forks

* Remote : is a git repository stored elsewhere (web, github ...etc)
* origin : is the default name git gives to the server I cloned from 

* Cloning a remote repository from URL, will fetch the whole repository, and make a local copy in my .git folder.

* I may have different privilege for a remote : Read/Write, Read Only

* Fork: is a copy for a repository that's stored in my Github account
* I can clone my fork to my local computer

* Pull request
* Upstream: is the base repository that I created a fork from, by adding an upstream remote, I can pull down changes that have been added to the original repo.
git remote add upstream https://github.com/Timoumi-mahmoud


* Tracking branch : to tie it to an upstream branch, I can use push/pull with no arguments
to checkout a remote branch with tracking:
    git checkout -t origin/feature
to tell Git which branch to track the first time I push:
    git psuh -u origin feature

* Fetch: keeping my local repository up to date with a remote, It pull down all the changes that happened on the server but it doesn't change my local repository.

* Pulling : will pull down the changes from the remote repository to my local repository, and merging them with a local branch
git pull = git fetch && git merge
if changes happened upstream, git will create a merge commit , otherwise , it will fast-forward. (I can use git rebase here)

* git pull --rebase will fetch, update my local branch to copy the upstream branch, then replay any commits I made via rebase ==> Bonus: not unsightly merge commits
    git pull origin/master --rebase

* TAGS
git doesn't automatically push local tags to a remote repository.
to push tags:
    git push <tagname>
    git push --tags

* Push: sends my changes to the remote repository, git only allows my to push if my changes won't cause a conflict.



* Continuous Integration
merging smaller commits frequently, instead of waiting until a project is done and doing one big merge.
this means features can be released quicker, CI only works when there are tests that ensure that new commits didn't break the build.
It's even possible to perform a deployment at the end of a CI build

(COMMIT)---->(TEST)---->(MERGE)---->(DEPLOY)









git cherry -v # see commits which haven't been pushed upstream yet
git branch -vv 


