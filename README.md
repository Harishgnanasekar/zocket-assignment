# Product Management System

## Overview
A modular and scalable product management system built with:
- **Go** for backend.
- **PostgreSQL** for the database.
- **Redis** for caching.
- **RabbitMQ** for message queuing.

## Prerequisites
- Install Go (1.19+)
- Install PostgreSQL

- Install Redis
- Install RabbitMQ

## Setup Instructions
1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd product-management-system
   ```

2. Configure the `.env` file with your environment variables.

3. Run database migrations:
   ```bash
   psql -U <username> -d <database> -f database_schema.sql
   ```

4. Install dependencies:
   ```bash
   go mod tidy
   ```

5. Start the server:
   ```bash