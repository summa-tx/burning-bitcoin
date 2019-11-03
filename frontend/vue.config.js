const path = require('path')

module.exports = {
  "transpileDependencies": [
    "vuetify"
  ],
  chainWebpack: config => {
    config.module
      .rule('js')
      .test(/\.js$/)
      .exclude
        .add(/bcoinstuff\/BcoinClient.js/)
        .end()
      .use('babel-loader')
      .loader('babel-loader')
  }
}
