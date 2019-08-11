const path = require('path');

module.exports = {
    publicPath: "/assets",
    outputDir: "../assets/",
    devServer: {
        disableHostCheck: true,
        proxy: {
            '^/api': {
                //target: 'http://10.8.0.1:9000',
                target: 'http://localhost:8080',
                ws: true
            },
            '^/ws': {
                //target: 'http://10.8.0.1:9000',
                target: 'http://localhost:8080',
                ws: true
            },
        },
        port: 9000
    },
    configureWebpack: {
        resolve: {
            alias: {
                '@': path.resolve('src')
            }
        }
    }
}