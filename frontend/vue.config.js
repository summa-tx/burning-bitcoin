const path = require('path')

module.exports = {
  "transpileDependencies": [
    "vuetify"
  ],
  chainWebpack: config => {
    config.module
      .rule('bcoin')
      .test(/\.js$/)
      .exclude
        .add(path.resolve(__dirname + '../utils'))
        .end()
      .use('babel-loader')
      .loader('babel-loader')
  }
}
