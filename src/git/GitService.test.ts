import { GitNotDefineException, GitService } from './GitService';
import { execCommand } from '../utils/Shell';

jest.mock('../utils/Shell', () => ({
  execCommand: jest.fn(),
}));

describe('GitService test suite', () => {
  let service: GitService;

  beforeEach(() => {
    (execCommand as jest.Mock).mockReturnValue(Buffer.from(''));
    service = new GitService(global['console-logger']);
  });

  afterEach(() => {
    jest.resetAllMocks();
  });

  describe('clean', () => {
    test('should check if there is something to clean', () => {
      // Prepare
      (execCommand as jest.Mock).mockReturnValueOnce(
        Buffer.from('git version 2.31.1'),
      );

      // Run
      service.clean();

      // Expect
      expect(execCommand).toHaveBeenCalledTimes(3);
      expect(execCommand).toHaveBeenNthCalledWith(1, 'git version');
      expect(execCommand).toHaveBeenNthCalledWith(2, 'git branch');
      expect(execCommand).toHaveBeenNthCalledWith(3, 'git remote');
    });

    test.each<{
      branches: string[];
      expected: string[];
    }>([
      {
        branches: ['main', 'feature/feature-1', '* feature/feature-2'],
        expected: ['feature/feature-1'],
      },
      {
        branches: ['* main', 'feature/feature-1', 'feature/feature-2'],
        expected: ['feature/feature-1', 'feature/feature-2'],
      },
      {
        branches: ['* main', 'develop'],
        expected: [],
      },
      {
        branches: ['main', '* feature/feature-1'],
        expected: [],
      },
      {
        branches: [],
        expected: [],
      },
      {
        branches: ['', '     '],
        expected: [],
      },
    ])('should remove unprotected branches: %j', ({ branches, expected }) => {
      // Prepare
      (execCommand as jest.Mock)
        .mockReturnValueOnce(Buffer.from('git version 2.31.1'))
        .mockReturnValueOnce(branches.join('\n'));

      // Run
      service.clean();

      // Expect
      expect(execCommand).toHaveBeenCalledTimes(expected.length > 0 ? 4 : 3);
      if (expected.length > 0) {
        expect(execCommand).toHaveBeenNthCalledWith(
          3,
          `git branch -d ${expected.join(' ')}`,
        );
      }
    });

    test.each<{
      remotes: string[];
      expected: string[];
    }>([
      {
        remotes: ['remote-1'],
        expected: ['remote-1'],
      },
      {
        remotes: ['remote-1', 'remote-2'],
        expected: ['remote-1', 'remote-2'],
      },
      {
        remotes: [],
        expected: [],
      },
      {
        remotes: ['', '     '],
        expected: [],
      },
    ])('should remove unprotected branches: %j', ({ remotes, expected }) => {
      // Prepare
      (execCommand as jest.Mock)
        .mockReturnValueOnce(Buffer.from('git version 2.31.1'))
        .mockReturnValueOnce('')
        .mockReturnValueOnce(remotes.join('\n'));

      // Run
      service.clean();

      // Expect
      expect(execCommand).toHaveBeenCalledTimes(expected.length > 0 ? 4 : 3);
      if (expected.length > 0) {
        expect(execCommand).toHaveBeenNthCalledWith(
          4,
          `git remote prune ${expected.join(' ')}`,
        );
      }
    });

    describe('check if git is define then', () => {
      test.each([
        undefined,
        null,
        Buffer.from(''),
        Buffer.from('unknown result'),
        Buffer.from('git version 1.0.0'),
        Buffer.from('git version 2.31.1'),
      ])(
        'should throw GitNotDefineException error if git version retun: "%s"',
        (mockResult) => {
          // Prepare
          (execCommand as jest.Mock).mockReturnValueOnce(mockResult);

          try {
            // Run
            service.clean();
          } catch (err) {
            // Expect
            expect(err).toBeInstanceOf(GitNotDefineException);
          }
        },
      );

      test('should throw GitNotDefineException error if git version throw an error', () => {
        // Prepare
        (execCommand as jest.Mock).mockImplementation(() => {
          throw new Error('test');
        });

        try {
          // Run
          service.clean();
        } catch (err) {
          // Expect
          expect(err).toBeInstanceOf(GitNotDefineException);
        }
      });
    });
  });
});
