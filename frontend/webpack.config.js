const path = require('path');
const NotifierPlugin = require('webpack-notifier');

module.exports = function(env) {
    return {

        entry: './src/main.tsx',

        output: {
            filename: 'build.js',
            path: path.join(__dirname, '/web/build'),
            library: 'moneypenny'
        },

        module: {
            rules: [
                {
                    loader: 'ts-loader',
                    test: /\.(js|ts|tsx)$/,
                    exclude: /node_modules/
                },
                {
                    loader: 'tslint-loader',
                    test: /\.(ts|tsx)$/,
                    exclude: /node_modules/,
                    options: {
                        emitErrors: true,
                        typeCheck: true
                    }
                },
                {
                    loader: 'raw-loader',
                    test: /\.(fsh|vsh)$/,
                    exclude: /node_modules/
                }
            ]
        },

        resolve: {
            extensions: ['.tsx', '.ts', '.js']
        },

        devtool: env === 'prod' ? 'source-map' : 'inline-source-map',

        plugins: [
            new NotifierPlugin({
                title: 'Moneypenny'
            })
        ]
    }
}
