import yargs from 'yargs';
import { hideBin } from 'yargs/helpers';

export default yargs(hideBin(process.argv))
  .command(
    '$0',
    `
  _    _           _             
  | \\ | |         | |            
  |  \\| | ___  ___| |_ ___  _ __ 
  | . \` |/ _ \\/ __| __/ _ \\| '__|
  | |\\  |  __/\\__ \\ || (_) | |   
  |_| \\_|\\___||___/\\__\\___/|_|   
  
  nestor <context> <command>`,
    (yargs) => {
      return yargs
        .positional('context', {
          describe: 'Context can be `git` or `docker`',
        })
        .positional('command', {
          describe: 'Command in the context',
        });
    },
  )
  .help('help');
