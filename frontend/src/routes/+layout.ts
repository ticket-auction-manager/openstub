import { getConfig, type Config } from '#lib/api/config.ts';

export const prerender = false;
export const ssr = false;
export const trailingSlash = 'always';

export async function load({ fetch }) {
  const config: Config = await getConfig(fetch);
  const deviceName: string = config.deviceName;
  return { deviceName }
}
