# RSS Aggregator: A High-Performance Backend RSS Subscription Service

A robust, multi-user RSS feed aggregation service built entirely in Go (Golang). This backend API allows users to create accounts, subscribe to multiple RSS feeds, and efficiently retrieve consolidated content.

## Features
List the core functionalities of your server. This immediately tells a potential user or contributor what the project does.

* User Management: Secure user registration
* Subscription Service: CRUD operations for managing user subscriptions to external RSS feeds.
* RESTful API: A clean and documented set of HTTP endpoints for all functionalities.

## Setup

1. **Clone repository**
    ```
    git clone https://github.com/Aleksandar-G/rss-aggregator.git
    cd rss-aggregator
    ```

2. Configure Environment Variables
    ```
    # Example .env content
    DB_URL="host=localhost user=gorss dbname=gorss sslmode=disable password=yourpass"
    PORT=8080
    ```
3. Setup Database

    The application makes use of a SQL database. You can use anytpe of SQL database that you want.

4. Run Database Migrations:

    `make migration-up`

5. Run the Server

    `make run`

## API Endpoints

| Method  | Enpoint             | Description                         |
|---------|---------------------|-------------------------------------|
| `POST`  | `/v1/users`         | Create a new user                   |
| `DELETE`| `/v1/users/{id}`    | Delete a user                       |
| `GET`   | `/v1/users/{id}`    | Get a user by ID                    |
| `GET`   | `/v1/users`         | Get all users                       |
| `POST`  | `/v1/feeds`         | Create a new feed                   |
| `DELETE`| `/v1/feeds/{id}`    | Delete a feed                       |
| `GET`   | `/v1/feeds/{id}`    | Get a feed by ID                    |
| `GET`   | `/v1/feeds`         | Get all feeds                       |
| `POST`  | `/v1/user_feed`     | Create a new user feed subscription |
| `DELETE`| `/v1/user_feed/{id}`| Delete a user feed subscription     |
| `GET`   | `/v1/feeds/{id}`    | Get a user feed subscription by ID  |
| `GET`   | `/v1/feeds`         | Get all user feed subscriptions     |
| `GET`   | `/v1/healthz`       | Probe Endpoint for readiness        |
| `GET`   | `/v1/err`           | Test endpoint to check error page   |

## License
