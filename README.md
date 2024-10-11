# CTRL - Chapter Version Control System

CTRL is a Go-based project that manages chapter versions for various companies and projects. It provides an API to retrieve chapter details and version information.

## Features

- Retrieve chapter versions and details
- Database-driven storage (PostgreSQL)
- RESTful API

## Prerequisites

- Go 1.16 or higher
- PostgreSQL database

## Setup

1. Clone the repository:
   ```
   git clone https://github.com/your-username/ctrl.git
   cd ctrl
   ```

2. Set up the database:
   - Create a PostgreSQL database `ctrl`
   - Load the sql script in `db/script/create-db.sql`
   - Update the database configuration in `src/env/config.local.json`

3. Build the project:
   ```
   cd src
   make build
   ```

## Running the Application

To run the application, use the following command from the `src` directory:

```
make run
```

The application will start a local server at `http://localhost:8080`.

## API Documentation

- Get all chapters:
  ```
  curl http://localhost:8080/chapter_versions/{chapter_id}
  ```
