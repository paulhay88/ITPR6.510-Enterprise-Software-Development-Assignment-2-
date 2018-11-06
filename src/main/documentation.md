# Documentation

---

- Routing options. Available routes and accepted input such as POST.
## Login 
### /login POST request
1. To Access the database a user must log in with credentials to be able to return data after which all data is verified by the users Cookie.
We have chosen to use the negroni package to handle the authentication of the Cookies with a function called logged in to check if the Cookie holds the "authUser" credentials matching the current session.
2. Login requires a password and a username which is parsed and checked against the database.
3. Firstly checks if the users Cookie is currently active and returns a "User is already logged in." message if the Cookie is currently active.
4. The information from the body is then decoded into a logInUser variable.
5. A query is then sent containing that name and password
6. A variable of a new user is created and the query is then scanned for data
7. Error check the scan and output "Your username or password was incorrect. Please try again."
8. Else create and set Cookie
9. Then redirect to the home page of / with a valid Cookie in place.
10. Doing a GET request on this page will tell you to log in

### /signup POST request
1. Sign up is a simple POST that creates a User by creating space on the Database. 
2. Signup take 5 fields to fill up 
+ UserName
+ Name
+ Phone
+ Email
+ Password
3. This information is taken then decoded into a variable of a User
4. A query is then sent on the inserted UserName
5. A scan of the information is then done
6. Error check for existing name
7. Redirect to /signup with a message "Username already exists, please choose a different one."
8. Else create a User and insert it into the table

### /logout GET request
1. Firstly this checks the current Cookie and if the response is nil it returns "User is already logged out."
2. After that it creates a Cookie of the of the authUser with no variables and sets it as the current Cookie
3. After which will redirect you to the home page / 

## Meetings
### /allMeetings GET request
1. First step is a query is sent and stored in a variable of type Meetings of ALL meetings.
2. That list is then iterated through to check the data that is relevant to this User.
3. In the loop a new meeting is created to temporarily store relevant meeting data.
4. The entire list of meetings is scanned and checked for errors.
5. The application then checks to see if the current user is a participant in all meetings according to their name and userID through a SQL query and stores it as a variable from the database.
6. The application parses a password locally to keep it secure.
7. The application then then appends all relevant data to myMeetiings and then outputs it.

### / GET request
1. Home page after you login will always appear with the related meetings to your userID
2. The application first checks the Cookie for an authorized user then creates a meeting and user variable to use
3. The UserName is taken from the Cookie from the first variable split by ":"
4. A query is then constructed based on that name to get the userID
5. Error check
6. Another query is then launched based on the participant’s userID called participants
7. Error check the return
8. Participants is then scanned to an error check
9. A new query is launched based on the meetingID
10. That is also error checked 
11. After all that checking the application then appended the meetings.Meetings
12. Then Outputs the results

### /meeting/create POST request
1. The application creates an instance of a new meeting. 
2. Then decodes the body of the request into a variable.
3. Then inserts it as an Exec INSERT directly into the Database.
4. That is all error checked.

### /meeting/:id/edit PUT request
1. The application created an instance of a new meeting.
2. Then decodes that information into a variable which is error checked.
3. An Exec on the Database is then preformed which takes all the variables from the body of the request and inserts them into the Database WHERE the id is matching.

### /meeting/:id/delete DELETE request
1. This fires two Exec on the database.
2. The first one DLETEs FOM participants WHERE the meetingID matched the variable.
3. The second part DELETEs FROM meetings WHERE id matched the variable.
4. Both are error checked.

### /meetings GET request
1. The Application starts by creating a few variables that need to be used.
2. The Application then goes into a loop based on an URL.Query which take 0-many arguments based on the fields in meetings.
+ dateAndTime 
+ topic
+ roomName
+ ownerName
4. Because roomName and ownerName are not used directly inside the query we have to do a sub query to match both those fields with their corresponding roomID and ownerID.
5. All the values that are used in the search are then stored as values in values.
6. Those values are stored as a string to be used in a query later on.
7. If the length of the values is 0 the Application redirects you to the home page.
8. A result query is sent based on the string created by the parameters in the search.
9. The result set is then iterated over and are scanned and error checked.
10. That data is then appended to a meetings.Meetings variable and then output.

## Users
### /users/:id/settings GET request

### /users/:id/settings/create POST request

### /users/:id/settings/edit PUT request

### /users/:id/settings/delete DLETE request

## AgendaSearch GET request
The Agenda search is a Function that uses RegEx to search through the database and check to see if there is valid data that matches the regular expression.
### /AgendaSearch?sentence=var1
Sample input: 
localhost:9090/agendaSearch?sentence=cat

What this does is goes through the database and matches up any agenda which starts with the preceding word and returns it if is either owned by the person searching or is actually a participant for that meeting. 
eg: searching for "KeyWord" will return all meetings with which agenda's start with the key word and is also a valid member of that meeting.

### /AgendaSearch?phoneNumber=var1
Sample input: 
localhost:9090/agendaSearch?phoneNumber=8449159 (explicit number find)

What this does is goes through the database and matches up any agenda which has the phone number in it any where and returns it if is either owned by the person searching or is actually a participant for that meeting.
eg: searching for "8449159" will return true if it is in the agenda but will not return if the number is in "844-9159" format if number has a '-' in it the search will need to look for that explicitly.

### /AgendaSearch?email=var1
Sample input: 
localhost:9090/agendaSearch?email=PaulD@eit.co.nz 

What this does is goes through the database and matches up any agenda which has the email in it anywhere and returns it if is either owned by the person searching or is actually a participant for that meeting.
eg: Searching for the email as a full search or only part of the email address should return the agenda that is related to the search.

### /AgendaSearch?keyWords=var1,var2,var3
Sample input: 
localhost:9090/agendaSearch?keyWords=cat&keyWords=mat&keyWords=pat

What this does is goes through the database and matches up any agenda which has the 3 key words in it in any order, it is limited to 3 words but can be made into as many as needed but 3 was where we stopped for functionality purposes; this then returns the meetings it if is either owned by the person searching or is actually a participant for that meeting.

### /AgendaSearch?dollar=var1
Sample input: localhost:9090/dollar=$4.50

- Discuss features of the MeetingPlanner that are candidates to be executed on the client
side instead of on the server. Clearly describe the pros and cons.

################################## haven't done this part YET!!!!!! 

- All persistent data (bookings, user accounts, etc.) are to be stored in a PostgreSQL
database. Explain the design choices you made to interact with the database.
## Database Model
### Object Based Data Model
In our Database we have 5 Tables to represent the entire structure of our application.
1. users
2. rooms
3. meetings
4. priorMeetings
5. participants
#### users
users consists of 6 attributes, id is the PRIMARY KEY and all other fields are generated at time of user creation.
1. id SERIAL PRIMARY KEY
2. userName VARCHAR(50)
3. name VARCHAR(50)
4. phone VARCHAR(20)
5. email VARCHAR(50)
6. password VARCHAR(20)
##### rooms
rooms consists of only two attributes one is the PRIMARY KEY and the other is the name of the room.
1. id SERIAL PRIMARY KEY
2. name VARCHAR(20)
#### meetings
meetings is our most relational Table as it used by most functions to retrieve data. meetings is made up of 6 attributes, half of which are ID's. The reason for this was that it is a 'Meeting Planner Application' and needless to say 'meetings' are what it’s about. The use of references in relation to the meetings makes the relationships easy to establish as every person has something to do with some meeting 'generally' speaking.
1. id SERIAL PRIMARY KEY
2. topic VARCHAR(20)
3. dateAndTime TIMESTAMP 
4. agenda VARCHAR(1000)
5. roomID INT REFERENCES rooms(id)
6. ownerID INT REFERENCES users (id)
#### priorMeetings
priorMeetings is a copy of a previous meeting, in which we store the same data as a previous meeting. This table is purely relational and is completely ID based.
1. id SERIAL PRIMARY KEY
2. meetingID INT REFERENCES meetings (id)
3. userID INT REFERENCES users (id)
#### participants
participants is also a relational Table and is mealy a set of ID's that relate to usersID's and meetingID's.
1. id SERIAL PRIMARY KEY
2. meetingID INT REFERENCES meetings (id)
3. userID INT REFERENCES users (id)
##### Seeders
To test our Database we have a set of randomly generated text behind Keywords in matching fields to test the database. Each seeder INSERT's data into corresponding Tables that return valid data sets for each field.

- The enterprise typically hosts a variety of operating systems and internet browsers.
Discuss how your solution copes with this variety.
1. API don’t need to worry :D 

- Provide a "Quick Start Guide" outlining the steps and details required to install your
Application on a new server.
1. Have a Database in place that matches the corresponding postgres dataset in the application in this case it is 
    postgres 
    user=postgres 
    password=password 
    dbname=meetingplannerdb 
    sslmode=disable
2. Open a CMD window by typing cmd.exe into the search portion of the Windows screen and navigate to the file location.
3. Run the following command in the 'main' folder: go install main
4. Followed by: main.exe
5. Install RESTer on your Browser (preferably Chrome)
6. Open RESTer to Test the Application
7. Logging in as a "Test User"
- Use a POST method
- Use the URL localhost:9090/login
- The formatting is in .json and so in the body of the request type:
{
    "userName": "Test1",
    "password": "Password1"
}

- Provide a document that lists the additional specifications that were missing but
Required to implement your solution.

