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
- [ ] fix bugs and test more on windows 
- [ ] add CRUD REST-API handler for all entities
- [X] use other datatypes for repositories
- [X] create the folder before saving the files
- [ ] Automatic CICD tests with this tutorial:
- https://octopus.com/blog/githubactions-running-unit-tests
- [ ] Add tests for Java Spring with these tutorial:
- https://techwithmaddy.com/testing-with-spring-boot#heading-testing-the-rest-controller
- https://docs.spring.io/spring-security/site/docs/5.0.x/reference/html/test-mockmvc.html
