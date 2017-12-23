AFTO 1 "2017" UNIX "AFTO Manual"
=======================================

NAME
----

`afto` - An automated command-line cydia repo generator/builder and server.

SYNOPSIS
--------

`afto` `--version`

`afto` command [`options`]

`afto` `-h` | `--help`

DESCRIPTION
-----------

`afto` is an automated command-line cydia repo generator/builder and server for Cydia tweak developers.

_The name 'afto' comes from the word 'automatic' in greek._

+ Automatic Cydia repo generation.
+ Automatic Cydia repo updating.
+ Cydia repo server testing.
+ Many more.

The only thing you need on your system is:

`dpkg`

*You also need at least 1 or more .deb files. So that you can test or host your repo.*

USAGE
-------

To generate a new Cydia repo, use the following command replacing 'example_repo' with your desired repo name:

`afto new example_repo`

To serve your new Cydia repo, use the following command replacing 'example_repo' with your desired repo name:

`afto serve example_repo`

If you choose to add another .deb file to your repo, you can instruct `afto` to automatically regenerate your repo with the `-w` option.

`afto serve -w example_repo`

You can visit `http://127.0.0.1:2468` to view your newly generated repo, and you can also put this in Cydia to view this in the Cydia iOS app.

COMMANDS
-------

`new`: New repository. (Use "." for the same directory)

`serve`: Serve the directory and optionally watch the repo with `-w`.

`update`: Update the deb file in the repo with `-r`.
   
    
OPTIONS
-------

`-f` | `--file`
  A file normally a *.deb file.

`-w` | `--watch`
  Watch the repo.
  
`-s` | `--sign`
  Sign the repo.

`-c` | `--control` 
  Specify control file to use.

`-p` | `--port` 
  Specify port number for `afto`.
  
`--h` | `--help`
  Help menu.
  
`--version`
  Show version.

BUGS
----

If you find any bugs please file an issue on GitHub or contact me. See the `AUTHOR` section.

AUTHOR
------

Wesley Hill <wesley@hakobaito.co.uk>


