# Switcheo Backend Internship - Problem 5 - CRUD Interface for Blockchain resource - Jun Han Huang

## Overview
This directory contains a full-stack CRUD implementation for blog posts. 
### blog
blog contains the blockchain itself. The blockchain leverages CosmosSDK and Ignite CLI extensively to conduct the manipulation of blog posts and the maintenance of transactions across nodes in the blockchain. 
### blogclient
blogclient written in golang, is the backend API for the blockchain. It recieves CRUD requests from the frontend and executes the command upon the blockchain accordingly.
### blogclient/frontend
blogclient/frontend contains a react frontend interface to allow the user to more easily manipulate the blockchain. A list of all posts is also displayed.

## How to Run

### blog
### blogclient
### frontend


## Challenges
    -first time using golang, learning what a context is and how to use it
    -becoming familar with golang deployment and package management

## Future Improvements
### blog
        introducing 2nd resource, or more complicated data
### blogclient
        lots of code repition
        features for multiple accounts
### frontend
        reactive list of posts
        clientside / serverside validation
        how i would go about filtering


