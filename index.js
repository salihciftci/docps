const express = require("express");
const path = require("path");
const logger = require("morgan");

const auth = require("./routes/middlewares/authenticate");

const usersRouter = require("./routes/api/users");
const containersRouter = require("./routes/api/containers");

const app = express();

app.use(logger("dev"));
app.use(express.json());
app.use(express.urlencoded({ extended: false }));

// Serving React files
app.use(express.static(path.join(__dirname, "build"))); //todo

// Using middleware
// app.use(auth);

// API Routes
app.use("/api/users", usersRouter);
app.use("/api/containers", containersRouter);

module.exports = app;