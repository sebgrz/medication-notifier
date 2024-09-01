/** @type {import('next').NextConfig} */
import nextswPkg from "next-sw";
import CopyPlugin from "copy-webpack-plugin";

const { withServiceWorker } = nextswPkg;

const nextConfig = withServiceWorker({
  name: "firebase-messaging-sw.js",
  entry: "config/service-worker/firebase-messaging-sw.js",
  livereload: true
})({
  webpack: (config) => {
    config.plugins.push(
      new CopyPlugin({
        patterns: [
          // path to 'public' is set to parent dir because build is running from '.next' directory
          // https://github.com/vercel/next.js/discussions/12844#discussioncomment-14380
          { from: "node_modules/firebase/firebase-app-compat.js", to: "../public/", force: true },
          { from: "node_modules/firebase/firebase-messaging-compat.js", to: "../public/", force: true },
        ],
      }),
    )
    return config;
  }
});

export default nextConfig;
