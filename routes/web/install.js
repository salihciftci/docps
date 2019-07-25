const express = require("express");
const router = express.Router();
const bcrypt = require("bcrypt");
const path = require("path");
const fs = require("fs");
const mkdirp = require("mkdirp");
const { generateKeyPairSync } = require("crypto");

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
            email = "example@example.com"; //todo fix in production
        }

        let encrypted = bcrypt.hashSync(password, 10);

        // Generating SQLite Database
        let installPath = path.join(__dirname, "../../data");
        if (!fs.existsSync(installPath + "/db")) {
            mkdirp.sync(installPath + "/db", 0o777);

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
                "admin": true,
                "created_at": knex.fn.now(),
                "updated_at": knex.fn.now()
            }]).then().catch((e) => {
                console.log(e);
            });
            console.log("Database succesfully created");
        }

        // Generating RSA Keys
        if (!fs.existsSync(installPath + "/keys")) {
            mkdirp.sync(installPath + "/keys", 0o777);
            const { publicKey, privateKey } = generateKeyPairSync("rsa", {
                modulusLength: 4096,
                publicKeyEncoding: {
                    type: "spki",
                    format: "pem"
                },
                privateKeyEncoding: {
                    type: "pkcs8",
                    format: "pem",
                }
            });

            fs.writeFileSync(installPath + "/keys/private.pem", privateKey);
            fs.writeFileSync(installPath + "/keys/public.pem", publicKey);
            console.log("RSA keys succesfully created");
        }

        console.log("Liman succesfully installed");
        res.status(301).redirect("/login");
    } catch (e) {
        console.log(e);
    }
});

module.exports = router;