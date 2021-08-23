# Zone Identifier Remover

A simple command line utility that removes the `file:ZoneIdentifier` files that get created when copying files between Windows and Windows Subsystem Linux (WSL).

## Installation

**Using `go get`**

`$ go get github.com/OskarLiew/rm-zone-id`

**For Linux**

Download the latest release of [rm-zone-id](https://github.com/OskarLiew/rm-zone-id/releases) and make sure it is executable. Then either add it to your path variable or to a directory on your path, e.g.

`$ cp /path/to/rm-zone-id /usr/local/bin`

**For Windows**

TODO (You should be able to run the same executable as linux users)

## Usage

Give an arbitrary number of paths to recursively search through

`$ rm-zone-id PATH ...`
