import config from 'config';
import { Container } from 'typedi';

import { createLogger } from './Logger';

const logger = createLogger('config');

logger.info({
  msg: 'File used to load configs',
  files: config.util.getConfigSources().map((source) => source.name),
});

export function getConfig<T>(configPath: string, required = true): T {
  const value = config.has(configPath) ? config.get<T>(configPath) : null;

  if (
    required &&
    (value === null || value === undefined || (value as any) === '')
  ) {
    throw new Error(
      `Required config '${configPath}'. Are you missing a value in your local.js?`,
    );
  }

  return value;
}

/**
 * Decorator for creating child config (scoped with module).
 * @param moduleName The module name
 */
export function InjectConfig<T>(
  configPath: string,
  required = true,
): ParameterDecorator {
  const value = getConfig<T>(configPath, required);

  logger.trace({ msg: 'Inject config', configPath, value });

  return (object: any, propertyName: string, index?: number): void => {
    Container.registerHandler({
      index,
      object,
      propertyName,
      value: (/* containerInstance */) => value,
    });
  };
}
