# [Medo - The Healthcare Application 💝🏥](https://docs.google.com/document/d/1Yv4IiNrY4KB18DJa-aIl8wjnyTxemw6FgNrCGpM124s/edit?usp=sharing)


## Problem Statement:

We've learned to seek engaging in-hand comfy experiences from every service provider as smartphones transform mundane chores like food ordering, taxi and more. And the healthcare industry is no exception. Patients are transforming into consumers who do not want to waste time booking or postponing healthcare visits. They want to discover healthcare professionals who can aid them promptly, read user ratings, have digital copies of their health data and schedule an appointment as soon as possible. That alone appears to be sufficient to persuade an application for healthcare.

## Features:
* **Scheduling appointments** anytime, anywhere.
* Special **Direct-Interaction Access** for Pro Users (VIP Patients).
* **Reminders** are sent automatically and **synchronised** with their digital calendars.
* A comprehensive **listing** of available healthcare experts for patients.
* Possibility of **uploading** any existing health-related documents prior to a consultation / **downloading** Prescriptions digitally.
* Payments & Billing

## Topics Implemented:
1. Design : **Domain Driven Design**
2. Database : **MongoDB**
3. APIs :     **REST API**
4. Documentation : **Swagger**
6. Logging : **Custom Logger**
7. Error Handling : **Custom Error Handling**
8. Containerization : **Docker**

## Functional Requirements : 
1. **Patient Profile (Registration & Login )**
    * Check if the Patient has Basic/Pro Subscription.
2. **Patient Health-Record Database**
    * Ability to upload the health record prior consulting
    * Ability to delete the health records (Patients)
3. **Patient Appointments & Reminders**
    * Ability to download the schedule as invite
    * Schedule recurring appointments (Pro Users)
4. **Doctor Profile  / Ratings**
    * Ability to provide ratings/comments post consultation.
    * Available list of doctors and their ratings
5. **Payments**
    * Choose between multiple payment options.


## Modules:
1. Login/Register Module
2. Search Available Doctors Module
3. Appointment Module
4. Health Record Module
5. Special Access Module (Pro)
6. Payments Module

## Few Screenshots:


### Login
![img.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1650723322587/SsIyO1Yaq.png)
### JWT Token  
![02-login-token.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1650723369633/Qd5koTG65.png)
### Home
![login-03-200OK-admin.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1650723513319/t2QjIBfW_.png)
### If no token present?
![login-04-nocookie.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1650723433813/oyL1rE2z-.png)
#### then, Login will not be authorized.
![login-05-nocookie-notauthorized.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1650723474176/E_dFlyvIR.png)
### Available List of Doctors

![doctors-list.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1650723565781/J7-OggQt_.png)
