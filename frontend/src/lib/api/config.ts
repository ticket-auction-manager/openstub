export type Config = {
  deviceName: string;
  deviceID: string;
}

export async function getConfig(fetch: Function): Promise<Config> {
  const res: Response = await fetch('/api/config');
  if (!res.ok) {
    const blank: Config = {} as Config;
    return blank;
  }
  const data: Config = await res.json();
  return data;
}
