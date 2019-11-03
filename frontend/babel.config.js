module.exports = {
  presets: [
    '@vue/app',
    "@babel/preset-env"
  ],
  "plugins": ["dynamic-import-node", "@babel/plugin-transform-modules-commonjs"]
}
