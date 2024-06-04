import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";

import * as path from "path";
import { fileURLToPath } from "url";

const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

export default defineConfig({
  build: {
    outDir: "./static",
  },
  resolve: {
    alias: {
      "@": path.resolve(__dirname, "src"),
    },
  },
  plugins: [vue()],
  define: {
    VITE_BACKEND_URL: process.env.VITE_BACKEND_URL,
  },
});
