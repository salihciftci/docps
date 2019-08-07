//During the test the env variable is set to test
process.env.NODE_ENV = "test";

const mkdirp = require("mkdirp");
const path = require("path");
const fs = require("fs");
const { generateKeyPairSync } = require("crypto");

const knex = require("../db").knex;

before(async () => {
    try {
        let installPath = path.join(__dirname, "../test-data");
        if (!fs.existsSync(installPath)) {
            mkdirp.sync(path.join(installPath, "/db"));
            await knex.schema.createTable("users", (table) => {
                table.increments().primary();
                table.string("username").unique();
                table.string("password");
                table.string("email");
                table.string("avatarURL");
                table.boolean("admin");
                table.timestamps();
            });

            await knex("users").insert([{
                "username": "test",
                "password": "$2y$10$Y0BjJwJDSEXZE7SvTpYsJe5JEzGqvU/uoAgj7bnZAUhXWTVV4I54e", //test
                "email": "test@test.com",
                "admin": true,
                "avatarURL": "test url",
                "created_at": knex.fn.now(),
                "updated_at": knex.fn.now()
            }]);

            // Generating RSA Keys
            mkdirp.sync(path.join(installPath + "/keys"));
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

            fs.writeFileSync(path.join(installPath, "/keys/private.pem"), privateKey);
            fs.writeFileSync(path.join(installPath, "/keys/public.pem"), publicKey);
        }
    } catch (e) {
        console.log(e);
    }
});

require("./users");
require("./containers");
require("./images");
require("./volumes");
require("./networks");
require("./stats");