const path = require('path');
const HtmlWebpackPlugin = require('html-webpack-plugin');

module.exports = {
  entry: path.resolve(__dirname, 'src/index'),
  output: {
    path: path.resolve(__dirname, 'build'),
    filename: 'bundle.js'
  },
  module: {
    rules: [{
      test: /\.js$/,
      include: path.resolve(__dirname, 'src'),
      exclude: path.resolve(__dirname, 'node_modules'),
      use: ['babel-loader', 'eslint-loader']
    },
    {
      test: /\.svg$/,
      use: [
        {
          loader: 'svg-url-loader',
          options: {
            limit: 10000,
          },
        },
      ]
    },
    {
      test: /\.css$/i,
      use: ['style-loader', 'css-loader'],
    }]
  },
  devServer: {
    contentBase:  path.resolve(__dirname, 'build'),
    port: 9000
  },
  plugins: [
    new HtmlWebpackPlugin({
      template: "src/index.html" //source html
    })
  ]
};