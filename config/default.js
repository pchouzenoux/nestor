module.exports = {
  logger: {
    pretty: {
      colorize: true,
      translateTime: 'yyyy-mm-dd HH:MM:ss',
      errorProps: 'err,error',
      messageKey: 'msg',
      ignore: 'hostname',
    },

    // minumum level of logging
    level: 'fatal',

    // Modules to log
    // e.g.
    //   '*,!cache,!method' => log everything except cache and method logging
    //   'cache' => only log cache
    modules: '*',
  },
};
