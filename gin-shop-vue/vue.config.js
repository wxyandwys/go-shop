module.exports = {
  devServer: {
    host: '0.0.0.0',
    port: 8080,
    disableHostCheck: true,
    proxy: {
      '/*': {
            target: 'http://127.0.0.1',
            secure: false,
            changeOrigin: true
      }
    }
  }
}