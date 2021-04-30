/* istanbul ignore file */
import config from 'config'; // can't use the helpers/Config because config depends on logger (circular dep)
import { Map } from 'immutable';
import pino from 'pino';
import { Container } from 'typedi';

const rootLogger = pino({
  level: config.get<string>('logger.level'),
  prettyPrint: config.get<any>('logger.pretty'),
});

const logModulesList: string[] = config
  .get<string>('logger.modules')
  .split(',')
  .map((m: string) => m.trim());

const logModules = {
  logAll: logModulesList.includes('*'),
  whitelist: Map<string, boolean>(
    logModulesList
      .filter((m) => !m.startsWith('!'))
      .map<[string, boolean]>((m) => [m, true]),
  ),
  blacklist: Map<string, boolean>(
    logModulesList
      .filter((m) => m.startsWith('!'))
      .map<[string, boolean]>((m) => [m.slice(1), true]),
  ),
};

function isLoggingEnabledForModule(moduleName: string): boolean {
  const fragmentCheck = (map: Map<string, boolean>): boolean => {
    let acc: string;
    for (const fragment of moduleName.split('.')) {
      acc = acc ? `${acc}.${fragment}` : fragment;
      if (map.has(acc)) {
        return true;
      }
    }
    return false;
  };

  return logModules.logAll
    ? !fragmentCheck(logModules.blacklist)
    : fragmentCheck(logModules.whitelist);
}

export type Logger = pino.Logger;

export const createLogger = (moduleName?: string): Logger => {
  if (!moduleName) {
    return rootLogger;
  }

  return rootLogger.child({
    level: isLoggingEnabledForModule(moduleName) ? undefined : 'silent',
    module: moduleName,
  });
};

/**
 * Decorator for creating child logger (scoped with module).
 * @param moduleName The module name
 */
export function InjectLogger(moduleName: string): ParameterDecorator {
  const logger = createLogger(moduleName);

  return (object: any, propertyName: string, index?: number): void => {
    Container.registerHandler({
      index,
      object,
      propertyName,
      value: (/* containerInstance */) => logger,
    });
  };
}
