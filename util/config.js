var util = require("./util");

module.exports.secretKey = util.generateSecret(512);