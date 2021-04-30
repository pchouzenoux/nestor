// needs to be imported first
import 'reflect-metadata';
import sourcemap from 'source-map-support';

sourcemap.install();

import Container from 'typedi';
import { createLogger } from './helpers';
import { NestorCli } from './NestorCli';

const logger = createLogger('root');

process.on('uncaughtException', (err) => {
  logger.error({ msg: 'UncaughtException', err });
});

process.on('unhandledRejection', (err) => {
  logger.error({ msg: 'UnhandledRejection', err });
});

process.on('SIGINT', () => {
  logger.error('SIGINT recieved, kill the app');
  process.exit();
});

const main = (): void => {
  const nestorCli = Container.get(NestorCli);
  nestorCli.start();
};

main();
