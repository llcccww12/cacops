const CssMinimizerPlugin = require("css-minimizer-webpack-plugin");
const fastGlob = require("fast-glob");
const MiniCssExtractPlugin = require("mini-css-extract-plugin");
const MonacoWebpackPlugin = require("monaco-editor-webpack-plugin");
const PostCSSPresetEnv = require("postcss-preset-env");
const PostCSSSafeParser = require("postcss-safe-parser");
const SpriteLoaderPlugin = require("svg-sprite-loader/plugin");
const TerserPlugin = require("terser-webpack-plugin");
const VueLoaderPlugin = require("vue-loader/lib/plugin");
const { statSync } = require("fs");
const { resolve, parse } = require("path");
const { SourceMapDevToolPlugin } = require("webpack");

const glob = (pattern) =>
  fastGlob.sync(pattern, { cwd: __dirname, absolute: true });

const themes = {};
for (const path of glob("web_src/less/themes/*.less")) {
  themes[parse(path).name] = [path];
}

const standalone = {};
const stadalonePaths = [
  ...glob("web_src/js/standalone/**/*.js"),
  ...glob("web_src/less/standalone/**/*.less"),
];
for (const path of stadalonePaths) {
  standalone[parse(path).name] = [path];
}

const vuePages = {};
for (const path of glob("web_src/vuepages/**/vp-*.js")) {
  vuePages[parse(path).name] = [path];
}

const isProduction = process.env.NODE_ENV !== "development";

module.exports = {
  mode: isProduction ? "production" : "development",
  entry: {
    index: [
      resolve(__dirname, "web_src/js/index.js"),
      resolve(__dirname, "web_src/less/index.less"),
    ],
    home: [resolve(__dirname, "web_src/js/home.js")],
    jquery: [resolve(__dirname, "web_src/js/jquery.js")],
    icons: glob("node_modules/@primer/octicons/build/svg/**/*.svg"),
    ...standalone,
    ...themes,
    ...vuePages,
  },
  devtool: false,
  output: {
    path: resolve(__dirname, "public"),
    filename: "js/[name].js",
    chunkFilename: "js/[name].js",
  },
  cache: {
    type: "filesystem",
    buildDependencies: {
      config: [__filename],
    },
  },
  optimization: {
    minimize: isProduction,
    minimizer: [
      new TerserPlugin({
        extractComments: false,
        terserOptions: {
          keep_fnames: /^(HTML|SVG)/, // https://github.com/fgnass/domino/issues/144
          compress: {
            drop_console: isProduction,
          },
          output: {
            comments: false,
          },
        },
      }),
      new CssMinimizerPlugin({
        minimizerOptions: {
          preset: ["default", { discardComments: { removeAll: true } }],
        },
      }),
    ],
    splitChunks: {
      chunks: "async",
      cacheGroups: {
        // bundle all monaco languages into one file instead of emitting 1-65.js files
        monaco: {
          test: /monaco-editor/,
          name: "monaco",
          chunks: "async",
        },
        vendors: {
          test: /node_modules/,
          name: "vendor-libs",
          chunks: "all",
          minSize: 0,
          minChunks: 2,
          maxInitialRequests: 3,
        },
      },
    },
  },
  module: {
    rules: [
      {
        test: /\.vue$/,
        exclude: /node_modules/,
        loader: "vue-loader",
      },
      {
        test: require.resolve("jquery-datetimepicker"),
        use: {
          loader: "imports-loader",
          options: {
            additionalCode: "var define = false; var exports = false;",
          },
        },
      },
      {
        test: /\.ts$/,
        use: [{ loader: "ts-loader" }],
        exclude: /node_modules/,
      },
      {
        test: /\.js$/,
        exclude: /node_modules/,
        use: [
          {
            loader: "babel-loader",
            options: {
              cacheDirectory: true,
              cacheCompression: false,
              cacheIdentifier: [
                resolve(__dirname, "package.json"),
                resolve(__dirname, "webpack.config.js"),
              ]
                .map((path) => statSync(path).mtime.getTime())
                .join(":"),
              sourceMaps: !isProduction,
              presets: [
                [
                  "@babel/preset-env",
                  {
                    useBuiltIns: "usage",
                    corejs: 3,
                  },
                ],
              ],
              plugins: [
                [
                  "@babel/plugin-transform-runtime",
                  {
                    regenerator: true,
                  },
                ],
                "@babel/plugin-proposal-object-rest-spread",
                "@babel/plugin-syntax-import-meta",
              ],
            },
          },
        ],
      },
      {
        test: /\.(less|css)$/i,
        use: [
          { loader: MiniCssExtractPlugin.loader },
          {
            loader: "css-loader",
            options: {
              importLoaders: 2,
              url: {
                filter: (_url, resourcePath) => {
                  // only resolve URLs for dependencies
                  return resourcePath.includes("node_modules");
                },
              },
            },
          },
          {
            loader: "postcss-loader",
            options: {
              postcssOptions: {
                plugins: [PostCSSPresetEnv()],
              },
            },
          },
          { loader: "less-loader" },
        ],
      },
      {
        test: /\.svg$/,
        use: [
          {
            loader: "svg-sprite-loader",
            options: {
              extract: true,
              spriteFilename: "img/svg/icons.svg",
              symbolId: (path) => {
                const { name } = parse(path);
                if (/@primer[/\\]octicons/.test(path)) {
                  return `octicon-${name}`;
                }
                return name;
              },
            },
          },
          { loader: "svgo-loader" },
        ],
      },
      {
        test: /\.(png|jpg|gif)$/,
        type: "asset/resource",
      },
      {
        test: /\.(ttf|woff2?)$/,
        type: "asset/resource",
        generator: {
          filename: "fonts/[name][ext]",
        },
      },
    ],
  },
  plugins: [
    new VueLoaderPlugin(),
    new MiniCssExtractPlugin({
      filename: "css/[name].css",
      chunkFilename: "css/[name].css",
      ignoreOrder: true,
    }),
    new SpriteLoaderPlugin({
      plainSprite: true,
    }),
    new MonacoWebpackPlugin({
      filename: "js/monaco-[name].worker.js",
    }),
  ],
  performance: {
    hints: false,
    maxEntrypointSize: Infinity,
    maxAssetSize: Infinity,
  },
  resolve: {
    symlinks: false,
    fallback: {
      fs: false,
    },
    alias: {
      vue$: "vue/dist/vue.esm.js", // needed because vue's default export is the runtime only
      "~": resolve(__dirname, "web_src/vuepages"),
    },
    extensions: [".tsx", ".ts", ".js"],
  },
  watchOptions: {
    ignored: ["node_modules/**"],
  },
  stats: {
    children: false,
  },
};
