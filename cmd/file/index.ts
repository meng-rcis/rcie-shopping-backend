import { getPath, addTimeToFail, mergeMetrics, generateMetrics } from "./src";
import {
  CPU_HEADER,
  LATENCY_HEADER,
  MEM_HEADER,
  NET_HEADER,
  RESPONSE_HEADER,
  TPS_HEADER,
} from "./src/constant/metrics-header";
import { MetricType } from "./src/enum/metrics-type";
import { IMetric } from "./src/interfaces/metric";
import { standardizeUnit } from "./src/unit";

const main = async () => {
  console.log("Starting getting path...");
  const path = getPath();

  console.log("Generating metrics file...");
  const metrics = generateMetrics(path);
  await mergeMetrics(metrics, path);

  console.log("Adding timestamp output to metrics file...");
  await addTimeToFail(metrics, path);

  console.log("Standardize unit...");
  await standardizeUnit(path);

  console.log("Done!");
};

main();
