// npm install express bcrypt dotenv

const loadEnv = require('./config/env');
// Call the function to load environment variables
loadEnv();
const express = require("express");
const usecases = require('./domain/usecases/user');
const jwtAuth = require('./domain/middlewares/jwt');
const passport = require('passport');
const { initializePassport, authenticatePassportJwt } = require('./domain/middlewares/passport-jwt');

const PORT = 7000;
const app = express();

// Middleware to parse JSON bodies (for Express 4.16.0 and above)
app.use(express.json());
// Initialize Passport
app.use(initializePassport());

// Create a router
const router = express.Router();

// Mount the router on the app under the /digistar base path
app.use('/digistar', router);

// Login API endpoint
router.post("/login", async (req, res, next) => {
  try {
    const { email, password } = req.body;
    const payload = {
      email,
      password
    };
    const token = await usecases.login(payload);
    res.status(200).json({ message: "Success login!", token });
  } catch (err) {
    res.status(500).json({ error: 'Internal Server Error', message: err.message});
  }
});

// Registration API endpoint
router.post("/register", async (req, res, next) => {
  try {

    // Validate request data (Implement validation logic as needed)
    const { username, email, password } = req.body;

    // Check if user already exists
    const existingUser = await usecases.findOneByEmail(email);
    if (existingUser) {
      return res.status(409).json({ message: "User already exists" });
    }
    
    if (!username || !email || !password) {
      return res.status(400).json({ message: "Username, Email, and password are required" });
    }

    const user = {
      username,
      email,
      password
    };

    // Create user
    const newUser = await usecases.register(user)

    // Respond with success message (Consider returning the created user ID or similar)
    res.status(201).json({ message: "User created successfully", userId: newUser.id });
  } catch (err) {
    res.status(500).json({ error: 'Internal Server Error', message: err.message});
  }
});

// Get all users API endpoint with JWT authentication
router.get("", jwtAuth, async (req, res, next) => {
  try {
    // Retrieve all users from the usecasessitory
    const users = await usecases.findAll();

    // Respond with the users array
    res.json(users);
  } catch (err) {
    res.status(500).json({ error: 'Internal Server Error', message: err.message});
  }
});

// Get user by ID API endpoint with JWT authentication
router.get("/:id", authenticatePassportJwt(), async (req, res, next) => {
  try {
    const userId = req.params.id;
    const user = await usecases.findOneByUserId(userId);
    if (!user) {
      return res.status(404).json({ message: "User not found" });
    }
    res.json(user);
  } catch (err) {
    res.status(404).json({ error: 'Error Not Found', message: err.message});
  }
});

// Error handling middleware
app.use((err, req, res, next) => {
  console.error(err.stack);
  res.status(500).send('Something went wrong!');
});

// Start the server
app.listen(PORT, () => {
  console.log(`Server is running on port ${PORT}`);
});