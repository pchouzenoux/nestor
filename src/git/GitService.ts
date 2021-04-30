import { Service } from 'typedi';
import { InjectLogger, Logger } from '../helpers';
import { ConsoleLogger } from '../utils/ConsoleLogger';
import { execCommand } from '../utils/Shell';

export class GitNotDefineException extends Error {
  constructor() {
    super('Git command is not defined, tried `git version` to validate it.');
  }
}
@Service()
export class GitService {
  private static PROTECTED_BRANCH_PATTERNS = [
    /^(main|master)$/,
    /^\*/,
    /^(dev|develop)$/,
  ];

  constructor(
    @InjectLogger('GitService') private logger: Logger,
    /* @Inject */ private consoleLogger: ConsoleLogger,
  ) {}

  public clean(): void {
    if (!this.isGitDefine()) {
      throw new GitNotDefineException();
    }
    this.removeNotProtectedBranches();
    this.pruneRemotes();
  }

  private isGitDefine(): boolean {
    try {
      const result = execCommand('git version');
      return !!result?.toString().match(/^git version 2\.\d+\.\d+/);
    } catch (err) {
      this.logger.error(err, { msg: 'Error checking git version' });
      return false;
    }
  }

  private removeNotProtectedBranches(): void {
    const branches = execCommand('git branch')
      .toString()
      .split('\n')
      .map((l) => l.trim())
      .filter((l) => l.length > 0)
      .filter((branch) => {
        for (const protectedPattern of GitService.PROTECTED_BRANCH_PATTERNS) {
          if (branch?.match(protectedPattern)) {
            return false;
          }
        }
        return true;
      });

    if (branches.length > 0) {
      const output = execCommand(`git branch -d ${branches.join(' ')}`);

      this.consoleLogger.log(
        'Some branche have been removed my Sir:\n',
        output,
      );
    } else {
      this.consoleLogger.log('No branche to delete.');
    }
  }

  private pruneRemotes(): void {
    const remotes = execCommand('git remote')
      .toString()
      .split('\n')
      .map((l) => l.trim())
      .filter((l) => l.length > 0);

    if (remotes.length > 0) {
      const output = execCommand(`git remote prune ${remotes.join(' ')}`);

      this.consoleLogger.log('Some remote have been pruned, Mr. :\n', output);
    } else {
      this.consoleLogger.log('No remote to prune.');
    }
  }
}
