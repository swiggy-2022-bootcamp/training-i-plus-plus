# Requirement 
- Function Requirement 
  - Customer
    - Login 
    - Register
    - Order Food
    - Cancel Food
    - .... 
  - Restaurant 
    - Login 
    - Register
    - Add Food 
    - Update Food 
    - Remove Food 
- This is never clear 

- Non Functional Requirement 
  - Internet Scale Application 
  - 20 Billion / Min 
  - Trun Around for each order < 1s 
  - .... 
  -  to order a food item on swiggy it takes - 2 sec 


- People have Smart Phone 
  - 2022 - 100% 
  - 2020 - 100% 
  - 2015 - 80% 
  - 2012 - 50% 
  - 2010 - < 20 


- TDD - Test Driven Development 
- BDD - Behavioural Driven Development 
 ```
    Given the web application is loaded 
    When the user enters valid username and password
    Then take the user to home page / dashboard 
    And show him/her the list of foods availale 
    And on top right corner his/her name Welcome: Name
 ```

- DDD - Domain Driven Development 
- MDD - Model Driven Development 



- SLA 
- SLI 
- SLO 



NFR SLAs	
Generic
Performance SLAs	
"Consistent Performance, aligned with the incremental scaling…

1. Response times < 2 ms
2. Throughput - Req/Sec - 25
3. Concurrent users - 15


Optionally, we could raise the performance bar incrementally, as per the needs... 

1. Response times < 1 s, 100 ms, ...
2. Throughput - Req/Sec -  50, 100...
3. Concurrent users - 25, 100,..."
Scalability	"Incremental Scale…

# of transactions expected to register per day - in an increasing trend
(10, 100, 1000, 5000...)

Availability	"Incremental HA
99.0, 99.9, 99.99 "
Robustness & Resilience	"1. Proper Error page should be displayed
1. Continuous failures shall be communicated at priority via alerts to the support teams to be acted upon"
Security	"1. HTTPS/TLS
2. Encrypt PII Data in the data-store
3. Encrypt Credentials in files"
Accessiblity	"AA
Reponsive - supported view port 767px (Mobile and Desktop)
Phone, and Desktop
latest verison of Chrome and IE
"
Quality	"100% Unit test success  (FE / BE)
>80% Unit Test coverage (FE /BE)
JSLint /CS lint
Controlled Blocker , critical (FE / BE)"


Agile -> DOJO (SDET)

- POST	/api/register	To register the user with basic details like 
Input: 
{
  id: Number, 
  username: String, 
  email: String, 
  password: String 
   address: {
     houseno:Number, 
     street:String, 
     city:String, 
     state:String,  
     zip:number
  }   
}

Return code : 201 
Return: 
{
  id: Number, 
  username: String, 
  email: String, 
  password: String 
   address: {
     houseno:Number, 
     street:String, 
     city:String, 
     state:String,  
     zip:number
  }   
}



- Confluence - Attlasian 
- Jira - Attlasian 
- Bitbucket - Attlasian 

- 2000 - Java + OOps Concepts (i)
- 2003 - Java + Servlets + Jsp 
- 2010 - Java + i + i + i 


```
/** 
* this is sample method which will take employee object and after inserting gives +ve number 
**/
    int storeEmployee(Emp emp);
    
    e = new Employeee();
    int retVal = mock(storeEmployee (e))
    if(retVal ==1){
        Test Case passed
    }else {
        Test Case Faile 
    }

```


- Naveen Srinivas
- Naveen Srinivasan 


- No Code 
- Low Code World 