# Shortest Driving Path Project

Given a starting location and a list of drop-off locations, this project finds the shortest driving path by making use of Google Maps Distance Matrix API and applying a Travelling Salesman algorithm.

## Installation

TODO: Write this

## Project Structure

The project is mainly separated into 4 parts:

- Front-Tier (`frontier`)
- Task Queue (`taskqueue`)
- Back-Tier (`backtier`)
- Garbage Collector (`garbagecollector`)

### Front-Tier

Front-Tier is an HTTP server responsible for accepting requests from the clients. After receiving a query from the client, it generates a token and registers the question in Task Queue. It sends back the token to the client so that it could come back for the result.

### Task Queue

Task Queue deals with 3 kinds of entities:

- Questions
- Answers
- Garbages

#### Question

Question is received from Front-Tier (and in turn, from the client). It contains the starting location and a list of drop-off locations.

#### Answer

Answer is received from Back-Tier. It contains the status of the query, the answer of the question, or error the occurred during the process.

#### Garbage

Garbage is received from Back-Tier. It contains information to help cleaning up entries no longer in use in Task Queue.

### Back-Tier

Back-Tier is responsible for finding the answer of the question. After receiving a question from Task Queue, it acts retrieves a distance matrix via the Google Maps Distance Matrix API. Then it calculates the shortest driving path applying a Travelling Salsman algorithm. After that, it sends the answer to Task Queue.

### Garbage Collector

TODO: Write this

## Directory Structure

    + (project root)
    ├── bitstring
    ├── cmd
    │   ├── backtier
    │   │   ├── internal
    │   │   │   └── getdm
    │   │   └── lib
    │   ├── frontier
    │   │   └── lib
    │   └── garbagecollector
    ├── common
    ├── constant
    ├── matrix
    ├── taskqueue
    └── tsp
        ├── bruteforce
        └── heldkarp

TODO: Write this (description)

## To Be Improved

TODO: Write this
