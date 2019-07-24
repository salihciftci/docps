const express = require("express");
const router = express.Router();
const bcrypt = require("bcrypt");
const path = require("path");
const fs = require("fs");
const mkdirp = require("mkdirp");

const db = require("../../db");
const knex = db.knex;

let error = "";

router.get("/", (req, res) => {
    try {
        res.render("install", {
            title: "Install",
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
        console.log("Trying to install Liman");
        let username = req.body.username;
        let password = req.body.password;
        let email = req.body.email;

        if (!username || !password) {
            throw new Error;
        }

        if (!email) {
            email = "example@example.com";
        }

        let encrypted = bcrypt.hashSync(password, 10);

        let installPath = path.join(__dirname, "../../data");
        if (!fs.existsSync(installPath)) {
            mkdirp.sync(installPath, 0o777);

            knex.schema.createTable("users", (table) => {
                table.increments().primary();
                table.string("username").unique();
                table.string("password");
                table.string("email");
                table.boolean("admin");
                table.timestamps();
            }).then().catch((e) => {
                console.log(e);
            });

            knex("users").insert([{
                "username": username,
                "password": encrypted,
                "email": email,
                "admin": false,
                "created_at": knex.fn.now(),
                "updated_at": knex.fn.now()
            }]).then().catch((e) => {
                console.log(e);
            });

            console.log("Liman succesfully installed");
            res.status(301).redirect("/login");
        }
    } catch (e) {
        console.log(e);
    }
});

module.exports = router;