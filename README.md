# Shortest Driving Path Project

Given a starting location and a list of drop-off locations, this project finds the shortest driving path by making use of a map provider (Google Maps Distance Matrix API) and applying a Travelling Salesman algorithm (brute force or Held & Karp).

## Table of Contents

- [Getting Started](#getting-started)
  - [Building](#building)
    - [Front-Tier](#front-tier)
    - [Back-Tier](#back-tier)
    - [Garbage Collector](#garbage-collector)
  - [Running](#running)
    - [Task Queue](#task-queue)
    - [Front-Tier](#front-tier-1)
    - [Backtier](#back-tier-1)
    - [Garbage Collector](#garbage-collector-1)
- [Project Structure](#project-structure)
  - [Front-Tier](#front-tier-2)
  - [Task Queue](#task-queue-1)
    - [Question](#question)
    - [Answer](#answer)
    - [Garbage](#garbage)
  - [Back-Tier](#back-tier-2)
  - [Garbage Collector](#garbage-collector-2)
- [Project Architecture](#project-architecture)
  - [Directory Structure](#directory-structure)
- [Dependencies](#dependencies)
  - [Gorilla Mux](#gorilla-mux)
  - [Google Logging Module](#google-logging-module)
  - [Beanstalk](#beanstalk)
  - [UUID package for Go language](#uuid-package-for-go-language)
  - [Google Maps Distance Matrix API](#google-maps-distance-matrix-api)
- [To Be Improved](#to-be-improved)

## Getting Started

### Building

#### Front-Tier

Create a new directory and change the current directory to it:

    mkdir temp1
    cd temp1

Download the Docker config file `Dockerfile`:

    curl -O https://raw.githubusercontent.com/asukakenji/151a48667a3852a43a2028024ffc102e/master/cmd/frontier/Dockerfile

Download the app config file `frontier.json`:

    curl -O https://raw.githubusercontent.com/asukakenji/151a48667a3852a43a2028024ffc102e/master/cmd/frontier/frontier.json

Edit the config file to fit the environment.

Build the Docker image:

    docker build -t frontier-image .

#### Back-Tier

Create a new directory and change the current directory to it:

    mkdir temp2
    cd temp2

Download the Docker config file `Dockerfile`:

    curl -O https://raw.githubusercontent.com/asukakenji/151a48667a3852a43a2028024ffc102e/master/cmd/backtier/Dockerfile

Download the app config file `backtier.json`:

    curl -O https://raw.githubusercontent.com/asukakenji/151a48667a3852a43a2028024ffc102e/master/cmd/backtier/backtier.json

Edit the config file to fit the environment.

Build the Docker image:

    docker build -t backtier-image .

#### Garbage Collector

Create a new directory and change the current directory to it:

    mkdir temp3
    cd temp3

Download the Docker config file `Dockerfile`:

    curl -O https://raw.githubusercontent.com/asukakenji/151a48667a3852a43a2028024ffc102e/master/cmd/garbagecollector/Dockerfile

Download the app config file `garbagecollector.json`:

    curl -O https://raw.githubusercontent.com/asukakenji/151a48667a3852a43a2028024ffc102e/master/cmd/garbagecollector/garbagecollector.json

Edit the config file to fit the environment.

Build the Docker image:

    docker build -t garbagecollector-image .

### Running

#### Task Queue

    docker run schickling/beanstalkd

### Front-Tier

    docker run frontier-image

### Back-Tier

    docker run backtier-image

### Garbage Collector

    docker run garbagecollector-image

## Project Structure

The project is mainly divided into 4 parts:

- Front-Tier (`frontier`)
- Task Queue (`taskqueue`)
- Back-Tier (`backtier`)
- Garbage Collector (`garbagecollector`)

### Front-Tier

Front-Tier is an HTTP server responsible for accepting requests from the clients. There are 2 kinds of requests.

The first kind contains a starting location and a list of drop-off locations. The client would like to know the shortest driving path which begins from the starting location, and visits all drop-off locations once. After receiving a request from the client, Front-Tier generates a token and sends the Question to Task Queue for being retrieved by Back-Tier. It then sends back the token to the client immediately so that it could come back for the result.

The second kind contains a token. Front-Tier looks up Task Queue and returns the Answer to the client.

### Task Queue

Task Queue deals with 3 kinds of entities:

#### Question

Question is received from Front-Tier (and in turn, from the request of the client). It contains the starting location and a list of drop-off locations.

#### Answer

Answer is received from Back-Tier. It contains the status of the query, the shortest driving path, or the error occurred during the process.

#### Garbage

Garbage is received from Back-Tier. It contains information to help cleaning up entries no longer useful in Task Queue.

### Back-Tier

Back-Tier is responsible for finding the solution for the query. After receiving a Question from Task Queue, it retrieves a distance matrix via the map provider. Then it calculates the shortest driving path by applying a Travelling Salsman algorithm. After that, it sends the Answer to Task Queue for being retrieved by Front-Tier.

### Garbage Collector

Garbage Collector is responsible for cleaning up Task Queue. Questions and Answers that are too old are deleted to free the memory they used.

## Project Architecture

### Directory Structure

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

## Dependencies

### Gorilla Mux

Gorilla Mux implements an HTTP request router and dispatcher. It is used to implement Front-Tier.

- Website: http://www.gorillatoolkit.org/pkg/mux
- GitHub: https://github.com/gorilla/mux

### Beanstalk

Beanstalk is a simple, fast work queue (task queue / job queue). It is used to implement Task Queue.

- Website: https://kr.github.io/beanstalkd/
- GitHub (Server): https://github.com/kr/beanstalkd
- GitHub (Client): https://github.com/kr/beanstalk

### Google Maps Distance Matrix API

Google Maps Distance Matrix API provides duration and distance values based on the recommended route between start and end points. It is used to implement Back-Tier.

- Website: https://developers.google.com/maps/documentation/distance-matrix/
- GitHub: https://github.com/googlemaps/google-maps-services-go

### UUID package for Go language

This package provides pure Go implementation of Universally Unique Identifier (UUID). It is used in the `common` package to generate new tokens.

- GitHub: https://github.com/satori/go.uuid

### Google Logging Module

Google Logging Module provides leveled execution logs for Go. It is used everywhere in the project. It is the main logging mechism.

- GitHub: https://github.com/golang/glog

## To Be Improved

- Write a connection pool for Task Queue.
- Orchestra to make the tiers to be-born if down.
- Implement Branch-and-Bound to support more locations, and use less memory, and use shorter execution time.
- Use the Golang `"context"` mechanism (new feature) to limit the runtime of each task.

TODO: Write this
