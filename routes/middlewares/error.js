module.exports = (err, req, res, next) => {
    if (err) {
        if (err instanceof SyntaxError) {
            console.log("Bad JSON syntax");
            res.sendStatus(400);
        } else {
            console.log(err);
        }
    } else {
        next();
    }
};