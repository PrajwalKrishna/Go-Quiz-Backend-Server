# Go-Quiz-Portal-Backend

*Coded by:*
**Prajwal Krishna**

This **README** file contains :
 1. Information About the Quiz-App
 2. How to run the backend server
 3. Controls for game play
 4. File structure
 5. List of APIs

----------


About The Quiz Portal
-------------

This a backend written in GoLang using **Gin** as microframework **Gorm** as ORM and **sqlite3** as database.

This server is caters to restful apis of a quiz app whose frontend is totally decoupled with it.

List of apis and a breif description about those are written in API section of README.md.


----------

## Running the program

- Install GoLang in your machine.
- Inside GoLang src folder clone this repo
         git clone https://github.com/PK1210/Go-Quiz-Backend-Server
- Go inside the folder
         cd Go-Quiz-Backend-Server
- Running the program is easy
         go run
- This starts a go-server at **localhost:8080**
- The database is intialized with some values to make it empty run following commands before go run
         rm gorm.db
         touch gorm.db

   This starts a backend server for quiz app on localhost:8080 which can cater to any service making a api request to it.

    For a sample frontend for this quiz-portal click [here](https://github.com/PK1210/React-Quiz-App-Frontend)
----------

_______________

#### Prajwal Krishna Maitin
