const path = require("path");
const fs = require("fs");

const installPath = path.join(__dirname, "../../data/db");

module.exports = (req, res, next) => {
    try {
        let path = req.path.split("/")[1];
        if (fs.existsSync(installPath)) {
            if (path === "install") {
                res.status(301).redirect("/");
                return;
            }
            next();
            return;
        }

        if (path === "install") {
            next();
            return;
        }

        res.status(301).redirect("/install");
    } catch (e) {
        console.log(e);
    }
};