const express = require("express");
const path = require("path");
const logger = require("morgan");

const auth = require("./routes/middlewares/authenticate");

const app = express();

app.use(logger("dev"));
app.use(express.json());
app.use(express.urlencoded({ extended: false }));

// Serving React files
app.use(express.static(path.join(__dirname, "build"))); //todo

// Using middleware
app.use(auth);

// API Routes
app.use("/api/users", require("./routes/api/users"));
app.use("/api/containers", require("./routes/api/containers"));
app.use("/api/images/", require("./routes/api/images"));
app.use("/api/volumes/", require("./routes/api/volumes"));
app.use("/api/networks/", require("./routes/api/networks"));

module.exports = app;