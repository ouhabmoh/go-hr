HR API
This document describes the API for an HR API.

## Technologies Used

Go (programming language)
Gin (web framework)
Gorm (ORM for Postgres)
Viper (configuration management)
JWT (JSON Web Token) for authentication

## Endpoints

User Management

GET /users/me (Requires authentication): Retrieve information about the currently logged-in user.
Authentication

POST /auth/register: Register a new user.
POST /auth/login: Login an existing user and obtain an access token.
GET /auth/refresh: Refresh an expired access token (requires a valid refresh token).
GET /auth/logout (Requires authentication): Logout the currently logged-in user.
Job Management

(Requires authentication for all job endpoints)
POST /jobs: Create a new job posting.
GET /jobs: Find all available jobs .
PUT /jobs/:jobID: Update an existing job posting (provide the job ID in the URL path).
GET /jobs/:jobID: Get details of a specific job posting (provide the job ID in the URL path).
DELETE /jobs/:jobID: Delete a job posting (provide the job ID in the URL path).
Application Management

(Requires authentication for all application endpoints)
POST /applications: Create a new application for a job opening.
GET /applications: Find all applications submitted by the logged-in user (candidate).
GET /applications/:applicationID: Get details of a specific application (provide the application ID in the URL path).
PATCH /applications/:applicationID: Update an existing application (provide the application ID in the URL path).

## Authentication

This API uses JWT for authentication. Users need to register or login to obtain an access token. Subsequent requests should include the access token in the authorization header to access protected resources.

Unimplemented Functionalities & Upcoming Features
RBAC (Role-Based Access Control):

Implement access control based on user roles (candidate/recruiter).
Candidates can only apply for jobs, manage applications, and update profiles.
Recruiters can create/manage jobs, manage applications, and potentially schedule interviews (future).
Candidate Profile Management:

Implement endpoint for updating candidate profiles (work experience, education, etc.).
Job Applications:

Implement endpoint for uploading resumes for specific job applications.
Interview Scheduling (Future):

Explore functionalities for managing interview scheduling between recruiters and candidates.

## Getting Started

(Replace these instructions with your specific setup steps)

Install Go and set up your development environment.
Clone the project repository.
Configure the database connection details and other settings in the configuration file (create a new file app.env and check the .env.example).
Run the migrations (run the migrate/migrate.go file).
Build and run the API server.
Use a client library or tool that supports sending HTTP requests with JSON bodies and authorization headers to interact with the API.
