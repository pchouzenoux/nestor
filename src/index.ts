// needs to be imported first
import 'reflect-metadata';
import sourcemap from 'source-map-support';

sourcemap.install();

import Container from 'typedi';
import { NestorCli } from './NestorCli';

process.on('uncaughtException', (err) => {
  // eslint-disable-next-line no-console
  console.error({ msg: 'UncaughtException', err });
});

process.on('unhandledRejection', (err) => {
  // eslint-disable-next-line no-console
  console.error({ msg: 'UnhandledRejection', err });
});

process.on('SIGINT', () => {
  // eslint-disable-next-line no-console
  console.error('SIGINT recieved, kill the app');
  process.exit();
});

const main = (): void => {
  const nestorCli = Container.get(NestorCli);
  nestorCli.start();
};

main();
