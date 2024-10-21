# Debitor-Case Application

## Overview

Debitor-Case is composed of three independent microservices: `People-Service`, `Contract-Service`, and `Property-Service`. Each service manages its own PostgreSQL database and communicates via RESTful APIs. The system is designed to manage people, their contracts, and associated property details.

## Prerequisites

Before running the project, ensure you have the following installed:

- **Golang**: Version 1.18+
- **Docker**: For containerizing the microservices.
- **Docker Compose**: To orchestrate multi-container environments.
- **PostgreSQL**: For database management.
- **Terraform**: To set up the infrastructure on Google Cloud Platform (GCP).
- **Google Cloud SDK**: For deploying to GCP.


## Environment Variables

Each microservice requires environment variables for database configuration. You can pass them via a `.env` file (for local development), Docker Compose, or Google Cloud Secret Manager in production. The required environment variables are:

- `DB_HOST`: The hostname for the database (e.g., `localhost` for local development).
- `DB_USER`: The database user.
- `DB_PASSWORD`: The database password.
- `DB_NAME`: The database name.
- `DB_PORT`: The port for PostgreSQL (default: 5432).

## Running the Application Locally

### Step 1: Set Up Environment Variables

To run the services locally, create a `.env` file in the root of your project and specify the database credentials. Here's an example:

```
DB_HOST=localhost 
DB_USER=yourusername 
DB_PASSWORD=yourpassword 
DB_NAME=people_service 
DB_PORT=5432
```


Repeat this for the `contract-service` and `property-service`, specifying their respective `DB_NAME` values.

### Step 2: Start the Services Using Docker Compose

You can use Docker Compose to run all microservices and a PostgreSQL container. Run the following command:

```bash
docker-compose up --build
```
This will start:

- People-Service on http://localhost:8081
- Contract-Service on http://localhost:8082
- Property-Service on http://localhost:8083

### Step 3: Running Unit Tests
To run the unit tests for each microservice, navigate to the respective service directory and run:
```bash
cd people-service
go test ./tests

cd ../contract-service
go test ./tests

cd ../property-service
go test ./tests
```

## Deploying to Google Cloud
### Step 1: Set Up Infrastructure Using Terraform
In the terraform/ directory, you will find the scripts to set up Google Cloud infrastructure. These scripts will provision Cloud SQL, Cloud Run, and necessary networking.

Run the following commands to deploy the infrastructure:
```bash
cd terraform
terraform init
terraform apply
```

### Step 2: Deploy the Services to Google Cloud
You can deploy the microservices to Google Cloud Run. If you are using a CI/CD pipeline (such as GitHub Actions), the deployment process can be automated. Ensure that the pipeline is configured to:

- Build the Docker images.
- Push the images to Google Container Registry (GCR).
- Deploy the services to Cloud Run.

Alternatively, you can manually deploy each service:

```bash
gcloud builds submit --tag gcr.io/YOUR_PROJECT_ID/people-service ./people-service
gcloud run deploy people-service --image gcr.io/YOUR_PROJECT_ID/people-service --platform managed --region us-central1
```

> Repeat the process for the contract-service and property-service.

## API Endpoints
Each microservice exposes the following RESTful API endpoints:

### People-Service
- GET /people: Retrieve all people.
- GET /people/
: Retrieve a specific person.
- POST /people: Create a new person.
- PUT /people/
: Update a person's information.
- DELETE /people/
: Delete a person.
### Contract-Service
- GET /contracts: Retrieve all contracts.
- GET /contracts/
: Retrieve a specific contract.
- POST /contracts: Create a new contract.
- PUT /contracts/
: Update a contract.
- DELETE /contracts/
: Delete a contract.
### Property-Service
- GET /properties: Retrieve all properties.
- GET /properties/
: Retrieve a specific property.
- POST /properties: Create a new property.
- PUT /properties/
: Update a property's details.
- DELETE /properties/
: Delete a property.


## WARNING and DISCLAIMER
This project has no security configured, and terraform configures an external load balancer.
Add any form of authentication to the REST layer before deploying this to production
