const path = require("path");
const HtmlWebpackPlugin = require("html-webpack-plugin");

module.exports = {
    entry: "./src/index.tsx",
    output: {
        path: path.resolve(__dirname, "dist"),
        filename: "bundle.js",
        publicPath: "/", // important for SPA routing
    },
    resolve: {
        extensions: [".tsx", ".ts", ".js"],
    },
    module: {
        rules: [
            {
                test: /\.tsx?$/,
                loader: "ts-loader",
                exclude: /node_modules/,
            },
        ],
    },
    plugins: [
        new HtmlWebpackPlugin({
            template: path.resolve(__dirname, "src/public", "index.html"),
        }),
    ],
    devServer: {
        static: path.join(__dirname, "dist"),
        port: 3000,
        proxy: [
            {
                context: ["/repo.json", "/api/plugins"],
                target: "http://localhost:8080",
            },
        ],
        hot: true,
    },
};
