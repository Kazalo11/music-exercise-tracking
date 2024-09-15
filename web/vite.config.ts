 
import react from '@vitejs/plugin-react'
import { defineConfig } from 'vite'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react()],
  preview: {
    port: 3000,
    strictPort: true
  },
  build: {
    minify: 'esbuild',
    sourcemap: true
  },
  server: {
    port: 3000,
    strictPort: true,
    host: true
    }
  })
