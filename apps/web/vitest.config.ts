import { configDefaults, defineConfig, } from 'vitest/config'

export default defineConfig({
  test: {
    projects: [
      {
        extends: true,
        test: {
          globals: true,
          environment: "jsdom",
          setupFiles: "./src/setupTests.ts",
          include: ['src/**.test.tsx'],
          name: 'unit'
        },        
      }
    ]

  },
})