This project is a slight modification of Shopping Cart
-------------------------------------------------------
Generally in Shopping Cart User buys different products like Mobile,Bags, etc

But, In this project User buys services provided by employees like Plumber, Painter, Carpenter, Daily Labour, Field Worker

-------------------------------------------------------

Definitions:

User   : Customer who buys a service
Expert : Person who sells his service (like Carpentry, plumber, painter)  to the User

-------------------------------------------------------

Features included in this Project:

1) Authentication and Authorisation
2) Book Expert to get services
3) Rating and Reviewing Expert
4) Signin and Signup
5) System : Automatic Expert Suggestion to the User
6) System : An algorithm is implemented based on which it favours the Expert with good rating to get more work
7) Payment after work is completed  //  Since there is no GUI it might be impossible but will try

-------------------------------------------------------

Technologies Used:

1) GoLang for Backend server
2) Gin for routing
3) MongoDB for Database


--------------------------------------------------------

API Requests:
=============

1) GET   /expert/service   => to get all the services available
2) GET   /expert/expert?skill=carpenter  =>Automatically CPU assigns an expert to the user
3) POST  /expert/addrating?expertid=3   => to add rating and review to an user
4) GET   /expert/get?skill=carpenter  => it shows all the experts who are carpenters
5) GET   /expert/done?expertid=3   => it releases the expert so system can assign him to other users
6) POST  /expert/signexpert  => creates a new expert 
7) GET   /expert/getexpert?expertid=4  =>get an expert based on his ID
8) GET   /expert/filter?skill=carpenter&rating=4   => filters experts and returns based on the rating
9) POST  /user/signuser  => creates an user
10)POST  /user/loginuser  => checks the user and returns JWT-token
11)GET   /expert/services  => returns all the available services in the system
12)POST  /user/isuserpresent  => returns boolean whether the user is present or not
13)POST  /user/getuser   =>  returns a specfic user