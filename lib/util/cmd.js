const util = require("util");
const exec = util.promisify(require("child_process").exec);

module.exports.execCommand = async (command) => {
    const { stdout } = await exec(command);
    let lines = stdout.split("\n");
    lines.pop();
    return lines;
};