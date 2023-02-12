// This file can be replaced during build by using the `fileReplacements` array.
// `ng build --prod` replaces `environment.ts` with `environment.prod.ts`.
// The list of file replacements can be found in `angular.json`.
let WS_BACKGROUND_OVERRIDE = {getWs() { return ''; }};
try {
  WS_BACKGROUND_OVERRIDE = require('../../env');
} catch (e) {}
export const environment = {
  production: false,
  WSServ: getUrl(),
  optimization: false,
  sourceMap: false,
  extractCss: true,
  namedChunks: false,
  extractLicenses: true,
  vendorChunk: false,
  buildOptimizer: false,
  configurations: {
      tsConfig: '../tsconfig.app.json'
  },
};

export function getUrl() {
  console.log(WS_BACKGROUND_OVERRIDE.getWs());
  const url = WS_BACKGROUND_OVERRIDE.getWs();
  return url || 'wss://' + window.location.hostname + (window.location.port ? ':' + window.location.port : '') + '/ws';
}
/*
 * For easier debugging in development mode, you can import the following file
 * to ignore zone related error stack frames such as `zone.run`, `zoneDelegate.invokeTask`.
 *
 * This import should be commented out in production mode because it will have a negative impact
 * on performance if an error is thrown.
 */
// import 'zone.js/dist/zone-error';  // Included with Angular CLI.
