import { CPU_HEADER, MEM_HEADER, NET_HEADER } from "./constant/metrics-header";
import { IMetric } from "./interfaces/metric";
import { IPath } from "./interfaces/path";
import { streamMergeMetrics } from "./stream/metric";

export const mergeMetrics = async (
  metrics: IMetric[],
  path: IPath
): Promise<void> => {
  await streamMergeMetrics(metrics, path.metricsFile);
};
