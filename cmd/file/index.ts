import { getPath, addTimeToFail, mergeMetrics } from "./src";
import {
  CPU_HEADER,
  MEM_HEADER,
  NET_HEADER,
} from "./src/constant/metrics-header";
import { IMetric } from "./src/interfaces/metric";

const main = async () => {
  console.log("Starting getting path...");
  const path = getPath();

  console.log("Generating metrics file...");
  const metrics: IMetric[] = [
    { path: path.cpuFile, headers: CPU_HEADER },
    { path: path.memoryFile, headers: MEM_HEADER },
    { path: path.networkFile, headers: NET_HEADER },
  ];
  await mergeMetrics(metrics, path.metricsFile);

  console.log("Adding timestamp output to metrics file...");
  await addTimeToFail(path, metrics);

  console.log("Done!");
};

main();
