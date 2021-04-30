module.exports = {
  roots: ['<rootDir>/src'],
  setupFiles: ['reflect-metadata', './test/globals.ts'],
  transform: {
    '^.+\\.ts?$': 'ts-jest',
  },
};
