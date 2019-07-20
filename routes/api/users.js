const bcrypt = require("bcrypt");
const express = require("express");
const router = express.Router();
const jwt = require("jsonwebtoken");
const os = require("os");
const uuid = require("uuid/v5");

const db = require("../../db/index");
const knex = db.knex;

// POST create user
router.post("/", async (req, res) => {
    try {
        let username = req.body.username;
        let password = req.body.password;

        if (!username || !password) {
            console.log("username or password not included in body");
            res.status(400);
            return;
        }

        let encryped = bcrypt.hashSync(password, 10);

        await knex("users").insert([{
            "username": username,
            "password": encryped,
            "created_at": knex.fn.now(),
            "updated_at": knex.fn.now()
        }]);

        res.sendStatus(200);
    } catch (e) {
        if (e.errno = 19) {
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
        let user = req.params.username;

        let result = await knex("users").count("username as count").where("username", user);
        let count = result[0].count;

        if (count !== 1) {
            console.log("User not found");
            res.sendStatus(404);
            return;
        }

        result = await knex.select("password").from("users").where("username", user);

        let match = bcrypt.compareSync(password, result[0].password);
        if (!match) {
            console.log("Passwords are not match");
            res.sendStatus(404);
            return;
        }

        let token = jwt.sign({}, uuid(os.hostname(), uuid.DNS), { expiresIn: "1w" });
        res.json({ "token": token });
    } catch (e) {
        console.log(e);
        res.sendStatus(500);
    }
});

// GET user info
router.get("/:username", async (req, res) => {

});

module.exports = router;