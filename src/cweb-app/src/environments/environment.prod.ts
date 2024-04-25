export const environment = {
  production: true,
  WSServ: 'wss://' + window.location.hostname + (window.location.port ? ':' + window.location.port : '') + '/ws',
};
