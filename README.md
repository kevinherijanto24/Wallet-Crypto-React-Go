Requirements
Backend:
Node.js and npm installed on your machine.
XAMPP or any other local server to initialize the MySQL database.
Frontend:
React (set up using create-react-app).
Any modern browser for accessing the frontend.
Installation and Setup
1. Initialize the Database
Start XAMPP and enable Apache and MySQL services.
Open phpMyAdmin in your browser.
Create a new database called crypto_wallet.
Execute the following SQL script to create the wallets table:
sql
Copy code
CREATE TABLE wallets (
  id INT AUTO_INCREMENT PRIMARY KEY,
  username VARCHAR(255) NOT NULL,
  address VARCHAR(255) NOT NULL,
  balance DECIMAL(10, 2) NOT NULL,
  currency VARCHAR(50) NOT NULL
);
2. Run the Backend
Navigate to the backend folder:
bash
Copy code
cd backend
Install the dependencies:
bash
Copy code
npm install
Configure the database connection:
Open the config.js file (or similar, depending on your setup).
Update the database settings to match your local MySQL configuration:
javascript
Copy code
module.exports = {
  host: 'localhost',
  user: 'root',
  password: '', // Your MySQL root password
  database: 'crypto_wallet',
};
Start the backend server:
bash
Copy code
npm start
The backend server should be running at http://localhost:3000.
3. Run the Frontend
Navigate to the frontend folder:
bash
Copy code
cd frontend
Install the dependencies:
bash
Copy code
npm install
Start the React development server:
bash
Copy code
npm start
The frontend application should open automatically in your browser at http://localhost:3001. If not, open the URL manually.
How to Use the Application
Enter Username: On the homepage, enter your username to fetch or manage wallets associated with your account.
Create Wallet: Use the "Create Wallet" form to add a new wallet.
View Wallets: All wallets associated with the entered username will be displayed.
Edit Wallet: Click on "Edit Wallet" to modify wallet details.
Delete Wallet: Use the delete button to remove a wallet.
Troubleshooting
If the backend does not connect to the database, ensure your MySQL server is running and the credentials in config.js are correct.
Ensure both backend and frontend servers are running simultaneously.
If there are CORS issues, you may need to enable CORS on the backend by adding a package like cors:
bash
Copy code
npm install cors
And update your backend server code:
javascript
Copy code
const cors = require('cors');
app.use(cors());
Tech Stack
Frontend: React
Backend: Node.js, Express.js
Database: MySQL (via XAMPP)
