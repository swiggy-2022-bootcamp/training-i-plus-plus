# Mini Project - Online Shopping Store

## Services

- **Auth**
- **Product**
- **Search**
- **User**
- **Cart**

## Prerequisties

- Sonarqube

## Steps to run Sonarqube

- Run the following command in root of the project/microservices to generate coverage file.<br>
  > `go test -short -coverprofile=./cov.out ./...`
- Run the following command in root of the project/microservices to generate gosec report.<br>
  Install gosec from [Gosec Github link](https://github.com/securego/gosec)<br>
  > `gosec -fmt=sonarqube -out report.json ./...`
- SonarQube properties file is already present in the root of each service directory.
- Set the project in SonarQube and provide the same name in `sonar.projectKey` field and running address of SonarQube in `sonar.host.url` field of `sonar-project.properties`.
- Provide the generated authentication token by SonarQube in `sonar.login` field of `sonar-project.properties`.
- Run the `sonar-scanner` command to run sonar scanner and visit to dashboard for the analysis of the codebase.

## References

- Requirement Document - [Online Shopping Store](https://docs.google.com/document/d/1cnCHEVkOgFDYSmZmSbxcDlZiLjCZXr1W9jHf62id7T8/edit?usp=sharing)
