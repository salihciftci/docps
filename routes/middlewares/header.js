let jwt = require("jsonwebtoken");
let os = require("os");
let uuid = require("uuid/v5");

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
        const bearer = bearerHeader.split(" ");
        token = bearer[1];

        try {
            let decoded = jwt.verify(token, uuid(os.hostname(), uuid.DNS));
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