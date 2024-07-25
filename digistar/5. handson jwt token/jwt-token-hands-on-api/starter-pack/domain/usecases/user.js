// npm install mongoose uuid

const mongoose = require('mongoose');
const User = require('../models/user');
const bcrypt = require('bcrypt');
const jwt = require('jsonwebtoken');
const { v4: uuidv4 } = require('uuid');

// Connect to MongoDB
mongoose.connect(process.env.MONGO_DB_URL, {
  useNewUrlParser: true,
  useUnifiedTopology: true,
});

// Function to find a user by their ID
async function findOneByUserId(userId) {
  try {
    // Find the user by ID
    const user = await User.findOne({ user_id: userId });
    return user;
  } catch (error) {
    console.error('Error finding user by ID:', error);
    throw error;
  }
}

// Function to find a user by their email
async function findOneByEmail(email) {
  try {
    // Find the user by email
    const user = await User.findOne({ email: email });
    return user;
  } catch (error) {
    console.error('Error finding user:', error);
    throw error;
  }
}

// Function to find all users
async function findAll() {
  try {
    // Find all users
    const users = await User.find();
    return users;
  } catch (error) {
    console.error('Error finding users:', error);
    throw error;
  }
}

// Function to register a new user
async function register(user) {
  try {

    // Generate a unique user_id
    user.user_id = uuidv4();

    // Hash password
    const hashedPassword = await bcrypt.hash(user.password, 10); // 10 is the saltRounds
    user.password = hashedPassword;

    // Create a new user
    const newUser = new User(user);

    // Save the user to the database
    const savedUser = await newUser.save();
    return savedUser;
  } catch (error) {
    console.error('Error registering user:', error);
    throw error;
  }
}

async function login(payload) {
  try {
    const checkUser = await findOneByEmail(payload.email);
    if (!checkUser) {
      throw new Error('Invalid email or password');
    }
    const user = {
      userId: checkUser.user_id,
      email: checkUser.email,
      password: checkUser.password
    };
    const isValidPassword = await bcrypt.compareSync(payload.password, user.password);
    if (!isValidPassword) {
      throw new Error('Invalid email or password');
    }
    const key = process.env.JWT_SECRET; // Use environment variable for the secret key or a default one
    const token = jwt.sign(user, key, { expiresIn: '15m' }); // Set token expiration to 15 minutes
    return token;
  } catch (error) {
    console.error('Error login: ', error);
    throw error;
  }
}

module.exports = { findOneByUserId, findOneByEmail, findAll, register, login };