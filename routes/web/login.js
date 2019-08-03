const express = require("express");
const router = express.Router();
const bcrypt = require("bcryptjs");
const jwt = require("jsonwebtoken");
const fs = require("fs");
const path = require("path");

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
        res.render("500", { title: "Error" });
    }
});

router.post("/", async (req, res) => {
    try {
        let password = req.body.password;
        let username = req.body.username;

        if (!password || !username) {
            console.log("Login Attemp: Username or password not included in body");
            error = "Invalid email or password";
            res.redirect("/login");
            return;
        }

        let result = await knex.select("password").from("users").where("username", username);

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

        result = await knex.select("admin").select("email").select("avatarURL").from("users").where("username", username);

        let user = {
            "username": username,
            "email": result[0].email,
            "admin": result[0].admin,
            "avatarURL": "https://www.gravatar.com/avatar/" + result[0].avatarURL
        };

        const privateKey = fs.readFileSync(path.join(__dirname, "../../data/keys/private.pem"));
        let token = jwt.sign({ user }, privateKey, { expiresIn: "1w", algorithm: "RS256" });
        res.cookie("liman", token, { "path": "/", maxAge: "604800000" }); // 1 week 
        res.redirect("/");
    } catch (e) {
        console.log(e);
        res.render("500", { title: "Error" });
    }
});

module.exports = router;