
# template-go

This repository defines a scaffolding project to kick-start a back-end application. It is meant to work best with a frontend service created using the [template-frontend](https://github.com/Knoblauchpilze/template-frontend).

# Why this project?

When building a back-end service, it is common to have to solve the same problems. This includes, among other things:
* starting a server listening to incoming connections to serve requests.
* interfacing with a database to get and save data.
* publishing the service's code to a docker image that can be deployed in a cluster
* etc.

Those operations are usually quite similar from one service to the next. This is where this repository comes into play: by using it, you can kick-start a new back-end service in an instance, by using a scaffolding that provides most of the boilerplate you will need.

# What does this repository provide?

By using this repository and adapting it for a new project you will get, out of the box:
* a working server which you can extend with your own endpoints.
* a working CI, which allows to verify your work, run your tests and publish the result of your implementation.
* a way to make the service available in `dockerhub` through a docker image.
* the possibility to deploy automatically a new version of the service to a cluster management system.

# Badges

[![Build services](https://github.com/Knoblauchpilze/template-go/actions/workflows/build-and-push.yml/badge.svg)](https://github.com/Knoblauchpilze/template-go/actions/workflows/build-and-push.yml)

[![codecov](https://codecov.io/gh/Knoblauchpilze/template-go/graph/badge.svg?token=D4V8SZODWV)](https://codecov.io/gh/Knoblauchpilze/template-go)

# Installation

## Prerequisites

The tools described below are directly used by the project. It is mandatory to install them in order to build the project locally.

See the following links:

- [golang](https://go.dev/doc/install): this project was developed using go `1.23.2`.
- [golang migrate](https://github.com/golang-migrate/migrate/blob/master/cmd/migrate/README.md): following the instructions there should be enough.
- [postgresql](https://www.postgresql.org/) which can be taken from the packages with `sudo apt-get install postgresql-14` for example.
- [docker](https://docs.docker.com/engine/install/ubuntu/): this is needed in case you want to build the docker image locally.

## Cloning the repository

Once the above steps have been completed, you can clone the repository with the following command:
```bash
git clone git@github.com:Knoblauchpilze/template-go.git
```

## Secrets in the CI

The CI workflows define several secrets that are expected to be created for the repository when cloned/forked/used. Each secret should be self-explanatory based on its name. Below is a description of the ones used by this project:

| Secret             | Description |
| ------------------ | ----------- |
| CODECOV_TOKEN      | A token generated from [codecov](https://about.codecov.io/) allowing to pubilsh the code coverage reports |
| DOCKERHUB_USERNAME | Obtained from your account over at [dockerhub](https://hub.docker.com/), allows to publish the docker image of the service |
| DOCKERHUB_TOKEN    | Also taken from [dockerhub](https://hub.docker.com), this authenticates the request to publish to the repositories |
| DEPLOYMENT_TOKEN   | Used to trigger the automatic deployment of the service's latest  version over at [ec2-deployment](https://github.com/Knoblauchpilze/ec2-deployment) |

# How does this project work?

## General structure

TODO: Project layout

## The CI

TODO

# How to use this project to kick-start a new service?

## Setting up the project

TODO

## Use the configure script

TODO

## What to modify next?

TODO

# How to extend this project

## Adding more packages

TODO

## Creating secrets

TODO

## Deploying the service

As mentioned in the [secrets](#secrets-in-the-ci) section, this project is configured to automatically update the deployment over at [ec2-deployment](https://github.com/Knoblauchpilze/ec2-deployment). This is achieved by:
* pushing the docker image of the service to `dockerhub`.
* triggering a commit to the `ec2-deployment` repository.

In case you have a similar workflow, it is easy to update the `DEPLOYMENT_TOKEN` to point to another repository and to modify the [CI workflow](.github/workflows/build-and-push.yml) in the `update-deployment` step to deploy to another repository.

In case you don't use `dockerhub` but another system to store the docker image (or don't use docker images at all), then you might need to rewrite part of the CI workflow (namely the `build-and-push-docker-image` step).

Once updated, this should automatically trigger a commit in the `ec2-deployment` (or any other repository) with the latest update. The commit message should be explicit enough:

![Example commit messages](resources/commit-messages-example.png)

## Working and developing on a new project

TODO
