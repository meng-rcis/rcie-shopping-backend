interface ICPU {
  iowait: string;
  irq: string;
  nice: string;
  softirq: string;
  steal: string;
  system: string;
  user: string;
}

interface IMemory {
  used: string;
}

interface INetwork {
  br: string;
  docker0: string;
  eth0: string;
  eth1: string;
  lo: string;
  veth: string;
}

export interface IStandardUnit {
  cpu: ICPU;
  memory: IMemory;
  network: INetwork;
}
