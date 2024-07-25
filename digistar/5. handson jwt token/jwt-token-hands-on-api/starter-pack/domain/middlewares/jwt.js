// npm install jsonwebtoken

const jwt = require('jsonwebtoken'); // Ensure you have required the jwt module

const verifyToken = (req, res, next) => {
    const token = req.headers.authorization?.split(' ')[1]; // Assumes "Bearer TOKEN" format

    if (!token) {
        return res.status(401).json({ message: 'Token is required', error: 'Unauthorized' });
    }

    jwt.verify(token, process.env.JWT_SECRET, (err, decoded) => {
        if (err) {
            // Check if the error is because the token has expired
            if (err.name === 'TokenExpiredError') {
                return res.status(401).json({ message: 'Token has expired!' });
            }
            return res.status(403).json({ message: 'Failed to authenticate token' });
        }
        req.user = decoded;
        next();
    });
};

module.exports = verifyToken;