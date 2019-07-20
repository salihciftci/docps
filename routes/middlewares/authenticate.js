let jwt = require("jsonwebtoken");
let os = require("os");
let uuid = require("uuid/v5");

module.exports = (req, res, next) => {
    console.log(req.path);
    if (req.method === "POST") {
        let split = req.path.split("/");
        if (split.length > 3 && split[2] === "users") {
            if (split[3]) {
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