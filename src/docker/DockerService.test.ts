import { DockerNotDefineException, DockerService } from './DockerService';
import { execCommand } from '../utils/Shell';

jest.mock('../utils/Shell', () => ({
  execCommand: jest.fn(),
}));

describe('DockerService test suite', () => {
  let service: DockerService;

  beforeEach(() => {
    (execCommand as jest.Mock).mockReturnValue(Buffer.from(''));
    service = new DockerService(global['console-logger']);
  });

  afterEach(() => {
    jest.resetAllMocks();
  });

  describe('clean', () => {
    test('should check if there is something to clean and prune', () => {
      // Prepare
      (execCommand as jest.Mock).mockReturnValueOnce(
        Buffer.from('Client: Docker Engine - Community'),
      );

      // Run
      service.clean();

      // Expect
      expect(execCommand).toHaveBeenCalledTimes(4);
      expect(execCommand).toHaveBeenNthCalledWith(1, 'docker version');
      expect(execCommand).toHaveBeenNthCalledWith(
        2,
        'docker ps -aq --no-trunc -f status=exited',
      );
      expect(execCommand).toHaveBeenNthCalledWith(3, 'docker system prune -f');
      expect(execCommand).toHaveBeenNthCalledWith(4, 'docker volume prune -f');
    });

    test.each<{
      containers: string[];
      expected: string[];
    }>([
      {
        containers: ['container-1'],
        expected: ['container-1'],
      },
      {
        containers: ['container-1', 'container-2'],
        expected: ['container-1', 'container-2'],
      },
      {
        containers: [],
        expected: [],
      },
      {
        containers: ['', '     '],
        expected: [],
      },
    ])(
      'should remove unprotected containers: %j',
      ({ containers, expected }) => {
        // Prepare
        (execCommand as jest.Mock)
          .mockReturnValueOnce(Buffer.from('Client: Docker Engine - Community'))
          .mockReturnValueOnce(containers.join('\n'));

        // Run
        service.clean();

        // Expect
        expect(execCommand).toHaveBeenCalledTimes(expected.length > 0 ? 5 : 4);
        expect(execCommand).toHaveBeenNthCalledWith(
          2,
          'docker ps -aq --no-trunc -f status=exited',
        );
        if (expected.length > 0) {
          expect(execCommand).toHaveBeenNthCalledWith(
            3,
            `docker rm ${expected.join(' ')}`,
          );
        }
      },
    );

    describe('check if docker is define then', () => {
      test.each([
        undefined,
        null,
        Buffer.from(''),
        Buffer.from('unknown result'),
        Buffer.from('Client: Docker Engine - Community'),
      ])(
        'should throw DockerNotDefineException error if docker version retun: "%s"',
        (mockResult) => {
          // Prepare
          (execCommand as jest.Mock).mockReturnValueOnce(mockResult);

          try {
            // Run
            service.clean();
          } catch (err) {
            // Expect
            expect(err).toBeInstanceOf(DockerNotDefineException);
          }
        },
      );

      test('should throw DockerNotDefineException error if docker version throw an error', () => {
        // Prepare
        (execCommand as jest.Mock).mockImplementation(() => {
          throw new Error('test');
        });

        try {
          // Run
          service.clean();
        } catch (err) {
          // Expect
          expect(err).toBeInstanceOf(DockerNotDefineException);
        }
      });
    });
  });
});
