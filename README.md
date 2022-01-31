# URL Shortener
The URL shortener is a backend service written in Go that provides the functionality to generate a shortened version of a url. Often employee may have a link to share; however, is to long to fit in their document, text, web form, etc. Via the URL shortener they can pass a "long url" as a parameter and retrive a shorter, fixed length URL for their purposes. Employee's may also delete short urls and view statistics on usage. 

# Solution Details
The following are inferred from the Overview and Expectation of the Cloudflare API Interview Coding Project.
## Requirements
- [X] **R1** Short url data model defined as: (**R1.1**) Has one long url (**R1.2**) No duplicates are allowed (**R1.3**) Short links can expire or remain indefinitely.  
- [X] **R2** Supports generating a short url from a long url. 
- [X] **R3** Supports redirecting a short url to a long url. 
- [X] **R4** Supports listing the number of times a short url has been accessed in the last 24 hours, past week, and all time. 
- [X] **R5** Supports data persistence (must survive computer restarts) 
- [X] **R6** Supports short links can be deleted
- [X] **R7** Project should be able to be runnable locally with some simple instructions
- [X] **R8** Project's documentation should include build and deploy instruction
- [X] **R9** Tests should be provided and able to be executed locally or within a test environment.

## Implementation
The following is a summary of the implementation of each of the items outlined in the `Requirements` section above:
- **R1.1**: In data model, entity for url defines column for LongUrl 
- **R1.2**: In data model, entity for url defines `short_url`, and `long_url` as having a UNIQUE constraint.
- **R1.3**: In data model, entity for url defines an additional field, `expiration_dt` allowing a date of expiration to be specified for future functionality.
- **R2**: Solution supports a creation endpoint, returning a short url.
- **R3**: Solution supports a get endpoint, performing a http redirect to the associated long url, if existing.
- **R4**: Solution supports a get endpoint, retrieving information from a view aggregating rows from a history table. When short urls are accessed the system records a row in the history table.
- **R5**: Solution persists data to a sqlite3 database (file remains on file system on an after start).
- **R6**: Solution supports a delete endpoint, deleting short url associated with the url id passed via path parameter.
- **R7**: Solution provides a Docker file for building and running a docker container locally on desktop.
- **R8**: Solution README.MD provides build information for how to create and run docker container hosting application.
- **R9**: Solution README.MD provides information on how to run provided unit tests.

## Assumptions
As defined by Overview and Expectations of Cloudflare API Interview Coding Project
1. Choose whichever languages and frameworks you would like for this exercise
2. Utilize any folder/file hierarchy that makes sense to you any folder/file hierarchy that makes sense to you. 
3. You can implement it in any way you see fit, including whatever backend functionality that you think makes sense

Confirmed by Rupalim on Jan 25, 2022

4. Short links can expire, but data model just needs to support this feature. No expiration functionality required at this time.
5. Front-end not necessary.

Un-verified

6. If the URL is already "short", it may not be a suitable use case for this solution. For example, "https://a.com" probably would not need to have it's URL shortened to most uses. 
7. Code coverage target is undefined
8. Security related requirements not defined at this time. See future considerations no. 7 for details 


# API
## Redirect short url to a long url
The short url behind the scenes is a hash that is passed to the base url via a **path parameter**. If the hash for the short url is found, the user will be redirected to the long url on record. A entry will also be made to the history table to record the shorl url being used. If no match is found, the API will return error code 400 with an error specifying that the url has not been found. 

### GET /{shorturl}
**Example Response**

_Redirected to_ `LongUrl` _on record_

**Example Error**
```
{
    "Error": "M003: https://localhost:9000/nonexisting does not exist, double-check to see if the URL is correct."
}
```
## Create a short url
To create a short url, you pass a JSON object containing a `LongUrl` and optionally `ExpirationDt`. If successful, a object containing `UrlId`, `LongUrl`, `ShortlUrl`, and `ExpirationDt` will be returned. If the `LongUrl` is null, empty or invalid - you will be send back an error response with a corresponding validation error message. Additionally, if a `ExpirationDt` is passed but has an invalid format, you will also receive a validation response. 
### POST /url/create
**Example Body**
```
{
    "LongUrl": "https://hawaiinewsnow.com",
    "ExpirationDt": "2023-09-01 12:00:00"
}
```
**Example Response**
```
{
    "UrlID": 4,
    "LongUrl": "https://hawaiinewsnow.com",
    "ShortUrl": "http://localhost:9000/4459c2ab",
    "ExpirationDt": "2023-09-01 12:00:00"
}
```
**Example Error(s)**
```
{
    "Error": "M001: Long Url is required."
}
```
```
{
    "Error": "M006: htt////hawaiinewsnow.com is not a valid url."
}
```
```
{
    "Error": "M007: Jan 21 21 120500 is not a validate datetime format. Exepecting YYYY-mm-dd HH:MM:SS"
}
```

## Get statistics for a short url
Statistics for when a short url is accessed within the list 24-hours, 7-days, and all time are recorded. To get the statistics for a short url, pass the `UrlId` of the short url as a path parameter. If the short url exists, you will be returned a JSON object containing statistics for the short url. If the short url does not exist, a error message will be returned. 
### GET /statistics/{urlid}
**Example Response**
```
{
    "UrlID": 1,
    "TwentyFourHours": 0,
    "LastSevenDays": 0,
    "AllTime": 0
}
```
**Example Error**
```
{
    "Error": "M005: Short Url: 24 does not exist."
}
```

## Delete a short url
To delete a url, pass the `UrlId` of the short url as a path parameter. When a short url is deleted, any associated rows in the history table are also deleted.
### POST /url/delete/{urlid}

# Build Information
🚨 These instuctions are based on a OSX workstation using [homebrew](https://brew.sh/). 

Clone the repo at your preferred workspace location

`git clone https://github.com/analogefficiency/cf-proposal.git`

This project will be using `Docker` to build and run the application. Docker can be installed via the following command:

`brew install --cask docker`

From the root directory of the project run the following command to create the `Docker` image to run:

`docker build --tag cf-proposal .`

In order for sqlite3 databases to persist after shutdown of a docker container, we'll need to create a docker volume:

`docker volume create cf-proposal-db`

Start the application container via the following command:

`docker run --publish 9000:9000 -v cf-proposal-db:/app/sqlite cf-proposal`

Press `Ctrl-C` to stop the container. 

# Testing
The following describes how to run the unit tests for this project.
1. If not already, install go using the documentation projected [here](https://go.dev/doc/install), and clone the repository into your preferred workspace.
2. From the project root, run the following command: `go test ./...`
3. Alternatively, for any package with unit test file `*_test.go`, from its parent directory run the command: `go test` 

# Design Decisions

## Application Layering
The application is split into 4 layers: 
- Controller: Route management and entry point to endpoints.
- Service: Where code related to domain knowledge is performed.
- Repo: Provides an interface to the actions in the datastore. 
- Datastore: Directly interacting with the database to manipulate the data model. Some of this code was autogenerated via sqlc.

Validation is handled via the DTO passed into the endpoints. 

## Data Model
[Data Dictionary](https://docs.google.com/spreadsheets/d/1lYeBe29FgTnOEaFF-xYTOj10ipwja7ZW6d8-eWqQOho/edit?usp=sharing)

The above linked Google Sheet is a depiction of the data model current in use. As summary of the more notable decisions is as follows:
- **Date fields using TEXT typing**: Sqlite does not have a native date type but the built-in date functions can infer date so long as the string or integer is in the correct format. 
- **Statistics View**: To reduce the code necessary at the service layer, the statistics view aggregates history table rows by defined date conditions. See `statistics.sql` for details on how the statistics are calculated. The view is referenced by datastore layer, accessible to the service layer via the repository.
- **Expiration Date**: To future proof short urls supporting expiration, an additional field is added to capture a expiration date if desired. 

## Stack
This section desribes the rationale behind why specific stack elements were used.
### Golang
Originally, I had intended to do this project in Java using Spring Boot; however, my personal machine was already set up for Go development. Weighing the options I felt that being able to start development quicker would be more beneficial and specific implementation info could be looked up. 

### Sqlite3
Provides a relational db without setup overhead, allowing focus on the development of MVP. In addition:
- Allows persistence of data after shutdown since can remain on the filesystem.
- Quick tear of DB via file deletion.
- Accessible via third party client for direct manipulation of sql and sandboxing.    

### [sqlc](https://github.com/kyleconroy/sqlc)
A Golang code generator which compiles sql. Leveraged this to write out the data model and queries via sql, and save time on writing the datastore layer code. 

### Docker
Leveraging Docker to simplify build environment for those checking out the project.

### Postman
For testing endpoints.

### [DB Browser for SQL Lite](https://sqlitebrowser.org/):
Provides a GUI interface to the sqlite database used for development locally. Primary used for debugging and verfication. 

## Future Condiderations
The following is a list of items I noted as area's for improvement, not within the scope of the solution requirements, but merit future consideration.
1. As part of the validation, confirming that the `LongUrl` being passed actually returns a response code of 200 - to prevent users from creating short urls to "dead" links. As a follow on, a scheduled service (yearly) that validates the `LongUrl` still works, potentially alerting the owner of the url (not implemented)
2. The API currently returns a error message for short urls that to not reference a long url. The redirect could go to a 404 page outlining potential reasons for error.
3. Documentation for the API is currently typed out in this README.md file. Preferably, the documentation could be managed by Swagger or Twilio.
4. I feel like a potential use case for getting the statistics on the endpoints would be sending the logs to a third-party service for aggregation and reporting instead, creating an interface collecting relevant information for end-user consumption. 
5. URL length checking, comparing the submitted `LongUrl` against what the resulting `ShortUrl` would be. I'd argue they should be able to create the short url regardless, but have a front-end system warn them that the short url would be longer. 
6. Currently passing the database connection as a global variable; I went with this approach to move forward with developing, but i'd like to review other patterns for what would be more appropirate.
7. Current docker file implements this service over HTTP instead of HTTPS. Short URL data is directly viewable to those with database access. Should the administrators of the service be able to view the contents of a long url? Given the scope of the proposal, this is more engineering than is required, but is a potential question to product owners when addessing how this data is managed. 
8. The code for "seeding" the database relies on relative paths, works for now but won't scale well. 
9. Code grabbing port information should pull from os environment variables instead of being hard coded. 
10. 
