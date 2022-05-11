![1337x Logo](https://duckduckgo.com/i/e4d3d1a0.png)

# Unofficial Discord Bot


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

# How to run via Azure VM with Terraform

## Prerequisites

- Install Terraform [**here**](https://learn.hashicorp.com/tutorials/terraform/install-cli)

- Install Azure CLI [**here**](https://docs.microsoft.com/en-us/cli/azure/install-azure-cli)

## Login to your Azure Instance with a Subscription

```sh
az login

# or

az login --use-device-code
```


## Go into /terraform/

- Edit versioning.tf as instructed by the comments

- Run these commands

```sh
terraform init
terraform apply
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

![Image of the Bot Working](/images/example.png?raw=true)
![Bot replying to Query with no results](/images/noresults.png?raw=true)
