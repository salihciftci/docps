let knex = require("knex");
let path = require("path");

let sqlPath = "";

if (process.env.NODE_ENV === "test") {
    sqlPath = path.join(__dirname, "../test-data");
} else {
    sqlPath = path.join(__dirname, "../data");
}

let knexInstance = knex({
    client: "sqlite3",
    connection: {
        filename: path.join(sqlPath, "/db/liman.sqlite3")
    },
    "useNullAsDefault": true
});

module.exports.knex = knexInstance;