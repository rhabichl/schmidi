# Schmidi
Simple helper tool to generate Java Spring Jpa classes with a cli tool.

## Install
Clone the git repo with:
```bash
git clone https://github.com/rhabichl/schmidi.git
```
Make sure you have golang installed. Then compile the programm with:
```bash 
go build .
```
## Use
To use the tool simply execute the model subcommand like this:
```bash
./schmidi model -p <path to your Jpa project>
```

## TODO
* fix bugs and test more on windows 
* add CRUD REST-API handler for all entities
* use other datatypes for repositories