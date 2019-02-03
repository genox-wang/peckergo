const path = require('path');

function resolve (dir) {
  return path.join(__dirname, dir);
}

module.exports = {
  chainWebpack: config => {
    config.module
      .rule('vue')
      .use('iview-loader')
      .loader('iview-loader')
      .options({
        prefix: false
      });

    config.resolve.alias
      .set('@', resolve('/src'));
  }
}
;
