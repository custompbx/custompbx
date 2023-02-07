export const environment = {
  production: true,
  WSServ: getUrl(),
};

export function getUrl() {
  return 'wss://' + window.location.hostname + (window.location.port ? ':' + window.location.port : '') + '/ws';

  // return 'wss://45.61.54.76:8080/ws';
}
