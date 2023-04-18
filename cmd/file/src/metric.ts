import {
  CPU_HEADER,
  LATENCY_HEADER,
  MEM_HEADER,
  NET_HEADER,
  RESPONSE_HEADER,
  TPS_HEADER,
} from "./constant/metrics-header";
import { MetricType } from "./enum/metrics-type";
import { IMetric } from "./interfaces/metric";
import { IPath } from "./interfaces/path";
import { streamMergeMetrics } from "./stream/metric";

export const generateMetrics = (path: IPath): IMetric[] => {
  const metrics: IMetric[] = [
    { type: MetricType.CPU, path: path.cpuFile, headers: CPU_HEADER },
    { type: MetricType.Memory, path: path.memoryFile, headers: MEM_HEADER },
    { type: MetricType.Network, path: path.networkFile, headers: NET_HEADER },
    { type: MetricType.TPS, path: path.tpsFile, headers: TPS_HEADER },
    {
      type: MetricType.Latency,
      path: path.latencyFile,
      headers: LATENCY_HEADER,
    },
    {
      type: MetricType.Response,
      path: path.responseFile,
      headers: RESPONSE_HEADER,
    },
  ];
  return metrics;
};

export const mergeMetrics = async (
  metrics: IMetric[],
  path: IPath
): Promise<void> => {
  await streamMergeMetrics(metrics, path.metricsFile);
};
