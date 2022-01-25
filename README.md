# cf-proposal
Sample URL Shortener Service

# Minimum Viable Product
- [X] POST Endpoint taking in a LONG URL returning a SHORT URL
- [X] GET Endpoint that redirects short urls to the requested long URL
- [X] GET Endpoint returning access statistics for each short URL (24 hours, past week, all time)
- [X] Data survives restarts
- [X] DELETE Endpoint allowing a short URL to be deleted
- [ ] Runnable locally with simple instructions
- [X] Documentation includes build and deploy instructions
- [ ] Tests able to be executed locally OR within a test environment 

# Build Information
🚨 These instuctions are based on a OSX workstation using [homebrew](https://brew.sh/). 

Clone the repo at your preferred workspace location

`git clone https://github.com/analogefficiency/cf-proposal.git`

This project will be using `Docker` to build and run the application. Docker can be installed via the following command:

`brew install --cask docker`

From the root directory of the project run the following command to create the `Docker` image to run:

`docker build --tag cf-proposal .`

Start the application container via the following command:

`docker run --publish 9000:9000 cf-proposal`

Press `Ctrl-C` to stop the container. 


# Assumptions

# Design Decisions
## Stack
**Golang: Jan 20, 2022**
- At first the intention was to go with a Springboot project; however, at present my current personal workstation isn't set up for it. While my experience with Go is quite basic, my personal machine is setup for it. Making the assumption that the time spent setting up / debugging is outweighed by leveraging existing capacity to start dev in go.

**Sqlite3: Jan 21, 2022**
- Relational db without much setup overhead, focus is MVP. 

**sqlc: Jan 22, 2022**
- Making up for lack of Go foundation using SQLC for boilerplate, db connections. Leverage existing SQL knowledge to focus time on weaker areas.

**Postman: Jan 22, 2022**
- Used for testing endpoints

**[DB Browser for SQL Lite](https://sqlitebrowser.org/): Jan 22, 2022**
- Used for confirming results on database.

## Data Model
[Data Dictionary](https://docs.google.com/spreadsheets/d/1lYeBe29FgTnOEaFF-xYTOj10ipwja7ZW6d8-eWqQOho/edit?usp=sharing)
```
01/22/2022 - Dates stored as INTEGER as Sqlite has no built-in date types. UNIX date seems the better route for 
             comparisons later.
```
```
01/23/2022 - Short URLS technically should be unique too. Updated dictionary.
```
## Trade-offs
1. A validator of some sort will be needed for the UrlDto payload; however, foregoing a validation layer until MVP is achived. Run risk of burning too many hours going into the different permutations for a URL and date. If time allows i'll circle around and add validation there. 
2. Passing database connection as a global variable. Would like to spend more time looking at implementing another pattern, but current implementation is working - will focus on remaining MVP items. 
