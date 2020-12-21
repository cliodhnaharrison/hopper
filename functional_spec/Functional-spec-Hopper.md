# CA326 Third Year Project  

## Functional Specification Requirement  

- **Project title:** Hopper
- **Student 1 Name:** Maciej Swierad **Student ID:** 18455726
- **Student 2 Name:** Clíodhna Harrison **Student ID:** 15440568
- **Supervisor:** Stephen Blott  


## Table of Contents  
1. **Introduction**  
    1.1. Overview  
    1.2. Business Context  
    1.3. Glossary  
  
2. **General Description**  
    2.1.  Product / System Functions  
    2.2.  User Characteristics and Objectives  
    2.3.  Operational Scenarios  
    2.4.  Constraints  
  
3. **Functional Requirements**  
    3.1.  Login  
    3.2.  Register  
    3.3.  Create container  
    3.4.  Discard container  
    3.5.  Add exercise  
    3.6.  Test solution  
    3.7.  Set preferences  
    3.8.  Delete account  
4. **System Architecture**  
5. **High-Level Design**  
6. **Preliminary Schedule**  
7. **Appendices**  

## 1. Introduction  

### 1.1 Overview  

Our idea is to provide users with discardable containers in which they have full access to the system in which their actions have no consequences on the system at large as each container can easily be discarded. Administrators can set exercises that can be found on an /exercises endpoint on the website, the users can then complete the exercises in the container and then their solution could be evaluated by a script mounted on a Docker volume. The users can request to have their logs sent to them via email after each container discard. Users would connect to the service by logging into a website in which they can request a container after which they will be taken to an endpoint with a container running WeTTY in their browser.  



### 1.2 Business Context  
The website could be used by anyone to learn from the exercises or to simply make use of a linux container in their browser.  

The backend of the application could be modified to be deployable in an extended project. The website could also be modified to become course/university/company specific and custom exercises could be added.


### 1.3 Glossary  

**Containers**  
Containers are a form of operating system virtualization.  

**Administrator**  
This is a super user with full permissions that has access to both the backend and frontend of the application.  

**Distros**  
These are Linux Distributions, we will be using distribution docker images.  

**WeTTY**  
A web based TTY written in TypeScript.  

**Docker Volume**  
This is a file system different to the standard Docker file system, it is used to preserve data generated by the current container.  

**OAuth**  
This is a way for users to login to our website using Google without granting us the use of their passwords.  

## 2. General Description  

### 2.1 Product / System Functions  
The necessary functions for Hopper are, but are not limited to, the following:  

- Login
- Register
- Create Container
- Discard container
- Add exercise
- Test solution
- Set preferences
- Delete account  

### 2.2 User Characteristics and Objectives  
The application will be hosted on a website and not all users will have access to all the features of the application. The permissions go as follows:  

**Guest**

These are unregistered or not signed in users, their permissions are limited to either registering an account or signing in to an account.  

**Administrators**

These are the maintainers of the website and application with elevated permissions. They have access to adding exercises, access to docker config, logs etc. and deleting other user accounts along with adding new distros.The administrators will also be responsible for the running of the website.  

**Members**

These users are registered with the website where the application instance is running, they have access to spin up containers in which they have root permissions.  

### 2.3 Operational Scenarios  

**Login**  

Users can login to the website.  

**Register**  

Users can register for an account on the website.  

**Create container**  

Users can request a new container and the system will create one.  

**Discard container**  

Users can request that their container be discarded. The system checks what the user's preferences are and depending on that it may email the logs and history of the container to the user. The system then destroys the container.  

**Add exercise**  

Administrators can add exercises that users can complete.  

**Test solution**  

Users can submit their solutions to the exercises to be tested.  

**Set preferences**  

Users can set their preference for what distribution they want their containers to be, set their preference for whether the logs and history of a container are emailed to them after the container has been discarded and set any other preference that may arise.  

**Delete account**

Users can delete their accounts.

### 2.4 Constraints
**Time Constraints**  
In this project we are going to be using a lot of technologies unfamiliar to us so we will need to allow for some time to research within our alloted time for a task. Overall as we are under a time constraint we will have to manage our time well as we are limited to 8 weeks of development time. To manage our time well we are going to use Agile methodology and employ weekly sprints and regular standups.  

**Code linting**  
In order to stay in tune with best coding practices we will be using various different linters such as ShellCheck for bash scripts, GoLint for Go and Hadolint for Dockerfile.  

**Accessibility**  
We are going to ensure that the website is readable by screen readers, and that the page can adjust to accommodate other visual impairments. We will be using a black and white colour scheme for the containers themselves.  

## 3. Functional Requirements   


### 3.1 Login  

**Description**  
The user is going to login using OAuth  based on the account linked in the register function.  

**Criticality**   
This requirement is of high importance for the end product however it is not critical to the essential function of the application.  

**Technical Issues**  
We must ensure that user data is stored securely at all times. For this we need to make sure our database is secure from any database injection attacks.  

**Inter-requirement Dependencies**  
The function depends on the user having registered previously.  

### 3.2 Register  

**Description** 
For a user to log in, they must first register an account. They can do so by navigating to the registration form provided on the website. Where they will be prompted to register using a Google account.  


**Criticality**  
This requirement is of high importance for the end product however it is not critical to the essential function of the application.  

**Technical Issues**  
We must ensure that user data is stored securely at all times. For this we need to make sure our database is secure from any database injection attacks.  

**Inter-requirement Dependencies**  
None.  

### 3.3 Create container  

**Description**  
The system will create a container using whatever distro image the user has chosen in their preferences. Prior to implementing preferences the default will be an ubuntu image.  

**Criticality**  
This requirement is the most important part of the application. It is the minimum viable product and will be implemented first.  

**Technical Issues**  
We must ensure that user data is stored securely at all times. For this we need to make sure our database is secure from any database injection attacks.  
 
**Inter-requirement Dependencies**  
None.  

### 3.4 Discard container  

**Description**  
The user is able to discard their container at will, the system will send the users logs and history based on their preferences and proceed to destroy the container.  


**Criticality**  
This requirement is a critical part of the application. It is the minimum viable product and will be implemented first along with creating containers.  

**Technical Issues**  
We must ensure that user data is stored securely at all times. For this we need to make sure our database is secure from any database injection attacks.  

**Inter-requirement Dependencies**  
Create container must have been called previously and there must be a container currently running.  

### 3.5 Add Exercise  

**Description**  
The administrator is able to add exercises to the system for the users to complete.  

**Criticality**  
This requirement is of moderate importance, this allows administrators set exercises for users to complete and to be marked on.  

**Technical Issues**  
We must make sure that the user has no access to the marking scheme of the given exercise, we will do this by using Docker volumes.  

**Inter-requirement Dependencies**  
None.  

### 3.6 Test Solution  

**Description**  
The user is able to submit their solution to the given exercise to be tested and graded by the system based on marking scheme provided by administrator.  

**Criticality**  
This requirement is of moderate importance, this allows users to complete exercises.  

**Technical Issues**  
We must make sure that the user has no access to the marking scheme of the given exercise, we will do this by using Docker volumes.  

**Inter-requirement Dependencies**  
Create container must have been called previously and there must be a container currently running and an exercise must have been previously added.  

### 3.7 Set Preferences  

**Description**  
The user is able to set their preferences such as what is their distro of choice, whether they want their logs and history sent via email when they discard a container etc.  

**Criticality**  
This requirement is of moderate importance, it's important for the end experience but the system can function without it.  

**Technical Issues**  
We must ensure that the user's preferences are correctly linked to their account and stored securely.  

**Inter-requirement Dependencies**  
The user must have registered previously and be currently logged in.  

### 3.8 Delete Account  

**Description**  
The user is able to delete their account and remove any link between their Google account and the Hopper application.  

**Criticality**  
This requirement is of moderate importance, it's important for the end experience but not vital for usability.  

**Technical Issues**  
We must ensure that the user's information is securely and timely deleted to comply with GDPR and storage of information.  

**Inter-requirement Dependencies**  
Create container must have been called previously and there must be a container currently running.  

## 4. System Architecture  
**Frontend**  
The Frontend will be a website built using Javascript, HTML and CSS. It will be receiving information from the backend via HTTPS.
It will be used to display the UI information received.
When a user presses a button such as "Request Container" a call will be made to the backend API written in GO.  

**Container**  
The container will be running WeTTY to provide an interface for the user. WeTTY is a web terminal run with NodeJS. Reverse proxy will be in place to map the containers to endpoints.  

**Backend**  
The backend will be a Linux box running the database, provisionally Redis, and running our Go code which will be handling the creation of containers via API. Traefik is a networking tool which we will be using to map container WeTTY ports to endpoints for users to hit.  

## 5. High Level Design  
In this section we illustrate how the program will flow using a Data Flow diagram. This diagram shows how exactly the system and the users interract and what flow the program can take.

![](https://i.imgur.com/IOFbdMt.png)



## 6. Preliminary Schedule  
Below is a preliminary Gantt chart illustrating our aim for task duration in days for the each of project's tasks and goals. They could possibly change along the way due to certain tasks being more or less difficult than we have estimated.

![](https://i.imgur.com/h5UahRw.png)


## 7. Appendices  
Docker: https://www.docker.com/  

Node: https://nodejs.org/en/  

OAuth: https://oauth.net/  

React: https://reactjs.org/  

Redis: https://redis.io/  

Traefik: https://traefik.io/  

Wetty: https://github.com/butlerx/wetty  






