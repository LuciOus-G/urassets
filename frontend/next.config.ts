import type { NextConfig } from "next";

let indodaxApi;
const nextConfig: NextConfig = {
  /* config options here */
  async rewrites() {
    return [
      {
        source: "/api/pairs",
        destination: process.env.INDODAX_API,
      },
      {
        source: "/user/login",
        destination: `${process.env.BACKEND_LOCAL}/user/login`,
      },
    ];
  },

  async redirects() {
    return [
      {
        source: "/",
        destination: "/dashboard/summary",
        permanent: true,
      },
    ];
  },

  webpack: (config) => {
    config.resolve.alias = {
      ...config.resolve.alias,
      "@services": "./src/services",
      "@components": "./src/components",
      "@toast": "./src/components/toast"
    };
    return config;
  },
};

export default nextConfig;
