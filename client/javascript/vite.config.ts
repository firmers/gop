import {defineConfig} from 'vite';
import {resolve} from "path";
import dts from 'vite-plugin-dts';
import tsc from '@firmer/tsc';
import react from "@vitejs/plugin-react";

// https://vitejs.dev/config/
export default defineConfig({
    plugins: [
        dts({
            insertTypesEntry: true,
        }),
        tsc(),
        react({
            jsxRuntime: 'automatic'
        }),
    ],
    resolve: {
        alias: {
            '@': resolve(__dirname, './src'),
        },
    },
    build: {
        lib: {
            // Could also be a dictionary or array of multiple entry points
            entry: resolve(__dirname, 'src/index.ts'),
            name: 'gop',
            fileName: 'gop'
        },
        rolldownOptions: {
            // 需要是外部依赖，否则会把mesh完全编译到独立的包中出现模块隔离
            external: ['@firmer/mesh', 'react', 'react-dom', 'react/jsx-runtime'],
        }
    },
})