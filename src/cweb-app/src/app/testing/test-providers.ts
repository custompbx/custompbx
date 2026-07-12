import {EnvironmentProviders, Provider} from '@angular/core';
import {provideHttpClient} from '@angular/common/http';
import {provideNoopAnimations} from '@angular/platform-browser/animations';
import {provideRouter} from '@angular/router';
import {combineReducers, provideStore} from '@ngrx/store';
import {config as websocketConfig} from '../services/websocket/websocket.config';
import {reducers} from '../store/app.states';

export function customPbxTestProviders(): Array<Provider | EnvironmentProviders> {
  return [
    provideHttpClient(),
    provideNoopAnimations(),
    provideRouter([]),
    provideStore({app: combineReducers(reducers)}),
    {
      provide: websocketConfig,
      useValue: {url: 'ws://localhost/ws', reconnectAttempts: 0},
    },
  ];
}
