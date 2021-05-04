# git-user-switch
command line tool for git user switching


## install
make


## Display current git user in bash prompt.
Add below code to .bachrc

```
export BASE_PS1=${PS1}
GETGITUSER () {
        export GITUSER=$(git-user-switch get)
        PS1="(${GITUSER})$BASE_PS1"
}
PROMPT_COMMAND=GETGITUSER
```

`(GituserprofileNickname)user@notepc:~/develop/git-user-switch$ `
