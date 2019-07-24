const express = require("express");
const path = require("path");
const logger = require("morgan");
const hbs = require("express-hbs");
const cookieParser = require("cookie-parser");

const header = require("./routes/middlewares/header");
const cookie = require("./routes/middlewares/cookie");
const install = require("./routes/middlewares/install");

const app = express();

app.use(logger("dev")); //todo change in production
app.use(express.json());
app.use(express.urlencoded({ extended: false }));
app.use(cookieParser());

// Handlebars configs
app.set("view engine", "tmpl");
app.set("views", __dirname + "/templates");
app.engine("tmpl", hbs.express4({
    partialsDir: __dirname + "/templates/partials/",
    extname: ".tmpl"
}));

// Serving Static files
app.use(express.static(path.join(__dirname, "public")));

// Using middlewares
app.use(header, cookie, install);

// API Routes
app.use("/api/users", require("./routes/api/users"));
app.use("/api/containers", require("./routes/api/containers"));
app.use("/api/images/", require("./routes/api/images"));
app.use("/api/volumes/", require("./routes/api/volumes"));
app.use("/api/networks/", require("./routes/api/networks"));
app.use("/api/", require("./routes/api/index"));

// SSR Routes
app.use("/", require("./routes/web/index"));
app.use("/install", require("./routes/web/install"));
app.use("/login", require("./routes/web/login"));

app.use("/logout", (req, res) => { res.clearCookie("liman"); res.redirect("/login"); });

module.exports = app;