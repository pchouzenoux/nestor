import { execSync } from 'child_process';

export class ShellException extends Error {
  constructor(
    public status: number,
    public stdout: Buffer,
    public stderr: Buffer,
  ) {
    super(stderr.toString());
  }
}
export const execCommand = (command: string): Buffer => {
  try {
    return execSync(command);
  } catch (err) {
    throw new ShellException(err.status, err.stdout, err.stderr);
  }
};
