This project is a simple crypto wallet application that allows users to create, edit, delete, and view their wallets. The project is divided into a **backend** and a **frontend**. Below are the steps to set up and run the application.

---

## Requirements

- **Backend**:
  - Node.js and npm installed on your machine.
  - XAMPP or any other local server to initialize the MySQL database.
  
- **Frontend**:
  - React (set up using `create-react-app`).
  - Any modern browser for accessing the frontend.

---

## Installation and Setup

### 1. **Initialize the Database**

1. Start **XAMPP** and enable **Apache** and **MySQL** services.
2. Open [phpMyAdmin](http://localhost/phpmyadmin) in your browser.
3. Create a new database called `crypto_wallet`.
4. Execute the following SQL script to create the `wallets` table:

```sql
CREATE TABLE wallets (
  id INT AUTO_INCREMENT PRIMARY KEY,
  username VARCHAR(255) NOT NULL,
  address VARCHAR(255) NOT NULL,
  balance DECIMAL(10, 2) NOT NULL,
  currency VARCHAR(50) NOT NULL
);
```

---

### 2. **Run the Backend**

The backend is responsible for handling API requests such as creating, updating, deleting, and retrieving wallet information. It is built using **Node.js** and **Express.js**.

1. Navigate to the backend folder:
   ```bash
   cd backend
   ```
2. Install the required dependencies:
   ```bash
   npm install
   ```
3. Configure the database connection:
   - Open the `config.js` file (or similar, depending on your setup).
   - Update the database settings to match your local MySQL configuration:
     ```javascript
     module.exports = {
       host: 'localhost',
       user: 'root',
       password: '', // Your MySQL root password
       database: 'crypto_wallet',
     };
     ```
4. Start the backend server:
   ```bash
   npm start
   ```
5. The backend server should be running at `http://localhost:3000`.

### Backend API Endpoints

- `GET /wallets/username/:username` – Fetch all wallets associated with the provided username.
- `POST /wallets` – Create a new wallet.
- `PUT /wallets/:id` – Update an existing wallet.
- `DELETE /wallets/:id` – Delete a wallet.

---

### 3. **Run the Frontend**

The frontend is a React-based web application that allows users to interact with the crypto wallet. It communicates with the backend API to display, create, update, and delete wallets.

1. Navigate to the frontend folder:
   ```bash
   cd frontend
   ```
2. Install the required dependencies:
   ```bash
   npm install
   ```
3. Start the React development server:
   ```bash
   npm start
   ```
4. The frontend application should open automatically in your browser at `http://localhost:3001`. If not, open the URL manually.

### Frontend Functionality

- **Enter Username**: The user must input their username to fetch and manage their wallets.
- **Create Wallet**: The user can create a new wallet by providing a wallet address, balance, and currency.
- **View Wallets**: All wallets associated with the entered username will be displayed.
- **Edit Wallet**: The user can edit the wallet details such as the address, balance, and currency.
- **Delete Wallet**: The user can delete a wallet.

---

## How to Use the Application

1. **Enter Username**: On the homepage, enter your username to fetch or manage wallets associated with your account.
2. **Create Wallet**: Use the "Create Wallet" form to add a new wallet. The wallet will be saved with the current username.
3. **View Wallets**: All wallets associated with the entered username will be displayed.
4. **Edit Wallet**: Click on "Edit Wallet" to modify wallet details.
5. **Delete Wallet**: Use the delete button to remove a wallet.

---

## Troubleshooting

- If the backend does not connect to the database, ensure your MySQL server is running and the credentials in `config.js` are correct.
- Ensure both backend and frontend servers are running simultaneously.
- If there are CORS issues, you may need to enable CORS on the backend by adding a package like `cors`:
  ```bash
  npm install cors
  ```
  And update your backend server code:
  ```javascript
  const cors = require('cors');
  app.use(cors());
  ```

---

## Tech Stack

- **Frontend**: React
- **Backend**: Node.js, Express.js
- **Database**: MySQL (via XAMPP)
