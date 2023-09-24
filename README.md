# Golang Fiber App Template

A template for building web applications in Golang using the Fiber web framework. Kickstart your web development project with this pre-configured template.

## Features

- **Fiber Web Framework**: Utilizes the popular Fiber web framework for building efficient and fast web applications.
- **Structured Directory Layout**: Follows a well-organized directory structure for your codebase.
- **Docker Support**: Easily containerize your application for deployment.
- **Configuration Management**: Manage your application configuration with ease.
- **Database Integration**: Includes an example of connecting to a PostgreSQL database.
- **Cache Integration**: Includes an example of connecting to a Redis Instance.
- **Authentication Middleware**: Basic authentication middleware included.
- **Error Handling**: Implements error handling to provide a better user experience.
- **Logging**: Configured logging to monitor your application effectively.
- **Testing**: Includes testing setup and example test cases.
- **Middleware Support**: Easily integrate and customize middleware to enhance your application's functionality.

## Getting Started

1. **Clone the Repository**:

   ```bash
   git clone https://github.com/your-username/Golang-Fiber-App-Template.git
   cd Golang-Fiber-App-Template
   ```
   
2. **Environment Variables**:

   Create a .env file in the root directory and add the necessary environment variables. You can use the provided .env.example as a starting point.

3. **Database Setup**:

   If you are using a database, set up the database connection in platform/database.

4. **Cache Database Setup**:

   If you are using a cache database like redi, set up the database connection in platform/cache.

5. **Run the Application**:

   Run the following command to start the server
   ```
   make run
   ```

6. **Access the Application**:

   Your application should now be running locally at http://localhost:3000

## Configuration

All application configuration can be found in the .env file. Customize it to suit your project's needs.

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments

- Fiber: A fast, Express.js-like web framework for Golang.
- GORM: The fantastic ORM library for Golang.
- fiber-go-template (https://github.com/create-go-app/fiber-go-template): An amazing repository for fiber-go-template which served us a great inspiration to improve it.
