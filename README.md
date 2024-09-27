# MiniTwitter

MiniTwitter is a lightweight social network allowing users to post short messages (twits), like twits, follow others, and interact with other users. This project uses a RESTful API architecture where users can follow each other, post twits, and like them.

Features

User creation and management: Users can create accounts and manage their profiles.

Twits: Users can post and share content through twits.

Likes: Users can like and unlike twits.

Follow system: Users can follow and unfollow other users.

Manage followers: Users can view and manage their followers and followings.

Technologies

Go: Primary programming language for the backend.

PostgreSQL: Database system used to store user, twit, follow, and like data.

Gin: Web framework used to build the HTTP RESTful API.

Docker: Used for containerization and running the project in an isolated environment.

Installation

To set up and run the project on your local machine, follow these steps:

1. Clone the repository


git clone git@github.com:axadjonovsardorbek/MiniTwitter.git

cd MiniTwitter

2. Start with Docker

Use Docker to start the project:

docker-compose up --build

docker-compose up

docker-compose up

This command will spin up PostgreSQL and other necessary services in containers.

3. Errors and Troubleshooting

In case of any errors, you can check the logs located in the logs directory for debugging. To view logs from a Docker container, use:

docker logs <container_name>