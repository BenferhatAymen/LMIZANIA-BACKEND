# Lmizania-Golang-Backend

**Lmizania-Golang-Backend** is a backend project built in Go, designed for managing transactions, goals, user authentication, and more. It leverages MongoDB for data storage, JWT for secure access, and OTP for verification.

## Features

- **User Management**: Includes user registration, login, and OTP-based verification.
- **Transaction Management**: Supports income, expense tracking, and wallet balance updates.
- **Goal Management**: Allows users to set financial goals and deposit towards them.
- **Authentication**: JWT-based authentication for secure user sessions.
- **OTP Verification**: Adds an additional layer of security with one-time password verification.
- **Role-Based Access Control (RBAC)**: Access to certain features based on user roles (if applicable).

## Project Structure

- `config/`: Configuration files, including OTP and database connection settings.
- `controllers/`: Contains the logic for handling HTTP requests and defining API endpoints.
- `database/`: MongoDB connection setup and database-related utilities.
- `middlewares/`: JWT and other middleware functions for request handling and validation.
- `models/`: Struct definitions for the application's data models, including User, Transaction, and Goal models.
- `pkg/`: Helper packages and utility functions.
- `repository/`: Database layer, where CRUD operations for models are defined.
- `routes/`: API route definitions for user, transaction, and goal management.

## Installation

1. **Clone the repository**:
   ```bash
   git clone https://github.com/yourusername/Lmizania-Golang-Backend.git
   cd Lmizania-Golang-Backend
   ```

2. **Set up environment variables**:
   Create a `.env` file in the root directory and add your MongoDB connection URI, JWT secret, and other required environment variables.
   
   Example:
   ```
   MONGO_URI=mongodb://localhost:27017
   JWT_SECRET=your_jwt_secret
   OTP_SECRET=your_otp_secret
   ```

3. **Install dependencies**:
   ```bash
   go mod tidy
   ```

4. **Run the application**:
   ```bash
   go run main.go
   ```

## API Endpoints

The backend provides several RESTful API endpoints, including:

- **User Endpoints**:
  - `POST /register`: Register a new user.
  - `POST /login`: User login with OTP verification.
  - `GET /user/:id`: Get user profile.

- **Transaction Endpoints**:
  - `POST /transaction`: Add a new transaction.
  - `GET /transaction/:id`: Get transaction details.
  - `DELETE /transaction/:id`: Delete a transaction.
  - `PUT /transaction/:id`: Update a transaction.

- **Goal Endpoints**:
  - `POST /goal`: Create a new financial goal.
  - `GET /goal/:id`: Get goal details.
  - `POST /goal/deposit`: Deposit into a goal.
  - `DELETE /goal/:id`: Delete a goal.

## Technologies Used

- **Golang**: Backend programming language.
- **MongoDB**: Database for storing application data.
- **Gorilla Mux**: HTTP router for handling API requests.
- **JWT**: JSON Web Token for authentication.
- **OTP**: One-time password for user verification.

## License

This project is licensed under the MIT License.

