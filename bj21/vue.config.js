module.exports = {
  chainWebpack: (config) => {
    config.plugin("html").tap((args) => {
      args[0].title = "BlackJack21";
      return args;
    });
  },
};
