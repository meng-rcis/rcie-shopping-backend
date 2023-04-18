export const TIME_HEADER = "Time";
export const TIMESTAMP_HEADER = "Timestamp";
export const TIME_TO_FAIL_HEADER = "TimeToFail";
export const CPU_HEADER = [
  "iowait",
  "irq",
  "nice",
  "softirq",
  "steal",
  "system",
  "user",
];
export const MEM_HEADER = ["Used", "Buffer", "Cached", "Free"];
export const NET_HEADER = [
  "br-10cd1f6f3899 receive",
  "docker0 receive",
  "eth0 receive",
  "eth1 receive",
  "lo receive",
  "veth6750e95 receive",
  "br-10cd1f6f3899 transmit",
  "docker0 transmit",
  "eth0 transmit",
  "eth1 transmit",
  "lo transmit",
  "veth6750e95 transmit",
];
export const TPS_HEADER = ["total-req-rate", "ok-req-rate", "error-req-rate"];
export const RESPONSE_HEADER = ["avg-res-time"];
export const LATENCY_HEADER = [
  "latency-50-percentile",
  "latency-95-percentile",
  "latency-99-percentile",
];
