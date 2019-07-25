let knex = require("knex");
let path = require("path");

const sqlPath = path.join(__dirname, "../data");

let knexInstance = knex({
    client: "sqlite3",
    connection: {
        filename: sqlPath + "/db/liman.sqlite3"
    },
    "useNullAsDefault": true
});

module.exports.knex = knexInstance;