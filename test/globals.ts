import { Logger } from 'pino';

const testLogger: Logger = {
  trace: jest.fn(),
  debug: jest.fn(),
  info: jest.fn(),
  warn: jest.fn(),
  error: jest.fn(),
};

global['test-logger'] = testLogger;
