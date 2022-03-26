#CI/CD

CI stands for continuous integration and CD refers to continuous delivery.

CI means that any substantial code changes are easily built and tested in a shared environment 

CD makes the changes done by the developer automatically go through test cases and deployment if code passes the checks

#Jenkins

Jenkins is an open source automation server. It helps automate the parts of software development related to building, testing, and deploying of software, helping achieve CI/CD.

Check setup

-sudo systemctl daemon-reload
-sudo systemctl start jenkins
-sudo systemctl status jenkins

#Jenkins Pipeline

Jenkins pipeline is a collection of events or jobs which are interlinked with one another in a sequence achieved by a combination of plugins that support the integration and implementation of continuous delivery pipelines.

Simply, Jenkins Pipeline is a collection of jobs that brings the software from version control to production by using automation tools.

# JenkinsFile

Jenkins Pipeline is configured by a text file called JenkinsFile.
