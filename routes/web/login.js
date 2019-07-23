const express = require("express");
const router = express.Router();
const bcrypt = require("bcrypt");
const jwt = require("jsonwebtoken");
const uuid = require("uuid/v5");
const os = require("os");

const db = require("../../db");
const knex = db.knex;

let error = "";

router.get("/", (req, res) => {
    try {
        res.render("login", {
            title: "Login",
            error: error
        });

        error = "";
    } catch (e) {
        console.log(e);
        res.sendStatus(500);
    }
});

router.post("/", async (req, res) => {
    try {
        let password = req.body.password;
        let email = req.body.email;

        if (!password || !email) {
            console.log("Login Attemp: Username or password not included in body");
            error = "Invalid email or password";
            res.redirect("/login");
            return;
        }

        let result = await knex.select("password").from("users").where("username", email);

        if (!result.length) {
            console.log("Login Attemp: User not found");
            error = "Invalid email or password";
            res.redirect("/login");
            return;
        }

        let dbPass = result[0].password;
        let match = bcrypt.compareSync(password, dbPass);

        if (!match) {
            console.log("Login Attemp: Invalid password");
            error = "Invalid email or password";
            res.redirect("/login");
            return;
        }

        let token = jwt.sign({}, uuid(os.hostname(), uuid.DNS), { expiresIn: "1w" });
        res.cookie("liman", token, { "path": "/", maxAge: "999999999999999" }); // todo: fix maxAge
        res.redirect("/");
    } catch (e) {
        console.log(e);
    }
});

module.exports = router;