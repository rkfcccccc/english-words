# English words learning app
This will be a full stack app for learning english words from movies (to watch them later, maybe).

The idea for this is quite simple: users can search for movies, pick the most interesting for them and then learn the words with something like *spaced repetition*.

## Backend (Go)
For backend **microservice architecture** written with Golang will be used. The current idea assumes the following list of services:  
1. API Gateway
2. User service
3. Verification service (by Email i guess)
4. Movie service
5. Dictionary service

For communication between services i'm going to use **gRPC** and **Apache Kafka**.

## Frontend (Flutter)
A mobile application with Flutter will play the role of a frontend for this app.

At the moment I see it mostly as a screen that scrolls down, giving you definitions of words and some challenges to test their knowledge.

Also there will be a list of movies that will represent a collection of words

## Future growth opportunities
Implement a language-dependent API for learning words not only in English.