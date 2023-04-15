import { CPU_HEADER, MEM_HEADER, NET_HEADER } from "./constant/metrics-header";
import { IMetric } from "./interfaces/metric";
import { IPath } from "./interfaces/path";
import { streamMergeMetrics } from "./stream/metric";

export const mergeMetrics = async (
  metrics: IMetric[],
  mergedPath: string
): Promise<void> => {
  await streamMergeMetrics(metrics, mergedPath);
};
