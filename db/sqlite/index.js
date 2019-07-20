const sqlite3 = require("sqlite3").verbose();
const path = require("path");

const sqlPath = path.join(__dirname, "../../data");

module.exports.install = (user, pass) => {
    return new Promise((resolve, reject) => {
        let db = new sqlite3.Database(sqlPath + "/liman.db");
        db.serialize(function () {
            db.run(`
                CREATE TABLE IF NOT EXISTS users
                (
                    id   INTEGER PRIMARY KEY AUTOINCREMENT,
                    user TEXT,
                    pass TEXT
                )
            `, (err) => {
                if (err) {
                    reject(err);
                }
            });

            let stmt = db.prepare("INSERT INTO users(user, pass) VALUES (?, ?)");
            stmt.run([user, pass], (err) => {
                if (err) {
                    reject(err);
                }
            });
            stmt.finalize();
        });
        db.close();
        resolve();
    });
};

module.exports.query = (query, params) => {
    let db = new sqlite3.Database(sqlPath + "/liman.db");
    return new Promise((resolve, reject) => {
        db.all(query, params, (err, rows) => {
            if (err) reject(err);
            resolve(rows);
        });
        db.close();
    });
};