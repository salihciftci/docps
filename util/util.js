module.exports.generateSecret = (len) => {
    let chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789";
    let secret = "";
    for (let i = 0; i < len; i++) {
        secret = secret + chars.charAt(Math.floor(Math.random() * chars.length));
    }
    return secret;
};