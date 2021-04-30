"use strict";
exports.__esModule = true;
var testLogger = {
    trace: jest.fn(),
    debug: jest.fn(),
    info: jest.fn(),
    warn: jest.fn(),
    error: jest.fn()
};
global['test-logger'] = testLogger;
