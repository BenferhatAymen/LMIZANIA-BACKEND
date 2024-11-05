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

### Authentication Endpoints
- `POST /login` - Log in a user.
- `POST /register` - Register a new user.
- `POST /verify/{id}` - Verify a user account.
- `POST /resetpassword` - Reset a user's password.

### Balance Endpoints
- `GET /balance/wallet` - Get user's wallet balance.
- `PUT /balance/wallet` - Set/update user's wallet balance.
- `GET /balance/target` - Get user's financial target.
- `PUT /balance/target` - Set/update user's financial target.
- `POST /balance/savings` - Deposit into user's savings.
- `GET /balance/savings` - Retrieve user's savings balance.
- `GET /balance/income` - Get user's income information.
- `GET /balance/expense` - Get user's expense information.

### Goals Endpoints
- `POST /goals` - Create a new financial goal.
- `PUT /goals/{id}` - Update a specific goal.
- `DELETE /goals/{id}` - Delete a specific goal.
- `GET /goals` - Retrieve all goals for the user.
- `POST /goals/{id}/deposit` - Deposit funds towards a specific goal.

### Transactions Endpoints
- `POST /transactions` - Add a new transaction.
- `PUT /transactions/{id}` - Update a specific transaction.
- `DELETE /transactions/{id}` - Delete a specific transaction.
- `GET /transactions` - Retrieve all transactions for the user.


## Technologies Used

- **Golang**: Backend programming language.
- **MongoDB**: Database for storing application data.
- **Gorilla Mux**: HTTP router for handling API requests.
- **JWT**: JSON Web Token for authentication.
- **OTP**: One-time password for user verification.

## License

This project is licensed under the MIT License.

