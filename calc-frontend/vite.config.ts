// import react from "@vitejs/plugin-react";
// import { defineConfig, loadEnv } from "vite";
//
// export default ({ mode }: { mode: string }) => {
//   process.env = { ...process.env, ...loadEnv(mode, process.cwd()) };
//
//   return defineConfig({
//     plugins: [react()],
//     server: {
//       host: "0.0.0.0",
//       port: 5173,
//       watch: { usePolling: true }, // only if hot reload doesn't trigger
//       proxy: {
//         "/api": {
//           target: process.env.BACKEND_URL,
//           changeOrigin: true,
//           rewrite: (path) => path.replace(/^\/api/, ""), // remove if backend serves under /api
//           ws: true, // for websockets
//         },
//       },
//     },
//   });
// };

import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// Mirrors the nginx proxy so `npm run dev` behaves like production.
export default defineConfig({
  plugins: [react()],
  server: {
    proxy: {
      '/api': {
        target: 'http://localhost:9090',
        changeOrigin: true,
        rewrite: (path) => path.replace(/^\/api/, ''),
      },
    },
  },
})
