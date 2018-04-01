const webpack = require('webpack');
const path = require('path');

/*
 * We've enabled UglifyJSPlugin for you! This minifies your app
 * in order to load faster and run less javascript.
 *
 * https://github.com/webpack-contrib/uglifyjs-webpack-plugin
 *
 */

const UglifyJSPlugin = require('uglifyjs-webpack-plugin');

/*
 * We've enabled ExtractTextPlugin for you. This allows your app to
 * use css modules that will be moved into a separate CSS file instead of inside
 * one of your module entries!
 *
 * https://github.com/webpack-contrib/extract-text-webpack-plugin
 *
 */

// BUG: https://github.com/webpack-contrib/extract-text-webpack-plugin/issues/731
// const ExtractTextPlugin = require('extract-text-webpack-plugin');

const MiniCssExtractPlugin = require('mini-css-extract-plugin');

const isDevMode = process.env.NODE_ENV !== 'production';
const devPort = process.env.CLIENT_PORT || 8080;
const apiPort = process.env.PORT || 3000;
const devPublicDomain = "http://localhost:" + devPort;
const apiPublicDomain = "http://localhost:" + apiPort;
const devPublicFullPath = devPublicDomain + "/dist/"

module.exports = {

    mode: isDevMode ? 'development' : 'production',

    devtool: isDevMode ? "source-map" : false,

    devServer: {
	publicPath: devPublicFullPath,
	port: devPort,
	hot: true,
	proxy: {
	    '/': {
		target: apiPublicDomain,
		// bug: required to work for Mac OSX
		pathRewrite: { '^/': '/' },
	    }
	},
    },

    entry: "./client/index.jsx",

    output: { filename: '[name].bundle.js',
	path: path.resolve(__dirname, 'dist')
    },

    resolve: {
	extensions: ['.js', '.jsx'],
	alias: {
	    react: path.resolve('./node_modules/react'),
	    'react-dom': path.resolve('./node_modules/react-dom'),
	},
    },

    module: {
	rules: [
	    {
		test: require.resolve('react'),
		loader: 'imports-loader?shim=es5-shim/es5-shim&sham=es5-shim/es5-sham',
	    },
	    {
		test: /\.(js|jsx)$/,
		exclude: /node_modules/,
		loader: 'babel-loader',
	    },
	    {
		test: /\.(less|css)$/,
		use: [
		    MiniCssExtractPlugin.loader,
		    'css-loader',
		    'less-loader'
		]

		// use: ExtractTextPlugin.extract({
		// 	use: [
		// 		{
		// 			loader: 'css-loader',
		// 			options: {
		// 				sourceMap: true
		// 			}
		// 		},
		// 		{
		// 			loader: 'less-loader',
		// 			options: {
		// 				sourceMap: true
		// 			}
		// 		}
		// 	],
		// 	fallback: 'style-loader'
		// })
	    }
	]
    },

    plugins: [
	new webpack.HotModuleReplacementPlugin(),
	new MiniCssExtractPlugin({
	    filename: "[name].css",
	    chunkFilename: "[id].css"
	})
    ]
};

