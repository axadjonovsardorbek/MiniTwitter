# MiniTwitter

MiniTwitter is a lightweight social network allowing users to post short messages (twits), like twits, follow others, and interact with other users. This project uses a RESTful API architecture where users can follow each other, post twits, and like them.

Features

User creation and management
Users can post twits and share content
Ability to like and unlike twits
Users can follow and unfollow other users
Manage followers and followings list
Technologies

Go - Primary programming language for the backend
PostgreSQL - Database system
Gin - Web framework for building the HTTP RESTful API
Docker - Used for containerization and running the project in an isolated environment
Database Structure

The project uses the following core tables:

users
Column	Type	Details
id	UUID	Primary key, unique
name	VARCHAR	User's full name
username	VARCHAR	Unique username
bio	VARCHAR	User bio (optional)
image_url	VARCHAR	Profile picture URL (optional)
password	VARCHAR	User's password
email	VARCHAR	User's email
created_at	TIMESTAMP	Creation time
updated_at	TIMESTAMP	Last update time
deleted_at	BIGINT	Soft delete flag
twits
Column	Type	Details
id	UUID	Primary key, unique
user_id	UUID	Reference to the user
twit_id	UUID	Reference to the original twit (for retwits)
content	VARCHAR	Twit content
image_url	VARCHAR	Image URL (optional)
created_at	TIMESTAMP	Creation time
updated_at	TIMESTAMP	Last update time
deleted_at	BIGINT	Soft delete flag
follows
Column	Type	Details
id	UUID	Primary key, unique
follower_id	UUID	ID of the user following
followed_id	UUID	ID of the user being followed
created_at	TIMESTAMP	Creation time
updated_at	TIMESTAMP	Last update time
deleted_at	BIGINT	Soft delete flag
likes
Column	Type	Details
id	UUID	Primary key, unique
user_id	UUID	ID of the user who liked the twit
twit_id	UUID	ID of the twit that was liked
created_at	TIMESTAMP	Creation time
updated_at	TIMESTAMP	Last update time
deleted_at	BIGINT	Soft delete flag
Installation

To set up and run the project on your local machine, follow these steps:

1. Clone the repository
bash
Copy code
git clone git@github.com:axadjonovsardorbek/MiniTwitter.git
cd MiniTwitter
2. Start with Docker
Use Docker to start the project:

Copy code
docker-compose up --build
docker-compose up
docker-compose up
This command will spin up PostgreSQL and other necessary services in containers.


Errors and Troubleshooting

In case of any errors, you can check the logs located in the logs directory for debugging. To view logs from a Docker container, use:

docker logs <container_name>