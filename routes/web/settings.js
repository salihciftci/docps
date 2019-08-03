const express = require("express");
const router = express.Router();
const bcrypt = require("bcryptjs");
const { createHash } = require("crypto");

const db = require("../../db");
const knex = db.knex;

let error = "";

router.get("/", async (req, res) => {
    try {
        res.render("settings", {
            title: "Settings",
            user: req.user,
            error: error
        });
        error = "";
    } catch (e) {
        console.log(e);
        res.render("500", { title: "Error" });
    }
});

router.post("/profile", async (req, res) => {
    try {
        let username = req.body.username;
        let email = req.body.email;

        let result = -1;

        if (!username || !email) {
            error = "Username and email are required";
        }

        if (username !== req.user.username) {
            result = await knex("users").where("username", req.user.username).update({
                username: username,
                "updated_at": knex.fn.now()
            });
        }

        if (email !== req.user.email) {
            result = await knex("users").where("email", req.user.email).update({
                "avatarURL": createHash("md5").update(email).digest("hex"),
                email: email,
                "updated_at": knex.fn.now()
            });
        }

        if (result !== -1) {
            res.redirect("/logout");
            return;
        }

        console.log(result);
        res.redirect("/settings");

    } catch (e) {
        if (e.errno === 19) {
            error = "Username already exist";
            res.redirect("/settings");
            return;
        }
        console.log(e);
        res.render("500", { title: "Error" });
    }
});

router.post("/password", async (req, res) => {
    try {
        let password = req.body.password;
        let newPassword = req.body.newPassword;
        let confirmPassword = req.body.confirmPassword;

        if (newPassword !== confirmPassword) {
            error = "Passwords are not match";
            res.redirect("/settings");
            return;
        }

        let result = await knex("users").select("password").where("username", req.user.username);
        let p = result[0].password;
        let match = bcrypt.compareSync(password, p);

        if (!match) {
            error = "Invalid password";
            res.redirect("/settings");
        }

        let encrypted = bcrypt.hashSync(newPassword, 10);
        await knex("users").where("username", req.user.username).update({
            "password": encrypted,
            "updated_at": knex.fn.now()
        });

        res.redirect("/logout");
    } catch (e) {
        console.log(e);
        res.render("500", { title: "Error" });
    }
});

module.exports = router;