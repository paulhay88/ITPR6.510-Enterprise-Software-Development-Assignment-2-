# Documentation

---

- Routing options. Available routes and accepted input such as POST.
## Login 
### /login POST request
1. To Access the database a user must log in with credentials to be able to return data after which all data is verified by the users Cookie.
We have chosen to use the negroni package to handle the authentication of the Cookies with a function called logged in to check if the Cookie holds the "authUser" credentials matching the current session.
2. Login requiers a password and a username which is parsed and checked against the database.
3. Firstly checks if the users Cookie is currently active and returns a "User is already logged in." message if the Cookie is currently active.
4. The information from the body is then decoded into a logInUser variable.
5. A query is then sent containing that name and password
6. A variable of a new user is created and the query is then scanned for data
7. Error check the scan and out put "Your username or password was incorrect. Please try again."
8. Else creater and set Cookie
9. Then redirect to the home page of / with a valid Cookie in place.
10. Doing a GET request on this page will tell you to log in

### /signup POST request
1. sign up is a simple POST that creates a User by creating space on the Database. 
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
2. That list is then iterated through to check the data that is relevent to this User.
3. In the loop a new meeting is created to temporarly store relevant meeting data.
4. The entire list of meetings is scanned and checked for errors.
5. The application then checks to see if the current user is a participant in all meetings according to their name and userID through a SQL query and stores it as a variable from the database.
6. The application parses a password locally to keep it secure.
7. The appliication then then appeneds all relevent data to myMeetiings and then outputs it.

### / GET request
1. Home page after you login will always appear with the related meetings to your userID
2. The application first checks the Cookie for an authorized user then creates a meeting and user variable to use
3. The UserName is taken from the Cookie from the first varible split by ":"
4. A query is then constructed based on that name to get the userID
5. Error check
6. Another query is then launched based on the participants userID called participants
7. Error check the return
8. Participants is then scaned to an error check
9. A new query is launcehed based on the meetingID
10. That is also error checked 
11. After all that checking the application then appendeds the meetings.Meetings
12. Then Outputs the results

### /meeting/create POST request
1. The application creates an instance of a new meeting 
2. Then decodes the body of the request into a variable
3. Then inserts it as an Exec INSERT directly into the Database.
4. That is all error checked.

### /meeting/:id/edit PUT request
1. The application created an instance of a new meeting.
2. then decodes that information into a variable which is error checked.
3. An Exec on the Database is then preformed which takes all the variables from the body of the request and inserts them intto the Database WHERE the id is matching.

### /meeting/:id/delete DELETE request
1. This fires two Exec on the database.
2. The first one DLETEs FOM participants WHERE the meetingID matched the variable.
3. The second part DELETEs FROM meetings WHERE id matched the variable.
4. Both are error checked.

### /meetings GET request
1. The Application starts by creating a few variables that need to be used.
2. The Application then goes into a loop based on an URL.Query which take 0-many arguements based on the fields in meetings.
+ dateAndTime 
+ topic
+ roomName
+ ownerName
4. Because roomName and ownerName are not used directly inside the query we have to do a sub query to match both those fields with thir corosponding roomID and ownerID.
5. All the values that are used in the search are then stored as values in values.
6. Those values are stored as astring to be used in a query later on.
7. If the length of the values is 0 the Application redirects you to the home page.
8. A result query is sent based on the string created by the paramiters in the search.
9. The result set is then itterated over and are scanned and error checked.
10. That data is then appeneded to a meetings.Meetings variable and then outped.

## Users
### /users/:id/settings GET request

### /users/:id/settings/create POST request

### /users/:id/settings/edit PUT request

### /users/:id/settings/delete DLETE request

- Discuss features of the MeetingPlanner that are candidates to be executed on the client
side instead of on the server. Clearly describe the pros and cons.
- All persistent data (bookings, user accounts, etc.) are to be stored in a PostgreSQL
database. Explain the design choices you made to interact with the database.
- The enterprise typically hosts a variety of operating systems and internet browsers.
Discuss how your solution copes with this variety.
- Provide a "Quick Start Guide" outlining the steps and details required to install your
application on a new server.
- Provide a document that lists the additional specifications that were missing but
required to implement your solution.