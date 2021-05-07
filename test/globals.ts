import { ConsoleLogger } from '../src/utils/ConsoleLogger';

const testLogger: ConsoleLogger = {
  log: jest.fn(),
  error: jest.fn(),
};

global['console-logger'] = testLogger;
