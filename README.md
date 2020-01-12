# nephew-pipelines

`nephew-pipelines` is the repository that will house all the pipelines and the endpoints necessary to commmunicate with the database.

### Setup

First, download the latest version of Go found [here](https://golang.org/dl/) or download it through `brew` which might be easier. 
Download gcc5 which is necessary for [mongo-client drivers](https://github.com/mongodb/mongo-go-driver).

Then, you can do a `go get -u github.com/clarencejychan/nephew-pipeline` which will get all the required dependencies for the project.

If you need to install any other dependencies as you're working, simply go to the root folder and do a `go get -u <package name>`.

Finally, run `go run main.go` to start up the server. The framework does not come with hot-reload out of the box, so that's something we can look into.

### DB Setup
As of now, we are planning on using MongoDB.

If you're using WSL (Windows Susbsystem for Linux) [follow these instructions to download MongoDB](https://github.com/michaeltreat/Windows-Subsystem-For-Linux-Setup-Guide/blob/master/readmes/installs/MongoDB.md).

If you're using a Mac [follow these ones](https://docs.mongodb.com/manual/installation/)

Remember to make sure you have the folder `/data/db` created.

#### Dumping Data from the DB
#### Importing Data into the DB

