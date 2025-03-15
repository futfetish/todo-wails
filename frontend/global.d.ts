
import {main} from './wailsjs/go/models';

declare module "wailsjs/go/main" {
    export function GetTodos(completed?: boolean): Promise<Array<Record<string, main.Todo>>>;
  }