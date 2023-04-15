import { getPath, addTimeToFail, mergeMetrics } from "./src";
import {
  CPU_HEADER,
  MEM_HEADER,
  NET_HEADER,
} from "./src/constant/metrics-header";
import { MetricType } from "./src/enum/metrics-type";
import { IMetric } from "./src/interfaces/metric";
import { standardizeUnit } from "./src/unit";

const main = async () => {
  console.log("Starting getting path...");
  const path = getPath();

  console.log("Generating metrics file...");
  const metrics: IMetric[] = [
    { type: MetricType.CPU, path: path.cpuFile, headers: CPU_HEADER },
    { type: MetricType.Memory, path: path.memoryFile, headers: MEM_HEADER },
    { type: MetricType.Network, path: path.networkFile, headers: NET_HEADER },
  ];
  await mergeMetrics(metrics, path);

  console.log("Adding timestamp output to metrics file...");
  await addTimeToFail(metrics, path);

  console.log("Standardize unit...");
  await standardizeUnit(path);

  console.log("Done!");
};

main();
