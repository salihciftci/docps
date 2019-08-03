const bcrypt = require("bcryptjs");
const express = require("express");
const router = express.Router();
const jwt = require("jsonwebtoken");
const fs = require("fs");
const path = require("path");
const { createHash } = require("crypto");

const db = require("../../db/index");
const knex = db.knex;

// POST create user
router.post("/", async (req, res) => {
    try {
        let username = req.body.username;
        let password = req.body.password;
        let email = req.body.email || "example@example.com"; //todo fix here

        if (!username || !password) {
            console.log("username or password not included in body");
            res.status(400);
            return;
        }

        let encryped = bcrypt.hashSync(password, 10);

        await knex("users").insert([{
            "username": username,
            "password": encryped,
            "email": email,
            "admin": true,
            "avatarURL": createHash("md5").update(email).digest("hex"),
            "created_at": knex.fn.now(),
            "updated_at": knex.fn.now()
        }]);

        res.sendStatus(200);
    } catch (e) {
        if (e.errno === 19) {
            console.log("user already exist");
            res.sendStatus(409);
            return;
        }
        console.log(e);
        res.sendStatus(500);
    }
});

// POST Login 
router.post("/:username", async (req, res) => {
    try {
        let password = req.body.password;
        let username = req.params.username;

        let result = await knex("users").count("username as count").where("username", username);
        let count = result[0].count;

        if (count !== 1) {
            console.log("User not found");
            res.sendStatus(404);
            return;
        }

        result = await knex.select("password").from("users").where("username", username);

        let match = bcrypt.compareSync(password, result[0].password);
        if (!match) {
            console.log("Passwords are not match");
            res.sendStatus(404);
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
        res.json({ "token": token });
    } catch (e) {
        console.log(e);
        res.sendStatus(500);
    }
});

// GET user info
router.get("/:username", async (req, res) => {
    res.json(req.user);
});

module.exports = router;