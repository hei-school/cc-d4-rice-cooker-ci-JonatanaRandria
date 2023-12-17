[![Review Assignment Due Date](https://classroom.github.com/assets/deadline-readme-button-24ddc0f5d75046c5622901739e7c5dd533143b0c8e959d652212380cedb1ea36.svg)](https://classroom.github.com/a/PHq8Kfj_)

# Ruby implementation of My Rice Cooker

Here you can find the implementation of __Vi'lany__ using __Ruby__.

## Requirements :computer:

On your machine you will have to install and setup ruby environment. You can refer to this [documentation](https://www.ruby-lang.org/fr/documentation/installation/) for installation.

## Installation :hammer_and_wrench:
Clone this repository in your local machine:
```shell
    git clone https://github.com/hei-school/cc-d4-rice-cooker-ci-JonatanaRandria
```

Checkout into the __feature/ruby__ branch:
```shell
    cd cc-d4-rice-cooker-ci-JonatanaRandria/
    git checkout origin/feature/ruby
```

## Running the application :flight_departure:

To run the application, just enter the following command: 
```sheel
    ruby ./src/service/mains.rb
```

## Bug :bug:
When leaving the application, it doesn't directly leave but still loop on the menu. 



## Workflows
I used [Circle CI](https://circleci.com/) to run my workflows.

The workflows configuration file is viewable in _.circleci/config.yml_. Here it only aims at running
tests.