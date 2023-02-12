export const environment = {
  production: true,
  WSServ: getUrl(),
};

export function getUrl() {
  return 'wss://' + window.location.hostname + (window.location.port ? ':' + window.location.port : '') + '/ws';
}
