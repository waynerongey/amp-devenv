const MiniCssExtractPlugin = require('mini-css-extract-plugin');
var path = require('path');

module.exports = {
  mode: 'production',
  entry: './src/index.scss',
  module: {
    rules: [
      {
      test: /\.s[ac]ss$/i,
      use: [{
          loader: 'style-loader'
        },
        {
          loader: MiniCssExtractPlugin.loader
        },
        {
          loader: 'css-loader'
        },
        {
          loader: 'postcss-loader',
          options: {
            plugins: function () {
              return [require('autoprefixer')];
            }
          }
        },
        {
          loader: 'sass-loader'
        },
      ],
    }],
  },
  plugins: [
    new MiniCssExtractPlugin({
      // Options similar to the same options in webpackOptions.output
      // both options are optional
      filename: '[name].css',
      chunkFilename: '[id].css',
    })
  ],
  devServer: {
    contentBase: path.join(__dirname, 'dist'),
    compress: true,
    hot: true,
    overlay: {
      errors: true
    }
  }
};