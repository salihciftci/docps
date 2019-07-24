const jwt = require("jsonwebtoken");
const os = require("os");
const uuid = require("uuid/v5");

module.exports = (req, res, next) => {
    let path = req.path.split("/")[1];
    if (path === "api" || path === "install" || path === "login") {
        next();
        return;
    }

    let cookie = req.cookies.liman;
    if (typeof cookie !== "undefined") {
        try {
            let decoded = jwt.verify(cookie, uuid(os.hostname(), uuid.DNS));
            if (!decoded) {
                console.log("Login Attemp: Invalid cookie");
                res.redirect("/login");
                next();
                return;
            }
            req.liman = decoded;
            next();
        } catch (e) {
            if (e.name === "TokenExpiredError") {
                console.log("Login Attemp: Token Expired");
            } else if (e.name === "JsonWebTokenError") {
                console.log("Login Attemp: JsonWebTokenError");
            } else if (e.name === "NotBeforeError") {
                console.log("Login Attemp: NotBeforeError");
            } else {
                console.log(e);
            }
            res.redirect("/login");
        }
    } else {
        res.redirect("/login");
    }
};