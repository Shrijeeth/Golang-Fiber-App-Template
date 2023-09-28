# Golang Fiber App Template

A template for building web applications in Golang using the Fiber web framework. Kickstart your web development project with this pre-configured template.

## Features

- **Fiber Web Framework**: Utilizes the popular Fiber web framework for building efficient and fast web applications.
- **Structured Directory Layout**: Follows a well-organized directory structure for your codebase.
- **Docker Support**: Easily containerize your application for deployment.
- **Configuration Management**: Manage your application configuration with ease.
- **Database Integration**: Incorporates robust database integration to efficiently handle a variety of SQL databases, ensuring seamless data management capabilities.
- **Cache Integration**: Incorporates robust cache integration to efficiently handle Redis Cache
- **Cloud Object Storage Integration**: Incorporates robust cloud object storage integration to efficiently handle a variety of unstructured data, ensuring seamless data management capabilities.
- **Authentication**: Basic authentication is included which supports both normal login mechanism and google login mechanism.
- **Error Handling**: Implements error handling to provide a better user experience.
- **Logging**: Configured logging to monitor your application effectively.
- **Testing**: Includes testing setup and example test cases.
- **Middleware Support**: Easily integrate and customize middleware to enhance your application's functionality.

## Getting Started

1. **Clone the Repository**:

   ```bash
   git clone https://github.com/Shrijeeth/Golang-Fiber-App-Template.git
   cd Golang-Fiber-App-Template
   ```
   
2. **Environment Variables**:

   Create a .env file in the root directory and add the necessary environment variables. You can use the provided .env.example as a starting point.

3. **Database Setup**:

   If you are using a database, set up the database connection in environment variables.
   Below are the currently supported SQL databases:
      - Postgresql
      - MySql

4. **Cache Database Setup**:

   If you are using a cache database like redis, set up the database connection in environment variables.
   Below are the currently supported Cache databases:
      - Redis

5. **Cloud Object Storage Setup**:

   If you are using a cloud object storage, set up the connection in environment variables.
   Below are the currently supported Cloud Object Storage platforms:
     - AWS S3

6. **Application Linter**:

   Run the following command to run the linter
   ```
   make lint
   ```

7. **Run the Application**:

   Run the following command to start the server
   ```
   make run
   ```

8. **Access the Application**:

   Your application should now be running locally at specified address in environment variable

## Configuration

All application configuration should be in the .env file. Customize it to suit your project's needs. Use .env.example file as a starting point.

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments

- Fiber: A fast, Express.js-like web framework for Golang.
- GORM: The fantastic ORM library for Golang.
- GolangCi-Lint (https://github.com/golangci/golangci-lint): golangci-lint is a fast Go linters runner.
- fiber-go-template (https://github.com/create-go-app/fiber-go-template): An amazing repository for fiber-go-template which served us a great inspiration to improve it additional functionalities.
