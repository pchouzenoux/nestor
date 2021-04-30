import { Service } from 'typedi';
import yargs from 'yargs';
import { DockerService } from './docker/DockerService';
import { GitService } from './git/GitService';
import { InjectLogger, Logger } from './helpers';
import args from './CliCommand';
import { ConsoleLogger } from './utils/ConsoleLogger';

enum Context {
  GIT = 'git',
  DOCKER = 'docker',
}
@Service()
export class NestorCli {
  constructor(
    @InjectLogger('NestorCli') private logger: Logger,
    /* @Inject() */ private consoleLogger: ConsoleLogger,
    /* @Inject() */ private gitService: GitService,
    /* @Inject() */ private dockerService: DockerService,
  ) {}

  public start(): void {
    try {
      const [context, command] = args.argv._;
      this.run(context as Context, command as string);
    } catch (err) {
      this.consoleLogger.error(err.message);
      yargs.showHelp();
    }
  }

  private run(context: Context, command: string): void {
    switch (context) {
      case Context.GIT:
        this.runGitCommand(command);
        break;
      case Context.DOCKER:
        this.runDockerCommand(command);
        break;
      default:
        throw new Error(`Unknown context: '${context}'`);
    }
  }

  private runGitCommand(command: string): void {
    switch (command) {
      case 'clean':
        this.gitService.clean();
        break;
      default:
        throw new Error(`Unknown command in git context: '${command}'`);
    }
  }

  private runDockerCommand(command: string): void {
    switch (command) {
      case 'clean':
        this.dockerService.clean();
        break;
      default:
        throw new Error(`Unknown command in docker context: '${command}'`);
    }
  }
}
