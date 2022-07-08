# Website Kantin Kejujuran
The website is deployed on [https://kantin-kejujuran-dimasfm.herokuapp.com/](https://kantin-kejujuran-dimasfm.herokuapp.com/)
## This is the backend repository of the project
Frontend code can be found [here](https://github.com/maspaitujaki/kantinkejujuran-compfest)
### Get Started
Use website on URL address stated above or follow the instruction below to run it locally
1. Fork this repository and clone to your local machine
2. Open terminal on the root folder
3. Make sure you have [Golang](https://go.dev/) installed
4. Run the program with following command
```
go run main.go
```
### Technology
- [Golang](https://nextjs.org/)
- PostgreSQL
### Notes
By default, it use a PostgreSQL database running on Heroku. The connection made with an URL provided as environment variable. You can use another postgres database you want by changing the url on /.env file with the following format `postgres://YourUserName:YourPassword@YourHostname:5432/YourDatabaseName`.
