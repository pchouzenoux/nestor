import { execCommand, ShellException } from './Shell';

describe('Shell utils test suite', () => {
  describe('execCommand', () => {
    test('should execute and return result buffer', () => {
      const result = execCommand('echo Hello World');

      expect(result).toBeInstanceOf(Buffer);
      expect(result.length).toEqual(12);
      expect(result.toString()).toEqual('Hello World\n');
    });

    test('should throw ShellException is command fails', () => {
      try {
        execCommand('UNKNOWN');
      } catch (err) {
        expect(err).toBeInstanceOf(ShellException);
        expect(err.message).toEqual('/bin/sh: UNKNOWN: command not found\n');
        expect(err.status).toEqual(127);
        expect(err.stdout).toBeInstanceOf(Buffer);
        expect(err.stdout.length).toEqual(0);
        expect(err.stderr).toBeInstanceOf(Buffer);
        expect(err.stderr.length).toEqual(36);
      }
    });
  });
});
