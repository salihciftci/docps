let jwt = require("jsonwebtoken");
const path = require("path");
const fs = require("fs");

module.exports = (req, res, next) => {
    let paths = req.path.split("/");
    if (paths[1] !== "api") {
        next();
        return;
    }

    if (req.method === "POST") {
        if (paths.length > 3 && paths[2] === "users") {
            if (paths[3]) {
                next();
                return;
            }
        }
    }

    const bearerHeader = req.headers["authorization"];
    let token;
    if (typeof bearerHeader !== "undefined") {
        try {
            let privateKey;
            if (process.env.NODE_ENV === "test") {
                privateKey = fs.readFileSync(path.join(__dirname, "../../test-data/keys/private.pem"));
            } else {
                privateKey = fs.readFileSync(path.join(__dirname, "../../data/keys/private.pem"));
            }
            const bearer = bearerHeader.split(" ");
            token = bearer[1];

            let decoded = jwt.verify(token, privateKey, { algorithms: "RS256" });
            if (!decoded) {
                res.sendStatus(403);
            }
            req.user = decoded.user;
        } catch (e) {
            if (e.name === "TokenExpiredError") {
                res.status(401).json({ "error": e.message });
                return;
            } else if (e.name === "JsonWebTokenError") {
                res.status(401).json({ "error": e.message });
                return;
            } else if (e.name === "NotBeforeError") {
                res.status(401).json({ "error": e.message });
                return;
            } else if (e.errno === -2) {
                res.status(501).json({ "error": "Liman is not installed yet" });
                return;
            }
            console.log(e);
            res.sendStatus(500);
            return;
        }
        next();
    } else {
        res.status(401).json({ "error": "token required" });
    }
};