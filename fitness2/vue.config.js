module.exports = {
  devServer: {
    proxy: {
      '^/api': {
        target: 'http://localhost:2379',
        changeOrigin: true,
        logLevel: 'debug',
        pathRewrite: { '^/api/': '/' },
      },
    },
  },
}