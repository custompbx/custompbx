// This file is required by karma.conf.js and loads recursively all the .spec and framework files

import 'zone.js';
import 'zone.js/testing';
import { getTestBed } from '@angular/core/testing';
import {
  BrowserDynamicTestingModule,
  platformBrowserDynamicTesting
} from '@angular/platform-browser-dynamic/testing';
import { customPbxTestProviders } from './app/testing/test-providers';

// First, initialize the Angular testing environment.
getTestBed().initTestEnvironment(
  BrowserDynamicTestingModule,
  platformBrowserDynamicTesting()
);

beforeEach(() => {
  getTestBed().configureTestingModule({
    providers: customPbxTestProviders(),
  });
});

// Then we find all the tests.
const context = (import.meta as any).webpackContext('./', {
  recursive: true,
  regExp: /\.spec\.ts$/,
});
// And load the modules.
context.keys().map(context);
