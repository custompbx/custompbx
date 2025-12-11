import { Action, createAction, FunctionWithParametersType } from '@ngrx/store';


export interface PayloadAction<T extends string, P> extends Action {
  type: T;
  payload: P;
}

export function createActionHelper(type: string):
  FunctionWithParametersType<[any], { payload: any } & Action<string>> & Action<string> {
  return createAction(type, (payload: any) => ({ payload }));
}
