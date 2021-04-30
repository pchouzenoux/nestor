/* eslint-disable no-console */
import chalk from 'chalk';
import { Service } from 'typedi';

@Service()
export class ConsoleLogger {
  public log(...args: any[]): void {
    console.log(args.join(' '));
  }

  public error(...args: any[]): void {
    console.log(chalk.red.bold('Error:'), chalk.red(args.join(' ')));
  }
}
