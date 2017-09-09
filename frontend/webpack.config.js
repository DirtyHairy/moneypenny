const path = require('path');
const NotifierPlugin = require('webpack-notifier');
const CleanWebpackPlugin = require('clean-webpack-plugin');
const webpack = require('webpack');
const MinifyPlugin = require('babel-minify-webpack-plugin');

module.exports = function(env) {
    const builddir = path.join(__dirname, '/web/build'),
        isProduction = env && env.indexOf('prod') === 0;

    return {

        entry: './src/main.tsx',

        output: {
            filename: 'build.js',
            path: builddir,
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

        devtool: isProduction ? false : 'eval-source-map',

        plugins: [
            new NotifierPlugin({
                title: 'Moneypenny',
                alwaysNotify: true
            }),
            new webpack.DefinePlugin({
                'process.env.NODE_ENV': isProduction ? "'production'" : "'development'"
            }),
            new CleanWebpackPlugin(
                [builddir],
                {
                    verbose: true
                }
            ),
            ...(isProduction ? [new MinifyPlugin(
                {},
                {
                    comments: false
                }
            )] : [])
        ]
    }
}
