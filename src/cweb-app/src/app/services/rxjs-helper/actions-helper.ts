import {Action, createAction, FunctionWithParametersType} from "@ngrx/store";
import {TypedAction} from "@ngrx/store/src/models";

export interface PayloadAction<T extends string, P> extends Action {
  type: T;
  payload: P;
}

export function createActionHelper(type: string): FunctionWithParametersType<[any], { payload: any } & TypedAction<string>> & TypedAction<string> {
  return createAction(type, (payload: any) => ({ payload }))
}
