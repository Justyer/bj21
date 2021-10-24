const path = require("path");
const resolve = dir => path.join(__dirname, dir);

module.exports = {
  lintOnSave: true,
  chainWebpack: (config) => {
    config.resolve.alias.set("@", resolve("src"))
    config.plugin("html").tap((args) => {
      args[0].title = "BlackJack21";
      return args;
    })
  },
  pluginOptions: {
    electronBuilder: {
      nodeIntegration: true
    }
  }
};
