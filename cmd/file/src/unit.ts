import { IStandardUnit } from "./interfaces/standard-unit";

const STANDARD_UNIT: IStandardUnit = {
  cpu: {
    iowait: "%",
    irq: "%",
    nice: "%",
    softirq: "%",
    steal: "%",
    system: "%",
    user: "%",
  },
  memory: {
    used: "MiB",
  },
  network: {
    br: "MB/s",
    docker0: "MB/s",
    eth0: "MB/s",
    eth1: "MB/s",
    lo: "MB/s",
    veth: "MB/s",
  },
};

export const convertUnit = async (path: string): Promise<string> => {
  const test = " ";
  return new Promise(() => "");
};
