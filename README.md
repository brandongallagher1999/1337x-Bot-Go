![1337x Logo](https://duckduckgo.com/i/e4d3d1a0.png)

# Unofficial Discord Bot

![deployment](https://github.com/brandongallagher1999/1337x-Bot-Go/actions/workflows/deploy.yaml/badge.svg)

## Description

- Uses a 1337x Microservice (Nest.JS) and Golang to create a Bot that allows users to search for content such as Movies, Games, etc. and returning a relevant
  list of Torrents with their respective shortened magnet URLs, Names, Seeders and File Sizes.

# Commands

```txt
// Search for a list of Torrents
.torrent <query>

// Display all available commands
.help

// Get the GitHub link to this project
.github
```

# How to run using Docker

## Pre-Requisites

- Install Docker [**here**](https://docs.docker.com/get-docker/)

## Clone repository

- Clone the project by running:

```txt
git clone https://github.com/brandongallagher1999/1337x-Bot-Go/
```

## Configuration File

1.
```sh
touch config/config.yaml
```

2. Edit config/config.yaml

```yaml
discord:
  token: "<Your Bot's Token Here>"
  prefix: "."
  command: "torrent"
  maxLinksPerQuery: 10
```

## Container

- Go into root folder and run

```txt
docker-compose build
docker-compose up -d
```

# Example / Usage

![Image of the Bot Working](/images/example.jpg?raw=true)
