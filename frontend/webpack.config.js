const path = require('path');
const NotifierPlugin = require('webpack-notifier');
const CleanWebpackPlugin = require('clean-webpack-plugin');
const webpack = require('webpack');
const MinifyPlugin = require('babel-minify-webpack-plugin');

module.exports = function(env) {
    const builddir = path.join(__dirname, '/web/build'),
        isProduction = typeof env === 'object' && env.prod,
        buildForEs5 = typeof env === 'object' && env.es5;

    return {
        entry: ['./src/main.tsx', ...(buildForEs5 ? [require.resolve('core-js/es6')] : [])],

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
                    exclude: /node_modules/,
                    options: {
                        compilerOptions: buildForEs5
                            ? {
                                  target: 'es5',
                                  lib: ['es6', 'dom'],
                                  downlevelIteration: true
                              }
                            : {}
                    }
                },
                {
                    loader: 'tslint-loader',
                    test: /\.(ts|tsx)$/,
                    exclude: /node_modules/,
                    options: {
                        emitErrors: true,
                        typeCheck: true
                    }
                }
            ]
        },

        resolve: {
            extensions: ['.tsx', '.ts', '.js'],
            alias: {
                moment$: 'moment/moment.js'
            }
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
            new CleanWebpackPlugin([builddir], {
                verbose: true
            }),
            ...(isProduction
                ? [
                      new MinifyPlugin(
                          {},
                          {
                              comments: false
                          }
                      )
                  ]
                : [])
        ]
    };
};
