import { Service } from 'typedi';
import { ConsoleLogger } from '../utils/ConsoleLogger';
import { execCommand } from '../utils/Shell';

export class DockerNotDefineException extends Error {
  constructor() {
    super(
      'Docker command is not defined, tried `docker version` to validate it.',
    );
  }
}

@Service()
export class DockerService {
  constructor(/* @Inject */ private consoleLogger: ConsoleLogger) {}

  public clean(): void {
    if (!this.isDockerDefine()) {
      throw new DockerNotDefineException();
    }

    this.removeUnnecessaryContainer();

    ['docker system prune -f', 'docker volume prune -f'].forEach((cmd) =>
      execCommand(cmd),
    );
    this.consoleLogger.log('I clean underlying docker system, Monsieur !');
  }

  private isDockerDefine(): boolean {
    try {
      const result = execCommand('docker version');
      return !!result?.toString().match(/^Client: Docker Engine - Community/);
    } catch (err) {
      return false;
    }
  }

  private removeUnnecessaryContainer(): void {
    const containers = execCommand('docker ps -aq --no-trunc -f status=exited')
      .toString()
      .split('\n')
      .map((l) => l.trim())
      .filter((l) => l.length > 0);

    if (containers.length > 0) {
      const output = execCommand(`docker rm ${containers.join(' ')}`);

      this.consoleLogger.log(
        'I found and removed some unused containers , my lord:',
        output,
      );
    } else {
      this.consoleLogger.log('No container to delete.');
    }
  }
}
