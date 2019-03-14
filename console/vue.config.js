const path = require('path');

function resolve (dir) {
  return path.join(__dirname, dir);
}

module.exports = {
  css: {
    loaderOptions: { // 向 CSS 相关的 loader 传递选项
      less: {
        javascriptEnabled: true
      }
    }
  },
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
