# STYLE BLITZ
Style blitz is an online platform that empowers service providers within the fashion and beauty niche to enlarger their online presence. Blitz gives users the necessary tools to automate booking, engaging larger customer base and adverting their products for free. 
Blitz Will further empower these service providers with the ability to create customised profiles and shop tailored their brands and services.

## Project structure
**Style Blitz** is structured to use modern solutions like GRAPHQL. Most of the technological decisions were made to. ease the development timeline, so that we can focus more on implementing core features than building solutions to already resolved problems.
To make the app mode decoupled, the project is structured base on its core defining technology. Eg lambda folder holds all things lambda related.

To push for massive code reuse which further reduces development time, core access to the database is abstracted into a different project. This removes the need for recreating these methods. NB: we do this since we're using the databases. The idea is to create a monolithic application that is highly decoupled and encourages massive code reuse. This solution could result in a single failure in the reusable code breaking the entire application. 
However comma the solution to this is to encourage massive unit and integration testing. This does not fix the issue of a major part of the code relying on the shared library, but it helps reduce the possibility of failure drastically.
- ### Folders:
    - Lambda: All things lambda
    - Shared: All shared code across modules and services
    - Docs: Well written documentation for the app
    - Migrations: Service that handles all things DB migration
    - Utils: App-wide command line utility that simplifies deploy. migration and app development. 

## Technologies
1. **Golang**
2. **Graphql**
3. **Mysql**
4. **Redis/Mongo**
5. **Swagger**
6. **Bash**
7. **AWS/ Lambda / Cloud formation**
8. **Docker**
9. **Serverless**

