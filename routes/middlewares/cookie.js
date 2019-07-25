const jwt = require("jsonwebtoken");
const nodePath = require("path");
const fs = require("fs");

module.exports = (req, res, next) => {
    let path = req.path.split("/")[1];
    if (path === "api" || path === "install" || path === "login") {
        next();
        return;
    }

    let cookie = req.cookies.liman;
    if (typeof cookie !== "undefined") {
        try {
            let privateKey = fs.readFileSync(nodePath.join(__dirname, "../../data/keys/private.pem"));
            let decoded = jwt.verify(cookie, privateKey, { algorithms: "RS256" });
            if (!decoded) {
                console.log("Login Attemp: Invalid cookie");
                res.redirect("/login");
                next();
                return;
            }
            req.user = decoded.user;
            next();
        } catch (e) {
            if (e.name === "TokenExpiredError") {
                console.log("Login Attemp: Token Expired");
            } else if (e.name === "JsonWebTokenError") {
                console.log("Login Attemp: JsonWebTokenError");
            } else if (e.name === "NotBeforeError") {
                console.log("Login Attemp: NotBeforeError");
            } else if (e.errno === -2) {
                console.log("Liman is not installed yet. Redirected to /install");
            } else {
                console.log(e);
            }
            res.redirect("/login");
        }
    } else {
        res.redirect("/login");
    }
};