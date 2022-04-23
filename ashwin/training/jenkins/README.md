**What is Jenkins?**

====================

[Jenkins](https://www.jenkins.io/) is a continuous open-source integration written in Java. It was forked from Hudson Project after a dispute with Oracle. Since the fork, Jenkins has grown to be much more than a continuous integration solution.

Jenkins is not just a [Continuous Integration](https://devopscube.com/continuous-integration-delivery-deployment/) tool anymore. Instead, it is a Continuous Integration and Continuous delivery tool. You can orchestrate application deployments using Jenkins with a wide range of freely available community plugins and native Jenkins workflows.

Jenkins also supports GitOps workflows with [Jenkins X](https://jenkins-x.io/). It helps you accelerate the continuous delivery pipeline on Kubernetes.

**Jenkins Use Cases**

=====================

The following are the main Jenkins use cases from my experience extensively working on Jenkins.

1\.  **Continuous Integration:** With Jenkins pipelines, we can achieve CI for both applications and infrastructure as code.

2\.  **Continuous Delivery:** You can set up well-defined and automated application delivery workflows with Jenkins pipelines.

3\.  **Automation & Ad-Hoc Tasks**: With Jenkins jobs and pipelines, you can automate infrastructure components and perform ad-hoc infrastructure tasks (Backups, remote executions, etc.).

**Jenkins Tutorials For Beginners**

===================================

In this collection of Jenkins tutorial posts, I will cover various Jenkins tutorials, which will help beginners get started with many of the Jenkins core functionalities.

I have categorized the list of Jenkins beginner tutorials into multiple sections. From Jenkins installation to advanced concepts like shared libraries are covered in this tutorial list.

It is a growing list of Jenkins step-by-step guides. I will add all the latest tutorials to this list.

### **How Does Jenkins Works?**

In this section, you will learn the very basics of Jenkins that every beginner should know. Overall you will learn **how Jenkins works** as a CI/CD tool.

1\.  [Jenkins Architecture Explained for Beginners](https://devopscube.com/jenkins-architecture-explained/): In this tutorial you will learn how Jenkins work, its core componets, how Jenkins data is organized and the protocols involved in Jenkins master -- agent communitcation.

### **Jenkins Administration Tutorials**

1\.  [Installing and configuring Jenkins](https://devopscube.com/install-configure-jenkins-2-0/) [Ubuntu VM & Docker]

2\.  [Install and Configure Jenkins on centos/Redhat](https://devopscube.com/install-configure-jenkins-2-centos-redhat-servers/)

3\.  [Setting up a distributed Jenkins architecture (Master and agents)](https://devopscube.com/setup-slaves-on-jenkins-2/): This tutorial will teach you show to configure Jenkins master and agents using both SSH and JNLP methods.

4\.  [Configure SSL on Jenkins Server](https://devopscube.com/configure-ssl-jenkins/)

5\.  [Running Jenkins on port 80](https://devopscube.com/access-run-jenkins-port-80/)

6\.  [Backing up Jenkins Data and Configurations](https://devopscube.com/jenkins-backup-data-configurations/)

7\.  [Setting up Custom UI for Jenkins](https://devopscube.com/setup-custom-materialized-ui-theme-jenkins/) -- A tutorial to change the default UI of Jenkins with custom css.

8\.  [Setting up Jenkins Master on Kubernetes Cluster](https://devopscube.com/setup-jenkins-on-kubernetes-cluster/) -- Tutorial to setup Jenkins infrastructure on containers.

### **Scaling Jenkins Agents**

1\.  [Configuring Docker Containers as Jenkins Build Agents](https://devopscube.com/docker-containers-as-build-slaves-jenkins/)

2\.  [Scaling Jenkins Agents with Kubernetes](https://devopscube.com/jenkins-build-agents-kubernetes/)

3\.  [Configuring ECS as Build Slave For Jenkins](https://devopscube.com/setup-ecs-cluster-as-build-slave-jenkins/) [Needs Update]

### **Jenkins Pipeline Tutorials**

1\.  [Jenkins Pipeline as Code Tutorial for Beginners](https://devopscube.com/jenkins-pipeline-as-code/): Jenkins pipeline as code helps you organized and manage all the CI/CD jenkins pipeline as code. This tutorial walks you through the important pipeline concepts.

2\.  [Getting Started With Jenkins Shared Libary:](https://devopscube.com/jenkins-shared-library-tutorial/) Shared library is a very is mporatnt concept in Jenkins. It helps you to create re-usable pipeline code. It is a must-learn concept.

3\.  [Creating Jenkins Shared Library](https://devopscube.com/create-jenkins-shared-library/): Step by step guide on practical implementation of Jenkins shared library

4\.  [Beginner Guide to Parameters in Declarative Pipeline](https://devopscube.com/declarative-pipeline-parameters/): In this tutorial, you will learn to manage static and dynamic Jenkins parameters using pipleine as code.

### **Jenkins CI/CD Tutorials**

1\.  [Java Continuous Integration with Jenkins](https://devopscube.com/java-continuos-integration-jenkins-beginners-guide/): This tutorial teaches you to setup a Java CI pipeline using maven build utility. [Needs Update]

2\.  [Jenkins PR based builds with Github Pull Request Builder Plugin](https://devopscube.com/jenkins-build-trigger-github-pull-request/)

3\.  [Jenkins Multi-branch Pipeline Detailed Guide for Beginners](https://devopscube.com/jenkins-build-trigger-github-pull-request/)

4\.  [Building Docker Images Using Kaniko & Jenkins on Kubernetes](https://devopscube.com/build-docker-image-kubernetes-pod/)

5\.  How to manage Secrets with Jenkins [Upcomming]

6\.  How to Build Docker Image using Jenkins [Upcomming]

**Jenkins Core Features**

=========================

Let's have look at the overview of key Jenkins 2.x features that you should know.

1\.  Pipeline as Code

2\.  Shared Libraries

3\.  Better UI and UX

4\.  Improvements in security and plugins

**Pipeline as Code**

====================

Jenkins introduced a DSL by which you can version your build, test, deploy pipelines as a code. Pipeline code is wrapped around a groovy script that is easy to write and manage. An example pipeline code is shown below.

```

node('linux'){

  git url: '<https://github.com/devopscube/simple-maven-pet-clinic-app.git>'

  def mvnHome = tool 'M2'

  env.PATH = "${MNHOME}/bin:${env.PATH}"

  sh 'mvn -B clean verify'

}

```

Using pipeline as a code you can run parallel builds on a single job on different slaves. Also, you have good programmatic control over how and what each Jenkins job should do.

Jenkinsfile is the best way to implement Pipeline as code. There are two types of **pipeline as code**.

1\.  Scripted Pipeline and

2\.  Declarative Pipeline.

I recommend using only a declarative pipeline for all your Jenkins-based CI/CD workflows, as you will have more control and customization over your pipelines.

Jenkins 2.0 has a better User interface. The pipeline design is also great in which the whole flow is visualized. Now you can configure the user, password, and plugins right from the moment you start the Jenkins instance through awesome UI.

Also, [Jenkins Blueocean](https://jenkins.io/projects/blueocean/) is a great plugin that gives a great view of pipeline jobs. You can even create a pipeline using the blue ocean visual pipeline editor. Blueocen looks like the following.

![https://devopscube.com/wp-content/uploads/2017/04/pipeline-run.png](https://devopscube.com/wp-content/uploads/2017/04/pipeline-run.png)

**Jenkins Shared Libraries**

============================

[Jenkins shared library](https://jenkins.io/doc/book/pipeline/shared-libraries/) is a great way to reuse the pipeline code. You can create your CI/CD code libraries, which can be referenced in your pipeline script. In addition, the extended shared libraries will allow you to write custom groovy code for more flexibility.

Here is how shared libraries work.

1\.  The common pipeline shared library code resides on a version control system like Github.

2\.  The shared library GitHub repo needs to be configured in Jenkins global configuration. It enables access to the library for all the Jenkins jobs.

3\.  You have to add an import statement to import the shared libraries and use it in the pipeline code in individual jobs. During job builds, first, the shared library files get checked out.

**GitOps With Jenkins X**

=========================

[Jenkins X](https://jenkins.io/projects/jenkins-x/) is a project from Jenkins for CI/CD on Kubernetes. This project is entirely different from normal Jenkins.

Jenkins X has the following key features.

1\.  GitOps based Tekton pipelines

2\.  Environment Promotion via GitOps

3\.  Pull Request Preview Environments

4\.  Feedback and chat on Issues and Pull Requests