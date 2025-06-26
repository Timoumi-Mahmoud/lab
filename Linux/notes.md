###
# Intro
The Unix os is a collection of programs, each with a special role:
    * kernel: medial access between user programs and system resources manage process,CPU scheduling, memory, storage, IO to computer hardware ...., programs request resources by making a syscall.

    * shell: is a computer program that can execute other programs from a text-based interface. Where I interact completely from my cli and text output.
    * utilities

-----------------------------------------------------------------------------------------
## Shortcuts
CTRL + A – takes you to the beginning of the line
CTRL + E – takes you to the end of the line
CTRL + K – "yank" everything after the cursor
CTRL + U – "yank" everything before the cursor
CTRL + Y - "paste" (paste in quotes because it doesn't actually go into your system clipboard) everything you yanked
CTRL + L - clear the screen
CTRL + R – reverse search through history

ls --ignore=file*
Signals: is a notification that I send to a program
CTRL + C - SIGINT interupt
SIGTERM
SIGKILL
SIGALARL
see kill -l

-----------------------------------------------------------------------------------------
## Tar

tape archive, not compresed
```bash
mkdir folder1
touch file1.txt file2.txt folder1/file3.txt
tar -cf archive.tar file1.txt file2.txt folder1
ls -lsah
```
-z flag it will automatically compress it
```bash
tar -zcf archive.tar.gz file1.txt file2.txt folder1
ls -lsah
```

```bash
mkdir destination-folder
tar -xzf archive.tar.gz -C destination-folder/
```



-----------------------------------------------------------------------------------------
## Wilde cards
*
? 
echo {a..z} # prints a to z
echo {z..a} # reverse order
echo {0..100..2} # prints every other (aka even) number from 0 to 100
echo {100..0..5} # prints every 5th number in reverse order from 100 to 0
echo {a..z}{1..5} # prints out a1 a2 a3 a4 a5 b1 b2 b3 <etc>



-----------------------------------------------------------------------------------------
## users
sudo usermod -aG sudo <user> 
sudo chow <group>:<user>
chmod (4:r, w:2, 1:x)


-----------------------------------------------------------------------------------------
## ENV
```bash
printenv
GREETING=hello
echo $GREETING # hello
```

/etc/environment : modify every user's environment
/etc/profile
/etc/bashrc

export VARIABLE=value

.bashrc_profile: is only run on login shells, run once for each time I log into my computer.
.bashrc: is run on every nonlogin shell, run on every ttab of bash I start up
!! Forget about .bashrc_profile and put this into it
```vim
if [ -f ~/.bashrc ]; then
    source ~/.bashrc
fi
```

-----------------------------------------------------------------------------------------

## Process
A process is any sort of command that's currently running.
```bash
sleep 100 &
ps # find the sleep pid
kill -s SIGKILL <the pid from above>
ps # notice the process has been killed
or kill -o <pid>
```
A Process can either run in the foreground or the backgound.
```bash
sleep 100
# hit CTRL + Z
jobs # notice process is stopped
bg 1 # it's the first and only item in the list, the number refers to that
jobs # notice process is running
fg 1 # reattch to the process
```

-----------------------------------------------------------------------------------------

## Exit code
```bash
date # show current date, runs successfully
echo $? # $? corresponds to the last exit code, in this case 0
yes # hit CTRL+C to stop it
echo $? # you stopped it so it exited with a non-zero code, 130

```
0: means it was successful. Anything other than 0 means it failed
1: a good general catch-all "there was an error"
2: a bash internal error, meaning you or the program tried to use bash in an incorrect way
126: Either you don't have permission or the file isn't executable
127: Command not found
128: The exit command itself had a problem, usually that you provided a non-integer exit code to it
130: You ended the program with CTRL+C
137: You ended the program with SIGKILL
255: Out-of-bounds, you tried to exit with a code larger than 255


Run if first one succeeds
    true && echo hi
Run if first one fails
    false || echo hi 
Always Run 
    false; true; echo hey


Subcommands
echo I think $(whoami) is a genies
The $() allows me to put bash commands inside of it that then I can use that output as part of an input to another command
example:
    echo $(date +%x) – $(uptime) >> log.txt


-----------------------------------------------------------------------------------------

## Apt
APT (advanced packaging tool) is based on dpkg and available to all in the Debian family of distros

sudo apt autoremove # will remove unused dependencies
sudo apt update # updates the list of available packages apt uses
apt list # everything installed
apt list --upgradable # everything with an update available
sudo apt upgrade # updates all your packages to their latest available versions
sudo apt full-upgrade # basically autoremove and upgrade together



-----------------------------------------------------------------------------------------
## Curl
```bash
curl http://localhost:8000 -o output.txt
curl -I http://localhost:8000  # or curl --head

curl -X POST -d'{}' http://localhost:8000
curl -X PUT http://localhost:8000
curl -X PATCH http://localhost:8000
curl -X DELETE http://localhost:8000
curl -X OPTIONS http://localhost:8000

curl -b "name=mahmoud" http://localhost:8000 ## send cookies if I have many cookies to use , I can use Cookie jar file -c <file-contains-cookie>

curl -L http://bit.ly/linux-cli ## Follow redirect
curl -H "'accept-language: en-US'" -H "Authorization: Bearer 1234" http://localhost:800

```

In Edge, Chrome, Firefox I can copy url as curl: network -> copy-> copy as cURL


-----------------------------------------------------------------------------------------

## Script
```vim
test -z ""
echo $? # 0, this is true
test -z "lol"
echo $? # 1, this is false

test 15 -eq 15 # 0
test bara = bara # 0
test bara != bara # 1
test 15 -gt 10 # 0 gt means greater than
test 15 -le 10 # 1 le means less than or equal to
test -e ~/some-file.txt # tests to see if a file exists
test -w ~/some-file.txt # tests to see if a file exists and you can write to it
```


```vim
if [ $1 -gt 10 ]; then
echo "greater than 10"
elif [ $1 -lt 10 ]; then
echo "less than 10"
else
echo "equals 10"
fi

case $1 in
"smile")
echo ":)"
;;
"sad")
echo ":("
;;
"laugh")
echo ":D"
;;
"sword")
echo "o()xxx[{::::::::::::::>"
;;
"surprise")
echo "O_O"
;;
*)
echo "I don't know that one yet!"
;;
esac

```

-----------------------------------------------------------------------------------------

## CRON
```bash
crontab -e
crontab -l

<minutes> <hours> <day of the month> <month> <day of the week>
```


-----------------------------------------------------------------------------------------

## SSH
secure shell, connect from one computer to another

1. 
```bash
ssh-keygen -t rsa
```

ssh key cryptographic key , it create a private and public version, the private version you never reveal to anyone, And I give the public version.
The public and private version let servers to do a handshake with each other without revealing there private key(only to verify that each party has the public key that they have said they have)


cp id-pub --> authorized key

-----------------------------------------------------------------------------------------

## Regex
is a pattern-matching language
System tools with regex
    grep
    sed
    perl
    vim
    less

```bash 
echo cat cabbage | sed 's/a/@/g'
c@t c@bbage
\:$echo cat cabbage | sed 's/a/*/g'
c*t c*bb*ge

\:$echo -e 'one\ntwo\nthree' | grep ^t
two
three

```

Flags:
/PATTERN/FLAGS
s/PATTERN/REP/FLAGS
*  i - case insensitive
*  g - match all occurences (global)
*  m - treat string as muliple lines
*  s - treat string as single line



* `.` matches any character
* `?` - zero or one time 
* `*` - zero or many times
* `+` - one or more times

* `[]` - character class
* `^` - anchor at the beginning
* `$` - anchor at the end

* `(a|b)` - match a or b
* `()` - capture group
* `(?:)` -non capture group
* `\d` - digit `[0-9]`
* `\D` - non-digit `[^0-9]`
* `\w` - word `[A-Za-z0-9]`
* `\W` - non-word  character`[^A-Za-z0-9]`
* `\s` - whitespace `[\t\r\n\f]`
* `\S` - non-whitespace `[^\t\r\n\f]`

```bash
echo hello beep boop | sed 's/b..p/XXXXXX/g'
echo dog and doge | sed 's/doge\?/DOGGY/g'
echo beep bp beeeeeeeep | sed 's/be*p/XXXXXXXXX/g'
echo beep bp beeeeeeeep | sed 's/be+p/XXXXXXXXX/g'

echo 'hey coooooool whatever cool' | sed -E 's/o+/@@/g'
// character class
echo 'dog days cats' | sed -E 's/d[ao]/DA/g'
echo 'dog days cats' | sed -E 's/d[a|o]/DA/g'
echo 'dooooog days cats' | sed -E 's/d[a|o]+/DA/g'
echo 'dooooooooog dxxx days cats ds' | sed -E 's/d[^aeiou]+/DA/g'

echo 'hello12345 what9999999 66666' | sed -E 's/[0-9]+//g'
echo 'hello12345 what9999999 66666' | sed -E 's/[^0-9]+//g'
echo 'hello12345 what9999999 66666' | sed -E 's/[A-Za-z]+//g'

echo '_hello12345 WHat999_9999 66666' | sed -E 's/[A-Za-z_]+//g'

echo 'the the whatever cats dogs dogs etc' | sed -E 's/(\w+) \1/\1/g'
sed 's/cool/bad/g' -i  beep.txt 
```

Handy reference:  perldoc perlreref
:set scrolloff=8
:set number  :set relativenumber
:set rnu

-----------------------------------------------


```bash
ls -1

echo mv moby-dick.txt{_,}
\:$echo cat b{ee,oo}p.txt
cat beep.txt boop.txt
\:$cat b{ee,oo}p.txt
BEEP !!!!
BOOP !!!!

\:$cat b*p.txt

:$(echo at the beginning; cat butterflies.txt) > butterflies.txt_ && mv butterflies.txt{_,}
\:$
\:$
\:$
\:$cat butterflies.txt
at the beginning
This is a butterFlies file
Blabla blabla adding

\:$date +'%M, %H, %Y'
24, 16, 2025
\:$date +'%F'
2025-06-24

\:$date +'%F %T'
head -c 10 # get first 10 char
\:$fold -w5 -s butterflies.txt
echo $((5+4))
TZ=America/Anchorage date
Tue Jun 24 08:15:47 AKDT 2025
TZ=America/Los_Angeles date
Tue Jun 24 09:16:24 PDT 2025

export
/:$while true ; do date ; sleep 1; done
/:$while true ; do node server.js ; done
\:$for x in {1..10};do echo $x; done
\:$while test -f cool.txt ; do date; sleep 1; done
\:$if test -f cool.txt; then  echo true; else echo false; fi
echo -e "this make a color e[32mgreen"
see awesome bash on github

```

## Vim 
```vim 
dt  <char>

:%s/dog/cat/g
V + >    : intend

:r fileToPast
:r! date

set -o vi


:set noic # ignore case
:set hls # highlight search after doing / (all result) also with :set ic for case
:nohls   
I can do /SIG.* to highlight every word that begin with SIG


':s/foo/baz<CR>' search and replace

dt (
fmt.Println("vim-go")  ==> ("vim-go")


V + "by   // yank in b register -> :reg then I can do "bp


f +  <firstCharacter>
F +  <firstCharacter> reverse

t +  <firstCharacter>
T +  <firstCharacter> reverse

{ and } vertical movement

'[m and ]m' move my curosor to { and }
'%' go the end of ( or { or [

```

https://git.github.com/substack/7745bb6ff9ad58d4805d

echo -e "this make a color e[32mgreen"
see awesome bash on github


## Script
```vim
test -z ""
echo $? # 0, this is true
test -z "lol"
echo $? # 1, this is false

test 15 -eq 15 # 0
test bara = bara # 0
test bara != bara # 1
test 15 -gt 10 # 0 gt means greater than
test 15 -le 10 # 1 le means less than or equal to
test -e ~/some-file.txt # tests to see if a file exists
test -w ~/some-file.txt # tests to see if a file exists and you can write to it
```


```vim
if [ $1 -gt 10 ]; then
echo "greater than 10"
elif [ $1 -lt 10 ]; then
echo "less than 10"
else
echo "equals 10"
fi

case $1 in
"smile")
echo ":)"
;;
"sad")
echo ":("
;;
"laugh")
echo ":D"
;;
"sword")
echo "o()xxx[{::::::::::::::>"
;;
"surprise")
echo "O_O"
;;
*)
echo "I don't know that one yet!"
;;
esac


## $1 $2 $3 positional arguments
## $* $@ all arguments
## chmod +x myfile
## cp script.sh /usr/local/bin
```

## CRON
```bash
crontab -e
crontab -l

<minutes> <hours> <day of the month> <month> <day of the week>
```


---------------------------------------------------------
## SSH
secure shell, connect from one computer to another

1. 
```bash
ssh-keygen -t rsa
```

ssh key cryptogrphic key , it create a private and public version, the private version you never revial to anyone, And I give the public version.
The public and private version let servers to do a handshake with each other without revialing there private key(only to verify that each party has the public key that they have said they have)


cp id-pub --> authorized key

---------------------------------------------------------
## AWK && SED

awk -F : '{print }' /etc/passwd
awk 'BEGIN{FS=:; OFS=_} {print ,,}' /etc/passwd
awk -F / '/^\// {print }' /etc/shells| uniq 
awk 'length() >15' /etc/shells 
df | awk '/\/usr\/lib/ {print t +}'
awk '{print }' file.txt
sed -i 's/blue/yellow/'  file.txt 

---------------------------------------------------------
## TMUX

ctrl + b  + d
tmux list-sessions
tmux attach
ctrl + b  + command
ctrl + b + % ;
ctrl+ b + doubleQuote 
ctrl+ b + up|down|right|left arrow 
ctrl+ b + x 
tmux new-window || ctrl+b+  + c 
ctrl + b + p|n 
ctrl + b + & 
tmux rename-window CustomName
tmux attach -t 1
tmux rename-session FirstSession || ctrl + b + $ 
ctrl + b + s 


