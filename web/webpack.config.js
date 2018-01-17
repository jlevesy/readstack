const path = require('path');
const HtmlWebpackPlugin = require('html-webpack-plugin');
const CleanWebpackPlugin = require('clean-webpack-plugin');
const process = require('process');

const defaultAppUrl = 'http://app:8080'

module.exports = {
  entry: './src/index.js',
  plugins: [
    new CleanWebpackPlugin(['dist']),
    new HtmlWebpackPlugin({
      title: 'Readstack'
    })
  ],
  devtool: 'inline-source-map',
  devServer: {
    port: 8081,
    host: "0.0.0.0",
    contentBase: './dist',
    proxy: {
      "/api/v1": process.env['READSTACK_BACKEND_URL'] || defaultAppUrl
    }
  },
  output: {
    filename: 'bundle.js',
    path: path.resolve(__dirname, 'dist')
  }
};

