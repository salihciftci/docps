const bcrypt = require("bcrypt");
const express = require("express");
const router = express.Router();
const jwt = require("jsonwebtoken");

const db = require("../../db/sqlite");
let { secretKey } = require("../../util/config");

// POST create user
router.post("/", async (req, res) => {
    //todo check user already exist or not
    try {
        let username = req.body.username;
        let password = req.body.password;

        if (!username || !password) {
            console.log("username or password not included in body");
            res.status(400);
            return;
        }

        let encryped = bcrypt.hashSync(password, 10);

        let params = [username, encryped];
        let query = "INSERT INTO users(user, pass) VALUES (?,?)";

        await db.query(query, params);

        res.sendStatus(200);
    } catch (e) {
        console.log(e);
        res.sendStatus(500);
    }
});

// POST Login 
router.post("/:username", async (req, res) => {
    try {
        let password = req.body.password;
        let user = req.params.username;

        let query = "SELECT COUNT(user) AS 'count' FROM users WHERE user = ?";
        let params = [user];
        let result = await db.query(query, params);

        if (result[0].count !== 1) {
            console.log("User not found");
            res.sendStatus(404);
            return;
        }

        query = "SELECT pass FROM users WHERE user = ?";
        result = await db.query(query, params);


        let match = bcrypt.compareSync(password, result[0].pass);
        if (!match) {
            console.log("Passwords are not match");
            res.status(404);
            return;
        }

        let token = jwt.sign({}, secretKey, { expiresIn: "1w" });
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