import CrowClient from '~/lib/api';

import useConfig from './useConfig';

let apiClient: CrowClient | undefined;

export default (): CrowClient => {
  if (!apiClient) {
    const config = useConfig();
    const server = config.rootPath;
    const token = null;
    const csrf = config.csrf ?? null;

    apiClient = new CrowClient(server, token, csrf);
  }

  return apiClient;
};
