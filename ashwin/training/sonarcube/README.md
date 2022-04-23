SonarQube | SonarScanner : a step by step guide.
================================================

![SonarQube Logo](https://miro.medium.com/max/1094/1*KOadiTidZoHtrfanE3Ck5A.png)

<https://www.sonarqube.org/logos/>

Introduction
============

SonarQube is an open-source platform developed by SonarSource for continuous inspection of code quality to perform automatic reviews with static analysis of code to detect [bugs](https://www.techopedia.com/definition/3758/bug), [code smells](https://en.wikipedia.org/wiki/Code_smell), and [security vulnerabilities ](https://owasp.org/www-community/vulnerabilities/)on 20+ programming languages.

> It can report duplicated code, coding standards, unit tests, code coverage, code complexity and comments.

> The only prerequisite for running SonarQube is to have Java ([Oracle JRE 11](https://www.oracle.com/java/technologies/javase-jre8-downloads.html) or [OpenJDK 11](https://www.oracle.com/java/technologies/javase/javase-jdk8-downloads.html)) installed on your machine. [Read More](https://docs.sonarqube.org/latest/requirements/requirements/)

Installation steps:
-------------------

Step 1:

[Download ](https://www.sonarqube.org/downloads/)the SonarQube Community Edition.

Step 2:

As a non-root user, unzip it, let's say in C:\sonarqube or /opt/sonarqube.

Step 3:

# On Windows, execute:`\
C:\sonarqube\bin\windows-x86-xx\StartSonar.bat`

# On other operating systems, as a non-root user execute:`\
/opt/sonarqube/bin/[OS]/sonar.sh console`

Step 4.

Open browser and <http://localhost:9000/> (9000 is default) you will be navigated to below window, with System Administrator credentials (login=admin, password=admin).

Note:
-----

For any configuration changes go to *conf* folder and *sonar.properties *file.

Here you can configure *database*, *LDAP*, *webserver*, *SSO authentication, logging, etc...*, e.g. for port --- under web-server section I have added *sonar.web.port=9001*

![SonarQube Local Dashboard](https://miro.medium.com/max/1094/1*MK4eCDxjW5mQ50aer8xKTw.png)

Using Docker --- (Optional)
-------------------------

Images of the Community, Developer, and Enterprise Editions are available on [Docker Hub](https://hub.docker.com/_/sonarqube/).

1.  Start the server by running:

$ docker run -d --name sonarqube -p 9000:9000 <image_name>

Step 5:

After login to the application, click the Create new project button to analyze your first project.

![SonarQube dashboard new project creation direction](https://miro.medium.com/max/1094/1*ORo1qsMI0wSLAkgV4WVcKg.png)

1.  Click on "+" icon on right-top corner on navigation bar
2.  Select *'Create new project'* option

Step 6:

![](https://miro.medium.com/max/1094/1*b7m95bU2WIC7zc4Y__vmWw.png)

Step 7:

![](https://miro.medium.com/max/1094/1*XSmsXcVTJwjoyj3h7hZsJA.png)

1.  Enter a token key (Enter your favorite word pairs)--- here *secret_key*
2.  After clicking the *generate *button, the application will provide a token. Which later is use for verification purpose before starting scan of specified project.

![](https://miro.medium.com/max/1094/1*dJLyvGtUDR0-0WZHg9XlnA.png)

Now click on *continue *button.

Step 8:

Select type of your project, mine is Angular in(JavaScript language)

![](https://miro.medium.com/max/1094/1*WDniyc9L4M-3iA3JU6OwPA.png)

The moment you click the button on step 3, it'll redirect you to download page for sonar scanner. Select the dist based on OS you're using.

![](https://miro.medium.com/max/1094/1*IkJkEyplX7YAktvAIr4jzQ.png)

<https://docs.sonarqube.org/latest/analysis/scan/sonarscanner/>

After completion of download of sonar scanner, extract the file. (I've extracted in the download folder)

Step 9:

Window --- Register the sonar-scanner path in environment variable.

![](https://miro.medium.com/max/1094/1*J8_Wp03Wi0Vx01QAZlc59w.png)

Mac --- [Setting up Environment Variables in MacOS Sierra](https://medium.com/@himanshuagarwal1395/setting-up-environment-variables-in-macos-sierra-f5978369b255)

Step 10:

Save the following properties in your project-folder ---

with file name sonar-project.properties (!important).

sonar.projectKey=TLH_PROJECT_SQ_V1\
sonar.projectName=TLH_PROJECT_SQ_V1\
sonar.login = ba4fd*******************\
sonar.scm.provider = svn\
sonar.projectVersion=1.0\
sonar.sources=src\
sonar.exclusions=node_modules/**,src/environments/**,**/*.spec.ts,dist/**,**/docs/**,**/*.js,e2e/**,coverage/**,TLH-distributions/**,src/bsb-theme/css/**\
sonar.ts.tslint.configPath=tslint.json\
sonar.typescript.lcov.reportPaths=coverage/lcov.info

*Feel free to change the above properties based on your project config.*

Add the* sonar-project.properties* at root level of project.

![](https://miro.medium.com/max/1094/1*oXxB3RGmkP-JoeFqGqzWkg.png)

Now open your project path in Terminal or CMD. Run the following command

sonar-scanner.bat

![](https://miro.medium.com/max/1094/1*wO1WkkDw63Mvv9Pm4iZUzQ.png)

*Sit back and relax, the scan will take a while. Go grab a coffee in a while. *☕

After the completion of scan go to the SonarQube dashboard (localhost:9000). Login if required. Select the project you'll able to view something as below.

![](https://miro.medium.com/max/1094/1*ULThMm00rfm9qHpXAAFuhw.png)

Yay! 🙌 If you've followed this along, then congratulations you have made it! and now you may share the report (*after correcting/fixing all the issues*) to your Project manager and other stakeholders.

Go to issues tab, select type of issues you want to fix and SonarQube will show the defined rule/guideline w.r.t to the issue.

![](https://miro.medium.com/max/1094/1*0R9v6CEGdIMLmDPjtEwUww.png)



Jenkins-SonarQube Integration
=============================

![](https://miro.medium.com/proxy/1*jRIb96bpeYbvyqKUNDgq9g.png)

Assume a Scenario : After I committed code to *GitHub*. I want to ensue my code quality, know bugs, vulnerabilities, code smells, etc. (static code analysis) for my code before I build my code automatically with *Jenkins *and I want this activity to perform every time I commit code.

In this scenario for Continuous Inspection and Continuous Integration of the code. We will follow the best practice using *GitHub-Jenkins-SonarQube* Integration for this scenario.

Flow : As soon as developer commits the code to *GitHub*, *Jenkins *will fetch/pull the code from repository and will perform static code analysis with help of *Sonar Scanner *and send analysis report to *SonarQube Server* then it will automatically build the project code.

Prerequisite :

1.  *Jenkins* is setup with *GitHub* with some build trigger (in my case its Poll SCM) if this is not done please follow this tutorial --- [*https://medium.com/@amitvermaa93/jenkins-github-with-java-maven-project-c17cdba7062*](https://medium.com/@amitvermaa93/jenkins-github-with-java-maven-project-c17cdba7062)
2.  *SonarQube* is running and you have *Sonar Scanner* setup in system. If not please follow the tutorial- [*https://medium.com/@amitvermaa93/sonarqube-setup-windows-e6a6c01be025*](https://medium.com/@amitvermaa93/sonarqube-setup-windows-e6a6c01be025)

Step 1. Open SonarQube server- Go to Administration > click on Security > Users > Click on Tokens (image 1)> Generate token with some name > Copy the token (image 2), it will be used in Jenkins for Sonar authentication.

![](https://miro.medium.com/max/1400/1*6QROeqXR8rxA36FcAhLCBw.png)

Image 1

![](https://miro.medium.com/max/1078/1*sBkaHFrHgTWWwMjiTpg2Qg.png)

Image 2

Step 2. Setup *SonarQube* with *Jenkins*- Go to *Manage Jenkins* > *Configure system *> *SonarQube* server section > *Add SonarQube* > Name it, provide Server Url as *http://<IP>:<port>* > and authentication token copied from SonarQube Server > *Apply* and *Save*

![](https://miro.medium.com/max/1400/1*VIPfmzWA5IJXyLF2pqAHOQ.png)

Step 3. Install *SonarQube plugin* to Jenkins. Go to *Manage Jenkins* > *Manage Plugins* > *Available *> Search for *SonarQube Scanner*> *Install.*

![](https://miro.medium.com/max/1400/1*dUbYg1JQEcCamPv5bgWVyA.png)

Download *SonarScanner* if you don't have [*https://docs.sonarqube.org/display/SCAN/Analyzing+with+SonarQube+Scanner*](https://docs.sonarqube.org/display/SCAN/Analyzing+with+SonarQube+Scanner)

Configure *Sonar Scanner* in *Jenkins* : Go to *Mange Jenkins* > *Global Tool Configuration* > Scroll for *SonarQube Scanner* > *Add* sonar scanner > name it, uncheck if you already have sonar else it will automatically download for you and your sonar scanner setup will be done(in my case I already have) > provide path to* sonar runner home *as in below image

![](https://miro.medium.com/max/1400/1*0h6C5ZolRGgBb80ghgaVfQ.png)

Step 4. Create a Job- *New Item* > Name and select a project type (in my case I am selecting *Maven* project you can opt for freestyle as well)

![](https://miro.medium.com/max/1400/1*ky8yX09XNgRLLcd9TP8AoA.png)

Step 5. Set *Git* under *SCM* section and use * * * * * for *Poll SCM* under B*uild Trigger *section. Under *Build Environment* section add pre-buid step > select *Execute SonarQube Scanner*

![](https://miro.medium.com/max/1400/1*91NtS40TL_1Y8S0bbdJC8A.png)

Step 6. Create a .properties file at any location and provide path on the task as below(I have created it in Jenkins workspace folder). This property file will be *project specific*. It contains certain sonar properties like which folder to scan, which folder to exclude in scanning, what is the project key and many more you can see it from [*https://docs.sonarqube.org/display/SCAN/Analyzing+with+SonarQube+Scanner*](https://docs.sonarqube.org/display/SCAN/Analyzing+with+SonarQube+Scanner)

![](https://miro.medium.com/max/1400/1*D6i4QxnwaVs_KugnxW2xww.png)

Inside sonar-scanner.properties write below code ---

*sonar.projectKey=github-jenkins-sonar\
sonar.sources=./src*

To keep it simple I have used only two properties(as above), *sonar.projectKey* property will create a project inside your *SonarQube server* with the same name if project don't exist else it will append analysis to it, *sonar.sources* defines that which folder to scan. You can provide either relative path from your Jenkins Job workspace or actual path to the folder you want to scan.

Since I have used *./src **(use / for windows path ) *thatmeans that I am currently on my Job workspace i.e. on *C:\Users\Amit Verma\.jenkins\workspace\Jenkins-GitHub-SonarQube* location and from here I am providing the path to the folder(src) I want to scan.

Step 7. Build the job. After successful build if you can see build logs it will show you the files and folder it has scanned and after scanning it has posted the analysis report to *SonarQube Server* you have integrated.

Step 8. From job dashboard, click on sonar icon or navigate to *Sonar server c*lick on Projects (on header) you will see a new project with same project key you have given in *sonar-scanner.properties* file. Now you can go inside your project and analyse the report

![](https://miro.medium.com/max/1400/1*hkNan7nrmrLiDn9i-ILZtQ.png)
